package models

import (
	"booking-api/config"
	"fmt"
)

func MakeUrlPathForObjectStorage(objectName string) string {
	return "https://" + GetObjectStorageUrl() + "/" + objectName
}

func GetObjectStorageUrl() string {
	env, err := config.LoadConfig(".")
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	return env.OBJECT_STORAGE_URL
}
