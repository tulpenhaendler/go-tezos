package goTezos
/*
Author: DefinitelyNotAGoat/MagicAglet
Version: 0.0.1
Description: This file contains specific functions for delegation services
License: MIT
*/

import (
  "math/rand"
  "time"
  "strconv"
  "errors"
  "./tezex"
)

/*
Description: Calculates the percentage share of a specific cycle for all delegated contracts on a range of cycles.
Param delegatedClients ([]DelegatedClient): A list of all the delegated contracts
Param cycleStart (int): The first cycle we are calculating
Param cycleEnd (int): The last cycle we are calculating
Returns delegatedClients ([]DelegatedClient): A list of all the delegated contracts
*/
func CalculateAllCommitmentsForCycles(delegatedClients []DelegatedClient, cycleStart int, cycleEnd int, rate float64) ([]DelegatedClient, error){
  var err error

  for cycleStart <= cycleEnd {
    delegatedClients, err = CalculateAllCommitmentsForCycle(delegatedClients, cycleStart, rate)
    if (err != nil){
      return delegatedClients, errors.New("Could not calculate all commitments for cycles " + strconv.Itoa(cycleStart) + "-" +  strconv.Itoa(cycleEnd) + ":CalculateAllCommitmentsForCycle(delegatedClients []DelegatedClient, cycle int, rate float64) failed: " + err.Error())
    }
    cycleStart = cycleStart + 1
  }
   return delegatedClients, nil
}

/*
Description: Calculates the percentage share of a specific cycle for all delegated contracts
Param delegatedClients ([]DelegatedClient): A list of all the delegated contracts
Param cycle (int): The cycle we are calculating
Returns delegatedClients ([]DelegatedClient): A list of all the delegated contracts
*/
func CalculateAllContractsForCycle(delegatedContracts []DelegatedContract, cycle int, rate float64) error {
  var sum float64
  sum = 0
  for index, delegation := range delegatedContracts{
    balance, err := GetBalanceAtSnapShotFor(delegation.Address, cycle)
    if (err != nil){
      return delegatedClients, errors.New("Could not calculate all commitments for cycle " + strconv.Itoa(cycle) + ":GetBalanceAtSnapShotFor(tezosAddr string, cycle int) failed: " + err.Error())
    }
    sum = sum + balance
    delegatedClients[index].Contracts = append(delegatedClients[index].Contracts, Contracts{Cycle:cycle, Amount:balance})
  }

  for index, delegation := range  delegatedContracts{
    counter := 0
    for i, contract := range delegatedContracts.Contracts {
      if (delegatedContracts[index].Contracts[i].Cycle == cycle){
        break
      }
      counter = counter + 1
    }
    delegatedContracts[index].Contracts[counter].SharePercentage = delegatedContracts[index].Contracts[counter].Amount / sum
    delegatedContracts[index].Contracts[counter] = CalculatePayoutForContract(delegatedContracts[index].Contracts[counter], rate, delegatedContracts[index].Delegator)
  }
}

/*
Description: Retrieves the list of addresses delegated to a delegate
Param SnapShot: A SnapShot object describing the desired snap shot.
Param delegateAddr: A string that represents a delegators tz address.
Returns []string: An array of contracts delegated to the delegator during the snap shot
*/
func GetDelegatedContractsForCycle(cycle int, delegateAddr string) ([]DelegatedContract, error){
  var delegatedContracts []DelegatedContract
  snapShot, err := GetSnapShot(cycle) //TODO. Dependant on RPC from personal node. FIgure out how to do this with tezex.info API
  if (err != nil){
    return errors.New("func GetDelegatedContractsForCycle(cycle int, delegateAddr string) failed: " + err.Error())
  }
  contracts := tezex.GetDelegationsToAccount(delegateAddr, snapShot.AssociatedBlock)
  for _, contract := range contracts {
    delegatedContracts = append(delegatedContracts, DelegatedContract{Address:contract.PublicKey, TimeStamp:contract.Time, Delegate:false})
  }
  delegatedContracts = append(delegatedContracts, DelegatedContract{Address:delegateAddr, Delegate:true})
  return delegatedContracts, nil
}

/*
Description: Gets a list of all of the delegated contacts to a delegator
Param delegateAddr (string): string representation of the address of a delegator
Returns ([]string): An array of addresses (delegated contracts) that are delegated to the delegator
*/
func GetAllDelegatedContracts(delegateAddr string) []DelegatedContract{
  contracts := tezex.GetDelegationsToAccount(delegateAddr)
  for _, contract := range contracts {
    delegatedContracts = append(delegatedContracts, DelegatedContract{Address:contract.PublicKey, TimeStamp:contract.Time, Delegate:false})
  }
  delegatedContracts = append(delegatedContracts, DelegatedContract{Address:delegateAddr, Delegate:true})
  return delegatedContracts
}

/*
Description: Takes a commitment, and calculates the GrossPayout, NetPayout, and Fee.
Param commitment (Commitment): The commitment we are doing the operation on.
Param rate (float64): The delegation percentage fee written as decimal.
Param totalNodeRewards: Total rewards for the cyle the commitment represents. //TODO Make function to get total rewards for delegate in cycle
Param delegate (bool): Is this the delegate
Returns (Commitment): Returns a commitment with the calculations made
Note: This function assumes Commitment.SharePercentage is already calculated.
*/
func CalculatePayoutForContract(contract Contract, rate float64, delegate bool) Commitment{
  ////-------------JUST FOR TESTING -------------////
  rand.Seed(time.Now().Unix())
  totalNodeRewards := rand.Intn(105000 - 70000) + 70000
 ////--------------END TESTING ------------------////

  grossRewards := contract.SharePercentage * float64(totalNodeRewards)
  contract.GrossPayout = grossRewards
  fee := rate * grossRewards
  contract.Fee = fee
  var netRewards float64
  if (delegate){
    netRewards = grossRewards
    contract.NetPayout = netRewards
  } else {
    netRewards = grossRewards - fee
    contract.NetPayout = netRewards
  }

  return contract
}

/*
Description: A function to Payout rewards for all contracts in delegatedClients
Param delegatedClients ([]DelegatedClient): List of all contracts to be paid out
Param alias (string): The alias name to your known delegation wallet on your node


****WARNING****
If not using the ledger there is nothing stopping this from actually sending Tezos.
With the ledger you have to physically confirm the transaction, without the ledger you don't.

BE CAREFUL WHEN CALLING THIS FUNCTION!!!!!
****WARNING****
*/
func PayoutDelegatedContracts(delegatedClients []DelegatedClient, alias string) error{
  for _, delegatedClient := range delegatedClients {
    err := SendTezos(delegatedClient.TotalPayout, delegatedClient.Address, alias)
    if (err != nil){
      return errors.New("Could not Payout Delegated Contracts: SendTezos(amount float64, toAddress string, alias string) failed: " + err.Error())
    }
  }
  return nil
}

/*
Description: Calculates the total payout in all commitments for a delegated contract
Param delegatedClients (DelegatedClient): the delegated contract to calulate over
Returns (DelegatedClient): return the contract with the Total Payout
*/
func CalculateTotalPayout(delegatedClient DelegatedClient) DelegatedClient{
  for _, commitment := range delegatedClient.Commitments{
    delegatedClient.TotalPayout = delegatedClient.TotalPayout + commitment.NetPayout
  }
  return delegatedClient
}

/*
Description: payout in all commitments for a delegated contract for all contracts
Param delegatedClients (DelegatedClient): the delegated contracts to calulate over
Returns (DelegatedClient): return the contract with the Total Payout for all contracts
*/
func CalculateAllTotalPayout(delegatedClients []DelegatedClient) []DelegatedClient{
  for index, delegatedClient := range delegatedClients{
    delegatedClients[index] = CalculateTotalPayout(delegatedClient)
  }

  return delegatedClients
}

/*
Description: A test function that loops through the commitments of each delegated contract for a specific cycle,
             then it computes the share value of each one. The output should be = 1. With my tests it was, so you
             can really just ignore this.
Param cycle (int): The cycle number to be queryed
Param delegatedClients ([]DelegatedClient): the group of delegated DelegatedContracts
Returns (float64): The sum of all shares
*/
func CheckPercentageSumForCycle(cycle int, delegatedClients []DelegatedClient) float64{
  var sum float64
  sum = 0
  for x := 0; x < len(delegatedClients); x++{
    counter := 0
    for y := 0; y < len(delegatedClients[x].Commitments); y++{
      if (delegatedClients[x].Commitments[y].Cycle == cycle){
        break
      }
      counter = counter + 1
    }

    sum = sum + delegatedClients[x].Commitments[counter].SharePercentage
  }
  return sum
}

/*
Description: Reverse the order of an array of DelegatedClient.
             Used when fisrt retreiving contracts because the
             Tezos RPC API returns the newest contract first.
Param delegatedClients ([]DelegatedClient) Delegated

*/
func SortDelegateContracts(delegatedClients []DelegatedClient) []DelegatedClient{
   for i, j := 0, len(delegatedClients)-1; i < j; i, j = i+1, j-1 {
       delegatedClients[i], delegatedClients[j] = delegatedClients[j], delegatedClients[i]
   }
   return delegatedClients
}
