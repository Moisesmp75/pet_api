package helpers

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"

	fbs "firebase.google.com/go/storage"
	"google.golang.org/api/option"
)

func initializeFirebaseStorage(ctx context.Context) (*fbs.Client, error) {
	credsPath := "serviceAccountKey.json"

	opt := option.WithCredentialsFile(credsPath)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing Firebase app: %v", err)
	}

	client, err := app.Storage(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing Firebase Storage client: %v", err)
	}
	return client, nil
}

func UploadFile(file *multipart.FileHeader, route string) (string, error) {
	ctx := context.Background()
	firebaseStorageClient, err := initializeFirebaseStorage(ctx)
	if err != nil {
		// log.Printf("Error inicializando Firebase Storage: %v", err)
		return "", err
		// return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
	}

	src, err := file.Open()
	if err != nil {
		return "", err
		// return c.Status(fiber.StatusInternalServerError).SendString("Error al abrir el archivo")
	}
	defer src.Close()

	filename := filepath.Base(GenerateUniqueFileName(file.Filename))
	bucketName := "hairypets.appspot.com"
	bucket, err := firebaseStorageClient.Bucket(bucketName)
	if err != nil {
		// log.Printf("Error obteniendo el bucket de Firebase Storage: %v", err)
		return "", err
		// return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
	}
	objectPath := route + filename
	obj := bucket.Object(objectPath)
	wc := obj.NewWriter(ctx)

	if _, err := io.Copy(wc, src); err != nil {
		// log.Printf("Error subiendo archivo a Firebase Storage: %v", err)
		return "", err
		// return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
	}

	if err := wc.Close(); err != nil {
		return "", err
		// log.Printf("Error cerrando escritor de Firebase Storage: %v", err)
		// return c.Status(fiber.StatusInternalServerError).SendString("Error interno del servidor")
	}

	url, err := bucket.SignedURL(obj.ObjectName(), &storage.SignedURLOptions{
		Expires: time.Now().AddDate(100, 0, 0),
		Method:  "GET",
	})
	if err != nil {
		return "", err
		// log.Fatalf("error getting download URL: %v\n", err)
	}

	return url, nil
}

func UploadFiles(files *multipart.Form, route string) ([]string, error) {
	var url_images []string
	if files == nil || len(files.File) == 0 {
		return nil, errors.New("no files provided to upload")
	}

	for _, headers := range files.File {
		for _, fileHeader := range headers {
			url, err := UploadFile(fileHeader, route)
			if err != nil {
				return nil, err
			}
			url_images = append(url_images, url)
		}
	}

	return url_images, nil
}
