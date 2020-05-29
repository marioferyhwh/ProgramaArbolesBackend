package models

//LoanState Estado del prestamo
type LoanState struct {
	Modelsmall
	State string `gorm:"type:varchar(20);NOT NULL;DEFAULT:'';unique"`
}

/*
DROP TABLE public.loanstates;

CREATE TABLE public.loanstates
(
  id smallserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  state varchar(20) NOT NULL DEFAULT '',

  CONSTRAINT pk_loan_s PRIMARY KEY(id),
  CONSTRAINT uk_loan_s_state UNIQUE(state)
);

*/
