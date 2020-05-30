package models

//ClientTel numero telefonico del cliente
type ClientTel struct {
	ModelBig
	CodClient     int  `gorm:"integer"`
	Phone         int  `gorm:"NUMERIC"`
	CodTelDescrip int8 `gorm:"SMALLINT"`
}

/*
-- tabla 12
DROP TABLE public.client_tels;

CREATE TABLE public.client_tels
(
  id bigserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  cod_client integer NOT NULL,
  phone NUMERIC(12) NOT NULL,
  cod_tel_descrip SMALLINT NOT NULL,

  CONSTRAINT pk_client_t PRIMARY KEY(id)
);

ALTER TABLE public.client_tels ADD
  CONSTRAINT fk_client_t_clients FOREIGN KEY(cod_client)
    REFERENCES public.clients(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.client_tels ADD
  CONSTRAINT fk_client_t_tel_d FOREIGN KEY(cod_tel_descrip)
    REFERENCES public.tel_descrips(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.client_tels ADD
  CONSTRAINT ck_client_t_phone CHECK(phone > 999999);
*/
