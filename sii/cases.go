package sii

type cases struct {
	repository Repository
}

func NewCases(repository Repository) Cases {
	return &cases{
		repository: repository,
	}
}

func (c *cases) GetContributorInfo(user *User) (ContributorInfo, error) {
	return c.repository.GetContributorInfo(user)
}
