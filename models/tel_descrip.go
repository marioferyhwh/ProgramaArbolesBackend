package models

//TelDescrip descripcion telefono
type TelDescrip struct {
	Modelsmall
	Descrip string `gorm:"type:varchar(20);not null;default:'';unique"`
}

/*
DROP TABLE public.teldescrips;

CREATE TABLE public.teldescrips
(
  id smallserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  descrip varchar(20) NOT NULL DEFAULT '',

  CONSTRAINT pk_tel_d PRIMARY KEY(id),
  CONSTRAINT uk_tel_d_descrip UNIQUE(descrip)
);
*/
