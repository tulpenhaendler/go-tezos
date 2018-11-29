package main

import (
	"github.com/DefinitelyNotAGoat/goTezos/goTezos"
	"fmt"
)

func main(){
	gt :=  goTezos.NewGoTezos()
	client3 := goTezos.NewTezosRPCClient("localhost",":8732")
	gt.AddNewClient(client3)
	block, err := gt.GetHighestBlock()
	if err != nil {
		return
	}
	fmt.Println(block.Hash)
}