package entity

import (
	"database/sql"
)

type UserEntity struct {
	ID sql.NullInt64 `db:"id;primaryKey;autoIncrement"`
}

type AccountEntity struct {
	UserID  sql.NullInt64 `db:"user_id;foreignKey:user(id)"` // 1 account per user
	Balance sql.NullInt64 `db:"balance;int(11)"`
	Limit   sql.NullInt64 `db:"limit;int(11)"`
}

type TransactionEntity struct {
	UserID      sql.NullInt64  `db:"user_id;int(11);foreignKey:user(id)"`
	Value       sql.NullInt64  `db:"value;int(11)"`
	Type        sql.NullString `db:"type;varchar(1)"`
	Description sql.NullString `db:"description;varchar(10)"`
	RealizedIn  sql.NullTime   `db:"realized_in;datetime"`
}
