package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/burakturnaa/mailoop.git/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type companyRepository struct {
	CompanyCollection *mongo.Collection
}

type CompanyRepository interface {
	GetAll() ([]models.Company, error)
	GetOne(id primitive.ObjectID) (models.Company, error)
	InsertCompany(company models.Company) (models.Company, error)
	UpdateCompany(company models.Company) (models.Company, error)
	DeleteCompany(id primitive.ObjectID) (bool, error)
	FindByCompanyID(companyId primitive.ObjectID) (models.Company, error)
	FindByCompanyEmail(companyEmail string) (models.Company, error)
}

func NewCompanyRepository(dbClient *mongo.Collection) CompanyRepository {
	return &companyRepository{CompanyCollection: dbClient}
}

func (mt *companyRepository) GetAll() ([]models.Company, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var companyies []models.Company
	cursor, err := mt.CompanyCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &companyies); err != nil {
		return nil, err
	}
	return companyies, nil
}

func (mt *companyRepository) GetOne(id primitive.ObjectID) (models.Company, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var companyies models.Company
	err := mt.CompanyCollection.FindOne(ctx, bson.M{"id": id}).Decode(&companyies)
	fmt.Println(err)
	if err != nil {
		return companyies, err
	}
	return companyies, nil
}

func (mt *companyRepository) InsertCompany(company models.Company) (models.Company, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	company.Id = primitive.NewObjectID()
	company.CreatedAt = time.Now()
	_, err := mt.CompanyCollection.InsertOne(ctx, company)
	if err != nil {
		return company, err
	}

	return company, nil
}

func (mt *companyRepository) UpdateCompany(company models.Company) (models.Company, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"id", company.Id}}
	update := bson.D{{"$set", bson.D{{"name", company.Name}, {"email", company.Email}, {"phone", company.Phone}, {"location", company.Location}, {"website", company.Website}, {"updated_at", time.Now()}}}}
	_, err := mt.CompanyCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return company, err
	}
	return company, nil
}

func (mt *companyRepository) DeleteCompany(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := mt.CompanyCollection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil || result.DeletedCount < 1 {
		return false, err
	}
	return true, nil
}

func (mt *companyRepository) FindByCompanyID(companyId primitive.ObjectID) (models.Company, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var company models.Company
	err := mt.CompanyCollection.FindOne(ctx, bson.M{"id": companyId}).Decode(&company)
	if err != nil {
		return company, err
	}
	return company, nil
}

func (mt *companyRepository) FindByCompanyEmail(companyEmail string) (models.Company, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var company models.Company
	err := mt.CompanyCollection.FindOne(ctx, bson.M{"email": companyEmail}).Decode(&company)
	if err != nil {
		return company, err
	}
	return company, nil
}
