package util

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var Run = true

func SetupCloseHandler() {
	SetupCloseHandlerWFunc(nil)
}

func SetupCloseHandlerWFunc(f func()) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Ctrl+C pressed in Terminal")
		if nil != f {
			f()
		}
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

func GetEnvAsInt(varName string, defaultValue int) (int, error) {
	variable := os.Getenv(varName)

	var value int

	if variable == "" {
		value = defaultValue
	} else {
		if val, err := strconv.ParseInt(variable, 10, 64); nil == err {
			return int(val), nil
		} else {
			return -1, err
		}
	}

	return value, nil
}

func GetEnvAsDuration(varName string, defaultValue time.Duration) (time.Duration, error) {
	variable := os.Getenv(varName)

	if variable == "" {
		return defaultValue, nil
	} else {
		return time.ParseDuration(variable)
	}
}

func GetEnvAsBool(varName string, defaultValue bool) (bool, error) {
	variable := os.Getenv(varName)

	if variable == "" {
		return defaultValue, nil
	} else {
		return strconv.ParseBool(variable)
	}
}
