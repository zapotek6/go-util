package util

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var Run = true

func SetupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Ctrl+C pressed in Terminal")
		Run = false
	}()
}

func GetEnv(varName string, defaultValue string) string {
	variable := os.Getenv(varName)

	if variable == "" {
		variable = defaultValue
	}

	return variable
}