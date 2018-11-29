package goTezos

import (
	"github.com/DefinitelyNotAGoat/goTezos/models"
	"encoding/json"
)


func (this *GoTezos) GetBlockChainHeads() ([][]string,error) {
	blocks := [][]string{}
	bres, err := this.GetResponse("/chains/main/blocks", "")
	if err == nil {
		json.Unmarshal(bres.Bytes, &blocks)
	}
	return blocks,err
}


func (this *GoTezos) GetHighestBlock() (models.Block, error) {
	block := models.Block{}
	blocks,err := this.GetBlockChainHeads()
	if err == nil {
		bhash := blocks[0][0]
		bres, err := this.GetResponse("/chains/main/blocks/"+bhash, "")
		json.Unmarshal(bres.Bytes, &block)
		if err != nil {
			return block, err
		}
		return block ,nil
	}
	return block,err
}


func (this *GoTezos) GetBlockByHash(blockid string) (models.Block, error) {
	bres, err := this.GetResponse("/chains/main/blocks/"+blockid, "")
	block := models.Block{}
	json.Unmarshal(bres.Bytes, &block)
	if err != nil {
		return models.Block{}, nil
	}
	return block, nil
}

