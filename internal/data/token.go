package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/subgroup/tokenproject/resources"
)

type TokenQ interface {
	New() TokenQ
	Get() (*Token, error)
	Select(query pgdb.OffsetPageParams) ([]Token, error)
	Insert(data Token) (Token, error)
	Update(data Token) (Token, error)
	FilterById(data int64) TokenQ
	FilterByAddress(name string, address string) TokenQ
	Delete(id int64) error
	DeleteByAddresses(name string, address string) error
}

type Token struct {
	Id      int64  `db:"id"        structs:"-"`
	Asset      string  `db:"asset"        structs:"asset"`
	Addresslp string `db:"addresslp"      structs:"addresslp"`
	Vault string `db:"vault"      structs:"vault"`
}

func (p *Token) Resource() resources.TokenResponse {
	return resources.TokenResponse{
		Data: resources.Token{
			Attributes: resources.TokenAttributes{
				Asset: p.Asset,
				Addresslp: p.Addresslp,
				Vault: p.Vault,
			},
			Key: resources.NewKeyInt64(p.Id, resources.TOKEN),
		},
	}
}
