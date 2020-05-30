package models

//UserTel telefono de usuario
type UserTel struct {
	ModelBig
	CodUser       uint16 `json:"id_user,omit"            gorm:"type:integer; NOT NULL"`
	Phone         string `json:"number,omitempty"        gorm:"type:NUMERIC(12); NOT NULL"`
	CodTelDescrip uint8  `json:"id_tel_description,omit" gorm:"type:SMALLINT; NOT NULL;DEFAULT:0"`

	TelDescrip TelDescrip `json:"description,omitempty"`
}

/*
DROP TABLE public.user_tels;

CREATE TABLE public.user_tels
(
  id bigserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  cod_user integer NOT NULL,
  phone NUMERIC(12) NOT NULL,
  cod_tel_descrip SMALLINT NOT NULL DEFAULT 0,
  CONSTRAINT pk_user_t PRIMARY KEY(id)
);

ALTER TABLE public.user_tels ADD
  CONSTRAINT fk_user_t_users FOREIGN KEY(cod_user)
    REFERENCES public.users(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT ;

ALTER TABLE public.user_tels ADD
  CONSTRAINT fk_user_t_tel_d FOREIGN KEY(cod_tel_descrip)
    REFERENCES public.tel_descrips(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT ;

ALTER TABLE public.user_tels ADD
  CONSTRAINT uk_user_t_phone UNIQUE(phone);

ALTER TABLE public.user_tels ADD
  CONSTRAINT ck_user_t_phone CHECK(phone > 999999);
*/
