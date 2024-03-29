package models

//Client  cliente al que se le presta
type Client struct {
	ModelBig
	Name                  string `json:"name,omitempty"             gorm:"type:varchar(50); NOT NULL"`
	Email                 string `json:"email,omitempty"            gorm:"type:varchar(100)"`
	CodDocumentType       string `json:"document_code,omitempty"    gorm:"type:varchar(3);default:'CC'"`
	Document              string `json:"document,omitempty"         gorm:"type:NUMERIC(11)"`
	Adress                string `json:"adress,omitempty"           gorm:"type:varchar(60); NOT NULL"`
	LoanNumber            uint8  `json:"number_loans,omitempty"     gorm:"type:SMALLINT; NOT NULL;default:0"`
	CodCollection         uint32 `json:"id_collection,omitempty"    gorm:"type:integer; NOT NULL"`
	CodLoanState          uint8  `json:"id_loan_state,omitempty"    gorm:"type:SMALLINT; NOT NULL;default:0"`
	CodBusinessType       uint16 `json:"id_type_business,omitempty" gorm:"type:SMALLINT; NOT NULL;default:0"`
	CodClientListLocation uint64 `json:"id_location,omitempty"      gorm:"type:BIGINT; NOT NULL;default:0"`
	CodUser               uint32 `json:"id_user,omitempty"          gorm:"type:integer; NOT NULL"`

	LoanState          LoanState          `json:"state,omitempty"       gorm:"-"`
	BusinessTypes      BusinessType       `json:"bussiness,omitempty"   gorm:"-"`
	ClientListLocation ClientListLocation `json:"location,omitempty"    gorm:"-"`
	User               User               `json:"user_create,omitempty" gorm:"-"`
	ClientTelDelete    []ClientTel        `json:"tels_delete,omitempty" gorm:"-"`
	ClientTelNew       []ClientTel        `json:"tels_new,omitempty"    gorm:"-"`
	ClientTel          []ClientTel        `json:"tels,omitempty"        gorm:"-"`
	Loan               []Loan             `json:"loans,omitempty"       gorm:"-"`
}

/*
-- tabla 11
DROP TABLE public.clients;

CREATE TABLE public.clients
(
  id bigserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  name varchar(50) NOT NULL,
  email varchar(100),
  cod_document_type varchar(3) DEFAULT 'CC',
  document NUMERIC(11),
  adress varchar(60) NOT NULL,
  loan_number SMALLINT NOT NULL DEFAULT 0,
  cod_collection integer NOT NULL,
  cod_loan_state SMALLINT NOT NULL DEFAULT 0,
  cod_business_type SMALLINT NOT NULL DEFAULT 0,
  cod_list_location BIGINT NOT NULL DEFAULT 0,
  cod_user integer NOT NULL,

  CONSTRAINT pk_clients PRIMARY KEY(id)
);

ALTER TABLE public.clients ADD
  CONSTRAINT uk_clients_cdocumentt_document UNIQUE(cod_document_type,document);

ALTER TABLE public.clients ADD
  CONSTRAINT fk_clients_document_t FOREIGN key(cod_document_type)
    REFERENCES public.document_types (id)
    ON DELETE RESTRICT ON UPDATE RESTRICT;

ALTER TABLE public.clients ADD
  CONSTRAINT fk_clients_collections FOREIGN KEY(cod_collection)
    REFERENCES public.collections(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.clients ADD
  CONSTRAINT fk_clients_loan_s FOREIGN KEY(cod_loan_state)
    REFERENCES public.loan_states(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.clients ADD
  CONSTRAINT fk_clients_business_t FOREIGN KEY(cod_business_type)
    REFERENCES public.business_types(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.clients ADD
  CONSTRAINT fk_clients_list_l FOREIGN KEY(cod_list_location)
    REFERENCES public.list_locations(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.clients ADD
  CONSTRAINT fk_clients_users FOREIGN KEY(cod_user)
    REFERENCES public.users(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.clients ADD
  CONSTRAINT ck_clients_loan_number CHECK(loan_number >= 0);

*/
