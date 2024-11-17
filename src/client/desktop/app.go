package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Login(username string, password string) {
	resp, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer([]byte(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password))))
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
}
