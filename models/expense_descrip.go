package models

/*
DROP TABLE public.expensedescrips;

CREATE TABLE public.expensedescrips
(
  id bigserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  cod_collection integer NOT NULL,
  descrip varchar(11) NOT NULL DEFAULT '',

  CONSTRAINT pk_expense_d PRIMARY KEY(id),
  CONSTRAINT fk_expense_d_collectiones FOREIGN KEY(cod_collection)
    REFERENCES public.collectiones(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT ,
  CONSTRAINT uk_expense_d_ccollection_descrip UNIQUE(cod_collection,descrip)
);
*/
