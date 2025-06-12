package config

import (
	"github.com/Bangdams/kelola-dokumen-BE/internal/delivery/http"
	"github.com/Bangdams/kelola-dokumen-BE/internal/delivery/http/route"
	"github.com/Bangdams/kelola-dokumen-BE/internal/repository"
	"github.com/Bangdams/kelola-dokumen-BE/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Validate *validator.Validate
}

func Bootstrap(config *BootstrapConfig) {
	// repo
	userRepo := repository.NewUserRepository()
	refreshTokenRepo := repository.NewRefreshTokenRepository()
	rwListRepo := repository.NewRwListRepository()
	statementTypeRepo := repository.NewStatementTypeRepository()
	statementTypeItemRepo := repository.NewStatementTypeItemRepository()
	mailRepo := repository.NewMailRepository()
	birthCertificateDocumentRepo := repository.NewBirthCertificateDocumentRepository()
	deathCertificateDocumentRepo := repository.NewDeathCertificateDocumentRepository()
	marriageStatementDocumentRepo := repository.NewMarriageStatementDocumentRepository()

	// usecase
	userUsecase := usecase.NewUserUsecase(userRepo, refreshTokenRepo, rwListRepo, config.DB, config.Validate)
	statementTypeUsecase := usecase.NewStatementTypeUsecase(statementTypeRepo, config.DB, config.Validate)
	statementTypeItemUsecase := usecase.NewStatementTypeItemUsecase(statementTypeRepo, statementTypeItemRepo, config.DB, config.Validate)
	mailUsecase := usecase.NewMailUsecase(statementTypeItemRepo, mailRepo, config.DB, config.Validate)
	birthCertificateDocumentUsecase := usecase.NewBirthCertificateDocumentUsecase(statementTypeItemRepo, birthCertificateDocumentRepo, config.DB, config.Validate)
	deathCertificateDocumentUsecase := usecase.NewDeathCertificateDocumentUsecase(statementTypeItemRepo, deathCertificateDocumentRepo, config.DB, config.Validate)
	marriageStatementDocumentUsecase := usecase.NewMarriageStatementDocumentUsecase(statementTypeItemRepo, marriageStatementDocumentRepo, config.DB, config.Validate)

	// controller
	userController := http.NewUserController(userUsecase)
	statementTypeController := http.NewStatementTypeController(statementTypeUsecase)
	statementTypeItemController := http.NewStatementTypeItemController(statementTypeItemUsecase)
	mailController := http.NewMailController(mailUsecase)
	birthCertificateDocumentController := http.NewBirthCertificateDocumentController(birthCertificateDocumentUsecase)
	deathCertificateDocumentController := http.NewDeathCertificateDocumentController(deathCertificateDocumentUsecase)
	marriageStatementDocumentController := http.NewMarriageStatementDocumentController(marriageStatementDocumentUsecase)

	routeConfig := route.RouteConfig{
		App:                                 config.App,
		UserController:                      userController,
		StatementTypeController:             statementTypeController,
		StatementTypeItemController:         statementTypeItemController,
		MailController:                      mailController,
		BirthCertificateDocumentController:  birthCertificateDocumentController,
		DeathCertificateDocumentController:  deathCertificateDocumentController,
		MarriageStatementDocumentController: marriageStatementDocumentController,
	}

	routeConfig.Setup()
}
