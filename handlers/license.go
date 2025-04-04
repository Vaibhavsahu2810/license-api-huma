package handlers

import (
	"context"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"github.com/yourusername/huma-license-api/models"
)

var licenseStore = map[string]models.License{}

type CreateLicenseInput struct {
	Body models.License `json:"body" doc:"New license to be created" required:"true"`
}

func CreateLicenseHandler(ctx context.Context, input *CreateLicenseInput) (*models.License, error) {
	license := input.Body
	licenseStore[license.ShortName] = license
	return &license, nil
}

type UpdateLicenseInput struct {
	Path struct {
		ShortName string `path:"shortname" doc:"Short name of the license to update" required:"true"`
	} `json:"-"`
	Body models.License `json:"body" doc:"Updated license data" required:"true"`
}

func UpdateLicenseHandler(ctx context.Context, input *UpdateLicenseInput) (*models.License, error) {
	shortName := input.Path.ShortName

	license, ok := licenseStore[shortName]
	if !ok {
		return nil, huma.NewError(404, fmt.Sprintf("License with shortname %s not found", shortName))
	}

	license = input.Body
	license.ShortName = shortName
	licenseStore[shortName] = license

	return &license, nil
}
