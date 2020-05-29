/*

DROP TABLE public.userlevels;

CREATE TABLE userlevels
(
  id smallserial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  level varchar(11) NOT NULL DEFAULT '',

  CONSTRAINT pk_user_l PRIMARY KEY(id),
  CONSTRAINT uk_user_l_level UNIQUE(level)
);
*/