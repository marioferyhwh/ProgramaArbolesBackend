package models

//GeneralValues listado de configuraciones base
type GeneralValues struct {
	BusinessTypes []BusinessType         `json:"business_types,omitempty" gorm:"-"`
	TelDescrips   []TelDescrip           `json:"tel_descriptions,omitempty" gorm:"-"`
	LoanStates    []LoanState            `json:"loan_states,omitempty" gorm:"-"`
	UserLevels    []UserLevel            `json:"user_levels,omitempty" gorm:"-"`
	DocumentTypes []DocumentType         `json:"document_types,omitempty" gorm:"-"`
	Levels        map[int]map[string]int `json:"levels,omitempty" gorm:"-"`
}
