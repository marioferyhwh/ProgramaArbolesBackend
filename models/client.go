package models

/*
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

  CONSTRAINT pk_clients PRIMARY KEY(id),

  CONSTRAINT uk_clients_cdocumentt_document UNIQUE(cod_document_type,document),
  CONSTRAINT fk_clients_document_t FOREIGN key(cod_document_type)
    REFERENCES public.documenttypes (id)
    ON DELETE RESTRICT ON UPDATE RESTRICT,

  CONSTRAINT fk_clients_collectiones FOREIGN KEY(cod_collection)
    REFERENCES public.collectiones(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT,

  CONSTRAINT fk_clients_loan_s FOREIGN KEY(cod_loan_state)
    REFERENCES public.loanstates(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT,

  CONSTRAINT fk_clients_business_t FOREIGN KEY(cod_business_type)
    REFERENCES public.businesstypes(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT,

  CONSTRAINT fk_clients_list_l FOREIGN KEY(cod_list_location)
    REFERENCES public.listlocationes(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT,

  CONSTRAINT fk_clients_users FOREIGN KEY(cod_user)
    REFERENCES public.users(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT ck_clients_loan_number CHECK(loan_number >= 0)

);
*/
