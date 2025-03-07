package user

import (
	"os"
)

func initUserJsonFile(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		initialData := []byte("[]")
		if err := os.WriteFile(filePath, initialData, 0644); err != nil {
			return errCreateUserFile
		}
	}
	return nil
}
