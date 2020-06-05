package models

//Loan prestamos
type Loan struct {
	ModelBig
	InitialValue  float32 `json:"initial_value,omitempty"       gorm:"type:numeric(6,1);NOT NULL"`
	Interest      uint8   `json:"interest,omitempty"            gorm:"type:numeric(2);DEFAULT 20 ;NOT NULL"`
	Quota         uint8   `json:"quota,omitempty"               gorm:"type:numeric(2);NOT NULL"`
	Balance       float32 `json:"balance,omitempty"             gorm:"type:numeric(6,1);NOT NULL"`
	CodLoanState  uint8   `json:"id_loan_state,omitempty"       gorm:"type:SMALLINT;NOT NULL"`
	CodClient     uint64  `json:"id_client,omitempty"           gorm:"type:BIGINT;NOT NULL"`
	CodCollection uint32  `json:"id_collection,omitempty"       gorm:"type:integer;NOT NULL"`
	CodUser       uint32  `json:"id_user,omitempty"             gorm:"type:integer;NOT NULL"`

	LoanState   LoanState     `json:"loan_state,omitempty"  gorm:"-"`
	LoanPayment []LoanPayment `json:"payments,omitempty"    gorm:"-"`
	Client      Client        `json:"client,omitempty"      gorm:"-"`
	Collection  Collection    `json:"collection,omitempty"  gorm:"-"`
	User        User          `json:"user,omitempty"        gorm:"-"`
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
