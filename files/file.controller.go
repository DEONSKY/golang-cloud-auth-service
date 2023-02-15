package files

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	_ "github.com/forfam/authentication-service/config"
	"github.com/forfam/authentication-service/log"
	"github.com/forfam/authentication-service/s3service"
)

var logger *log.Logger
var defaultBucket string

func UploadFileEndpoint(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")

	if err != nil {
		logger.Error("Something went wrong while file upload: " + err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "Something went wrong!",
		})
	}

	key := uuid.New()
	if err != nil {
		logger.Error("UUID generation is failed: " + err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Something went wrong!",
		})
	}

	fileName, err := s3service.MultipartUpload(defaultBucket, key.String(), file)

	if err != nil {
		logger.Error("Something went wrong during file upload: " + err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Something went wrong!",
		})
	}

	logger.Info("File uploaded successfuly: " + fileName)
	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "File uploaded successfuly",
		"path":    fileName,
	})
}

func init() {
	logger = log.New("FileController")
	defaultBucket = os.Getenv("AWS_S3_BUCKET")
}
