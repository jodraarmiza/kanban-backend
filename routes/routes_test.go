package routes_test

import (
	"kanban-backend/routes"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestInitRoutes(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		e *echo.Echo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			routes.InitRoutes(tt.e)
		})
	}
}
