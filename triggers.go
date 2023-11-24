// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"fmt"
	unmarshal "github.com/eucrypt/unmarshal-go-sdk/pkg"
	conf "github.com/eucrypt/unmarshal-go-sdk/pkg/config"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	sdkTokenDetails "github.com/eucrypt/unmarshal-go-sdk/pkg/token_details"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/token_details/types"
	tokenPriceTypes "github.com/eucrypt/unmarshal-go-sdk/pkg/token_price/types"
	"github.com/onrik/ethrpc"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
)

var (
	_   = types.TokenDetails{}
	_   = prometheus.Labels{}
	_   = decimal.Decimal{}
	_   = big.NewInt
	_   = ethrpc.Transaction{}
	sdk = unmarshal.Unmarshal{}
	_   = sdkTokenDetails.TokenDetailsOptions{}
	_   = gorm.Model{}
	_   = time.Time{}
	_   = strconv.NumError{}
	_   = math.NaN()
	_   = log.Error
	_   = tokenPriceTypes.TokenPrice{}
	_   = fmt.Sprint()
)

func initUnmarshalSDK(cfg IndexerConfig) {
	key := strings.TrimSpace(cfg.ApiKey)
	if key == "" {
		panic("missing api key")
	}
	sdk = unmarshal.NewWithConfig(conf.Config{
		AuthKey:     key,
		Environment: constants.Prod,
	})
}

func (entity *TransferEvent) BeforeCreateHook(tx *gorm.DB) error {

	return nil
}

func (entity *TransferEvent) AfterCreateHook(tx *gorm.DB) error {
	return nil
}

func (entity *TransferMethod) BeforeCreateHook(tx *gorm.DB) error {

	return nil
}

func (entity *TransferMethod) AfterCreateHook(tx *gorm.DB) error {
	return nil
}
