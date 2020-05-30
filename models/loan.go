package models

//Loan prestamos
type Loan struct {
	ModelBig
	InitialValue  int `gorm:"type:numeric(6,1);NOT NULL"`
	Interest      int `gorm:"type:numeric(2);DEFAULT 20 ;NOT NULL"`
	Quota         int `gorm:"type:numeric(2);NOT NULL"`
	Balance       int `gorm:"type:numeric(6,1);NOT NULL"`
	CodLoanState  int `gorm:"type:SMALLINT;NOT NULL"`
	CodClient     int `gorm:"type:BIGINT;NOT NULL"`
	CodCollection int `gorm:"type:integer;NOT NULL"`
	CodUser       int `gorm:"type:integer;NOT NULL"`
	LoanState     LoanState
	Payments      []Payment
	Client        Client
	Collection    Collection
	User          User
}

/*
DROP TABLE public.loans;

CREATE TABLE public.loans
(
  id bigserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp NOT NULL,
  initial_value numeric(6,1) NOT NULL,
  interest numeric(2) NOT NULL DEFAULT 20,
  quota numeric(2) NOT NULL,
  balance numeric(6,1) NOT NULL,
  cod_loan_state SMALLINT NOT NULL,
  cod_client BIGINT NOT NULL,
  cod_collection integer NOT NULL,
  cod_user integer NOT NULL,

  CONSTRAINT pk_loans PRIMARY KEY(id)
);

ALTER TABLE public.loans ADD
  CONSTRAINT fk_loans_loan_s FOREIGN KEY(cod_loan_state)
    REFERENCES public.loan_states(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.loans ADD
  CONSTRAINT fk_loans_clients FOREIGN KEY(cod_client)
    REFERENCES public.clients(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.loans ADD
  CONSTRAINT fk_loans_collections FOREIGN KEY(cod_collection)
    REFERENCES public.collections(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.loans ADD
  CONSTRAINT fk_loans_users FOREIGN KEY(cod_user)
    REFERENCES public.users(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.loans ADD
  CONSTRAINT ck_loans_initialv CHECK( initial_value > 0 AND initial_value%5 = 0);

ALTER TABLE public.loans ADD
  CONSTRAINT ck_loans_interest CHECK( interest > 0);

ALTER TABLE public.loans ADD
  CONSTRAINT ck_loans_quota CHECK( quota > 0);

ALTER TABLE public.loans ADD
  CONSTRAINT ck_loans_balance CHECK( balance >= 0);
*/
