package utils

import "github.com/labstack/echo/v4"

// Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}

	return "./config/config-local"
}

// Read request body and validate
func ReadRequest(ctx echo.Context, request interface{}) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}

	return validate.StructCtx(ctx.Request().Context(), request)
}
