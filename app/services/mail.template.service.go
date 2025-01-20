package services

import (
	"github.com/burakturnaa/mailoop.git/app/dto"
	"github.com/burakturnaa/mailoop.git/app/models"
	"github.com/burakturnaa/mailoop.git/app/repository"
	_mailTemplate "github.com/burakturnaa/mailoop.git/app/services/mail_template"

	"github.com/mashingan/smapping"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MailTemplateService interface {
	GetAll() (*[]_mailTemplate.MailTemplateResponse, error)
	GetOne(id primitive.ObjectID) (*_mailTemplate.MailTemplateResponse, error)
	CreateMailTemplate(mailTemplateRequest dto.MailTemplateBody) (*_mailTemplate.MailTemplateResponse, error)
	UpdateMailTemplate(UpdateMailTemplateRequest dto.UpdateMailTemplateBody) (*_mailTemplate.MailTemplateResponse, error)
	FindMailTemplateByID(mailTemplateId primitive.ObjectID) (*_mailTemplate.MailTemplateResponse, error)
}

type mailTemplateService struct {
	mailTemplateRepo repository.MailTemplateRepository
}

func NewMailTemplateService(mailTemplateRepo repository.MailTemplateRepository) MailTemplateService {
	return &mailTemplateService{mailTemplateRepo: mailTemplateRepo}
}

func (mt *mailTemplateService) GetAll() (*[]_mailTemplate.MailTemplateResponse, error) {
	mailTemplates, err := mt.mailTemplateRepo.GetAll()
	if err != nil {
		return nil, err
	}

	response := _mailTemplate.NewMailTemplateArrayResponse(mailTemplates)
	return &response, nil
}

func (mt *mailTemplateService) GetOne(id primitive.ObjectID) (*_mailTemplate.MailTemplateResponse, error) {
	mailTemplates, err := mt.mailTemplateRepo.GetOne(id)
	if err != nil {
		return nil, err
	}

	response := _mailTemplate.NewMailTemplateResponse(mailTemplates)
	return &response, nil
}

func (mt *mailTemplateService) CreateMailTemplate(mailTemplateRequest dto.MailTemplateBody) (*_mailTemplate.MailTemplateResponse, error) {

	var mailTemplate models.MailTemplate
	err := smapping.FillStruct(&mailTemplate, smapping.MapFields(&mailTemplateRequest))
	if err != nil {
		return nil, err
	}
	mailTemplate, _ = mt.mailTemplateRepo.InsertMailTemplate(mailTemplate)

	res := _mailTemplate.NewMailTemplateResponse(mailTemplate)
	return &res, nil
}

func (mt *mailTemplateService) UpdateMailTemplate(mailTemplateRequest dto.UpdateMailTemplateBody) (*_mailTemplate.MailTemplateResponse, error) {
	var mailTemplate models.MailTemplate
	err := smapping.FillStruct(&mailTemplate, smapping.MapFields(&mailTemplateRequest))
	if err != nil {
		return nil, err
	}

	mailTemplate, err = mt.mailTemplateRepo.UpdateMailTemplate(mailTemplate)
	if err != nil {
		return nil, err
	}

	res := _mailTemplate.NewMailTemplateResponse(mailTemplate)
	return &res, nil
}

func (mt *mailTemplateService) FindMailTemplateByID(mailTemplateId primitive.ObjectID) (*_mailTemplate.MailTemplateResponse, error) {
	mailTemplate, err := mt.mailTemplateRepo.FindByMailTemplateID(mailTemplateId)
	if err != nil {
		return nil, err
	}

	mailTemplateResponse := _mailTemplate.MailTemplateResponse{}
	err = smapping.FillStruct(&mailTemplateResponse, smapping.MapFields(&mailTemplate))
	if err != nil {
		return nil, err
	}
	return &mailTemplateResponse, nil
}
