package util

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SaveValidatedFile(ctx *fiber.Ctx, file *multipart.FileHeader, fieldName string) (string, error) {
	allowedExtensions := map[string]bool{
		".jpg": true, ".jpeg": true, ".png": true, ".pdf": true,
	}
	ext := strings.ToLower(filepath.Ext(file.Filename))

	if !allowedExtensions[ext] {
		return "", fmt.Errorf("file format not allowed for %s", fieldName)
	}

	filename := filepath.Base(file.Filename)
	generateFilename := GenerateRandomFilename(filename)
	savePath := filepath.Join("./uploads", generateFilename)

	if err := ctx.SaveFile(file, savePath); err != nil {
		return "", fmt.Errorf("failed to save %s file", fieldName)
	}

	return generateFilename, nil
}

func CleanupFiles(savedFiles map[string]string) {
	for _, filePath := range savedFiles {
		_ = os.Remove(filePath)
	}
}
