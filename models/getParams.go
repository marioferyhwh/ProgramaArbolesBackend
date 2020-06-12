package models

import "github.com/labstack/echo"

//GetParams estructura para traer parametros en la url
type GetParams struct {
	C echo.Context
	P string
}
