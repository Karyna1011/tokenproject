package token

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3"
	"io/ioutil"
	"math/big"
)

func GetBody(log *logan.Entry) []string {
	    var	Addresses [] string
		Datan:=Addresses

		dataInBytes, err := ioutil.ReadFile("test.json")
		if err != nil {
			log.WithError(err).Error("failed to read json file")
		}

		err= json.Unmarshal(dataInBytes, &Datan)
		if err!= nil {
			log.WithError(err).Error("failed to unmarshal json file")
		}

	   return Datan
	}


func GetAddresses(Contract *bind.BoundContract, log *logan.Entry, address common.Address)  (string, error) {
	var 	result = make([]interface{}, 0)
	log.Info(address.String())

	err := Contract.Call(&bind.CallOpts{}, &result, "get_registry")

	log.Info("result=",result)

	if err != nil {
		log.WithError(err).Error("error during calling contract")
	}
	return fmt.Sprintf("%v", result[0]),err
}

func GetCount(Contract *bind.BoundContract, log *logan.Entry)  (string, error) {

	var 	countResult = make([]interface{}, 0)

	err := Contract.Call(&bind.CallOpts{}, &countResult, "pool_count")

	log.Info("result=",countResult)

	if err != nil {
		log.WithError(err).Error("error during calling contract function count")
	}
	return fmt.Sprintf("%v", countResult[0]),err
}

func GetPool(Contract *bind.BoundContract, log *logan.Entry, count int)  (string, error) {

	var 	poolResult = make([]interface{}, 0)
    log.Info("count=",count)

	err := Contract.Call(&bind.CallOpts{}, &poolResult, "pool_list", big.NewInt(int64(count)))

	log.Info("result=",poolResult)

	if err != nil {
		log.WithError(err).Error("error during calling contract function pool_list")
	}
	return fmt.Sprintf("%v", poolResult[0]),err
}

func GetCoinsNumber(Contract *bind.BoundContract, log *logan.Entry, address string)  (string, error) {

	var 	poolResult = make([]interface{}, 0)

	err := Contract.Call(&bind.CallOpts{}, &poolResult, "get_n_coins", common.HexToAddress(address))

	log.Info("result=",poolResult[0])

	if err != nil {
		log.WithError(err).Error("error during calling contract function get_n_coins")
	}
	return fmt.Sprintf("%v", poolResult[0]),err
}

func GetUnderlyingCoins(Contract *bind.BoundContract, log *logan.Entry, address string)  (string, error) {

	var 	poolResult = make([]interface{}, 0)

	err := Contract.Call(&bind.CallOpts{}, &poolResult, "get_underlying_coins", common.HexToAddress(address))

	log.Info("result=",poolResult[0])

	if err != nil {
		log.WithError(err).Error("error during calling contract function get_coins")
	}
	return fmt.Sprintf("%v", poolResult[0]),err
}

func GetLpAddress(Contract *bind.BoundContract, log *logan.Entry, address string)  (string, error) {

	var 	poolResult = make([]interface{}, 0)

	err := Contract.Call(&bind.CallOpts{}, &poolResult, "get_lp_token", common.HexToAddress(address))

	log.Info("result=",poolResult)

	if err != nil {
		log.WithError(err).Error("error during calling contract function get_lp_token")
		return "",err
	}
	return fmt.Sprintf("%v", poolResult[0]),err
}


func GetLatestVault(Contract *bind.BoundContract, log *logan.Entry, address string)  (string, error) {

	var 	poolResult = make([]interface{}, 0)

	err := Contract.Call(&bind.CallOpts{}, &poolResult, "latestVault", common.HexToAddress(address))

	log.Info("result=",poolResult)

	if err != nil {
		log.WithError(err).Error("error during calling contract function latestVault")
		return "",err
	}
	return fmt.Sprintf("%v", poolResult[0]), err
}
