# go-sii-info

## Table of Contents

- [go-sii-info](#go-sii-info)
  - [Table of Contents](#table-of-contents)
  - [About](#about)
  - [Localhost](#localhost)
  - [Usage](#usage)
    - [Http GET](#http-get)
      - [Result](#result)

## About

This is a simple Go library for getting contributor information from the [SII](https://www.sii.cl/) (Servicio de Impuestos Internos) of Chile.

## Localhost

```bash
go run main.go
```

## Usage

### Http GET

```bash
localhost:3000/sii/contributor-info/RUT/76192083-9
```
#### Result

```json
{
    "identifier_type": "RUT",
    "identifier_number": "76192083",
    "verification_code": "9",
    "commerce_name": "SASCO SPA ",
    "is_initiated_activities": true,
    "is_available_to_pay_tax_in_foreign_currency": false,
    "is_smaller_company": true,
    "commercial_activities": [
        {
            "name": "OTRAS ACTIVIDADES DE TELECOMUNICACIONES N.C.P.",
            "code": 619090,
            "category": 1,
            "is_vat_affected": false
        },
        {
            "name": "ACTIVIDADES DE CONSULTORIA DE INFORMATICA Y DE GESTION DE INSTALACIONE",
            "code": 620200,
            "category": 1,
            "is_vat_affected": false
        }
    ],
    "stamped_documents": [
        {
            "name": "Factura Electronica",
            "last_year_stamp": 2023
        },
        {
            "name": "Factura No Afecta O Exenta Electronica",
            "last_year_stamp": 2015
        },
        {
            "name": "Boleta Electronica",
            "last_year_stamp": 2023
        },
        {
            "name": "Boleta Exenta Electronica",
            "last_year_stamp": 2016
        },
        {
            "name": "Factura Compra Electronica",
            "last_year_stamp": 2023
        },
        {
            "name": "Guia Despacho Electronica",
            "last_year_stamp": 2016
        },
        {
            "name": "Nota Debito Electronica",
            "last_year_stamp": 2023
        },
        {
            "name": "Nota Credito Electronica",
            "last_year_stamp": 2023
        },
        {
            "name": "Boleta De Terceros Electronicas",
            "last_year_stamp": 2023
        },
        {
            "name": "Factura De Exportacion Electronica",
            "last_year_stamp": 2021
        }
    ]
}
```
