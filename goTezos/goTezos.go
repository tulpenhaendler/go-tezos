package goTezos

import (
	"github.com/DefinitelyNotAGoat/goTezos/models"
	"fmt"
	"math/rand"
	"time"
	"sync"
	"encoding/json"
)

type TezClientWrapper struct {
	healthy bool // isHealthy
	client  *TezosRPCClient
}

/*
 * GoTezos manages multiple Clients
 * each Client represents a Connection to a Tezos Node
 * GoTezos manages failover if one Node is down, there
 * are 2 Strategies:
 * failover: always use the same unless it is down -> go to the next - default
 * random: send to each Node equally
 */
type GoTezos struct {
	clientLock sync.Mutex
	RpcClients []*TezClientWrapper
	ActiveRPCCient *TezClientWrapper
	balancerStrategy string
	logfunction func(level, client, message string)
	rand *rand.Rand
}


func NewGoTezos() *GoTezos {
	a := GoTezos{}
	a.UseBalancerStrategyFailover()
	a.logfunction = func(level, client, message string){
		fmt.Printf("%8s %s: %s\n", level, client, message)
	}
	a.rand = rand.New(rand.NewSource(time.Now().Unix()))
	return &a
}

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



func (this *GoTezos) AddNewClient(client *TezosRPCClient) {
	this.clientLock.Lock()
	client.SetLogFunction(func(level, msg string) {
		this.logfunction(level, client.Host + client.Port, msg)
	})
	this.RpcClients = append(this.RpcClients, &TezClientWrapper{true, client})
	this.clientLock.Unlock()
}

func (this *GoTezos) UseBalancerStrategyFailover(){
	this.balancerStrategy = "failover"
}

func (this *GoTezos) UseBalancerStrategyRandom(){
	this.balancerStrategy = "random"
}

func (this *GoTezos) checkHealthStatus(){
	this.clientLock.Lock()
	for _,a := range this.RpcClients {
		res := a.client.Healthcheck()
		if a.healthy && res == false {
			this.logfunction("WARNING","Client State switched to unhealthy", this.ActiveRPCCient.client.Host + this.ActiveRPCCient.client.Port)
		}
		if !a.healthy && res {
			this.logfunction("INFO","Client State switched to healthy", this.ActiveRPCCient.client.Host + this.ActiveRPCCient.client.Port)
		}
		a.healthy = res
	}
	this.clientLock.Unlock()
}

func (this *GoTezos) getFirstHealthyClient() *TezClientWrapper{
	this.clientLock.Lock()
	defer this.clientLock.Unlock()
	for _,a := range this.RpcClients {
		if a.healthy == true {
			return a
		}
	}
	return nil
}

func (this *GoTezos) getRandomHealthyClient() *TezClientWrapper{
	this.clientLock.Lock()
	defer this.clientLock.Unlock()
	clients := []int{}
	for i,_ := range this.RpcClients {
		clients = append(clients, i)
	}
	for _, i := range this.rand.Perm(len(clients)) {
		return this.RpcClients[i]
	}
	return nil
}


func (this *GoTezos) setClient() error {
	if this.balancerStrategy == "failover" {
		c := this.getFirstHealthyClient()
		if c == nil {
			this.checkHealthStatus()
			c = this.getFirstHealthyClient()
			if c == nil {
				return NoClientError{}
			}
		}
		this.ActiveRPCCient = c
	}

	if this.balancerStrategy == "random" {
		c := this.getRandomHealthyClient()
		if c == nil {
			this.checkHealthStatus()
			c = this.getRandomHealthyClient()
			if c == nil {
				return NoClientError{}
			} else {
				this.ActiveRPCCient = c
			}
		}
		this.ActiveRPCCient = c

	}
	return nil
}

func (this *GoTezos) GetResponse(method string, args string) (models.ResponseRaw, error) {
	e := this.setClient()
	if e != nil {
		this.logfunction("FATAL","goTezos","could not find any healthy Clients")
		return models.ResponseRaw{},e
	}


	r, err := this.ActiveRPCCient.client.GetResponse(method,args)
	if err != nil {
		this.ActiveRPCCient.healthy = false
		this.logfunction("WARNING", this.ActiveRPCCient.client.Host + this.ActiveRPCCient.client.Port, "Client State switched to unhealthy")
		return this.GetResponse(method,args)
	}
	return r,err
}