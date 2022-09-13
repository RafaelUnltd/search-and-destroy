package start

import (
	handlers "poly-shooters/app/src/handlers/http"
	"poly-shooters/app/src/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func StartHttpServer(db *gorm.DB, s services.ServicesContainer) (*echo.Echo, error) {
	server := echo.New()

	httpHandlersContainer := handlers.NewHandlersContainer(s)
	httpHandlersContainer.RegisterRoutes(server)

	return server, nil
}
