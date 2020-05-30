package models

//ExpenseDescrip descripcion de gastos
type ExpenseDescrip struct {
	ModelBig
	CodCollection int32  `json:"id_collection,omit"     gorm:"type:integer;NOT NULL"`
	Descrip       string `json:"description,omitempty"  gorm:"type:varchar(11);NOT NULL;DEFAULT:''"`
}

/*
DROP TABLE public.expense_descrips;

CREATE TABLE public.expense_descrips
(
  id bigserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  cod_collection integer NOT NULL,
  descrip varchar(11) NOT NULL DEFAULT '',

  CONSTRAINT pk_expense_d PRIMARY KEY(id)
);

ALTER TABLE public.expense_descrips ADD
  CONSTRAINT fk_expense_d_collections FOREIGN KEY(cod_collection)
    REFERENCES public.collections(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT ;

ALTER TABLE public.expense_descrips ADD
  CONSTRAINT uk_expense_d_ccollection_descrip UNIQUE(cod_collection,descrip);
*/
