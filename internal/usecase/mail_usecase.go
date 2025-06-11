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

type MailUsecase interface {
	Create(ctx context.Context, request *model.MailRequest) (*model.MailResponse, error)
	Delete(ctx context.Context, mailId uint) error
	FindById(ctx context.Context, mailId uint) (*model.MailResponse, error)
	FindAll(ctx context.Context) (*[]model.MailResponse, error)
}

type MailUsecaseImpl struct {
	MailRepo              repository.MailRepository
	StatementTypeItemRepo repository.StatementTypeItemRepository
	DB                    *gorm.DB
	Validate              *validator.Validate
}

func NewMailUsecase(statementTypeItemRepo repository.StatementTypeItemRepository, mailRepo repository.MailRepository, DB *gorm.DB, validate *validator.Validate) MailUsecase {
	return &MailUsecaseImpl{
		MailRepo:              mailRepo,
		StatementTypeItemRepo: statementTypeItemRepo,
		DB:                    DB,
		Validate:              validate,
	}
}

// FindAll implements UserUsecase.
func (mailUsecase *MailUsecaseImpl) FindAll(ctx context.Context) (*[]model.MailResponse, error) {
	tx := mailUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var mails = &[]entity.Mail{}
	err := mailUsecase.MailRepo.FindAll(tx, mails)
	if err != nil {
		log.Println("failed when find all repo mail : ", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success find all from usecase mail")

	return converter.MailToResponses(mails), nil
}

// Create implements UserUsecase.
func (mailUsecase *MailUsecaseImpl) Create(ctx context.Context, request *model.MailRequest) (*model.MailResponse, error) {
	tx := mailUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	errorResponse := &model.ErrorResponse{}

	err := mailUsecase.Validate.Struct(request)
	if err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Field '%s' failed on '%s' rule", e.Field(), e.Tag())
			validationErrors = append(validationErrors, msg)
		}

		errorResponse.Message = "invalid request parameter"
		errorResponse.Details = validationErrors

		jsonString, _ := json.Marshal(errorResponse)

		log.Println("error create mail : ", err)

		return nil, fiber.NewError(fiber.ErrBadRequest.Code, string(jsonString))
	}

	err = mailUsecase.StatementTypeItemRepo.FindById(tx, &entity.StatementTypeItem{
		ID: request.StatementTypeItemId,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse := model.ErrorResponse{
				Message: "mail type data was not found",
				Details: []string{},
			}
			jsonString, _ := json.Marshal(errorResponse)

			log.Println("error delete mail : ", err)

			return nil, fiber.NewError(fiber.ErrNotFound.Code, string(jsonString))
		}

		log.Println("error delete mail : ", err)
		return nil, fiber.ErrInternalServerError
	}

	mail := &entity.Mail{
		StatementTypeItemId: request.StatementTypeItemId,
		UserId:              request.UserId,
		KtpFilePath:         request.KtpFilePath,
		KkFilePath:          request.KkFilePath,
		RtRwFilePath:        request.RtRwFilePath,
	}

	err = mailUsecase.MailRepo.Create(tx, mail)
	if err != nil {
		log.Println("failed when create repo mail : ", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success create from usecase mail")

	return converter.MailToResponse(mail), nil
}

// Delete implements UserUsecase.
func (mailUsecase *MailUsecaseImpl) Delete(ctx context.Context, mailId uint) error {
	tx := mailUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	mail := &entity.Mail{}
	mail.ID = mailId

	err := mailUsecase.MailRepo.FindById(tx, mail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse := model.ErrorResponse{
				Message: "mail type data was not found",
				Details: []string{},
			}
			jsonString, _ := json.Marshal(errorResponse)

			log.Println("error delete mail : ", err)

			return fiber.NewError(fiber.ErrNotFound.Code, string(jsonString))
		}

		log.Println("error delete mail : ", err)
		return fiber.ErrInternalServerError
	}

	err = mailUsecase.MailRepo.Delete(tx, mail)
	if err != nil {
		log.Println("failed when delete repo mail : ", err)
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return fiber.ErrInternalServerError
	}

	log.Println("success delete from usecase mail")

	return nil
}

// FindById implements MailUsecase.
func (mailUsecase *MailUsecaseImpl) FindById(ctx context.Context, mailId uint) (*model.MailResponse, error) {
	tx := mailUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	mail := new(entity.Mail)
	mail.ID = mailId

	if err := mailUsecase.MailRepo.FindById(tx, mail); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse := model.ErrorResponse{
				Message: "mail data was not found",
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

	log.Println("success find by id from usecase mail")

	return converter.MailToResponse(mail), nil
}
