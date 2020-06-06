package models

//DocumentType tipo de documento de identificacion
type DocumentType struct {
	Modelsmall
	NameShort string `json:"name_short,omitempty" gorm:"type:varchar(3);DEFAULT:'CC';NOT NULL"`
	Descrip   string `json:"description,omitempty" gorm:"type:varchar(30);DEFAULT:'';NOT NULL"`
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
