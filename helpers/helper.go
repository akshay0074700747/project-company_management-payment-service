package helpers

import (
	"fmt"

	"github.com/google/uuid"
)

func PrintErr(err error, messge string) {
	fmt.Println(messge, err)
}

func PrintMsg(msg string) {
	fmt.Println(msg)
}

func GenUuid() string {
	return uuid.New().String()
}
