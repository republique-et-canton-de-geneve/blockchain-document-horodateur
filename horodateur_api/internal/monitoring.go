package internal

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"math"
	"math/big"
)

type Sonde struct {
	ethereumActive                  bool
	balanceErrorThresholdExceeded   bool
	balanceWarningThresholdExceeded bool
	persistenceActive               bool
}

type TestStruct struct {
	Test string
}

type MonitoringEnv struct {
	NodeAddress      string
	LockedAddress    string
	PrivateKey       string
	ErrorThreshold   float64
	WarningThreshold float64
}

var mn MonitoringEnv

func GetNodeSignal(ctx context.Context) bool {
	blkCtx, ok := BLKFromContext(ctx)
	if !ok {
		log.Fatalf("Could not obtain ClientConnector from context\n", ok)
		return false
	}
	//txHash:= common.HexToHash("75139f2e9f045987f67ab04541d03d7cd872e663b5efd758c20da42c89e652eb")
	//Above this comment is the line of code used for production version, the hash used is from the main net.
	//Below this comment is the line of code used for development version, the hash used is from the rinkeby testnet.
	txHash := common.HexToHash("d3851f8ee9bbd79a4cf332999a89a4b2c6b8d5c4c0c001ea85e95ab7997843c0")
	_, _, err := blkCtx.NC.GetTransaction(ctx, txHash)
	if err != nil {
		return false
	}
	return true
}

func GetDBTests() (bool, error) {
	var err error

	if err = DB.AutoMigrate(&TestStruct{}).Error; err != nil {
		DB.Close()
		return false, err
	}
	if err = DB.Create(&TestStruct{Test: "Database test"}).Error; err != nil {
		fmt.Println(" Having troubles creating in the Database")
		return false, err
	}

	var testStruct TestStruct
	//Testing with a real entry
	if err = DB.Where("test= ?", "Database test").First(&testStruct).Error; err != nil {
		return false, err
	}
	//Testing with a fake entry
	if err = DB.Where("test= ?", "Database tset").First(&testStruct).Error; err != nil {
		return true, nil
	}
	return true, nil
}

func GetEthereumBalance() (bool, bool) {

	client, err := ethclient.Dial(mn.NodeAddress)
	account := common.HexToAddress(mn.LockedAddress)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatalf("getting balance error ",err)
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	errThre, warThre := big.NewFloat(mn.ErrorThreshold), big.NewFloat(mn.WarningThreshold)
	var errBool, warnBool bool
	errBool = true
	if ethValue.Cmp(errThre) == -1 {
		errBool = false
	}
	warnBool = true
	if ethValue.Cmp(warThre) == -1 {
		warnBool = false
	}

	return errBool, warnBool
}

func InitMonitoring(nodeAddress string, lockedAddress string, privateKey string, errorThreshold float64, warningThreshold float64) MonitoringEnv {
	mn = MonitoringEnv{NodeAddress: nodeAddress,
		LockedAddress:    lockedAddress,
		PrivateKey:       privateKey,
		ErrorThreshold:   errorThreshold,
		WarningThreshold: warningThreshold}
	return mn
}
