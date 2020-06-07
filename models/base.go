package models

import "time"

//TimeModel tiempo
type TimeModel struct {
	CreatedAt time.Time  `json:"create_at,omitempty" gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"update_at,omitempty"`
	DeletedAt *time.Time `json:"delete_at,omitempty" sql:"index"`
}

// Modelsmall base modelsmall
//    type User struct {
//      gorm.Modelsmall
//    }
type Modelsmall struct {
	ID uint8 `json:"id,omitempty" gorm:"type:smallserial;NOT NULL;primary_key" `
	TimeModel
}

// Model base model
//    type User struct {
//      gorm.Model
//    }
type Model struct {
	ID uint32 `json:"id,omitempty" gorm:"type:serial;NOT NULL;primary_key"`
	TimeModel
}

// ModelBig base modelBig
//    type User struct {
//      gorm.ModelBig
//    }
type ModelBig struct {
	ID uint64 `json:"id,omitempty" gorm:"type:bigserial;NOT NULL;primary_key"`
	TimeModel
}

type TimeValidator struct {
	I  time.Time
	E  time.Time
	Zh int
	M  int
}
