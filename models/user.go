package models

//User usuario
type User struct {
	Model
	Active          bool   `json:"active,omitempty"           gorm:"type:bool;not null;default:true"`
	NickName        string `json:"nick_name,omitempty"        gorm:"type:varchar(50);not null;default:''"`
	Email           string `json:"email,omitempty"            gorm:"type:varchar(100);not null"`
	Password        string `json:"password,omitempty"         gorm:"type:varchar(256);not null"`
	CodDocumentType string `json:"document_code,omitempty"  gorm:"type:varchar(3);not null;default:''"`
	Document        string `json:"document,omitempty"  gorm:"type:numeric(11);not null"`
	Name            string `json:"name,omitempty"             gorm:"type:varchar(50);not null;default:''"`
	ConfirmPassword string `json:"confirm_password,omitempty" gorm:"-"`

	DocumentType interface{} `json:"document_description,omitempty" gorm:"-"`
	ListUser     []ListUser  `json:"collections,omitempty"          gorm:"-"`
	UserTel      []UserTel   `json:"tels,omitempty"                 gorm:"-"`
	Expense      []Expense   `json:"expenses,omitempty"             gorm:"-"`

	GetDocumentType bool `json:"-" gorm:"-"`
	GetListUser     bool `json:"-" gorm:"-"`
	GetUserTel      bool `json:"-" gorm:"-"`
	GetExpense      bool `json:"-" gorm:"-"`
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
