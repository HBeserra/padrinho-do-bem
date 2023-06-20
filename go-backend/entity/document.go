package entity

import (
	"errors"
	"fmt"
	"github.com/klassmann/cpfcnpj"
)

type DocumentType string

const (
	CPF  DocumentType = "CPF"
	CNPJ              = "CNPJ"
)

type Document struct {
	number       string
	documentType DocumentType
}

func NewDocument(rawDocNumber string) (*Document, error) {

	doc := new(Document)

	doc.number = cpfcnpj.Clean(rawDocNumber)

	switch len(doc.number) {
	case 11:
		if !cpfcnpj.ValidateCPF(doc.number) {
			return nil, errors.New("invalid CPF document")
		}

		doc.documentType = CPF
		break
	case 14:
		if !cpfcnpj.ValidateCNPJ(doc.number) {
			return nil, errors.New("invalid CNPJ document")
		}

		doc.documentType = CNPJ
		break
	default:
		return nil, errors.New("invalid document number")
	}

	return doc, nil
}

func (d Document) GetDocumentStr() string {

	var str = fmt.Sprintf("%d", d.number)

	switch d.documentType {
	case CPF:
		return fmt.Sprintf("%s.%s.%s-%s", str[0:3], str[3:6], str[6:9], str[9:])
	case CNPJ:
		return fmt.Sprintf("%s.%s.%s/%s-%s", str[:2], str[2:5], str[5:8], str[8:12], str[12:])
	default:
		return ""
	}
}
