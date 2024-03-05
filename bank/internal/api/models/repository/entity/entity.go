package entity

import (
	"database/sql"
)

type ClientEntity struct {
	ID   sql.NullInt64  `db:"id;primaryKey;autoIncrement"`
	Nome sql.NullString `db:"nome;varchar(50)"`
}

type AccountEntity struct {
	ID        sql.NullInt64 `db:"id;primaryKey;autoIncrement"`
	ClienteID sql.NullInt64 `db:"cliente_id;foreignKey:user(id)"`
	Saldo     sql.NullInt64 `db:"saldo;int(11)"`
	Limite    sql.NullInt64 `db:"limite;int(11)"`
}

type TransactionEntity struct {
	ID          sql.NullInt64  `db:"id;primaryKey;autoIncrement"`
	ClienteID   sql.NullInt64  `db:"cliente_id;int(11);foreignKey:user(id)"`
	Valor       sql.NullInt64  `db:"valor;int(11)"`
	Tipo        sql.NullString `db:"tipo;enum('c','d')"`
	Descricao   sql.NullString `db:"descricao;varchar(10)"`
	RealizadaEm sql.NullTime   `db:"realizada_em;datetime"`
}
