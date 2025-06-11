package http

import (
	"log"

	"github.com/Bangdams/kelola-dokumen-BE/internal/model"
	"github.com/Bangdams/kelola-dokumen-BE/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type StatementTypeController interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	FindId(ctx *fiber.Ctx) error
}

type StatementTypeControllerImpl struct {
	StatementTypeUsecase usecase.StatementTypeUsecase
}

func NewStatementTypeController(statementTypeUsecase usecase.StatementTypeUsecase) StatementTypeController {
	return &StatementTypeControllerImpl{
		StatementTypeUsecase: statementTypeUsecase,
	}
}

// Create implements StatementTypeController.
func (controller *StatementTypeControllerImpl) Create(ctx *fiber.Ctx) error {
	request := new(model.StatementTypeRequest)

	if err := ctx.BodyParser(request); err != nil {
		log.Println("failed to parse request : ", err)
		return fiber.ErrBadRequest
	}

	response, err := controller.StatementTypeUsecase.Create(ctx.UserContext(), request)
	if err != nil {
		log.Println("failed to create StatementType")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.StatementTypeResponse]{Data: response})
}

// Delete implements StatementTypeController.
func (controller *StatementTypeControllerImpl) Delete(ctx *fiber.Ctx) error {
	statementTypeId, err := ctx.ParamsInt("statementTypeId")
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := controller.StatementTypeUsecase.Delete(ctx.UserContext(), uint(statementTypeId)); err != nil {
		log.Println("failed to delete StatementType")
		return err
	}

	return nil
}

// FindAll implements StatementTypeController.
func (controller *StatementTypeControllerImpl) FindAll(ctx *fiber.Ctx) error {
	var responses *[]model.StatementTypeResponse
	var err error

	responses, err = controller.StatementTypeUsecase.FindAll(ctx.UserContext())
	if err != nil {
		log.Println("failed to find all StatementType")
		return err
	}

	return ctx.JSON(model.WebResponses[model.StatementTypeResponse]{Data: responses})
}

// FindId implements StatementTypeController.
func (controller *StatementTypeControllerImpl) FindId(ctx *fiber.Ctx) error {
	statementTypeId, err := ctx.ParamsInt("statementTypeId")
	if err != nil {
		return fiber.ErrBadRequest
	}

	response, err := controller.StatementTypeUsecase.FindById(ctx.UserContext(), uint(statementTypeId))
	if err != nil {
		log.Println("failed to find by Id StatementType")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.StatementTypeResponse]{Data: response})
}

// Update implements StatementTypeController.
func (controller *StatementTypeControllerImpl) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateStatementTypeRequest)

	if err := ctx.BodyParser(request); err != nil {
		log.Println("failed to parse request : ", err)
		return fiber.ErrBadRequest
	}

	response, err := controller.StatementTypeUsecase.Update(ctx.UserContext(), request)
	if err != nil {
		log.Println("failed to update StatementType")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.StatementTypeResponse]{Data: response})
}
