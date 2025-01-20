package services

import (
	"github.com/burakturnaa/mailoop.git/app/dto"
	"github.com/burakturnaa/mailoop.git/app/models"
	"github.com/burakturnaa/mailoop.git/app/repository"
	_company "github.com/burakturnaa/mailoop.git/app/services/company"

	"github.com/mashingan/smapping"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyService interface {
	GetAll() (*[]_company.CompanyResponse, error)
	GetOne(id primitive.ObjectID) (*_company.CompanyResponse, error)
	CreateCompany(companyRequest dto.CompanyBody) (*_company.CompanyResponse, error)
	UpdateCompany(UpdateCompanyRequest dto.UpdateCompanyBody) (*_company.CompanyResponse, error)
	DeleteCompany(id primitive.ObjectID) (bool, error)
	FindCompanyByID(companyId primitive.ObjectID) (*_company.CompanyResponse, error)
}

type companyService struct {
	companyRepo repository.CompanyRepository
}

func NewCompanyService(companyRepo repository.CompanyRepository) CompanyService {
	return &companyService{companyRepo: companyRepo}
}

func (mt *companyService) GetAll() (*[]_company.CompanyResponse, error) {
	companys, err := mt.companyRepo.GetAll()
	if err != nil {
		return nil, err
	}

	response := _company.NewCompanyArrayResponse(companys)
	return &response, nil
}

func (mt *companyService) GetOne(id primitive.ObjectID) (*_company.CompanyResponse, error) {
	companys, err := mt.companyRepo.GetOne(id)
	if err != nil {
		return nil, err
	}

	response := _company.NewCompanyResponse(companys)
	return &response, nil
}

func (mt *companyService) CreateCompany(companyRequest dto.CompanyBody) (*_company.CompanyResponse, error) {

	var company models.Company
	err := smapping.FillStruct(&company, smapping.MapFields(&companyRequest))
	if err != nil {
		return nil, err
	}
	company, _ = mt.companyRepo.InsertCompany(company)

	res := _company.NewCompanyResponse(company)
	return &res, nil
}

func (mt *companyService) UpdateCompany(companyRequest dto.UpdateCompanyBody) (*_company.CompanyResponse, error) {
	var company models.Company
	err := smapping.FillStruct(&company, smapping.MapFields(&companyRequest))
	if err != nil {
		return nil, err
	}

	company, err = mt.companyRepo.UpdateCompany(company)
	if err != nil {
		return nil, err
	}

	res := _company.NewCompanyResponse(company)
	return &res, nil
}

func (mt *companyService) FindCompanyByID(companyId primitive.ObjectID) (*_company.CompanyResponse, error) {
	company, err := mt.companyRepo.FindByCompanyID(companyId)
	if err != nil {
		return nil, err
	}

	companyResponse := _company.CompanyResponse{}
	err = smapping.FillStruct(&companyResponse, smapping.MapFields(&company))
	if err != nil {
		return nil, err
	}
	return &companyResponse, nil
}

func (mt *companyService) DeleteCompany(id primitive.ObjectID) (bool, error) {
	result, err := mt.companyRepo.DeleteCompany(id)
	if err != nil {
		return result, err
	}
	return result, nil
}
