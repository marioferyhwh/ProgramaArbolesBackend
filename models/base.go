package models

import "time"

//TimeModel tiempo
type TimeModel struct {
	CreatedAt time.Time `gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Modelsmall base modelsmall
//    type User struct {
//      gorm.Modelsmall
//    }
type Modelsmall struct {
	ID uint `gorm:"type:smallserial;NOT NULL;primary_key"`
	TimeModel
}

// Model base model
//    type User struct {
//      gorm.Model
//    }
type Model struct {
	ID uint `gorm:"type:serial;NOT NULL;primary_key"`
	TimeModel
}

// ModelBig base modelBig
//    type User struct {
//      gorm.ModelBig
//    }
type ModelBig struct {
	ID uint `gorm:"type:bigserial;NOT NULL;primary_key"`
	TimeModel
}
