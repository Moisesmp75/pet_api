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

func UploadFile(file *multipart.FileHeader, route string, generateName bool) (string, string, error) {
	ctx := context.Background()
	firebaseStorageClient, err := initializeFirebaseStorage(ctx)
	if err != nil {
		return "", file.Filename, err
	}

	src, err := file.Open()
	if err != nil {
		return "", file.Filename, err
	}
	defer src.Close()
	filename := file.Filename
	if generateName {
		filename = filepath.Base(GenerateUniqueFileName(file.Filename))
	}

	bucketName := "hairypets.appspot.com"
	bucket, err := firebaseStorageClient.Bucket(bucketName)
	if err != nil {
		return "", filename, err
	}
	objectPath := route + filename
	obj := bucket.Object(objectPath)
	wc := obj.NewWriter(ctx)

	if _, err := io.Copy(wc, src); err != nil {
		return "", filename, err
	}

	if err := wc.Close(); err != nil {
		return "", filename, err
	}

	url, err := bucket.SignedURL(obj.ObjectName(), &storage.SignedURLOptions{
		Expires: time.Now().AddDate(100, 0, 0),
		Method:  "GET",
	})
	if err != nil {
		return "", filename, err
	}

	return url, filename, nil
}

func UploadFiles(files *multipart.Form, route string) ([]string, []string, error) {
	var url_images []string
	var filenames []string
	if files == nil || len(files.File) == 0 {
		return nil, nil, errors.New("no files provided to upload")
	}

	for _, headers := range files.File {
		for _, fileHeader := range headers {
			url, filename, err := UploadFile(fileHeader, route, true)
			if err != nil {
				return nil, nil, err
			}
			url_images = append(url_images, url)
			filenames = append(filenames, filename)
		}
	}

	return url_images, filenames, nil
}
