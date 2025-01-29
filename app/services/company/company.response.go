package _company

import (
	"github.com/burakturnaa/mailoop.git/app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyResponse struct {
	Id       *primitive.ObjectID `json:"id,omitempty"`
	Name     string              `json:"name,omitempty"`
	Email    string              `json:"email,omitempty"`
	Phone    string              `json:"phone,omitempty"`
	Location string              `json:"location,omitempty"`
	Website  string              `json:"website,omitempty"`
}

func NewCompanyResponse(company models.Company) CompanyResponse {
	return CompanyResponse{
		Id:       &company.Id,
		Name:     company.Name,
		Email:    company.Email,
		Phone:    company.Phone,
		Location: company.Location,
		Website:  company.Website,
	}
}

func NewCompanyArrayResponse(companies []models.Company) []CompanyResponse {
	companyRes := []CompanyResponse{}
	for _, v := range companies {
		Id = v.Id
		p := CompanyResponse{
			Id:       &Id,
			Name:     v.Name,
			Email:    v.Email,
			Phone:    v.Phone,
			Location: v.Location,
			Website:  v.Website,
		}
		companyRes = append(companyRes, p)
	}
	return companyRes
}
