package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"github.com/Bangdams/kelola-dokumen-BE/internal/model"
	"github.com/Bangdams/kelola-dokumen-BE/internal/model/converter"
	"github.com/Bangdams/kelola-dokumen-BE/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type StatementTypeUsecase interface {
	Create(ctx context.Context, request *model.StatementTypeRequest) (*model.StatementTypeResponse, error)
	Update(ctx context.Context, request *model.UpdateStatementTypeRequest) (*model.StatementTypeResponse, error)
	Delete(ctx context.Context, statementTypeId uint) error
	FindById(ctx context.Context, statementTypeId uint) (*model.StatementTypeResponse, error)
	FindAll(ctx context.Context) (*[]model.StatementTypeResponse, error)
}

type StatementTypeUsecaseImpl struct {
	StatementTypeRepo repository.StatementTypeRepository
	DB                *gorm.DB
	Validate          *validator.Validate
}

func NewStatementTypeUsecase(statementTypeRepo repository.StatementTypeRepository, DB *gorm.DB, validate *validator.Validate) StatementTypeUsecase {
	return &StatementTypeUsecaseImpl{
		StatementTypeRepo: statementTypeRepo,
		DB:                DB,
		Validate:          validate,
	}
}

// FindAll implements UserUsecase.
func (statementTypeUsecase *StatementTypeUsecaseImpl) FindAll(ctx context.Context) (*[]model.StatementTypeResponse, error) {
	tx := statementTypeUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var statementType = &[]entity.StatementType{}
	err := statementTypeUsecase.StatementTypeRepo.FindAll(tx, statementType)
	if err != nil {
		log.Println("failed when find all repo StatementType : ", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success find all from usecase StatementType")

	return converter.StatementTypeToResponses(statementType), nil
}

// Update implements UserUsecase.
func (statementTypeUsecase *StatementTypeUsecaseImpl) Update(ctx context.Context, request *model.UpdateStatementTypeRequest) (*model.StatementTypeResponse, error) {
	tx := statementTypeUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	errorResponse := &model.ErrorResponse{}

	err := statementTypeUsecase.Validate.Struct(request)
	if err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Field '%s' failed on '%s' rule", e.Field(), e.Tag())
			validationErrors = append(validationErrors, msg)
		}

		errorResponse.Message = "invalid request parameter"
		errorResponse.Details = validationErrors

		jsonString, _ := json.Marshal(errorResponse)

		log.Println("error update statementType : ", err)

		return nil, fiber.NewError(fiber.ErrBadRequest.Code, string(jsonString))
	}

	statementType := &entity.StatementType{
		ID: request.ID,
	}

	err = statementTypeUsecase.StatementTypeRepo.FindById(tx, statementType)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse.Message = "statement type data was not found"
			errorResponse.Details = []string{}

			jsonString, _ := json.Marshal(errorResponse)
			log.Println("Data not found")

			return nil, fiber.NewError(fiber.ErrNotFound.Code, string(jsonString))
		}

		log.Println("error find by id : ", err)
		return nil, fiber.ErrInternalServerError
	}

	statementType.Name = request.Name

	err = statementTypeUsecase.StatementTypeRepo.Update(tx, statementType)
	if err != nil {
		mysqlErr := err.(*mysql.MySQLError)
		log.Println("failed when update repo statementType : ", err)

		var errorField string
		parts := strings.Split(mysqlErr.Message, "'")
		if len(parts) > 2 {
			errorField = parts[1]
		}

		if mysqlErr.Number == 1062 {
			errorResponse.Message = "Duplicate entry"
			errorResponse.Details = []string{errorField + " already exists in the database."}

			jsonString, _ := json.Marshal(errorResponse)

			return nil, fiber.NewError(fiber.ErrConflict.Code, string(jsonString))
		}

		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success update from usecase statementType")

	return converter.StatementTypeToResponse(statementType), nil
}

// Create implements UserUsecase.
func (statementTypeUsecase *StatementTypeUsecaseImpl) Create(ctx context.Context, request *model.StatementTypeRequest) (*model.StatementTypeResponse, error) {
	tx := statementTypeUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	errorResponse := &model.ErrorResponse{}

	err := statementTypeUsecase.Validate.Struct(request)
	if err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Field '%s' failed on '%s' rule", e.Field(), e.Tag())
			validationErrors = append(validationErrors, msg)
		}

		errorResponse.Message = "invalid request parameter"
		errorResponse.Details = validationErrors

		jsonString, _ := json.Marshal(errorResponse)

		log.Println("error create statementType : ", err)

		return nil, fiber.NewError(fiber.ErrBadRequest.Code, string(jsonString))
	}

	statementType := &entity.StatementType{
		Name: request.Name,
	}

	if err := statementTypeUsecase.StatementTypeRepo.FindByName(tx, statementType); err == nil {
		errorResponse.Message = "Duplicate entry"
		errorResponse.Details = []string{"name already exists in the database."}

		jsonString, _ := json.Marshal(errorResponse)

		return nil, fiber.NewError(fiber.ErrConflict.Code, string(jsonString))
	}

	err = statementTypeUsecase.StatementTypeRepo.Create(tx, statementType)
	if err != nil {
		log.Println("failed when create repo statementType : ", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success create from usecase statementType")

	return converter.StatementTypeToResponse(statementType), nil
}

// Delete implements UserUsecase.
func (statementTypeUsecase *StatementTypeUsecaseImpl) Delete(ctx context.Context, statementTypeId uint) error {
	tx := statementTypeUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	statementType := &entity.StatementType{}
	statementType.ID = statementTypeId

	err := statementTypeUsecase.StatementTypeRepo.FindById(tx, statementType)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse := model.ErrorResponse{
				Message: "Statement type data was not found",
				Details: []string{},
			}
			jsonString, _ := json.Marshal(errorResponse)

			log.Println("error delete statementType : ", err)

			return fiber.NewError(fiber.ErrNotFound.Code, string(jsonString))
		}

		log.Println("error delete statementType : ", err)
		return fiber.ErrInternalServerError
	}

	err = statementTypeUsecase.StatementTypeRepo.Delete(tx, statementType)
	if err != nil {
		log.Println("failed when delete repo statementType : ", err)
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return fiber.ErrInternalServerError
	}

	log.Println("success delete from usecase statementType")

	return nil
}

// FindById implements StatementTypeUsecase.
func (statementTypeUsecase *StatementTypeUsecaseImpl) FindById(ctx context.Context, statementTypeId uint) (*model.StatementTypeResponse, error) {
	tx := statementTypeUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	statementType := new(entity.StatementType)
	statementType.ID = statementTypeId

	if err := statementTypeUsecase.StatementTypeRepo.FindById(tx, statementType); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse := model.ErrorResponse{
				Message: "StatementType data was not found",
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

	log.Println("success find by id from usecase statementType")

	return converter.StatementTypeToResponse(statementType), nil
}
