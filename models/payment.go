package models

//Payment pagos de prestamo
type Payment struct {
	ModelBig
	CodLoan       uint64  `json:"id_loan,omitempty"       gorm:"type:BIGINT ;NOT NULL"`
	Cash          float32 `json:"money,omitempty"         gorm:"type:numeric(6,1) ;NOT NULL"`
	CodUser       uint32  `json:"id_user,omitempty"       gorm:"type:integer ;NOT NULL"`
	CodCollection uint32  `json:"id_collection,omitempty" gorm:"type:integer ;NOT ;NULL"`

	User User `json:"user_create,omitempty"`
}

/*
DROP TABLE public.payments;F

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

  CONSTRAINT pk_payments PRIMARY KEY(id)
);

ALTER TABLE public.payments ADD
  CONSTRAINT fk_payments_loans FOREIGN KEY(cod_loan)
    REFERENCES public.loans(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.payments ADD
  CONSTRAINT fk_payments_users FOREIGN KEY(cod_user)
    REFERENCES public.users(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE public.payments ADD

  CONSTRAINT fk_payments_collections FOREIGN KEY(cod_collection)
    REFERENCES public.collections(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.payments ADD
  CONSTRAINT ck_payments_cash CHECK(cash > 0);
*/
