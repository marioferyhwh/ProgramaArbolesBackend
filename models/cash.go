package models

//Cash lista de egresos e ingresoso dee cobro
type Cash struct {
	ModelBig
	CodCollection uint32  `json:"id_collection,omitempty" gorm:"type:integer;NOT NULL"`
	CodUser       uint32  `json:"id_user,omitempty"       gorm:"type:integer;NOT NULL"`
	Cash          float32 `json:"money,omitempty"         gorm:"type:numeric(6,1);NOT NULL"`
}

/*
-- tabla 15
DROP TABLE public.cashes;

CREATE TABLE public.cashes
(
  id bigserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp NOT NULL,
  cod_collection integer NOT NULL,
  cod_user integer NOT NULL,
  cash numeric(6,1) NOT NULL,

  CONSTRAINT pk_cashes PRIMARY KEY(id)
);

ALTER TABLE public.cashes ADD
  CONSTRAINT fk_cashes_collections FOREIGN KEY(cod_collection)
    REFERENCES public.collections(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.cashes ADD
  CONSTRAINT fk_cashes_users FOREIGN KEY(cod_user)
    REFERENCES public.users(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT;

ALTER TABLE public.cashes ADD
  CONSTRAINT ck_cashes_cash CHECK(cash != 0);
*/
