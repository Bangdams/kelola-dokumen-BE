package route

import (
	"os"
	"path/filepath"

	"github.com/Bangdams/kelola-dokumen-BE/internal/delivery/http"
	"github.com/Bangdams/kelola-dokumen-BE/internal/util"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                                 *fiber.App
	UserController                      http.UserController
	StatementTypeController             http.StatementTypeController
	StatementTypeItemController         http.StatementTypeItemController
	MailController                      http.MailController
	BirthCertificateDocumentController  http.BirthCertificateDocumentController
	DeathCertificateDocumentController  http.DeathCertificateDocumentController
	MarriageStatementDocumentController http.MarriageStatementDocumentController
}

func (config *RouteConfig) Setup() {
	// API ADMIN
	// config.App.Post("/users", config.UserController.Create)

	admin := config.App.Group("/api-admin", util.CheckLevel("admin"))

	// API for user
	admin.Get("/users", config.UserController.FindAll)
	admin.Get("/users/:userId", config.UserController.FindId)
	admin.Post("/users", config.UserController.Create)
	admin.Delete("/users/:userId", config.UserController.Delete)
	admin.Put("/users", config.UserController.Update)

	// API for statement type
	admin.Get("/statement-types", config.StatementTypeController.FindAll)
	admin.Get("/statement-types/:statementTypeId", config.StatementTypeController.FindId)
	admin.Post("/statement-types", config.StatementTypeController.Create)
	admin.Delete("/statement-types/:statementTypeId", config.StatementTypeController.Delete)
	admin.Put("/statement-types", config.StatementTypeController.Update)

	// API for statement type item
	admin.Get("/statement-type-items/:statementTypeId/statementType", config.StatementTypeItemController.FindAll)
	admin.Get("/statement-type-items/:statementTypeItemId", config.StatementTypeItemController.FindId)
	admin.Post("/statement-type-items", config.StatementTypeItemController.Create)
	admin.Delete("/statement-type-items/:statementTypeItemId", config.StatementTypeItemController.Delete)
	admin.Put("/statement-type-items", config.StatementTypeItemController.Update)

	// API for mail
	admin.Get("/mails", config.MailController.FindAll)
	admin.Get("/mails/:mailId", config.MailController.FindId)
	admin.Post("/mails", config.MailController.Create)
	admin.Delete("/mails/:mailId", config.MailController.Delete)

	// API for birth certificate document
	admin.Get("/mails/birth-certificate-document", config.BirthCertificateDocumentController.FindAll)
	admin.Get("/mails/birth-certificate-document/:birthDocumentId", config.BirthCertificateDocumentController.FindId)
	admin.Post("/mails/birth-certificate-document", config.BirthCertificateDocumentController.Create)
	admin.Delete("/mails/birth-certificate-document/:birthDocumentId", config.BirthCertificateDocumentController.Delete)

	// API for death certificate document
	admin.Get("/mails/death-certificate-document", config.DeathCertificateDocumentController.FindAll)
	admin.Get("/mails/death-certificate-document/:deathDocumentId", config.DeathCertificateDocumentController.FindId)
	admin.Post("/mails/death-certificate-document", config.DeathCertificateDocumentController.Create)
	admin.Delete("/mails/death-certificate-document/:deathDocumentId", config.DeathCertificateDocumentController.Delete)

	// API for marriage statement document
	admin.Get("/mails/marriage-statement-document", config.MarriageStatementDocumentController.FindAll)
	admin.Get("/mails/marriage-statement-document/:marriageDocumentId", config.MarriageStatementDocumentController.FindId)
	admin.Post("/mails/marriage-statement-document", config.MarriageStatementDocumentController.Create)
	admin.Delete("/mails/marriage-statement-document/:marriageDocumentId", config.MarriageStatementDocumentController.Delete)

	// API USER
	user := config.App.Group("/api-user", util.CheckLevel("user"))

	user.Get("/mails", config.MailController.FindAll)
	user.Get("/mails/:mailId", config.MailController.FindId)
	user.Post("/mails", config.MailController.Create)
	user.Delete("/mails/:mailId", config.MailController.Delete)

	// API for pdf
	config.App.Get("/api/download/:filename", func(ctx *fiber.Ctx) error {
		filename := ctx.Params("filename")

		safeFilename := filepath.Base(filename)
		filepath := filepath.Join("./uploads", safeFilename)

		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "file not found!"})
		}
		return ctx.SendFile(filepath)
	})

	// Api for login
	config.App.Post("/login", config.UserController.Login)
	config.App.Post("/logout", config.UserController.Logout)
	config.App.Post("/refresh", config.UserController.Refresh)
	config.App.Get("/api/status-login", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"message": "success"})
	})
}
