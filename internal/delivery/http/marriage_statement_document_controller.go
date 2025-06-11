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

type MarriageStatementDocumentController interface {
	Create(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	FindId(ctx *fiber.Ctx) error
}

type MarriageStatementDocumentControllerImpl struct {
	MarriageStatementDocumentUsecase usecase.MarriageStatementDocumentUsecase
}

func NewMarriageStatementDocumentController(marriageStatementDocumentUsecase usecase.MarriageStatementDocumentUsecase) MarriageStatementDocumentController {
	return &MarriageStatementDocumentControllerImpl{
		MarriageStatementDocumentUsecase: marriageStatementDocumentUsecase,
	}
}

// Create implements MarriageStatementDocumentController.
func (controller *MarriageStatementDocumentControllerImpl) Create(ctx *fiber.Ctx) error {
	request := new(model.MarriageStatementDocumentRequest)
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

	// Upload files
	fileFields := []string{
		"kkFile",
		"ktpFile",
		"ijazahFile",
		"aktaFile",
		"ktpSaksiFile",
	}
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

	request.KkFilePath = savedFiles["kkFile"]
	request.KtpFilePath = savedFiles["ktpFile"]
	request.IjazahFilePath = savedFiles["ijazahFile"]
	request.AktaFilePath = savedFiles["aktaFile"]
	request.KtpSaksiFilePath = savedFiles["ktpSaksiFile"]

	response, err := controller.MarriageStatementDocumentUsecase.Create(ctx.UserContext(), request)
	if err != nil {
		log.Println("failed to create MarriageDocument")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.MarriageStatementDocumentResponse]{Data: response})
}

// Delete implements MarriageStatementDocumentController.
func (controller *MarriageStatementDocumentControllerImpl) Delete(ctx *fiber.Ctx) error {
	marriageDocumentId, err := ctx.ParamsInt("marriageDocumentId")
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := controller.MarriageStatementDocumentUsecase.Delete(ctx.UserContext(), uint(marriageDocumentId)); err != nil {
		log.Println("failed to delete MarriageDocument")
		return err
	}

	return nil
}

// FindAll implements MarriageStatementDocumentController.
func (controller *MarriageStatementDocumentControllerImpl) FindAll(ctx *fiber.Ctx) error {
	var responses *[]model.MarriageStatementDocumentResponse
	var err error

	responses, err = controller.MarriageStatementDocumentUsecase.FindAll(ctx.UserContext())
	if err != nil {
		log.Println("failed to find all MarriageDocument")
		return err
	}

	return ctx.JSON(model.WebResponses[model.MarriageStatementDocumentResponse]{Data: responses})
}

// FindId implements MarriageStatementDocumentController.
func (controller *MarriageStatementDocumentControllerImpl) FindId(ctx *fiber.Ctx) error {
	marriageDocumentId, err := ctx.ParamsInt("marriageDocumentId")
	if err != nil {
		return fiber.ErrBadRequest
	}

	response, err := controller.MarriageStatementDocumentUsecase.FindById(ctx.UserContext(), uint(marriageDocumentId))
	if err != nil {
		log.Println("failed to find by Id MarriageDocument")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.MarriageStatementDocumentResponse]{Data: response})
}
