package models

//LoanState Estado del prestamo
type LoanState struct {
	Modelsmall
	State string `json:"state,omitempty" gorm:"type:varchar(20);NOT NULL;DEFAULT:''"`
}

/*
DROP TABLE public.loan_states;

CREATE TABLE public.loan_states
(
  id smallserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  state varchar(20) NOT NULL DEFAULT '',

  CONSTRAINT pk_loan_s PRIMARY KEY(id)
);

ALTER TABLE public.loan_states ADD
  CONSTRAINT uk_loan_s_state UNIQUE(state);
*/
