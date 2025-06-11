package http

import (
	"log"

	"github.com/Bangdams/kelola-dokumen-BE/internal/model"
	"github.com/Bangdams/kelola-dokumen-BE/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type StatementTypeItemController interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	FindId(ctx *fiber.Ctx) error
}

type StatementTypeItemControllerImpl struct {
	StatementTypeItemUsecase usecase.StatementTypeItemUsecase
}

func NewStatementTypeItemController(statementTypeItemUsecase usecase.StatementTypeItemUsecase) StatementTypeItemController {
	return &StatementTypeItemControllerImpl{
		StatementTypeItemUsecase: statementTypeItemUsecase,
	}
}

// Create implements StatementTypeItemController.
func (controller *StatementTypeItemControllerImpl) Create(ctx *fiber.Ctx) error {
	request := new(model.StatementTypeItemRequest)

	if err := ctx.BodyParser(request); err != nil {
		log.Println("failed to parse request : ", err)
		return fiber.ErrBadRequest
	}

	response, err := controller.StatementTypeItemUsecase.Create(ctx.UserContext(), request)
	if err != nil {
		log.Println("failed to create StatementType")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.StatementTypeItemResponse]{Data: response})
}

// Delete implements StatementTypeItemController.
func (controller *StatementTypeItemControllerImpl) Delete(ctx *fiber.Ctx) error {
	statementTypeItemId, err := ctx.ParamsInt("statementTypeItemId")
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := controller.StatementTypeItemUsecase.Delete(ctx.UserContext(), uint(statementTypeItemId)); err != nil {
		log.Println("failed to delete StatementType")
		return err
	}

	return nil
}

// FindAll implements StatementTypeItemController.
func (controller *StatementTypeItemControllerImpl) FindAll(ctx *fiber.Ctx) error {
	var responses *[]model.StatementTypeItemResponse
	var err error

	statementTypeId, err := ctx.ParamsInt("statementTypeId")
	if err != nil {
		return fiber.ErrBadRequest
	}

	responses, err = controller.StatementTypeItemUsecase.FindAll(ctx.UserContext(), uint(statementTypeId))
	if err != nil {
		log.Println("failed to find all StatementTypeItem")
		return err
	}

	return ctx.JSON(model.WebResponses[model.StatementTypeItemResponse]{Data: responses})
}

// FindId implements StatementTypeItemController.
func (controller *StatementTypeItemControllerImpl) FindId(ctx *fiber.Ctx) error {
	statementTypeItemId, err := ctx.ParamsInt("statementTypeItemId")
	if err != nil {
		return fiber.ErrBadRequest
	}

	response, err := controller.StatementTypeItemUsecase.FindById(ctx.UserContext(), uint(statementTypeItemId))
	if err != nil {
		log.Println("failed to find by Id StatementType")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.StatementTypeItemResponse]{Data: response})
}

// Update implements StatementTypeItemController.
func (controller *StatementTypeItemControllerImpl) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateStatementTypeItemRequest)

	if err := ctx.BodyParser(request); err != nil {
		log.Println("failed to parse request : ", err)
		return fiber.ErrBadRequest
	}

	response, err := controller.StatementTypeItemUsecase.Update(ctx.UserContext(), request)
	if err != nil {
		log.Println("failed to update StatementTypeItem")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.StatementTypeItemResponse]{Data: response})
}
