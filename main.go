package main

import (
	"fmt"
	"log"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
	"github.com/yourusername/huma-license-api/handlers"
	"github.com/yourusername/huma-license-api/models"
)

func main() {
	r := gin.Default()

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	config := huma.DefaultConfig("License API", "1.0.0")
	config.OpenAPI.Servers = []*huma.Server{
		{
			URL:         fmt.Sprintf("http://%s:%s", host, port),
			Description: "Dynamic Server",
		},
	}

	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"BearerAuth": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}

	api := humagin.New(r, config)

	huma.Register[handlers.CreateLicenseInput, models.License](api, huma.Operation{
		Method:      "POST",
		Path:        "/licenses",
		Summary:     "Create a license",
		Description: "Create a new software license",
		Security: []map[string][]string{
			{"BearerAuth": {}},
		},
	}, handlers.CreateLicenseHandler)

	huma.Register[handlers.UpdateLicenseInput, models.License](api, huma.Operation{
		Method:      "PATCH",
		Path:        "/licenses/{shortname}",
		Summary:     "Update a license",
		Description: "Update an existing license by its short name",
		Security: []map[string][]string{
			{"BearerAuth": {}},
		},
	}, handlers.UpdateLicenseHandler)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
