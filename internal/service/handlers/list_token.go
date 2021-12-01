package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/subgroup/tokenproject/internal/data"
	"gitlab.com/tokend/subgroup/tokenproject/internal/service/requests"
	"gitlab.com/tokend/subgroup/tokenproject/resources"
	"net/http"
	"strconv"
)

func List(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetPersonListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	tokens, err := Token(r).Select(req.OffsetPageParams)
	if err != nil {
		Log(r).WithError(err).Error("error selecting tokens")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := resources.TokenListResponse{
		Data: newPeopleList(tokens),
	}

	ape.Render(w, response)
}

func newPeopleList(tokens []data.Token) []resources.Token {
	result := make([]resources.Token, len(tokens))
	for i, token := range tokens {
		result[i] = newTokenModel(token)
	}
	return result
}

func newTokenModel(token data.Token) resources.Token {
	return resources.Token{
		Key: resources.Key{
			ID:   strconv.FormatInt(token.Id, 10),
			Type: resources.TOKEN,
		},
		Attributes: resources.TokenAttributes{
			Addresslp: token.Addresslp,
			Vault: token.Vault,
			Asset: token.Asset,
		},
	}
}
