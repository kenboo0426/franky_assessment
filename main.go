package main

import (
	"github.com/joho/godotenv"
	"github.com/kenboo0426/franky_assessment/interfaces"
)

func main() {
	godotenv.Load()
	interfaces.InitializeHTTPServer()
}
