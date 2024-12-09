package main

import (
	"fmt"
	"github.com/anhao/fiber-bootstrap/bootstrap"
	"github.com/anhao/fiber-bootstrap/pkg/env"
	"log"
)

func main() {
	app := bootstrap.NewApplication()
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", env.GetEnv("APP_HOST", "localhost"), env.GetEnv("APP_PORT", "3000"))))
}
