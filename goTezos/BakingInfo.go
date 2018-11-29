package goTezos

import (
	"github.com/DefinitelyNotAGoat/goTezos/models"
	"strconv"
	"fmt"
	"encoding/json"
)

func (this *GoTezos) GetAllEndorsingRightsForCycle(cycle int) (models.CycleEndorseRights, error) {
	resp, err := this.GetResponse("/chains/main/blocks/head/helpers/endorsing_rights?cycle=" + strconv.Itoa(cycle), "")

	if err != nil {
		fmt.Println("Could not get Response from Tezos Node, Node down?", )
		return models.CycleEndorseRights{}, err
	}

	ret := models.CycleEndorseRights{}
	json.Unmarshal(resp.Bytes,&ret)
	return ret, nil
}