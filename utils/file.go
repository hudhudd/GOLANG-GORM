package utils

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const DefaultPathAssetImage = "./public/covers/"

func HandleSingleFile(ctx *fiber.Ctx) error {

	//Handle File
	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		log.Println("Error File = ", errFile)
	}

	var filename *string
	if file != nil {

		errCheckContentType := checkContentType(file, "image/jpeg", "image/jpg", "image/png", "image/gif")
		if errCheckContentType != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": errCheckContentType.Error(),
			})
		}

		filename = &file.Filename
		extentionFile := filepath.Ext(*filename)
		newFilename := fmt.Sprintf("gambar-satu%s", extentionFile)

		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", newFilename))
		if errSaveFile != nil {
			log.Println("Failed to Save File")
		}

	} else {
		log.Println("No cover to be uploaded")
	}

	if filename != nil {
		ctx.Locals("filename", *filename)
	} else {
		ctx.Locals("filename", nil)
	}

	return ctx.Next()
}

func HandleMultipleFile(ctx *fiber.Ctx) error {
	form, errForm := ctx.MultipartForm()
	if errForm != nil {
		log.Println(errForm)
	}

	files := form.File["photos"]

	var filenames []string

	for i, file := range files {
		var filename string
		if file != nil {
			filename = fmt.Sprintf("%d-%s", i, file.Filename)

			errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", filename))
			if errSaveFile != nil {
				log.Println("Failed to Save File")
			}
		} else {
			log.Println("No cover to be uploaded")
		}

		if filename != "" {
			filenames = append(filenames, filename)
		}
		ctx.Locals("filenames", filenames)
	}

	return ctx.Next()
}

func HandleRemoveFile(filename string, pathFile ...string) error {
	if len(pathFile) > 0 {
		err := os.Remove(pathFile[0] + filename)
		if err != nil {
			log.Println("failed to delete file")
			return err
		}

	} else {
		err := os.Remove(DefaultPathAssetImage + filename)
		if err != nil {
			log.Println("failed to delete file")
			return err
		}
	}

	return nil
}

func checkContentType(file *multipart.FileHeader, contentTypes ...string) error {
	if len(contentTypes) > 0 {
		for _, contentType := range contentTypes {
			contentTypeFile := file.Header.Get("Content-Type")
			log.Println(contentTypeFile)
			if contentTypeFile == contentType {
				return nil
			}
		}

		return errors.New("File Type doesn't supported")
	} else {
		return errors.New("content type not founded to be checked")
	}
}
