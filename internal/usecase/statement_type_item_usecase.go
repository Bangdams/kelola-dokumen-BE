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

type StatementTypeItemUsecase interface {
	Create(ctx context.Context, request *model.StatementTypeItemRequest) (*model.StatementTypeItemResponse, error)
	Update(ctx context.Context, request *model.UpdateStatementTypeItemRequest) (*model.StatementTypeItemResponse, error)
	Delete(ctx context.Context, statementTypeId uint) error
	FindById(ctx context.Context, statementTypeId uint) (*model.StatementTypeItemResponse, error)
	FindAll(ctx context.Context, statementTypeId uint) (*[]model.StatementTypeItemResponse, error)
}

type StatementTypeUsecItemaseImpl struct {
	StatementTypeItem repository.StatementTypeItemRepository
	StatementType     repository.StatementTypeRepository
	DB                *gorm.DB
	Validate          *validator.Validate
}

func NewStatementTypeItemUsecase(statementType repository.StatementTypeRepository, statementTypeRepo repository.StatementTypeItemRepository, DB *gorm.DB, validate *validator.Validate) StatementTypeItemUsecase {
	return &StatementTypeUsecItemaseImpl{
		StatementTypeItem: statementTypeRepo,
		StatementType:     statementType,
		DB:                DB,
		Validate:          validate,
	}
}

// FindAll implements UserUsecase.
func (statementTypeItemUsecase *StatementTypeUsecItemaseImpl) FindAll(ctx context.Context, statementTypeId uint) (*[]model.StatementTypeItemResponse, error) {
	tx := statementTypeItemUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var statementTypeItems = &[]entity.StatementTypeItem{}
	err := statementTypeItemUsecase.StatementTypeItem.FindAll(tx, statementTypeItems, statementTypeId)
	if err != nil {
		log.Println("failed when find all repo statementTypeItem : ", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success find all from usecase statementTypeItem")

	return converter.StatementTypeItemToResponses(statementTypeItems), nil
}

// Update implements UserUsecase.
func (statementTypeItemUsecase *StatementTypeUsecItemaseImpl) Update(ctx context.Context, request *model.UpdateStatementTypeItemRequest) (*model.StatementTypeItemResponse, error) {
	tx := statementTypeItemUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	errorResponse := &model.ErrorResponse{}

	err := statementTypeItemUsecase.Validate.Struct(request)
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

	statementTypeItem := &entity.StatementTypeItem{
		ID: request.ID,
	}

	err = statementTypeItemUsecase.StatementTypeItem.FindById(tx, statementTypeItem)
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

	err = statementTypeItemUsecase.StatementType.FindById(tx, &entity.StatementType{ID: request.StatementTypeId})
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

	statementTypeItem.Name = request.Name
	statementTypeItem.StatementTypeId = request.StatementTypeId

	err = statementTypeItemUsecase.StatementTypeItem.Update(tx, statementTypeItem)
	if err != nil {
		mysqlErr := err.(*mysql.MySQLError)
		log.Println("failed when update repo statementTypeItem : ", err)

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

	log.Println("success update from usecase statementTypeItem")

	return converter.StatementTypeItemToResponse(statementTypeItem), nil
}

// Create implements UserUsecase.
func (statementTypeItemUsecase *StatementTypeUsecItemaseImpl) Create(ctx context.Context, request *model.StatementTypeItemRequest) (*model.StatementTypeItemResponse, error) {
	tx := statementTypeItemUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	errorResponse := &model.ErrorResponse{}

	err := statementTypeItemUsecase.Validate.Struct(request)
	if err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("Field '%s' failed on '%s' rule", e.Field(), e.Tag())
			validationErrors = append(validationErrors, msg)
		}

		errorResponse.Message = "invalid request parameter"
		errorResponse.Details = validationErrors

		jsonString, _ := json.Marshal(errorResponse)

		log.Println("error create statementTypeItem : ", err)

		return nil, fiber.NewError(fiber.ErrBadRequest.Code, string(jsonString))
	}

	statementTypeItem := &entity.StatementTypeItem{
		StatementTypeId: request.StatementTypeId,
		Name:            request.Name,
	}

	if err := statementTypeItemUsecase.StatementTypeItem.FindByNameStatementTypeId(tx, statementTypeItem); err == nil {
		errorResponse.Message = "Duplicate entry"
		errorResponse.Details = []string{"name already exists in the database."}

		jsonString, _ := json.Marshal(errorResponse)

		return nil, fiber.NewError(fiber.ErrConflict.Code, string(jsonString))
	}

	err = statementTypeItemUsecase.StatementTypeItem.Create(tx, statementTypeItem)
	if err != nil {
		log.Println("failed when create repo statementTypeItem : ", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return nil, fiber.ErrInternalServerError
	}

	log.Println("success create from usecase statementTypeItem")

	return converter.StatementTypeItemToResponse(statementTypeItem), nil
}

// Delete implements UserUsecase.
func (statementTypeItemUsecase *StatementTypeUsecItemaseImpl) Delete(ctx context.Context, statementTypeItemId uint) error {
	tx := statementTypeItemUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	statementTypeItem := &entity.StatementTypeItem{}
	statementTypeItem.ID = statementTypeItemId

	err := statementTypeItemUsecase.StatementTypeItem.FindById(tx, statementTypeItem)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse := model.ErrorResponse{
				Message: "Statement type data was not found",
				Details: []string{},
			}
			jsonString, _ := json.Marshal(errorResponse)

			log.Println("error delete statementTypeItem : ", err)

			return fiber.NewError(fiber.ErrNotFound.Code, string(jsonString))
		}

		log.Println("error delete statementTypeItem : ", err)
		return fiber.ErrInternalServerError
	}

	err = statementTypeItemUsecase.StatementTypeItem.Delete(tx, statementTypeItem)
	if err != nil {
		log.Println("failed when delete repo statementTypeItem : ", err)
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println("Failed commit transaction : ", err)
		return fiber.ErrInternalServerError
	}

	log.Println("success delete from usecase statementTypeItem")

	return nil
}

// FindById implements StatementTypeItemUsecase.
func (statementTypeItemUsecase *StatementTypeUsecItemaseImpl) FindById(ctx context.Context, statementTypeItemId uint) (*model.StatementTypeItemResponse, error) {
	tx := statementTypeItemUsecase.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	statementTypeItem := new(entity.StatementTypeItem)
	statementTypeItem.ID = statementTypeItemId

	if err := statementTypeItemUsecase.StatementTypeItem.FindById(tx, statementTypeItem); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorResponse := model.ErrorResponse{
				Message: "StatemItementType data was not found",
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

	log.Println("success find by id from usecase statementTypeItem")

	return converter.StatementTypeItemToResponse(statementTypeItem), nil
}
