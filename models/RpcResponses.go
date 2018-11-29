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