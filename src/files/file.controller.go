package files

import (
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"

	_ "github.com/forfam/authentication-service/src/config"
	"github.com/forfam/authentication-service/src/s3service"
	"github.com/forfam/authentication-service/src/utils/logger"
)

var log *logger.Logger
var defaultBucket string

func UploadFileEndpoint(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")

	if err != nil {
		log.Error("Something went wrong while file upload: " + err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "Something went wrong!",
		})
	}

	key, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Error("UUID generation is failed: " + err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Something went wrong!",
		})
	}

	log.Info("AWS Bucket is: " + defaultBucket)
	_, err = s3service.MultipartUpload(defaultBucket, string(key), file, log)

	if err != nil {
		log.Error("Something went wrong during file upload: " + err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Something went wrong!",
		})
	}

	log.Info("File uploaded successfuly: " + string(key))
	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "File uploaded successfuly",
		"path":    key,
	})
}

func init() {
	log = logger.New("AUTHENTICATION_SERVICE", "FileController")
	defaultBucket = os.Getenv("AWS_S3_BUCKET")
}
