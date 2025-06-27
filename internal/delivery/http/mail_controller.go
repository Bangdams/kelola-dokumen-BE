package http

import (
	"log"
	"strconv"

	"github.com/Bangdams/kelola-dokumen-BE/internal/model"
	"github.com/Bangdams/kelola-dokumen-BE/internal/usecase"
	"github.com/Bangdams/kelola-dokumen-BE/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type MailController interface {
	Create(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	FindId(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	FindAllIncompliteMail(ctx *fiber.Ctx) error
	FindAllCompliteMail(ctx *fiber.Ctx) error
	FindAllMailForUser(ctx *fiber.Ctx) error
}

type MailControllerImpl struct {
	MailUsecase usecase.MailUsecase
}

func NewMailController(mailUsecase usecase.MailUsecase) MailController {
	return &MailControllerImpl{
		MailUsecase: mailUsecase,
	}
}

// FindAllMailForUser implements MailController.
func (controller *MailControllerImpl) FindAllMailForUser(ctx *fiber.Ctx) error {
	userToken := ctx.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	response, err := controller.MailUsecase.FindAllMailForUser(ctx.UserContext(), uint(userId))
	if err != nil {
		log.Println("failed to find mail")
		return err
	}

	return ctx.JSON(model.WebResponses[model.AllMailItemForUserResponse]{Data: response})

}

// FindAllCompliteMail implements MailController.
func (controller *MailControllerImpl) FindAllCompliteMail(ctx *fiber.Ctx) error {
	response, err := controller.MailUsecase.FindAllCompliteMail(ctx.UserContext())
	if err != nil {
		log.Println("failed to find mail")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.AllMailResponse]{Data: response})
}

// FindAllIncompliteMail implements MailController.
func (controller *MailControllerImpl) FindAllIncompliteMail(ctx *fiber.Ctx) error {
	response, err := controller.MailUsecase.FindAllIncompliteMail(ctx.UserContext())
	if err != nil {
		log.Println("failed to find mail")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.AllMailResponse]{Data: response})
}

// Update implements MailController.
func (controller *MailControllerImpl) Update(ctx *fiber.Ctx) error {
	request := new(model.BalasanRequest)
	if err := ctx.BodyParser(request); err != nil {
		log.Println("failed to parse request:", err)
		return fiber.ErrBadRequest
	}

	mailId, err := strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		log.Println("error badrequest")
		return fiber.ErrBadRequest
	}

	request.ID = uint(mailId)

	fileFields := []string{"suratBalasan"}
	savedFiles := make(map[string]string)

	for _, field := range fileFields {
		file, err := ctx.FormFile(field)
		if err != nil {
			log.Println("failed to parse", field, "file:", err)
			return fiber.ErrBadRequest
		}

		filePath, err := util.SaveValidatedFile(ctx, file, field)
		if err != nil {
			log.Println("failed to save", field, "file:", err)

			util.CleanupFiles(savedFiles)

			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		savedFiles[field] = filePath
	}
	request.SuratBalasanFilePath = savedFiles["suratBalasan"]

	response, err := controller.MailUsecase.Update(ctx.UserContext(), request.ID, request.SuratBalasanFilePath)
	if err != nil {
		log.Println("failed to update mail")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.MailResponse]{Data: response})
}

// Create implements MailController.
func (controller *MailControllerImpl) Create(ctx *fiber.Ctx) error {
	request := new(model.MailRequest)
	if err := ctx.BodyParser(request); err != nil {
		log.Println("failed to parse request:", err)
		return fiber.ErrBadRequest
	}

	userToken := ctx.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	statementTypeId, err := strconv.Atoi(ctx.FormValue("statement_type_item_id"))
	if err != nil {
		log.Println("error badrequest")
		return fiber.ErrBadRequest
	}

	request.StatementTypeItemId = uint(statementTypeId)
	request.UserId = uint(userId)

	// Upload files (KTP, KK, Pengantar RT/RW)
	fileFields := []string{"ktpFile", "kkFile", "pengantarRtRwFile"}
	savedFiles := make(map[string]string)

	for _, field := range fileFields {
		file, err := ctx.FormFile(field)
		if err != nil {
			log.Println("failed to parse", field, "file:", err)
			return fiber.ErrBadRequest
		}

		filePath, err := util.SaveValidatedFile(ctx, file, field)
		if err != nil {
			log.Println("failed to save", field, "file:", err)

			util.CleanupFiles(savedFiles)

			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		savedFiles[field] = filePath
	}

	request.KtpFilePath = savedFiles["ktpFile"]
	request.KkFilePath = savedFiles["kkFile"]
	request.RtRwFilePath = savedFiles["pengantarRtRwFile"]

	response, err := controller.MailUsecase.Create(ctx.UserContext(), request)
	if err != nil {
		log.Println("failed to create mail")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.MailResponse]{Data: response})
}

// Delete implements MailController.
func (controller *MailControllerImpl) Delete(ctx *fiber.Ctx) error {
	mailId, err := ctx.ParamsInt("mailId")
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := controller.MailUsecase.Delete(ctx.UserContext(), uint(mailId)); err != nil {
		log.Println("failed to delete mail")
		return err
	}

	return nil
}

// FindAll implements MailController.
func (controller *MailControllerImpl) FindAll(ctx *fiber.Ctx) error {
	var responses *[]model.MailResponse
	var err error

	responses, err = controller.MailUsecase.FindAll(ctx.UserContext())
	if err != nil {
		log.Println("failed to find all mail")
		return err
	}

	return ctx.JSON(model.WebResponses[model.MailResponse]{Data: responses})
}

// FindId implements MailController.
func (controller *MailControllerImpl) FindId(ctx *fiber.Ctx) error {
	mailId, err := ctx.ParamsInt("mailId")
	if err != nil {
		return fiber.ErrBadRequest
	}

	response, err := controller.MailUsecase.FindById(ctx.UserContext(), uint(mailId))
	if err != nil {
		log.Println("failed to find by Id StatementType")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.MailResponse]{Data: response})
}
