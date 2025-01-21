package services

import (
	"github.com/burakturnaa/mailoop.git/app/dto"
	"github.com/burakturnaa/mailoop.git/app/models"
	"github.com/burakturnaa/mailoop.git/app/repository"
	_log "github.com/burakturnaa/mailoop.git/app/services/log"

	"github.com/mashingan/smapping"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogService interface {
	GetAll() (*[]_log.LogResponse, error)
	GetOne(id primitive.ObjectID) (*_log.LogResponse, error)
	CreateLog(logRequest dto.MailSenderBody) (*_log.LogResponse, error)
}

type logService struct {
	logRepo repository.LogRepository
}

func NewLogService(logRepo repository.LogRepository) LogService {
	return &logService{logRepo: logRepo}
}

func (mt *logService) GetAll() (*[]_log.LogResponse, error) {
	logs, err := mt.logRepo.GetAll()
	if err != nil {
		return nil, err
	}

	response := _log.NewLogArrayResponse(logs)
	return &response, nil
}

func (mt *logService) GetOne(id primitive.ObjectID) (*_log.LogResponse, error) {
	logs, err := mt.logRepo.GetOne(id)
	if err != nil {
		return nil, err
	}

	response := _log.NewLogResponse(logs)
	return &response, nil
}

func (mt *logService) CreateLog(logRequest dto.MailSenderBody) (*_log.LogResponse, error) {

	var log models.Log
	err := smapping.FillStruct(&log, smapping.MapFields(&logRequest))
	if err != nil {
		return nil, err
	}
	log, _ = mt.logRepo.InsertLog(log)

	res := _log.NewLogResponse(log)
	return &res, nil
}
