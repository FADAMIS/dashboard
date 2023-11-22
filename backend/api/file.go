package api

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UploadImage(ctx *gin.Context) string {
	formFile, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusUnsupportedMediaType, gin.H{
			"message": "mime type must be multipart/form-data",
		})

		return ""
	}

	// read file
	file, _ := formFile.Open()
	defer file.Close()

	// read raw data
	data := make([]byte, 512)
	io.ReadFull(file, data)

	// detect mime type
	mimeType := http.DetectContentType(data)
	fileExtension := strings.Split(mimeType, "/")[1]

	if mimeType != "image/jpeg" && mimeType != "image/png" {
		ctx.JSON(http.StatusUnsupportedMediaType, gin.H{
			"message": "file must be a jpeg or a png",
		})

		return ""
	}

	// rewind file pointer to starting position
	file.Seek(0, io.SeekStart)

	// calculate sha256 checksum of the file
	h := sha256.New()
	io.Copy(h, file)
	hashed := hex.EncodeToString(h.Sum(nil))

	// update filename based on hash and mimetype
	formFile.Filename = hashed + "." + fileExtension
	ctx.SaveUploadedFile(formFile, "./uploads/"+formFile.Filename)

	return formFile.Filename
}
