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
	ethereumActive						bool
	balanceErrorThresholdExceeded		bool
	balanceWarningThresholdExceeded 	bool
	persistenceActive					bool
}

type TestStruct struct {
	Test 								string
}

type MonitoringEnv struct {
	NodeAddress                        	string
	LockedAddress                       string
	PrivateKey							string
	ErrorThreshold                      float64
	WarningThreshold                    float64
}

var mn MonitoringEnv

func GetNodeSignal(ctx context.Context) (bool){
	blkCtx, ok := BLKFromContext(ctx)
	if !ok {
		log.Println("Could not obtain ClientConnector from context\n")
		return false
	}
	// test connexion
	i, err := blkCtx.NC.SuggestGasPrice(ctx)
	if err != nil {
		return false
	}
	log.Println(i)
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
	fmt.Println(mn.NodeAddress)
	fmt.Println(mn.LockedAddress)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	errThre, warThre := big.NewFloat(mn.ErrorThreshold), big.NewFloat(mn.WarningThreshold)
	var errBool, warnBool bool

	if (ethValue.Cmp(errThre) == -1) {
		errBool = false
	} else {
		errBool = true
	}
	if (ethValue.Cmp(warThre) == -1) {
		warnBool = false
	} else {
		warnBool = true
	}
	return errBool, warnBool
}

func InitMonitoring(nodeAddress string, lockedAddress string, privateKey string, errorThreshold float64, warningThreshold float64) MonitoringEnv {
	mn = MonitoringEnv{NodeAddress: nodeAddress,
		LockedAddress: lockedAddress,
		PrivateKey: privateKey,
		ErrorThreshold: errorThreshold,
		WarningThreshold: warningThreshold}
	return mn
}
