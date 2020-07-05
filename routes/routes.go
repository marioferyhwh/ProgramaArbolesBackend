package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/marioferyhwh/IMFBackend_forest/commons"
	"github.com/marioferyhwh/IMFBackend_forest/controllers"
	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//InitRoutes en rutamiento de la API
func InitRoutes(e *echo.Echo) {
	const prefixAPI = "/api/v1/"

	const prefixGeneralValuesGet = prefixAPI + "general"
	e.GET(prefixGeneralValuesGet, SetGeneralValuesGetRoutes, controllers.ValidateJWT)

	const prefixBusinessType = prefixAPI + "busindess"
	e.GET(prefixBusinessType+"/:id", SetBusinessTypeGetRoutes, controllers.ValidateJWT)
	e.GET(prefixBusinessType+"/list/:id", SetBusinessTypeGetListRoutes, controllers.ValidateJWT)
	e.POST(prefixBusinessType, SetBusinessTypeCreateRoutes, controllers.ValidateJWT)
	e.PUT(prefixBusinessType+"/:id", SetBusinessTypeEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixBusinessType+"/:id", SetBusinessTypeDeleteRoutes, controllers.ValidateJWT)

	const prefixClient = prefixAPI + "client"
	const prefixClientTel = prefixClient + "tel"
	e.GET(prefixClientTel+"/:id", SetClientTelGetRoutes, controllers.ValidateJWT)
	e.GET(prefixClientTel+"/list/:id", SetClientTelGetListRoutes, controllers.ValidateJWT)
	e.POST(prefixClientTel, SetClientTelCreateRoutes, controllers.ValidateJWT)
	e.PUT(prefixClientTel+"/:id", SetClientTelEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixClientTel+"/:id", SetClientTelDeleteRoutes, controllers.ValidateJWT)
	const prefixClientListLocation = prefixClient + "location"
	e.GET(prefixClientListLocation+"/:id", SetClientListLocationGetRoutes, controllers.ValidateJWT)
	e.GET(prefixClientListLocation+"/list/:id", SetClientListLocationGetListRoutes, controllers.ValidateJWT)
	e.POST(prefixClientListLocation, SetClientListLocationCreateRoutes, controllers.ValidateJWT)
	e.PUT(prefixClientListLocation+"/:id", SetClientListLocationEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixClientListLocation+"/:id", SetClientListLocationDeleteRoutes, controllers.ValidateJWT)
	e.GET(prefixClient+"/:id", SetClientGetRoutes, controllers.ValidateJWT)
	e.GET(prefixClient+"/list/:id", SetClientGetListRoutes, controllers.ValidateJWT)
	e.POST(prefixClient, SetClientCreateRoutes, controllers.ValidateJWT)
	e.PUT(prefixClient+"/:id", SetClientEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixClient+"/:id", SetClientDeleteRoutes, controllers.ValidateJWT)

	const prefixCollection = prefixAPI + "collection"
	const prefixCollectionCash = prefixCollection + "cash"
	e.GET(prefixCollectionCash+"/:id", SetCollectionCashGetRoutes, controllers.ValidateJWT)
	e.GET(prefixCollectionCash+"/list/:id", SetCollectionCashGetListRoutes, controllers.ValidateJWT)
	e.POST(prefixCollectionCash, SetCollectionCashCreateRoutes, controllers.ValidateJWT)
	e.PUT(prefixCollectionCash+"/:id", SetCollectionCashEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixCollectionCash+"/:id", SetCollectionCashDeleteRoutes, controllers.ValidateJWT)
	e.GET(prefixCollection+"/:id", SetCollectionGetRoutes, controllers.ValidateJWT)
	e.GET(prefixCollection+"/list/:id", SetCollectionGetListRoutes, controllers.ValidateJWT)
	e.POST(prefixCollection, SetCollectionCreateRoutes, controllers.ValidateJWT)
	e.PUT(prefixCollection+"/:id", SetCollectionEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixCollection+"/:id", SetCollectionDeleteRoutes, controllers.ValidateJWT)

	const prefixDocumentType = prefixAPI + "document"
	e.GET(prefixDocumentType+"/:id", SetDocumentTypeGetRoutes, controllers.ValidateJWT)
	e.GET(prefixDocumentType+"/list/:id", SetDocumentTypeGetListRoutes, controllers.ValidateJWT)
	e.POST(prefixDocumentType, SetDocumentTypeCreateRoutes, controllers.ValidateJWT)
	e.PUT(prefixDocumentType+"/:id", SetDocumentTypeEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixDocumentType+"/:id", SetDocumentTypeDeleteRoutes, controllers.ValidateJWT)

	const prefixExpense = prefixAPI + "expense"
	const prefixExpenseDescrip = prefixExpense + "descrip"
	e.GET(prefixExpenseDescrip+"/:id", SetExpenseDescripGetRoutes, controllers.ValidateJWT)
	e.GET(prefixExpenseDescrip+"/list/:id", SetExpenseDescripGetListRoutes, controllers.ValidateJWT)
	e.POST(prefixExpenseDescrip, SetExpenseDescripCreateRoutes, controllers.ValidateJWT)
	e.PUT(prefixExpenseDescrip+"/:id", SetExpenseDescripEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixExpenseDescrip+"/:id", SetExpenseDescripDeleteRoutes, controllers.ValidateJWT)
	e.GET(prefixExpense+"/:id", SetExpenseGetRoutes, controllers.ValidateJWT)
	e.GET(prefixExpense+"/list/:id", SetExpenseGetListRoutes, controllers.ValidateJWT)
	e.POST(prefixExpense, SetExpenseCreateRoutes, controllers.ValidateJWT)
	e.PUT(prefixExpense+"/:id", SetExpenseEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixExpense+"/:id", SetExpenseDeleteRoutes, controllers.ValidateJWT)

	const prefixLoan = prefixAPI + "loan"
	const prefixLoanState = prefixLoan + "loan_state"
	e.GET(prefixLoanState+"/:id", SetLoanStateGetRoutes, controllers.ValidateJWT)
	e.GET(prefixLoanState+"/list/:id", SetLoanStateGetListRoutes, controllers.ValidateJWT)
	e.POST(prefixLoanState, SetLoanStateCreateRoutes, controllers.ValidateJWT)
	e.PUT(prefixLoanState+"/:id", SetLoanStateEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixLoanState+"/:id", SetLoanStateDeleteRoutes, controllers.ValidateJWT)
	const prefixLoanPayment = prefixLoan + "loan_payment"
	e.GET(prefixLoanPayment+"/:id", SetLoanPaymentGetRoutes, controllers.ValidateJWT)
	e.GET(prefixLoanPayment+"/list/:id", SetLoanPaymentGetListRoutes, controllers.ValidateJWT)
	e.POST(prefixLoanPayment, SetLoanPaymentCreateRoutes, controllers.ValidateJWT)
	e.PUT(prefixLoanPayment+"/:id", SetLoanPaymentEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixLoanPayment+"/:id", SetLoanPaymentDeleteRoutes, controllers.ValidateJWT)
	e.GET(prefixLoan+"/:id", SetLoanGetRoutes, controllers.ValidateJWT)
	e.GET(prefixLoan+"/list/:id", SetLoanGetListRoutes, controllers.ValidateJWT)
	e.POST(prefixLoan, SetLoanCreateRoutes, controllers.ValidateJWT)
	e.PUT(prefixLoan+"/:id", SetLoanEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixLoan+"/:id", SetLoanDeleteRoutes, controllers.ValidateJWT)

	const prefixTelDescription = prefixAPI + "tel"
	e.GET(prefixTelDescription+"/:id", SetTelDescripGetRoutes, controllers.ValidateJWT)
	e.GET(prefixTelDescription+"/list/:id", SetTelDescripGetListRoutes, controllers.ValidateJWT)
	e.POST(prefixTelDescription, SetTelDescripCreateRoutes, controllers.ValidateJWT)
	e.PUT(prefixTelDescription+"/:id", SetTelDescripEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixTelDescription+"/:id", SetTelDescripDeleteRoutes, controllers.ValidateJWT)

	const prefixUser = prefixAPI + "user"
	const prefixUserTel = prefixUser + "tel"
	e.GET(prefixUserTel+"/:id", SetUserTelGetRoutes, controllers.ValidateJWT)
	e.GET(prefixUserTel+"/list/:id", SetUserTelGetListRoutes, controllers.ValidateJWT)
	e.POST(prefixUserTel, SetUserTelCreateRoutes, controllers.ValidateJWT)
	e.PUT(prefixUserTel+"/:id", SetUserTelEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixUserTel+"/:id", SetUserTelDeleteRoutes, controllers.ValidateJWT)
	const prefixUserLevel = prefixUser + "level"
	e.GET(prefixUserLevel+"/:id", SetUserLevelGetRoutes, controllers.ValidateJWT)
	e.GET(prefixUserLevel+"/list/:id", SetUserLevelGetListRoutes, controllers.ValidateJWT)
	e.POST(prefixUserLevel, SetUserLevelCreateRoutes, controllers.ValidateJWT)
	e.PUT(prefixUserLevel+"/:id", SetUserLevelEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixUserLevel+"/:id", SetUserLevelDeleteRoutes, controllers.ValidateJWT)
	const prefixUserCollection = prefixUser + "collection"
	//e.GET(prefixUserCollection+"/:user/:collection", SetUserCollectionGetRoutes, controllers.ValidateJWT)

	e.GET(prefixUserCollection+"/:id", SetUserCollectionGetRoutes, controllers.ValidateJWT)
	e.GET(prefixUserCollection+"/list/:id", SetUserCollectionGetListRoutes, controllers.ValidateJWT)
	e.POST(prefixUserCollection, SetUserCollectionCreateRoutes, controllers.ValidateJWT)
	e.PUT(prefixUserCollection+"/:id", SetUserCollectionEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixUserCollection+"/:id", SetUserCollectionDeleteRoutes, controllers.ValidateJWT)
	e.POST(prefixUser, SetUserCreateRoutes, controllers.ValidateJWT)
	e.GET(prefixUser+"/:id", SetUserGetRoutes, controllers.ValidateJWT)
	e.POST(prefixUser+"/login", SetLoginRoutes)
	e.PUT(prefixUser+"/:id", SetUserEditRoutes, controllers.ValidateJWT)
	e.DELETE(prefixUser+"/:id", SetUserDeleteRoutes, controllers.ValidateJWT)
	e.GET("/*", SetDefaultRoutes)
	e.POST("/*", SetDefaultRoutes)
	e.DELETE("/*", SetDefaultRoutes)
	e.PUT("/*", SetDefaultRoutes)

	fmt.Println(prefixGeneralValuesGet)
	fmt.Println(prefixBusinessType)
	fmt.Println(prefixClient)
	fmt.Println(prefixClientTel)
	fmt.Println(prefixClientListLocation)
	fmt.Println(prefixCollection)
	fmt.Println(prefixCollectionCash)
	fmt.Println(prefixDocumentType)
	fmt.Println(prefixExpense)
	fmt.Println(prefixExpenseDescrip)
	fmt.Println(prefixLoan)
	fmt.Println(prefixLoanState)
	fmt.Println(prefixLoanPayment)
	fmt.Println(prefixTelDescription)
	fmt.Println(prefixUser)
	fmt.Println(prefixUserTel)
	fmt.Println(prefixUserLevel)
	fmt.Println(prefixUserCollection)
}

//SetDefaultRoutes Creacion de usuario
func SetDefaultRoutes(c echo.Context) error {
	m := models.Message{}
	m.Code = http.StatusBadRequest
	m.Message = "no existe esa pagina"
	return commons.DisplayMessage(c, &m)
}
