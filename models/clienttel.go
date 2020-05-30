package models

//ClientTel numero telefonico del cliente
type ClientTel struct {
	ModelBig
}

/*
DROP TABLE public.clienttels;

CREATE TABLE public.clienttels
(
  id bigserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  cod_client integer NOT NULL,
  phone NUMERIC(12) NOT NULL,
  cod_tel_descrip SMALLINT NOT NULL,
  CONSTRAINT pk_client_t PRIMARY KEY(id),

  CONSTRAINT fk_client_t_clients FOREIGN KEY(cod_client)
    REFERENCES public.clients(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT ,

  CONSTRAINT fk_client_t_tel_d FOREIGN KEY(cod_tel_descrip)
    REFERENCES public.teldescrips(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT ,
  CONSTRAINT ck_client_t_phone CHECK(phone > 999999)
);
*/
