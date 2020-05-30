package models

import "github.com/jinzhu/gorm"

//ListUser lista de usarios
type ListUser struct {
	gorm.Model
	ModelBig
	Active        bool  `gorm:"type:bool;NOT NULL;DEFAULT:true"`
	CodUser       uint  `gorm:"type:integer;not null"`
	CodCollection uint  `gorm:"type:integer;not null"`
	CodUserLevel  uint8 `gorm:"type:smallint;not null;default:1"`
	Cash          int   `gorm:"type:numeric(6,1);not null;default: 0"`

	UserLevel UserLevel
}

/*
DROP TABLE public.list_users;

CREATE TABLE public.list_users
(
  id bigserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  active bool DEFAULT TRUE,
  cod_user integer NOT NULL,
  cod_collection integer NOT NULL,
  cod_user_level SMALLINT NOT NULL DEFAULT 1,
  cash NUMERIC(6,1) NOT NULL DEFAULT 0,

  CONSTRAINT pk_list_u PRIMARY KEY(id)
);

ALTER TABLE public.list_users ADD
  CONSTRAINT fk_list_u_users FOREIGN KEY(cod_user)
    REFERENCES public.users(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.list_users ADD
  CONSTRAINT fk_list_u_collections FOREIGN KEY(cod_collection)
    REFERENCES public.collections(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.list_users ADD
  CONSTRAINT fk_list_u_user_l FOREIGN KEY(cod_user_level)
    REFERENCES public.user_levels(id)
    ON UPDATE CASCADE ON DELETE RESTRICT;

ALTER TABLE public.list_users ADD
  CONSTRAINT uk_list_u_cuser_ccollection UNIQUE(cod_user,cod_collection);
*/
