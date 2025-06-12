package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"github.com/Bangdams/kelola-dokumen-BE/internal/model"
	"github.com/Bangdams/kelola-dokumen-BE/internal/model/converter"
	"github.com/Bangdams/kelola-dokumen-BE/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DeathCertificateDocumentUsecase interface {
	Create(ctx context.Context, request *model.DeathCertificateDocumentRequest) (*model.DeathCertificateDocumentResponse, error)
	Delete(ctx context.Context, deathDocumentId uint) error
	FindById(ctx context.Context, deathDocumentId uint) (*model.DeathCertificateDocumentResponse, error)
	FindAll(ctx context.Context) (*[]model.DeathCertificateDocumentResponse, error)
	Update(ctx context.Context, deathDocumentId uint, suratBalasan string) (*model.DeathCertificateDocumentResponse, error)
}

type DeathCertificateDocumentUsecaseImpl struct {
	DeathCertificateDocumentRepo repository.DeathCertificateDocumentRepository
	StatementTypeItemRepo        repository.StatementTypeItemRepository
	DB                           *gorm.DB
	Validate                     *validator.Validate
}

func NewDeathCertificateDocumentUsecase(statementTypeItemRepo repository.StatementTypeItemRepository, deathCertificateDocumentRepository repository.DeathCertificateDocumentRepository, DB *gorm.DB, validate *validator.Validate) DeathCertificateDocumentUsecase {
	return &DeathCertificateDocumentUsecaseImpl{
		DeathCertificateDocumentRepo: deathCertificateDocumentRepository,
		StatementTypeItemRepo:        statementTypeItemRepo,
		DB:                           DB,
		Validate:                     validate,
	}
}

// Update implements DeathCertificateDocumentUsecase.
func (deathDocumentUsecase *DeathCertificateDocumentUsecaseImpl) Update(ctx context.Context, deathDocumentId uint, suratBalasan string) (*model.DeathCertificateDocumentResponse, error) {
	tx := deathDocumentUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	errorResponse := &model.ErrorResponse{}

	deathDocument := &entity.DeathCertificateDocument{
		ID: deathDocumentId,
	}

	err := deathDocumentUsecase.DeathCertificateDocumentRepo.FindById(tx, deathDocument)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse.Message = "statement type item data was not found"
			errorResponse.Details = []string{}

			jsonString, _ := json.Marshal(errorResponse)
			log.Println("Data not found")

			return nil, fiber.NewError(fiber.ErrNotFound.Code, string(jsonString))
		}

		log.Println("error find by id : ", err)
		return nil, fiber.ErrInternalServerError
	}

	deathDocument.Status = "selesai"
	deathDocument.SuratBalasanFilePath = suratBalasan

	err = deathDocumentUsecase.DeathCertificateDocumentRepo.Update(tx, deathDocument)
	if err != nil {
		log.Println("Failed to update : ", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success update from usecase deathDocument")

	return converter.DeathCertificateDocumentToResponse(deathDocument), nil
}

// Create implements DeathCertificateDocumentUsecase.
func (deathDocumentUsecase *DeathCertificateDocumentUsecaseImpl) Create(ctx context.Context, request *model.DeathCertificateDocumentRequest) (*model.DeathCertificateDocumentResponse, error) {
	tx := deathDocumentUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	errorResponse := &model.ErrorResponse{}

	err := deathDocumentUsecase.Validate.Struct(request)
	if err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Field '%s' failed on '%s' rule", e.Field(), e.Tag())
			validationErrors = append(validationErrors, msg)
		}

		errorResponse.Message = "invalid request parameter"
		errorResponse.Details = validationErrors

		jsonString, _ := json.Marshal(errorResponse)

		log.Println("error create deathDocumentUsecase : ", err)

		return nil, fiber.NewError(fiber.ErrBadRequest.Code, string(jsonString))
	}

	err = deathDocumentUsecase.StatementTypeItemRepo.FindById(tx, &entity.StatementTypeItem{
		ID: request.StatementTypeItemId,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse := model.ErrorResponse{
				Message: "StatementTypeItem type data was not found",
				Details: []string{},
			}
			jsonString, _ := json.Marshal(errorResponse)

			log.Println("error StatementTypeItem : ", err)

			return nil, fiber.NewError(fiber.ErrNotFound.Code, string(jsonString))
		}

		log.Println("error StatementTypeItem : ", err)
		return nil, fiber.ErrInternalServerError
	}

	deathDocument := &entity.DeathCertificateDocument{
		StatementTypeItemId:   request.StatementTypeItemId,
		UserId:                request.UserId,
		RtRwFilePath:          request.RtRwFilePath,
		FormulirFilePath:      request.FormulirFilePath,
		SuratKematianFilePath: request.SuratKematianFilePath,
		KtpFilePath:           request.KtpFilePath,
		KkFilePath:            request.KkFilePath,
		KtpPelaporFilePath:    request.KtpPelaporFilePath,
		KtpSaksi1FilePath:     request.KtpSaksi1FilePath,
		KtpSaksi2FilePath:     request.KtpSaksi2FilePath,
		BukuNikahFilePath:     request.BukuNikahFilePath,
	}

	err = deathDocumentUsecase.DeathCertificateDocumentRepo.Create(tx, deathDocument)
	if err != nil {
		log.Println("failed when create repo deathDocumentUsecase : ", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success create from usecase deathDocumentUsecase")

	return converter.DeathCertificateDocumentToResponse(deathDocument), nil
}

// Delete implements DeathCertificateDocumentUsecase.
func (deathDocumentUsecase *DeathCertificateDocumentUsecaseImpl) Delete(ctx context.Context, deathDocumentId uint) error {
	tx := deathDocumentUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	deathDocument := &entity.DeathCertificateDocument{}
	deathDocument.ID = deathDocumentId

	err := deathDocumentUsecase.DeathCertificateDocumentRepo.FindById(tx, deathDocument)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse := model.ErrorResponse{
				Message: "deathDocument data was not found",
				Details: []string{},
			}
			jsonString, _ := json.Marshal(errorResponse)

			log.Println("error deathDocument : ", err)

			return fiber.NewError(fiber.ErrNotFound.Code, string(jsonString))
		}

		log.Println("error deathDocument : ", err)
		return fiber.ErrInternalServerError
	}

	err = deathDocumentUsecase.DeathCertificateDocumentRepo.Delete(tx, deathDocument)
	if err != nil {
		log.Println("failed when delete repo deathDocument : ", err)
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return fiber.ErrInternalServerError
	}

	log.Println("success delete from usecase deathDocument")

	return nil
}

// FindAll implements DeathCertificateDocumentUsecase.
func (deathDocumentUsecase *DeathCertificateDocumentUsecaseImpl) FindAll(ctx context.Context) (*[]model.DeathCertificateDocumentResponse, error) {
	tx := deathDocumentUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var deathDocuments = &[]entity.DeathCertificateDocument{}
	err := deathDocumentUsecase.DeathCertificateDocumentRepo.FindAll(tx, deathDocuments)
	if err != nil {
		log.Println("failed when find all repo deathDocument : ", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success find all from usecase deathDocument")

	return converter.DeathCertificateDocumentToResponses(deathDocuments), nil
}

// FindById implements DeathCertificateDocumentUsecase.
func (deathDocumentUsecase *DeathCertificateDocumentUsecaseImpl) FindById(ctx context.Context, deathDocumentId uint) (*model.DeathCertificateDocumentResponse, error) {
	tx := deathDocumentUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	deathDocument := new(entity.DeathCertificateDocument)
	deathDocument.ID = deathDocumentId

	if err := deathDocumentUsecase.DeathCertificateDocumentRepo.FindById(tx, deathDocument); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse := model.ErrorResponse{
				Message: "deathDocument data was not found",
				Details: []string{},
			}

			jsonString, _ := json.Marshal(errorResponse)

			log.Println("error find by id usecase : ", err)

			return nil, fiber.NewError(fiber.ErrNotFound.Code, string(jsonString))
		} else {
			log.Println("Error find by id usecase:", err)
			return nil, fiber.ErrInternalServerError
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success find by id from usecase deathDocument")

	return converter.DeathCertificateDocumentToResponse(deathDocument), nil
}
