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

type DeathCertificateDocumentController interface {
	Create(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	FindId(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type DeathCertificateDocumentControllerImpl struct {
	DeathCertificateDocumentUsecase usecase.DeathCertificateDocumentUsecase
}

func NewDeathCertificateDocumentController(deathCertificateDocumentUsecase usecase.DeathCertificateDocumentUsecase) DeathCertificateDocumentController {
	return &DeathCertificateDocumentControllerImpl{
		DeathCertificateDocumentUsecase: deathCertificateDocumentUsecase,
	}
}

// Update implements DeathCertificateDocumentController.
func (controller *DeathCertificateDocumentControllerImpl) Update(ctx *fiber.Ctx) error {
	request := new(model.BalasanRequest)
	if err := ctx.BodyParser(request); err != nil {
		log.Println("failed to parse request:", err)
		return fiber.ErrBadRequest
	}

	deathDocumentId, err := strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		log.Println("error badrequest")
		return fiber.ErrBadRequest
	}

	request.ID = uint(deathDocumentId)

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

	response, err := controller.DeathCertificateDocumentUsecase.Update(ctx.UserContext(), request.ID, request.SuratBalasanFilePath)
	if err != nil {
		log.Println("failed to update mail")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.DeathCertificateDocumentResponse]{Data: response})
}

// Create implements DeathCertificateDocumentController.
func (controller *DeathCertificateDocumentControllerImpl) Create(ctx *fiber.Ctx) error {
	request := new(model.DeathCertificateDocumentRequest)
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
		"pengantarRtRwFile",
		"formulirFile",
		"suratKematianFile",
		"ktpFile",
		"kkFile",
		"ktpPelaporFile",
		"ktpSaksi1File",
		"ktpSaksi2File",
		"bukuNikahFile",
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

	request.RtRwFilePath = savedFiles["pengantarRtRwFile"]
	request.FormulirFilePath = savedFiles["formulirFile"]
	request.SuratKematianFilePath = savedFiles["suratKematianFile"]
	request.KtpFilePath = savedFiles["ktpFile"]
	request.KkFilePath = savedFiles["kkFile"]
	request.KtpPelaporFilePath = savedFiles["ktpPelaporFile"]
	request.KtpSaksi1FilePath = savedFiles["ktpSaksi1File"]
	request.KtpSaksi2FilePath = savedFiles["ktpSaksi2File"]
	request.BukuNikahFilePath = savedFiles["bukuNikahFile"]

	response, err := controller.DeathCertificateDocumentUsecase.Create(ctx.UserContext(), request)
	if err != nil {
		log.Println("failed to create DeathDocument")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.DeathCertificateDocumentResponse]{Data: response})
}

// Delete implements DeathCertificateDocumentController.
func (controller *DeathCertificateDocumentControllerImpl) Delete(ctx *fiber.Ctx) error {
	deathDocumentId, err := ctx.ParamsInt("deathDocumentId")
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := controller.DeathCertificateDocumentUsecase.Delete(ctx.UserContext(), uint(deathDocumentId)); err != nil {
		log.Println("failed to delete DeathDocument")
		return err
	}

	return nil
}

// FindAll implements DeathCertificateDocumentController.
func (controller *DeathCertificateDocumentControllerImpl) FindAll(ctx *fiber.Ctx) error {
	var responses *[]model.DeathCertificateDocumentResponse
	var err error

	responses, err = controller.DeathCertificateDocumentUsecase.FindAll(ctx.UserContext())
	if err != nil {
		log.Println("failed to find all DeathDocument")
		return err
	}

	return ctx.JSON(model.WebResponses[model.DeathCertificateDocumentResponse]{Data: responses})
}

// FindId implements DeathCertificateDocumentController.
func (controller *DeathCertificateDocumentControllerImpl) FindId(ctx *fiber.Ctx) error {
	deathDocumentId, err := ctx.ParamsInt("deathDocumentId")
	if err != nil {
		return fiber.ErrBadRequest
	}

	response, err := controller.DeathCertificateDocumentUsecase.FindById(ctx.UserContext(), uint(deathDocumentId))
	if err != nil {
		log.Println("failed to find by Id DeathDocument")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.DeathCertificateDocumentResponse]{Data: response})
}
