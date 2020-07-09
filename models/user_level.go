package models

//UserLevel nivel de acceso del usario a la base de datos
type UserLevel struct {
	Modelsmall
	Level string `json:"level,omitempty" gorm:"type:varchar(11);NOT NULL;DEFAULT:''"`
}

/*
DROP TABLE public.user_levels;

CREATE TABLE user_levels
(
  id smallserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  level varchar(11) NOT NULL DEFAULT '',

  CONSTRAINT pk_user_l PRIMARY KEY(id)
);

ALTER TABLE public.user_levels ADD
  CONSTRAINT uk_user_l_level UNIQUE(level);

*/
