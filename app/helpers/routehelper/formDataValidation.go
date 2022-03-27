package routehelper

import (
	"mime/multipart"
	"net/http"
)

var allowedMimeType = map[string]bool{
	"image/jpeg":   true,
	"image/png":    true,
	"image/x-icon": true,
	"image/apng":   true,
	"image/ico":    true,
}

func ValidateImageForm(file multipart.File) (isImage bool, mimeType string) {
	// Create a buffer to store the header of the file in
	fileHeader := make([]byte, 512)

	// Copy the headers into the FileHeader buffer
	if _, err := file.Read(fileHeader); err != nil {
		return false, ""
	}

	contentType := http.DetectContentType(fileHeader)
	_, isImage = allowedMimeType[contentType]
	if !isImage {
		return false, contentType
	}

	return true, contentType
}
