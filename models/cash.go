/*
DROP TABLE public.cashs;

CREATE TABLE public.cashs
(
  id bigserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp NOT NULL,
  cod_collection integer NOT NULL,
  cod_user integer NOT NULL,
  cash numeric(6,1) NOT NULL,
  
  CONSTRAINT pk_cashs PRIMARY KEY(id),

  CONSTRAINT fk_cashs_collectiones FOREIGN KEY(cod_collection) 
    REFERENCES public.collectiones(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT ,

  CONSTRAINT fk_cashs_users FOREIGN KEY(cod_user) 
    REFERENCES public.users(id)
    ON UPDATE RESTRICT ON DELETE RESTRICT ,

  CONSTRAINT ck_cashs_cash CHECK(cash != 0)

);
*/