package models

//DocumentType tipo de documento de identificacion
type DocumentType struct {
	id string `gorm:"type:varchar(3);DEFAULT:'CC';NOT NULL;primary_key"`
	TimeModel
	descrip string `gorm:"type:varchar(20);DEFAULT:'':NOT NULL;unique"`
}

/*
DROP TABLE public.documenttypes;

CREATE TABLE public.documenttypes
(
  id varchar(3) NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  descrip varchar(20) NOT NULL DEFAULT '',

  CONSTRAINT pk_document_t PRIMARY KEY(id),
  CONSTRAINT uk_document_t_descrip UNIQUE(descrip)
);
*/
