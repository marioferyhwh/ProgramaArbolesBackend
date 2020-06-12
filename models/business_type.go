package models

// BusinessType tipo de negocio que tien el cliente
type BusinessType struct {
	Model
	Descrip string `json:"type_business,omitempty" gorm:"type:varchar(45);not null;default:''"`
}

/*
-- tabla 7
DROP TABLE public.business_types;

CREATE TABLE public.business_types
(
  id smallserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  descrip varchar(45) NOT NULL DEFAULT '',

  CONSTRAINT pk_business_t PRIMARY KEY(id)
);

ALTER TABLE public.business_types ADD
  CONSTRAINT uk_business_t_descrip UNIQUE(descrip);


*/
