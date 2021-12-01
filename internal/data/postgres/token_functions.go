package postgres

import (
	"database/sql"
	"github.com/fatih/structs"
	"gitlab.com/tokend/subgroup/tokenproject/internal/data"

	"github.com/Masterminds/squirrel"

	"gitlab.com/distributed_lab/kit/pgdb"
)

const tableToken = "token"

type tokenQ struct {
	db  *pgdb.DB
	sql squirrel.SelectBuilder
}

func NewTokenQ(db *pgdb.DB) data.TokenQ {
	return &tokenQ{
		db:  db.Clone(),
		sql: squirrel.Select("*").From(tableToken),
	}
}

func (d *tokenQ) New() data.TokenQ {
	return NewTokenQ(d.db)
}

func (d *tokenQ) Get() (*data.Token, error) {
	var result data.Token

	err := d.db.Get(&result, d.sql)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (d *tokenQ) Select(query pgdb.OffsetPageParams) ([]data.Token, error) {
	var result []data.Token

	err := d.db.Select(&result, query.ApplyTo(d.sql, "id"))
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return result, err
}

func (d *tokenQ) Insert(token data.Token) (data.Token, error) {
	clauses := structs.Map(token)

	query := squirrel.Insert(tableToken).SetMap(clauses).Suffix("returning *")

	err := d.db.Get(&token, query)
	if err != nil {
		return data.Token{}, err
	}

	return token, err

}

func (d *tokenQ) Update(token data.Token) (data.Token, error) {
	clauses := structs.Map(token)

	query := squirrel.Update(tableToken).Where(squirrel.Eq{"id": token.Id}).SetMap(clauses).Suffix("returning *")

	err := d.db.Get(&token, query)
	if err != nil {
		return data.Token{}, err
	}

	return token, err
}

func (d tokenQ) FilterById(Id int64) data.TokenQ {
	d.sql = d.sql.Where(squirrel.Eq{"id": Id})

	return &d
}

func (d tokenQ) FilterByAddress(name string, address string) data.TokenQ {
	d.sql = d.sql.Where(squirrel.Eq{"asset":name,"addresslp": address})

	return &d
}

func (d *tokenQ) Delete(id int64) error {
	query := squirrel.Delete(tableToken).Where(squirrel.Eq{"id": id})
	err := d.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (d *tokenQ) DeleteByAddresses(name string, address string) error {
	query := squirrel.Delete(tableToken).Where(squirrel.Eq{"asset": name, "addresslp" : address})
	err := d.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
