package models

import "github.com/jinzhu/gorm"

//Collection
type Collection struct {
	gorm.Model
	Descrip      string
	Active       bool
	BalanceTotal int
}

/*
CREATE TABLE public.collectiones
(
  id serial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  descrip varchar(256) DEFAULT '',
  active bool NOT NULL DEFAULT TRUE,
  balance_total numeric(7,1) NOT NULL DEFAULT 0,

  CONSTRAINT pk_collectiones PRIMARY KEY(id)
);
*/
