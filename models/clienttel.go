package models

//ClientTel numero telefonico del cliente
type ClientTel struct {
	ModelBig
	CodClient     uint32 `json:"id_cliente,omit"      gorm:"type:integer"`
	Phone         string `json:"number,omitempty"     gorm:"type:NUMERIC(12)"`
	CodTelDescrip int8   `json:"id_tel_descript,omit" gorm:"type:SMALLINT"`

	TelDescrip TelDescrip `json:"type_tel,omitempty"`
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
