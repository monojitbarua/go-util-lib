package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const LOG_FILE string = `/application.log`

// info exported
var info *log.Logger

// warning exported
var warn *log.Logger

// error exported
var error *log.Logger

func init() {
	absPath, err := filepath.Abs("../log")
	if err != nil {
		fmt.Println("Error reading given path:", err)
	}

	logFile, err := os.OpenFile(absPath+LOG_FILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	info = log.New(logFile, "INFO:\t", log.Ldate|log.Ltime|log.Lshortfile)
	warn = log.New(logFile, "WARN:\t", log.Ldate|log.Ltime|log.Lshortfile)
	error = log.New(logFile, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(message string) {
	info.Println(message)
}
func Warn(message string) {
	warn.Println(message)
}
func Error(message string) {
	error.Println(message)
}
