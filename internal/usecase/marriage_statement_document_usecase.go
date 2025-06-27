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

type MarriageStatementDocumentUsecase interface {
	Create(ctx context.Context, request *model.MarriageStatementDocumentRequest) (*model.MarriageStatementDocumentResponse, error)
	Delete(ctx context.Context, marriageDocumentId uint) error
	FindById(ctx context.Context, marriageDocumentId uint) (*model.MarriageStatementDocumentResponse, error)
	FindAll(ctx context.Context) (*[]model.MarriageStatementDocumentResponse, error)
	Update(ctx context.Context, marriageDocumentId uint, suratBalasan string) (*model.MarriageStatementDocumentResponse, error)
}

type MarriageStatementDocumentUsecaseImpl struct {
	MarriageStatementDocumentRepo repository.MarriageStatementDocumentRepository
	StatementTypeItemRepo         repository.StatementTypeItemRepository
	DB                            *gorm.DB
	Validate                      *validator.Validate
}

func NewMarriageStatementDocumentUsecase(statementTypeItemRepo repository.StatementTypeItemRepository, deathCertificateDocumentRepository repository.MarriageStatementDocumentRepository, DB *gorm.DB, validate *validator.Validate) MarriageStatementDocumentUsecase {
	return &MarriageStatementDocumentUsecaseImpl{
		MarriageStatementDocumentRepo: deathCertificateDocumentRepository,
		StatementTypeItemRepo:         statementTypeItemRepo,
		DB:                            DB,
		Validate:                      validate,
	}
}

// Update implements MarriageStatementDocumentUsecase.
func (marriageDocumentUsecase *MarriageStatementDocumentUsecaseImpl) Update(ctx context.Context, marriageDocumentId uint, suratBalasan string) (*model.MarriageStatementDocumentResponse, error) {
	tx := marriageDocumentUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	errorResponse := &model.ErrorResponse{}

	marriageDocument := &entity.MarriageStatementDocument{
		ID: marriageDocumentId,
	}

	err := marriageDocumentUsecase.MarriageStatementDocumentRepo.FindById(tx, marriageDocument)
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

	marriageDocument.Status = "selesai"
	marriageDocument.SuratBalasanFilePath = suratBalasan

	err = marriageDocumentUsecase.MarriageStatementDocumentRepo.Update(tx, marriageDocument)
	if err != nil {
		log.Println("Failed to update : ", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success update from usecase marriageDocument")

	return converter.MarriageStatementDocumentToResponse(marriageDocument), nil
}

// Create implements MarriageStatementDocumentUsecase.
func (marriageDocumentUsecase *MarriageStatementDocumentUsecaseImpl) Create(ctx context.Context, request *model.MarriageStatementDocumentRequest) (*model.MarriageStatementDocumentResponse, error) {
	tx := marriageDocumentUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	errorResponse := &model.ErrorResponse{}

	err := marriageDocumentUsecase.Validate.Struct(request)
	if err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Field '%s' failed on '%s' rule", e.Field(), e.Tag())
			validationErrors = append(validationErrors, msg)
		}

		errorResponse.Message = "invalid request parameter"
		errorResponse.Details = validationErrors

		jsonString, _ := json.Marshal(errorResponse)

		log.Println("error create marriageDocumentUsecase : ", err)

		return nil, fiber.NewError(fiber.ErrBadRequest.Code, string(jsonString))
	}

	err = marriageDocumentUsecase.StatementTypeItemRepo.FindById(tx, &entity.StatementTypeItem{
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

	marriageDocument := &entity.MarriageStatementDocument{
		StatementTypeItemId: request.StatementTypeItemId,
		UserId:              request.UserId,
		KkFilePath:          request.KkFilePath,
		KtpFilePath:         request.KtpFilePath,
		IjazahFilePath:      request.IjazahFilePath,
		AktaFilePath:        request.AktaFilePath,
		KtpSaksiFilePath:    request.KtpSaksiFilePath,
		Status:              "ditunggu",
	}

	err = marriageDocumentUsecase.MarriageStatementDocumentRepo.Create(tx, marriageDocument)
	if err != nil {
		log.Println("failed when create repo marriageDocumentUsecase : ", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success create from usecase marriageDocumentUsecase")

	return converter.MarriageStatementDocumentToResponse(marriageDocument), nil
}

// Delete implements MarriageStatementDocumentUsecase.
func (marriageDocumentUsecase *MarriageStatementDocumentUsecaseImpl) Delete(ctx context.Context, marriageDocumentId uint) error {
	tx := marriageDocumentUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	birthDocument := &entity.MarriageStatementDocument{}
	birthDocument.ID = marriageDocumentId

	err := marriageDocumentUsecase.MarriageStatementDocumentRepo.FindById(tx, birthDocument)
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

	err = marriageDocumentUsecase.MarriageStatementDocumentRepo.Delete(tx, birthDocument)
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

// FindAll implements MarriageStatementDocumentUsecase.
func (marriageDocumentUsecase *MarriageStatementDocumentUsecaseImpl) FindAll(ctx context.Context) (*[]model.MarriageStatementDocumentResponse, error) {
	tx := marriageDocumentUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var birthDocuments = &[]entity.MarriageStatementDocument{}
	err := marriageDocumentUsecase.MarriageStatementDocumentRepo.FindAll(tx, birthDocuments)
	if err != nil {
		log.Println("failed when find all repo birthDocument : ", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success find all from usecase birthDocument")

	return converter.MarriageStatementDocumentToResponses(birthDocuments), nil
}

// FindById implements MarriageStatementDocumentUsecase.
func (marriageDocumentUsecase *MarriageStatementDocumentUsecaseImpl) FindById(ctx context.Context, marriageDocumentId uint) (*model.MarriageStatementDocumentResponse, error) {
	tx := marriageDocumentUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	birthDocument := new(entity.MarriageStatementDocument)
	birthDocument.ID = marriageDocumentId

	if err := marriageDocumentUsecase.MarriageStatementDocumentRepo.FindById(tx, birthDocument); err != nil {
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

	return converter.MarriageStatementDocumentToResponse(birthDocument), nil
}
