package models

//Expense lista de gastos
type Expense struct {
	ModelBig
	Cash              float32        `json:"money,omitempty"    gorm:"type:numeric(6,1);NOT NULL"`
	CodExpenseDescrip uint32         `json:"id_expense,omit"    gorm:"type:BIGINT;NOT NULL"`
	CodUser           uint32         `json:"id_user,omit"       gorm:"type:integer;NOT NULL"`
	CodCollection     uint32         `json:"id_collection,omit" gorm:"type:integer;NOT NULL"`
	ExpenseDescrip    ExpenseDescrip `json:"description,omitempty"`
}

/*
DROP TABLE public.expenses;

CREATE TABLE public.expenses
(
  id bigserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp NOT NULL,
  cash numeric(6,1) NOT NULL,
  cod_expense_descrip BIGINT NOT NULL,
  cod_user integer NOT NULL,
  cod_collection integer NOT NULL,

  CONSTRAINT pk_expense PRIMARY KEY(id)
);

ALTER TABLE public.expenses ADD
  CONSTRAINT fk_expense_expensed FOREIGN KEY(cod_expense_descrip)
    REFERENCES public.expense_descrips(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT ;

ALTER TABLE public.expenses ADD
  CONSTRAINT fk_expense_users FOREIGN KEY(cod_user)
    REFERENCES public.users(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT ;

ALTER TABLE public.expenses ADD
  CONSTRAINT fk_expense_collections FOREIGN KEY(cod_collection)
    REFERENCES public.collections(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.expenses ADD
  CONSTRAINT ck_expense_cash CHECK(cash > 0);
*/
