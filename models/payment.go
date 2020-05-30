package models

//Payment pagos de prestamo
type Payment struct {
	ModelBig
}

/*
DROP TABLE public.payments;

CREATE TABLE public.payments
(
  id bigserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp NOT NULL,
  cod_loan BIGINT NOT NULL,
  cash numeric(6,1) NOT NULL,
  cod_user integer NOT NULL,
  cod_collection integer NOT NULL,

  CONSTRAINT pk_payments PRIMARY KEY(id),

  CONSTRAINT fk_payments_loans FOREIGN KEY(cod_loan)
    REFERENCES public.loans(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT ,

  CONSTRAINT fk_payments_users FOREIGN KEY(cod_user)
    REFERENCES public.users(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT ,

  CONSTRAINT fk_payments_collectiones FOREIGN KEY(cod_collection)
    REFERENCES public.collectiones(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT ,

  CONSTRAINT ck_payments_cash CHECK(cash > 0)

);

*/
