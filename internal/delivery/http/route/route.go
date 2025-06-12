package route

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

	// API for dashboard admin
	admin.Get("/dashboard", config.UserController.DashboardAdmin)

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
	admin.Put("/mails", config.MailController.Update)
	admin.Delete("/mails/:mailId", config.MailController.Delete)

	// Find all mail incomplite
	admin.Get("/mails-incomplite", config.MailController.FindAllIncompliteMail)
	admin.Get("/mails-complite", config.MailController.FindAllCompliteMail)

	// API for birth certificate document
	admin.Get("/birth-certificate-documents", config.BirthCertificateDocumentController.FindAll)
	admin.Get("/mails/birth-certificate-document/:birthDocumentId", config.BirthCertificateDocumentController.FindId)
	admin.Post("/mails/birth-certificate-document", config.BirthCertificateDocumentController.Create)
	admin.Put("/mails/birth-certificate-document", config.BirthCertificateDocumentController.Update)
	admin.Delete("/mails/birth-certificate-document/:birthDocumentId", config.BirthCertificateDocumentController.Delete)

	// API for death certificate document
	admin.Get("/death-certificate-documents", config.DeathCertificateDocumentController.FindAll)
	admin.Get("/mails/death-certificate-document/:deathDocumentId", config.DeathCertificateDocumentController.FindId)
	admin.Post("/mails/death-certificate-document", config.DeathCertificateDocumentController.Create)
	admin.Put("/mails/death-certificate-document", config.DeathCertificateDocumentController.Update)
	admin.Delete("/mails/death-certificate-document/:deathDocumentId", config.DeathCertificateDocumentController.Delete)

	// API for marriage statement document
	admin.Get("/marriage-statement-documents", config.MarriageStatementDocumentController.FindAll)
	admin.Get("/mails/marriage-statement-document/:marriageDocumentId", config.MarriageStatementDocumentController.FindId)
	admin.Post("/mails/marriage-statement-document", config.MarriageStatementDocumentController.Create)
	admin.Put("/mails/marriage-statement-document", config.MarriageStatementDocumentController.Update)
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
		// ambil nama baru dari query, jika ada
		ext := strings.ToLower(filepath.Ext(safeFilename))
		if ext != ".png" && ext != ".jpg" && ext != ".jpeg" && ext != ".pdf" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "file type not allowed",
			})
		}

		customName := "dokumen" + ext

		filepath := filepath.Join("./uploads", safeFilename)
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "file not found!"})
		}

		ctx.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", customName))

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
