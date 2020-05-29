package models

//Collection
type Collection struct {
	Model
	Descrip      string `gorm:"type:varchar(256);DEFAULT:'' "`
	Active       bool   `gorm:"type:bool;NOT NULL;DEFAULT:true"`
	BalanceTotal int    `gorm:"type:numeric(7,1);NOT NULL;DEFAULT: 0"`
}

/*
CREATE TABLE public.collections
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
