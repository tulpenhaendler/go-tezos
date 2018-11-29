package main

import (
	"github.com/DefinitelyNotAGoat/goTezos/goTezos"
	"fmt"
	"github.com/DefinitelyNotAGoat/goTezos/models"
)

func main(){
	gt :=  goTezos.NewGoTezos()
	client := goTezos.NewTezosRPCClient("localhost",":8732")
	gt.AddNewClient(client)

	block, err := gt.GetHighestBlock()
	if err != nil {
		return
	}
	fmt.Println(block.Hash)


	contract,err := gt.GetContractDetails("KT1Um7ieBEytZtumecLqGeL56iY6BuWoBgio")
	if err != nil {
		return
	}
	fmt.Println(contract.Balance)

	printaddr := func(a models.KeyPair) {
		fmt.Print("\nPrivate Key: ", a.Pk)
		fmt.Print("\nSecret Key: ", a.Sk)
		fmt.Println("\nAddress: ", a.Address)
	}

	printaddr(gt.GenerateAddress())
	printaddr(gt.VanityAddressPrefix("tz1XXX"))


}