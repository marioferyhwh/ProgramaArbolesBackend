package models

// BusinessTypes tipo de negocio que tien el cliente
type BusinessTypes struct {
	Modelsmall
	Descrip string `gorm:"type:varchar(45);not null;default:'';unique"`
}

/*
DROP TABLE public.businesstypes;

CREATE TABLE public.businesstypes
(
  id smallserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  descrip varchar(45) NOT NULL DEFAULT '',

  CONSTRAINT pk_business_t PRIMARY KEY(id),
  CONSTRAINT uk_business_t_descrip UNIQUE(descrip)
);
*/
