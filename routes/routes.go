package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/controllers"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//InitRoutes en rutamiento de la API
func InitRoutes(e *echo.Echo) {
	const prefixAPI = "/api/v1/"
	const prefixBusinessType = prefixAPI + "busindess"
	e.POST(prefixBusinessType, SetBusinessTypeCreateRoutes, controllers.ValidateJWT)
	e.GET(prefixBusinessType+"/:id", SetBusinessTypeGetRoutes, controllers.ValidateJWT)
	e.PUT(prefixBusinessType+"/:id", SetBusinessTypeEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixBusinessType+"/:id", SetBusinessTypeDeleteRoutes, controllers.ValidateJWT)
	const prefixClient = prefixAPI + "client"
	e.POST(prefixClient, SetUserCreateRoutes, controllers.ValidateJWT)
	e.GET(prefixClient+"/:id", SetUserGetRoutes, controllers.ValidateJWT)
	e.PUT(prefixClient+"/:id", SetUserEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixClient+"/:id", SetUserDeleteRoutes, controllers.ValidateJWT)
	const prefixCollection = prefixAPI + "collection"
	e.POST(prefixCollection, SetUserCreateRoutes, controllers.ValidateJWT)
	e.GET(prefixCollection+"/:id", SetUserGetRoutes, controllers.ValidateJWT)
	e.PUT(prefixCollection+"/:id", SetUserEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixCollection+"/:id", SetUserDeleteRoutes, controllers.ValidateJWT)
	const prefixDocumentType = prefixAPI + "document"
	e.POST(prefixDocumentType, SetUserCreateRoutes, controllers.ValidateJWT)
	e.GET(prefixDocumentType+"/:id", SetUserGetRoutes, controllers.ValidateJWT)
	e.PUT(prefixDocumentType+"/:id", SetUserEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixDocumentType+"/:id", SetUserDeleteRoutes, controllers.ValidateJWT)
	const prefixExpense = prefixAPI + "expense"
	e.POST(prefixExpense, SetUserCreateRoutes, controllers.ValidateJWT)
	e.GET(prefixExpense+"/:id", SetUserGetRoutes, controllers.ValidateJWT)
	e.PUT(prefixExpense+"/:id", SetUserEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixExpense+"/:id", SetUserDeleteRoutes, controllers.ValidateJWT)
	const prefixLoan = prefixAPI + "loan"
	e.POST(prefixLoan, SetUserCreateRoutes, controllers.ValidateJWT)
	e.GET(prefixLoan+"/:id", SetUserGetRoutes, controllers.ValidateJWT)
	e.PUT(prefixLoan+"/:id", SetUserEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixLoan+"/:id", SetUserDeleteRoutes, controllers.ValidateJWT)
	const prefixTelDescription = prefixAPI + "tel"
	e.POST(prefixTelDescription, SetUserCreateRoutes, controllers.ValidateJWT)
	e.GET(prefixTelDescription+"/:id", SetUserGetRoutes, controllers.ValidateJWT)
	e.PUT(prefixTelDescription+"/:id", SetUserEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixTelDescription+"/:id", SetUserDeleteRoutes, controllers.ValidateJWT)
	const prefixUser = prefixAPI + "user"
	e.GET("login", SetLoginRoutes)
	e.POST(prefixUser, SetUserCreateRoutes, controllers.ValidateJWT)
	e.GET(prefixUser+"/:id", SetUserGetRoutes, controllers.ValidateJWT)
	e.PUT(prefixUser+"/:id", SetUserEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixUser+"/:id", SetUserDeleteRoutes, controllers.ValidateJWT)
	e.GET("/*", SetDefaultRoutes)
	e.POST("/*", SetDefaultRoutes)
	e.DELETE("/*", SetDefaultRoutes)
	e.PUT("/*", SetDefaultRoutes)

}

//SetDefaultRoutes Creacion de usuario
func SetDefaultRoutes(c echo.Context) error {

	m := models.Message{}

	m.Code = http.StatusBadRequest
	m.Message = "no existe esa pagina"

	return commons.DisplayMessage(c, &m)
}
