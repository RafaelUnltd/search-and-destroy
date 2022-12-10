package start

import (
	handlers "search-and-destroy/app/src/handlers/http"
	"search-and-destroy/app/src/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func StartHttpServer(db *gorm.DB, s services.ServicesContainer) (*echo.Echo, error) {
	server := echo.New()

	httpHandlersContainer := handlers.NewHandlersContainer(s)
	httpHandlersContainer.RegisterRoutes(server)

	return server, nil
}
