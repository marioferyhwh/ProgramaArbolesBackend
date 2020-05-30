package models

//Client  cliente al que se le presta
type Client struct {
	ModelBig
	Name            string `json:"name,omitempty" gorm:"type:varchar(50); NOT NULL"`
	Email           string `json:"email,omitempty" gorm:"type:varchar(100)"`
	CodDocumentType string `json:"document_codige,omitempty" gorm:"type:varchar(3);default:'CC'"`
	Document        uint   `json:"document_number,omitempty" gorm:"type:NUMERIC(11)"`
	Adress          string `json:"adress,omitempty" gorm:"type:varchar(60); NOT NULL"`
	LoanNumber      uint   `json:"number_loans,omitempty" gorm:"type:SMALLINT; NOT NULL;default:0"`
	CodCollection   uint   `json:"id_collection,omit" gorm:"type:integer; NOT NULL"`
	CodLoanState    uint8  `json:"id_loan_state,omit" gorm:"type:SMALLINT; NOT NULL;default:0"`
	CodBusinessType uint   `json:"id_type_business,omit" gorm:"type:SMALLINT; NOT NULL;default:0"`
	CodListLocation uint   `json:"id_location,omit" gorm:"type:BIGINT; NOT NULL;default:0"`
	CodUser         uint   `json:"id_user,omit" gorm:"type:integer; NOT NULL"`

	DocumentType  DocumentType `json:"document_description,omitempty"`
	LoanState     LoanState    `json:"state,omitempty"`
	BusinessTypes BusinessType `json:"bussiness,omitempty"`
	ListLocation  ListLocation `json:"location,omitempty"`
	User          User         `json:"user_create,omitempty"`
	ClientTel     []ClientTel  `json:"tels,omitempty"`
	Loan          []Loan       `json:"loans,omitempty"`
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
