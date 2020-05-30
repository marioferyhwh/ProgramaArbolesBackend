package models

// ListLocation ubicaion del cliente
type ListLocation struct {
	ModelBig
	CodCollection uint   `json:"id_collection,omit" gorm:"type:integer;not null"`
	Descrip       string `json:"name,omitempty" gorm:"type:varchar(11);not null;default:''"`
}

/*
DROP TABLE public.list_locations;

CREATE TABLE public.list_locations
(
  id bigserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  cod_collection integer NOT NULL,
  descrip varchar(11) NOT NULL DEFAULT '',

  CONSTRAINT pk_list_l PRIMARY KEY(id)
);


ALTER TABLE public.list_locations ADD
  CONSTRAINT fk_list_l_collections FOREIGN KEY(cod_collection)
    REFERENCES public.collections(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.list_locations ADD
  CONSTRAINT uk_list_l_ccollection_descrip UNIQUE(cod_collection,descrip);
*/
