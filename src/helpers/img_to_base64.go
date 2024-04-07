package common

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
)

func ConvertToBase64(img *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(img.Filename)
	if ext != ".png" && ext != ".jpg" {
		return "",fmt.Errorf("solo se permiten archivos PNG o JPG: %s", img.Filename)
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

func ImagesToBase64(files *multipart.Form) ([]string, error) {
	var base64Images []string

	if files == nil || len(files.File) == 0 {
		return nil, errors.New("no se han proporcionado archivos para convertir")
	}
	
	for _, headers := range files.File {
		for _, fileHeader := range headers {
			encodedImage, err := ConvertToBase64(fileHeader)
			if err != nil {
				return nil, err
			}
			base64Images = append(base64Images, encodedImage)
		}
	}
	return base64Images, nil
}