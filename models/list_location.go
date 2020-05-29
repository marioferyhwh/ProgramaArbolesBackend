package models

/*
DROP TABLE public.listlocationes;

CREATE TABLE public.listlocationes
(
  id bigserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  cod_collection integer NOT NULL,
  descrip varchar(11) NOT NULL DEFAULT '',

  CONSTRAINT pk_list_l PRIMARY KEY(id),
  CONSTRAINT fk_list_l_collectiones FOREIGN KEY(cod_collection)
    REFERENCES public.collectiones(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT ,
  CONSTRAINT uk_list_l_ccollection_descrip UNIQUE(cod_collection,descrip)
);
*/
