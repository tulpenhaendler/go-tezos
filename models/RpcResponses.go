package models

import (
	"time"
	"encoding/json"
)

type ResponseRaw struct {
	Bytes []byte
}

type Block struct {
	UpdatedAt time.Time
	DeletedAt *time.Time
	Protocol  string `json:"protocol"`
	ChainID   string `json:"chain_id"`
	Hash      string `gorm:"primary_key:yes,unique_index",json:"hash"`
	Header struct {
		Level            int64     `json:"level"`
		Proto            int       `json:"proto"`
		Predecessor      string    `json:"predecessor"`
		Timestamp        time.Time `json:"timestamp"`
		ValidationPass   int       `json:"validation_pass"`
		OperationsHash   string    `json:"operations_hash"`
		Fitness          []string  `gorm:"-",json:"fitness"`
		Context          string    `json:"context"`
		Priority         int       `json:"priority"`
		ProofOfWorkNonce string    `json:"proof_of_work_nonce"`
		Signature        string    `json:"signature"`
	} `gorm:"embedded;embedded_prefix:header_",json:"header"`
	Metadata struct {
		Protocol     string `json:"protocol"`
		NextProtocol string `json:"next_protocol"`
		TestChainStatus struct {
			Status string `json:"status"`
		} `gorm:"embedded",json:"test_chain_status"`
		MaxOperationsTTL       int `json:"max_operations_ttl"`
		MaxOperationDataLength int `json:"max_operation_data_length"`
		MaxBlockHeaderLength   int `json:"max_block_header_length"`
		MaxOperationListLength []struct {
			MaxSize int `json:"max_size"`
			MaxOp   int `json:"max_op,omitempty"`
		} `gorm:"-",json:"max_operation_list_length"`
		Baker string `json:"baker"`
		Level struct {
			Level                uint64 `json:"level"`
			LevelPosition        int    `json:"level_position"`
			Cycle                int    `json:"cycle"`
			CyclePosition        int    `json:"cycle_position"`
			VotingPeriod         int    `json:"voting_period"`
			VotingPeriodPosition int    `json:"voting_period_position"`
			ExpectedCommitment   bool   `json:"expected_commitment"`
		} `gorm:"embedded;embedded_prefix:level_",json:"level"`
		VotingPeriodKind string `json:"voting_period_kind"`
	} `gorm:"embedded;embedded_prefix:level_",json:"metadata"`
	Operations [][] json.RawMessage `gorm:"-",json:"operations"`
	Fee        uint64
	OperationsJSON string
	OperationsCount int64
}


type Peer struct {
	Id      string `gorm:"primary_key:yes,unique_index"`
	Score   int    `json:"score"`
	Trusted bool   `json:"trusted"`
	ConnMetadata struct {
		DisableMempool bool `json:"disable_mempool"`
		PrivateNode    bool `json:"private_node"`
	} `gorm:"-",json:"conn_metadata"`
	State string `json:"state"`
	ReachableAt struct {
		Addr string `json:"addr"`
		Port int    `json:"port"`
	} `json:"reachable_at";gorm:"-"`
	Addr string
	Port int
	Stat struct {
		TotalSent      string `json:"total_sent"`
		TotalRecv      string `json:"total_recv"`
		CurrentInflow  int    `json:"current_inflow"`
		CurrentOutflow int    `json:"current_outflow"`
	} `gorm:"-",json:"stat"`
	TotalSent                 string        `json:"total_sent"`
	TotalRecv                 string        `json:"total_recv"`
	CurrentInflow             int           `json:"current_inflow"`
	CurrentOutflow            int           `json:"current_outflow"`
	LastEstablishedConnection []interface{} `gorm:"-",json:"last_established_connection"`
	LastSeen                  []interface{} `gorm:"-",json:"last_seen"`
	LastCheck                 time.Time
	LastRunning               time.Time
	City                      string
	Country                   string
	Continent                 string
	Latitude                  float64
	Longitude                 float64
	Timezone                  string
	Version					  string
}

type ContractDetails struct {
	UpdatedAt time.Time
	DeletedAt *time.Time
	Contract  string `gorm:"primary_key:yes,unique_index"`
	Manager   string `json:"manager",gorm:"index:idx_manager"`
	Balance   uint64 `json:"balance,string",gorm:"type:BigInt[]"`
	Spendable bool   `json:"spendable"`
	Counter uint64 `json:"counter,string",gorm:"type:decimal"`
}

type DelegateDetails struct {
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	Delegate      string `gorm:"primary_key:yes,unique_index"`
	Balance       uint64 `json:"balance,string" gorm:"type:decimal"`
	FrozenBalance uint64 `json:"frozen_balance,string" gorm:"type:decimal"`
	FrozenBalanceByCycle []struct {
		Cycle   int    `json:"cycle"`
		Deposit string `json:"deposit"`
		Fees    string `json:"fees"`
		Rewards string `json:"rewards"`
	} `json:"frozen_balance_by_cycle" gorm:"-"`
	StakingBalance     uint64   `json:"staking_balance" gorm:"type:decimal"`
	DelegatedContracts []string `gorm:"-" json:"delegated_contracts"`
	DelegatedBalance   uint64   `json:"delegated_balance,string",gorm:"type:decimal"`
	Deactivated        bool     `json:"deactivated"`
	GracePeriod        int      `json:"grace_period"`
}