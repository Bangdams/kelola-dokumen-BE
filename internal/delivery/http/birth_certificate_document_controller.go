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

type BirthCertificateDocumentController interface {
	Create(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	FindId(ctx *fiber.Ctx) error
}

type BirthCertificateDocumentControllerImpl struct {
	BirthCertificateDocumentUsecase usecase.BirthCertificateDocumentUsecase
}

func NewBirthCertificateDocumentController(birthCertificateDocumentUsecase usecase.BirthCertificateDocumentUsecase) BirthCertificateDocumentController {
	return &BirthCertificateDocumentControllerImpl{
		BirthCertificateDocumentUsecase: birthCertificateDocumentUsecase,
	}
}

// Create implements BirthCertificateDocumentController.
func (controller *BirthCertificateDocumentControllerImpl) Create(ctx *fiber.Ctx) error {
	request := new(model.BirthCertificateDocumentRequest)
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
		"suratKelahiranFile",
		"bukuNikahFile",
		"kkFile",
		"ktpPelaporFile",
		"ktpSaksi1File",
		"ktpSaksi2File",
		"ktpOrangTuaFile",
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
	request.SuratKelahiranFilePath = savedFiles["suratKelahiranFile"]
	request.BukuNikahFilePath = savedFiles["bukuNikahFile"]
	request.KkFilePath = savedFiles["kkFile"]
	request.KtpPelaporFilePath = savedFiles["ktpPelaporFile"]
	request.KtpSaksi1FilePath = savedFiles["ktpSaksi1File"]
	request.KtpSaksi2FilePath = savedFiles["ktpSaksi2File"]
	request.KtpOrangTuaFilePath = savedFiles["ktpOrangTuaFile"]

	response, err := controller.BirthCertificateDocumentUsecase.Create(ctx.UserContext(), request)
	if err != nil {
		log.Println("failed to create BirthDocument")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.BirthCertificateDocumentResponse]{Data: response})
}

// Delete implements BirthCertificateDocumentController.
func (controller *BirthCertificateDocumentControllerImpl) Delete(ctx *fiber.Ctx) error {
	birthDocumentId, err := ctx.ParamsInt("birthDocumentId")
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := controller.BirthCertificateDocumentUsecase.Delete(ctx.UserContext(), uint(birthDocumentId)); err != nil {
		log.Println("failed to delete BirthDocument")
		return err
	}

	return nil
}

// FindAll implements BirthCertificateDocumentController.
func (controller *BirthCertificateDocumentControllerImpl) FindAll(ctx *fiber.Ctx) error {
	var responses *[]model.BirthCertificateDocumentResponse
	var err error

	responses, err = controller.BirthCertificateDocumentUsecase.FindAll(ctx.UserContext())
	if err != nil {
		log.Println("failed to find all BirthDocument")
		return err
	}

	return ctx.JSON(model.WebResponses[model.BirthCertificateDocumentResponse]{Data: responses})
}

// FindId implements BirthCertificateDocumentController.
func (controller *BirthCertificateDocumentControllerImpl) FindId(ctx *fiber.Ctx) error {
	birthDocumentId, err := ctx.ParamsInt("birthDocumentId")
	if err != nil {
		return fiber.ErrBadRequest
	}

	response, err := controller.BirthCertificateDocumentUsecase.FindById(ctx.UserContext(), uint(birthDocumentId))
	if err != nil {
		log.Println("failed to find by Id BirthDocument")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.BirthCertificateDocumentResponse]{Data: response})
}
