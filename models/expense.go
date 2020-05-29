/*
DROP TABLE public.expense;

CREATE TABLE public.expense
(
  id bigserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp NOT NULL,
  cash numeric(6,1) NOT NULL,
  cod_expense_descrip BIGINT NOT NULL,
  cod_user integer NOT NULL,
  cod_collection integer NOT NULL,
  CONSTRAINT pk_expense PRIMARY KEY(id),

  CONSTRAINT fk_expense_expensed FOREIGN KEY(cod_expense_descrip) 
    REFERENCES public.expensedescrips(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT ,

  CONSTRAINT fk_expense_users FOREIGN KEY(cod_user) 
    REFERENCES public.users(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT ,

  CONSTRAINT fk_expense_collectiones FOREIGN KEY(cod_collection) 
    REFERENCES public.collectiones(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT,

  CONSTRAINT ck_expense_cash CHECK(cash > 0)

);
*/