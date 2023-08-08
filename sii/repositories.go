package sii

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/xbizzybone/go-sii-info/utils"
	"go.uber.org/zap"
)

type repository struct {
	logger *zap.Logger
}

func NewRepository(logger *zap.Logger) Repository {
	return &repository{
		logger: logger,
	}
}

/*
 * GetContributorInfo fetches the contributor information from SII. Using the provided user information.
 */
func (r *repository) GetContributorInfo(user *User) (ContributorInfo, error) {
	captchaCode, captcha, err := r.fetchCaptcha()
	if err != nil {
		r.logger.Error("Error fetching captcha", zap.Error(err))
		return ContributorInfo{}, err
	}

	form := setForm(user, captchaCode, captcha)

	doc, err := r.postContributorInfoForm(form)
	if err != nil {
		r.logger.Error("Error posting contributor info form", zap.Error(err))
		return ContributorInfo{}, err
	}

	return r.extractContributorInfo(doc, user), nil
}

func setForm(user *User, captchaCode string, captcha string) url.Values {
	form := url.Values{}
	form.Add("RUT", user.IdentifierNumber)
	form.Add("DV", user.VerificationCode)
	form.Add("PRG", "STC")
	form.Add("OPC", "NOR")
	form.Add("txt_code", captchaCode)
	form.Add("txt_captcha", captcha)
	return form
}

func (r *repository) postContributorInfoForm(form url.Values) (*goquery.Document, error) {
	resp, err := http.PostForm(os.Getenv("CONTRIBUTOR_INFO_URL"), form)
	if err != nil {
		r.logger.Error("Error posting contributor info form", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	return goquery.NewDocumentFromReader(resp.Body)
}

func (r *repository) extractContributorInfo(doc *goquery.Document, user *User) ContributorInfo {
	result := ContributorInfo{
		IdentifierType:   user.IdentifierType,
		IdentifierNumber: user.IdentifierNumber,
		VerificationCode: user.VerificationCode,
	}

	initActivitiesSpan := utils.ExtractLastCharacters(doc.Find("span").Eq(1).Text(), 2)
	foreignTaxSpan := utils.ExtractLastCharacters(doc.Find("span").Eq(3).Text(), 3)
	smallerCompanySpan := utils.ExtractLastCharacters(doc.Find("span").Eq(4).Text(), 3)

	result.CommerceName = doc.Find("div").Eq(5).Text()
	result.IsInitiatedActivities = utils.GetBooleanFromYesNoString(initActivitiesSpan)
	result.IsAvailableToPayTaxInForeignCurrency = utils.GetBooleanFromYesNoString(foreignTaxSpan)
	result.IsSmallerCompany = utils.GetBooleanFromYesNoString(smallerCompanySpan)

	result.CommercialActivities = r.GetCommercialActivitiesFromDocument(doc)
	result.StampedDocuments = r.GetStampedDocumentsFromDocument(doc)

	return result
}

/*
Fetch captcha from SII.
This is necessary to make the request to get the contributor info. The captcha is a base64 encoded image. The code is the last 4 digits of the captcha.
*/
func (r *repository) fetchCaptcha() (string, string, error) {
	resp, err := http.PostForm(os.Getenv("CAPTCHA_URL"), url.Values{"oper": {"0"}})
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	var data CaptchaResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", "", err
	}

	code, err := base64.StdEncoding.DecodeString(data.TxtCaptcha)

	if err != nil {
		return "", "", err
	}

	return string(code[36:40]), data.TxtCaptcha, nil
}

func (r *repository) GetCommercialActivitiesFromDocument(doc *goquery.Document) []CommercialActivity {
	var commercialActivities []CommercialActivity

	doc.Find("table").Eq(0).Find("tr").Each(func(indexTr int, rowHtml *goquery.Selection) {
		if indexTr == 0 { // Skip header row
			return
		}

		var commercialActivity CommercialActivity
		rowHtml.Find("td").Each(func(indexTd int, cellHtml *goquery.Selection) {
			text := strings.TrimSpace(cellHtml.Text())
			switch indexTd {
			case 0: // Name
				commercialActivity.Name = text
			case 1: // Code
				code, err := strconv.Atoi(text)
				if err != nil {
					r.logger.Error("Error converting code to int", zap.Error(err))
					break
				}
				commercialActivity.Code = code
			case 2: // Category
				commercialActivity.Category = utils.FormatCategory(text)
			case 3: // Is VAT affected
				commercialActivity.IsVATAffected = (text == "SI")
			}
		})

		commercialActivities = append(commercialActivities, commercialActivity)
	})

	return commercialActivities
}

func (r *repository) GetStampedDocumentsFromDocument(doc *goquery.Document) []StampedDocument {
	var stampedDocuments []StampedDocument

	doc.Find("table").Eq(2).Find("tr").Each(func(indexTr int, rowHtml *goquery.Selection) {
		if indexTr == 0 { // Skip header row
			return
		}

		var stampedDocument StampedDocument
		rowHtml.Find("td").Each(func(indexTd int, cellHtml *goquery.Selection) {
			text := strings.TrimSpace(cellHtml.Text())
			switch indexTd {
			case 0: // Document type
				stampedDocument.Name = text
			case 1: // Document year
				year, err := strconv.Atoi(text)
				if err != nil {
					r.logger.Error("Error converting year to int", zap.Error(err))
					break
				}
				stampedDocument.LastYearStamp = year
			}
		})

		stampedDocuments = append(stampedDocuments, stampedDocument)
	})

	return stampedDocuments
}
