package models

//User usuario
type User struct {
	Model
	Active          bool   `json:"active,omit" gorm:"type:bool;not null;default:true"`
	NickName        string `json:"nick_name,omitempty" gorm:"type:varchar(50);not null;default:''"`
	Email           string `json:"email,omitempty" gorm:"type:varchar(100);not null"`
	Password        string `json:"password,omit" gorm:"type:varchar(256);not null"`
	CodDocumentType string `json:"document_codige,omitempty" gorm:"type:varchar(3);not null;default:''"`
	Document        int    `json:"document_number,omitempty" gorm:"type:numeric(11);not null"`
	Name            string `json:"name,omitempty" gorm:"type:varchar(50);not null;default:''"`

	DocumentType DocumentType `json:"document_description,omitempty"`
	ListUser     []ListUser   `json:"collections,omitempty"`
	UserTel      []UserTel    `json:"tels,omitempty"`
	Expense      []Expense    `json:"expenses,omitempty"`
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
