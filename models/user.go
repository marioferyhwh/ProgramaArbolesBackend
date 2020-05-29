package models

/*
DROP TABLE public.users;

CREATE TABLE public.users
(
  id serial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  active bool DEFAULT TRUE,
  nickname varchar(50) NOT NULL DEFAULT '',
  email varchar(100) NOT NULL,
  password varchar(256) NOT NULL,
  cod_document_type varchar(3) NOT NULL DEFAULT 'CC',
  document NUMERIC(11) NOT NULL ,
  name varchar(50) NOT NULL DEFAULT '',

  CONSTRAINT pk_users PRIMARY KEY(id),
  CONSTRAINT uk_users_nickname UNIQUE(nickname),
  CONSTRAINT uk_users_email UNIQUE(email),
  CONSTRAINT uk_users_cdocumentt_document UNIQUE(cod_document_type,document),
  CONSTRAINT fk_users_document_t FOREIGN key(cod_document_type)
    REFERENCES public.documenttypes (id)
    ON DELETE RESTRICT ON UPDATE RESTRICT
);
*/
