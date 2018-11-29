package goTezos

import (
	"github.com/DefinitelyNotAGoat/goTezos/models"
	"encoding/json"
)



func (this *GoTezos) GetContractDetails(contract string) (models.ContractDetails, error) {
	retstruct := models.ContractDetails{}
	respb, _ := this.GetResponse("/chains/main/blocks/head/context/contracts/"+contract, "")
	err := json.Unmarshal(respb.Bytes, &retstruct)
	return retstruct, err
}

func (this *GoTezos) GetContractDetailsAtBlock(contract string, blockhash string) (models.ContractDetails, error) {
	retstruct := models.ContractDetails{}
	respb, _ := this.GetResponse("/chains/main/blocks/"+blockhash+"/context/contracts/"+contract, "")
	err := json.Unmarshal(respb.Bytes, &retstruct)
	return retstruct, err
}


func (this *GoTezos) GetDelegateDetails(contract string) (models.DelegateDetails, error) {
	retstruct := models.DelegateDetails{}
	respb, _ := this.GetResponse("/chains/main/blocks/head/context/delegates/"+contract, "")
	err := json.Unmarshal(respb.Bytes, &retstruct)
	return retstruct, err
}