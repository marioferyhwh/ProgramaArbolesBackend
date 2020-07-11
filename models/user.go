package models

import "time"

//User usuario
type User struct {
	Model
	Actived         bool       `json:"actived,omitempty"          gorm:"type:bool;not null;default:true"`
	NickName        string     `json:"nick_name,omitempty"        gorm:"type:varchar(50);not null;default:''"`
	Email           string     `json:"email,omitempty"            gorm:"type:varchar(100);not null"`
	Password        string     `json:"password,omitempty"         gorm:"type:varchar(256);not null"`
	CodDocumentType string     `json:"document_code,omitempty"    gorm:"type:varchar(3);not null;default:''"`
	Document        string     `json:"document,omitempty"         gorm:"type:numeric(11);not null"`
	Name            string     `json:"name,omitempty"             gorm:"type:varchar(50);not null;default:''"`
	Admin           bool       `json:"admin,omitempty"            gorm:"type:bool;not null;default:false"`
	TimeZone        int8       `json:"time_zone,omitempty"        gorm:"type:smallint;not null;default:-5"`
	PassewordAt     *time.Time `json:"password_at,omitempty"      gorm:"NOT NULL;DEFAULT:CURRENT_TIMESTAMP"`
	ResetPassword   bool       `json:"change-password,omitempty"  gorm:"type:bool;not null;default:true"`
	ConfirmPassword string     `json:"confirm_password,omitempty" gorm:"-"`

	UserCollection []UserCollection `json:"collections,omitempty"          gorm:"foreignkey:CodUser;association_foreignkey:id"`
	UserTel        []UserTel        `json:"tels,omitempty"                 gorm:"foreignkey:CodUser;association_foreignkey:id"`
	UserTelNew     []UserTel        `json:"tels_new,omitempty"             gorm:"foreignkey:CodUser;association_foreignkey:id"`
	UserTelDelete  []UserTel        `json:"tels_delete,omitempty"          gorm:"foreignkey:CodUser;association_foreignkey:id"`
	Expense        []Expense        `json:"expenses,omitempty"             gorm:"foreignkey:CodUser;association_foreignkey:id"`

	CodCollection   uint32 `json:"-" gorm:"-"`
	GetDocumentType bool   `json:"-" gorm:"-"`
	GetListUser     bool   `json:"-" gorm:"-"`
	GetUserTel      bool   `json:"-" gorm:"-"`
	GetExpense      bool   `json:"-" gorm:"-"`
	ChangePassword  bool   `json:"-" gorm:"-"`
	ChangeActived   bool   `json:"-" gorm:"-"`
}

/*
DROP TABLE public.users;

CREATE TABLE public.users
(
  id serial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  active bool DEFAULT TRUE,
  nick_name varchar(50) NOT NULL DEFAULT '',
  email varchar(100) NOT NULL,
  password varchar(256) NOT NULL,
  cod_document_type varchar(3) NOT NULL DEFAULT 'CC',
  document NUMERIC(11) NOT NULL ,
  name varchar(50) NOT NULL DEFAULT '',

  CONSTRAINT pk_users PRIMARY KEY(id)
);

ALTER TABLE public.users ADD
  CONSTRAINT uk_users_nickname UNIQUE(nickname);

ALTER TABLE public.users ADD
  CONSTRAINT uk_users_email UNIQUE(email);

ALTER TABLE public.users ADD
  CONSTRAINT uk_users_cdocumentt_document UNIQUE(cod_document_type,document);

ALTER TABLE public.users ADD
  CONSTRAINT fk_users_document_t FOREIGN key(cod_document_type)
    REFERENCES public.document_types (id)
    ON DELETE RESTRICT ON UPDATE RESTRICT;
*/
