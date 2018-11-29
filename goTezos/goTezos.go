package goTezos

import (
	"github.com/DefinitelyNotAGoat/goTezos/models"
	"fmt"
	"math/rand"
	"time"
	"sync"
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
	go func(a *GoTezos){
		for {
			time.Sleep(15*time.Second)
			a.checkUnhealthyClients()
		}
	}(&a)
	return &a
}



func (this *GoTezos) SetLoggingFunction(a func(level, client, message string)) {
	this.logfunction = a
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
	wg := sync.WaitGroup{}
	for _,a := range this.RpcClients {
		wg.Add(1)
		go func(wg *sync.WaitGroup, client *TezClientWrapper){
			res := a.client.Healthcheck()
			if a.healthy && res == false {
				this.logfunction("WARNING","Client State swithished to unhealthy", this.ActiveRPCCient.client.Host + this.ActiveRPCCient.client.Port)
			}
			if !a.healthy && res {
				this.logfunction("INFO","Client State swithished to healthy", this.ActiveRPCCient.client.Host + this.ActiveRPCCient.client.Port)
			}
			a.healthy = res
			wg.Done()
		}(&wg, a)
	}
	wg.Wait()
	this.clientLock.Unlock()
}

func (this *GoTezos) checkUnhealthyClients(){
	this.clientLock.Lock()
	wg := sync.WaitGroup{}
	for _,a := range this.RpcClients {
		wg.Add(1)
		go func(wg *sync.WaitGroup, client *TezClientWrapper){
			if a.healthy == false {
				res := a.client.Healthcheck()
				if !a.healthy && res {
					this.logfunction("INFO", "Client State swithished to healthy", this.ActiveRPCCient.client.Host+this.ActiveRPCCient.client.Port)
				}
				a.healthy = res
			}
			wg.Done()
		}(&wg, a)
	}
	wg.Wait()
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


func (this *GoTezos) sethislient() error {
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
	e := this.sethislient()
	if e != nil {
		this.logfunction("FATAL","goTezos","could not find any healthy Clients")
		return models.ResponseRaw{},e
	}


	r, err := this.ActiveRPCCient.client.GetResponse(method,args)
	if err != nil {
		this.ActiveRPCCient.healthy = false
		this.logfunction("WARNING", this.ActiveRPCCient.client.Host + this.ActiveRPCCient.client.Port, "Client State swithished to unhealthy")
		return this.GetResponse(method,args)
	}
	return r,err
}