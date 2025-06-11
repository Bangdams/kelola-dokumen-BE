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

type BirthCertificateDocumentUsecase interface {
	Create(ctx context.Context, request *model.BirthCertificateDocumentRequest) (*model.BirthCertificateDocumentResponse, error)
	Delete(ctx context.Context, birthDocumentId uint) error
	FindById(ctx context.Context, birthDocumentId uint) (*model.BirthCertificateDocumentResponse, error)
	FindAll(ctx context.Context) (*[]model.BirthCertificateDocumentResponse, error)
}

type BirthCertificateDocumentUsecaseImpl struct {
	BirthCertificateDocumentRepo repository.BirthCertificateDocumentRepository
	StatementTypeItemRepo        repository.StatementTypeItemRepository
	DB                           *gorm.DB
	Validate                     *validator.Validate
}

func NewBirthCertificateDocumentUsecase(statementTypeItemRepo repository.StatementTypeItemRepository, birthCertificateDocumentRepo repository.BirthCertificateDocumentRepository, DB *gorm.DB, validate *validator.Validate) BirthCertificateDocumentUsecase {
	return &BirthCertificateDocumentUsecaseImpl{
		BirthCertificateDocumentRepo: birthCertificateDocumentRepo,
		StatementTypeItemRepo:        statementTypeItemRepo,
		DB:                           DB,
		Validate:                     validate,
	}
}

// Create implements BirthCertificateDocumentUsecase.
func (birthDocumentUsecase *BirthCertificateDocumentUsecaseImpl) Create(ctx context.Context, request *model.BirthCertificateDocumentRequest) (*model.BirthCertificateDocumentResponse, error) {
	tx := birthDocumentUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	errorResponse := &model.ErrorResponse{}

	err := birthDocumentUsecase.Validate.Struct(request)
	if err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Field '%s' failed on '%s' rule", e.Field(), e.Tag())
			validationErrors = append(validationErrors, msg)
		}

		errorResponse.Message = "invalid request parameter"
		errorResponse.Details = validationErrors

		jsonString, _ := json.Marshal(errorResponse)

		log.Println("error create birthDocumentUsecase : ", err)

		return nil, fiber.NewError(fiber.ErrBadRequest.Code, string(jsonString))
	}

	err = birthDocumentUsecase.StatementTypeItemRepo.FindById(tx, &entity.StatementTypeItem{
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

	birthDocument := &entity.BirthCertificateDocument{
		StatementTypeItemId:    request.StatementTypeItemId,
		UserId:                 request.UserId,
		RtRwFilePath:           request.RtRwFilePath,
		FormulirFilePath:       request.FormulirFilePath,
		SuratKelahiranFilePath: request.SuratKelahiranFilePath,
		BukuNikahFilePath:      request.BukuNikahFilePath,
		KkFilePath:             request.KkFilePath,
		KtpPelaporFilePath:     request.KtpPelaporFilePath,
		KtpSaksi1FilePath:      request.KtpSaksi1FilePath,
		KtpSaksi2FilePath:      request.KtpSaksi2FilePath,
		KtpOrangTuaFilePath:    request.KtpOrangTuaFilePath,
	}

	err = birthDocumentUsecase.BirthCertificateDocumentRepo.Create(tx, birthDocument)
	if err != nil {
		log.Println("failed when create repo birthDocumentUsecase : ", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success create from usecase birthDocumentUsecase")

	return converter.BirthCertificateDocumentToResponse(birthDocument), nil
}

// Delete implements BirthCertificateDocumentUsecase.
func (birthDocumentUsecase *BirthCertificateDocumentUsecaseImpl) Delete(ctx context.Context, birthDocumentId uint) error {
	tx := birthDocumentUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	birthDocument := &entity.BirthCertificateDocument{}
	birthDocument.ID = birthDocumentId

	err := birthDocumentUsecase.BirthCertificateDocumentRepo.FindById(tx, birthDocument)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse := model.ErrorResponse{
				Message: "birthDocument data was not found",
				Details: []string{},
			}
			jsonString, _ := json.Marshal(errorResponse)

			log.Println("error birthDocument : ", err)

			return fiber.NewError(fiber.ErrNotFound.Code, string(jsonString))
		}

		log.Println("error birthDocument : ", err)
		return fiber.ErrInternalServerError
	}

	err = birthDocumentUsecase.BirthCertificateDocumentRepo.Delete(tx, birthDocument)
	if err != nil {
		log.Println("failed when delete repo birthDocument : ", err)
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return fiber.ErrInternalServerError
	}

	log.Println("success delete from usecase birthDocument")

	return nil
}

// FindAll implements BirthCertificateDocumentUsecase.
func (birthDocumentUsecase *BirthCertificateDocumentUsecaseImpl) FindAll(ctx context.Context) (*[]model.BirthCertificateDocumentResponse, error) {
	tx := birthDocumentUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var birthDocuments = &[]entity.BirthCertificateDocument{}
	err := birthDocumentUsecase.BirthCertificateDocumentRepo.FindAll(tx, birthDocuments)
	if err != nil {
		log.Println("failed when find all repo birthDocument : ", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success find all from usecase birthDocument")

	return converter.BirthCertificateDocumentToResponses(birthDocuments), nil
}

// FindById implements BirthCertificateDocumentUsecase.
func (birthDocumentUsecase *BirthCertificateDocumentUsecaseImpl) FindById(ctx context.Context, birthDocumentId uint) (*model.BirthCertificateDocumentResponse, error) {
	tx := birthDocumentUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	birthDocument := new(entity.BirthCertificateDocument)
	birthDocument.ID = birthDocumentId

	if err := birthDocumentUsecase.BirthCertificateDocumentRepo.FindById(tx, birthDocument); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse := model.ErrorResponse{
				Message: "birthDocument data was not found",
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

	log.Println("success find by id from usecase birthDocument")

	return converter.BirthCertificateDocumentToResponse(birthDocument), nil
}
