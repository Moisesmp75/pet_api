package common

import (
	"encoding/base64"
	"errors"
	"io"
	"mime/multipart"
	"path/filepath"
)

func ConvertToBase64(img *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(img.Filename)
	if ext != ".png" && ext != ".jpg" {
		return "", errors.New("solo se permiten archivos PNG o JPG")
	}
	src, err := img.Open()
	if err != nil {
		return "", err
	}

	data, err := io.ReadAll(src)
	if err != nil {
		return "", err
	}

	encodedImage := base64.StdEncoding.EncodeToString(data)

	return encodedImage, nil
}