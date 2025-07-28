package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

const controllerTemplate = `package controllers

import "github.com/gin-gonic/gin"

func {{.ControllerName}}(c *gin.Context) {
	c.JSON(200, gin.H{"message": "{{.ControllerName}} is working!"})
}
`

const modelTemplate = `package models

	import (
		"time"

		"github.com/uptrace/bun"
	)

	type {{.ModelName}} struct {
		bun.BaseModel ` + "`bun:\"table:{{.TableName}}s\"`" + `
		ID        int64     ` + "`bun:\"id,pk,autoincrement\"`" + `
		Name      string    ` + "`bun:\"name,notnull\"`" + `
		CreatedAt time.Time ` + "`bun:\"created_at,default:now()\"`" + `
		UpdatedAt time.Time ` + "`bun:\"updated_at,default:now()\"`" + `
	}
`

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run cmd/commands/make.go controller ControllerName")
		return
	}

	command := os.Args[1]
	name := os.Args[2]

	switch command {
	case "controller":
		createController(name)
	case "model":
		createModel(name)
	default:
		fmt.Println("Unknown command:", command)
	}
}

func createController(name string) {
	controllerName := "Index" // will be method name
	fileName := strings.ToLower(name) + ".go"
	filePath := "internal/api/http/controllers/" + fileName

	_ = os.MkdirAll("internal/api/http/controllers", os.ModePerm)

	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("⚠️ Controller already exists:", filePath)
		return
	}

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("❌ Error creating controller:", err)
		return
	}
	defer f.Close()

	tmpl, _ := template.New("controller").Parse(controllerTemplate)
	tmpl.Execute(f, map[string]string{
		"ControllerName": controllerName,
	})

	fmt.Println("✅ Controller created at:", filePath)
}

func createModel(name string) {
	modelName := strings.Title(name)
	tableName := strings.ToLower(name)
	fileName := strings.ToLower(name) + ".go"
	filePath := "internal/models/" + fileName

	_ = os.MkdirAll("internal/models", os.ModePerm)

	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("⚠️ Model already exists:", filePath)
		return
	}

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("❌ Error creating model:", err)
		return
	}
	defer f.Close()

	tmpl, _ := template.New("model").Parse(modelTemplate)
	tmpl.Execute(f, map[string]string{
		"ModelName": modelName,
		"TableName": tableName,
	})

	fmt.Println("✅ Model created at:", filePath)
}
