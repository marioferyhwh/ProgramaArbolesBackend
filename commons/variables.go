package commons

//Port puerto del servidor web
var Port int

//User nombre due usuario traido en el token
var User = "user"

//UserLevel nivel de usuario
//1 ver
//2 editar
//3 administrar
var UserLevel = map[int]map[string]int{
	1: {
		"user":                     3,
		"user_active":              3,
		"user_change_password":     3,
		"user_collection_active":   3,
		"user_collection":          3,
		"user_collection_password": 3,
		"collection":               3,
		"collection_active":        3,
		"locations":                3,
		"client":                   3,
		"loan":                     3,
		"payment":                  3,
		"payment_change":           3,
		"cash":                     3,
		"expense":                  3,
		"expense_change":           3,
		"document_type":            3,
		"business_type":            3,
		"tel_descript":             3,
	},
	2: {
		"user":                   2,
		"user_active":            1,
		"user_change_password":   2,
		"user_collection_active": 2,
		"user_collection":        1,
		//"user_collection_password": 0,
		"collection":        2,
		"collection_active": 1,
		"locations":         2,
		"client":            2,
		"loan":              2,
		"payment":           2,
		"payment_change":    2,
		"cash":              2,
		"expense":           2,
		"expense_change":    2,
		"document_type":     1,
		"business_type":     1,
		"tel_descript":      1,
	},
	3: {
		"user":                   2,
		"user_active":            1,
		"user_change_password":   2,
		"user_collection_active": 1,
		//"user_collection":          0,
		//"user_collection_password": 0,
		"collection": 1,
		//"collection_active":        0,
		"locations": 2,
		"client":    2,
		"loan":      2,
		"payment":   2,
		//"payment_change":           0,
		"cash":    1,
		"expense": 2,
		//"expense_change":           0,
		"document_type": 1,
		"business_type": 1,
		"tel_descript":  1,
	},
}
