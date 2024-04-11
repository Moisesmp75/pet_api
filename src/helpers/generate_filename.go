package helpers

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func GenerateUniqueFileName(originalName string) string {
	extension := filepath.Ext(originalName)
	name := strings.TrimSuffix(originalName, extension)
	cleanName := strings.ReplaceAll(strings.ToLower(name), " ", "_")
	uuid := uuid.New().String()
	uniqueName := fmt.Sprintf("%s_%s%s", cleanName, uuid, extension)
	return uniqueName
}
