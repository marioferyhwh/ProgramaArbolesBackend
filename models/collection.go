package models

//Collection prestamos
type Collection struct {
	Model
	Descrip      string  `json:"description,omitempty" gorm:"type:varchar(256);DEFAULT:'' "`
	Actived      bool    `json:"actived,omitempty" gorm:"type:bool;NOT NULL;DEFAULT:true"`
	BalanceTotal float32 `json:"balance_total,omitempty" gorm:"type:numeric(7,1);NOT NULL;DEFAULT: 0"`

	CollectionCash []CollectionCash `json:"cashes,omitempty"    gorm:"-"`
	Expense        []Expense        `json:"expenses,omitempty"  gorm:"-"`
	Client         []Client         `json:"clients,omitempty"   gorm:"-"`
	UserCollection []UserCollection `json:"users,omitempty"     gorm:"-"`
	CodUser        uint32           `json:"-" gorm:"-"`
}

/*
DROP TABLE public.collections;

CREATE TABLE public.collections
(
  id serial NOT NULL ,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp,
  delete_at timestamp,
  descrip varchar(256) DEFAULT '',
  active bool NOT NULL DEFAULT TRUE,
  balance_total numeric(7,1) NOT NULL DEFAULT 0,

  CONSTRAINT pk_collections PRIMARY KEY(id)
);
*/
