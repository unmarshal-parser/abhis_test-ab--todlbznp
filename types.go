// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"math/big"
	"sync"
	"time"
)

var (
	_ = decimal.Decimal{}
	_ = big.NewInt
	_ = datatypes.JSON{}
	_ = time.Time{}
)

func GetTransferEventHash() string {
	return "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
}

type TransferEvent struct {
	EventFrom  string
	EventTo    string
	EventValue decimal.Decimal `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:f8a23122-538c-465a-a453-3fdd0f21b595,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:f8a23122-538c-465a-a453-3fdd0f21b595,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:f8a23122-538c-465a-a453-3fdd0f21b595,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetTransferMethodHash() string {
	return "a9059cbb"
}

type TransferMethod struct {
	MethodRecipient string
	MethodAmount    decimal.Decimal `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:3b71ed49-311d-4e66-85fa-00d813bf9c08,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:3b71ed49-311d-4e66-85fa-00d813bf9c08,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

type LastSyncedBlock struct {
	Contract    string `gorm:"primaryKey"`
	ChainID     string `gorm:"primaryKey"`
	SyncType    string `gorm:"primaryKey"`
	BlockNumber uint64
}

// Plugin Models
type TokenDetails struct {
	ID      int
	Address string `gorm:"uniqueIndex:address_and_chain"`
	Symbol  string
	ChainID string `gorm:"uniqueIndex:address_and_chain"`
	Decimal int
	Name    string
}

var tokenCache = sync.Map{}

// Config
type PostgresConfig struct {
	ConnectionString string `mapstructure:"connection_string"`
	TablePrefix      string `mapstructure:"table_prefix"`
	CreateBatchSize  int    `mapstructure:"create_batch_size"`
	MaxConnections   int    `mapstructure:"max_connections"`
}

type IndexerConfig struct {
	EthEndpoint       string `mapstructure:"eth_endpoint"`
	ContractAddress   string `mapstructure:"contract_address"`
	StartBlock        int    `mapstructure:"start_block"`
	ApiKey            string `mapstructure:"api_key"`
	PostgresConfig    `mapstructure:"postgres_config"`
	LagToHighestBlock int `mapstructure:"lag_to_highest_block"`
	StepSize          int `mapstructure:"step_size"`
	ParallelCalls     int `mapstructure:"parallel_calls_for_logs"`
	PrometheusConfig  `mapstructure:"prometheus_config"`
	Metrics           *Metrics
}

type PrometheusConfig struct {
	PrometheusNameSpace string //PrometheusNameSpace The prometheus namespace metrics will be reported on. Default -> custom_indexer
	PrometheusSubsystem string //PrometheusSubsystem The prometheus subsystem metrics will be reported on. Default -> parser
	MetricsPort         int    `mapstructure:"metrics_port"`    //MetricsPort port metrics is scraped on. Default -> 9091
	PrometheusHost      string `mapstructure:"prometheus_host"` //PrometheusHost Host for the metrics server. Default -> 0.0.0.0
}

type Metrics struct {
	//List of metrics
	NodeRestarts         *prometheus.CounterVec   //NodeRestarts node restarts specific to this parser
	NodeLatency          *prometheus.HistogramVec //NodeLatency Response latency for node
	UnmarshalAPIRestarts *prometheus.CounterVec   //UnmarshalAPIRestarts api gateway restarts specific to this parser
	UnmarshalAPILatency  *prometheus.HistogramVec //UnmarshalAPILatency Response Latency for unmarshal API calls

}

func (i *IndexerConfig) AssignDefaults() {
	if i.PostgresConfig.CreateBatchSize == 0 {
		i.PostgresConfig.CreateBatchSize = 100
	}
	if i.StepSize == 0 {
		i.StepSize = 50
	}
	if i.LagToHighestBlock == 0 {
		i.LagToHighestBlock = 10
	}
	if i.ParallelCalls == 0 {
		i.ParallelCalls = 1
	}
	if i.MetricsPort == 0 {
		i.MetricsPort = 9091
	}
	if i.PrometheusHost == "" {
		i.PrometheusHost = "0.0.0.0"
	}

	i.PrometheusNameSpace = "custom_indexer"
	i.PrometheusSubsystem = "parser"
}
