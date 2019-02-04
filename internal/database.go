package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Magicking/rc-ge-ch-pdf/merkle"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"math"
	"math/big"
	"time"
)

type Receipt struct {
	gorm.Model
	TargetHash      					string
	TransactionHash 					string
	Filename        					string
	Date            					time.Time
	JSONData        					[]byte
}

type Sonde struct {
	ethereumActive						bool
	balanceErrorThresholdExceeded		bool
	balanceWarningThresholdExceeded 	bool
	persistenceActive					bool
}

type TestStruct struct {
	Test 								string
}

var NodeAddress 						string
var LockedAddress						string
//Maybe error is warning and vice-versa
var ErrorThreshold						*big.Float
var WarningThreshold					*big.Float
var DB									*gorm.DB

func GetNodeSignal(ctx context.Context) (bool){
	fmt.Println(NodeAddress, " URI")
	ctxt := NewCCToContext(ctx, NodeAddress)
	_, ok := CCFromContext(ctxt)
	if !ok {
		log.Fatalf("Could not obtain ClientConnector from context\n")
	}
	return ok
}

func GetDBTests() (bool, error) {
	var err error

	if err = DB.AutoMigrate(&TestStruct{}).Error; err != nil {
		DB.Close()
		return false, err
	}
	if err = DB.Create(&TestStruct{Test: "Database test"}).Error; err != nil {
		fmt.Println(" Du mal à create dans la DB")
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
	client, err := ethclient.Dial(NodeAddress)
	account := common.HexToAddress(LockedAddress)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
//	resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
	var errBool, warnBool bool
	if (ethValue.Cmp(ErrorThreshold) == -1) {
		errBool = false
	} else {
		errBool = true
	}
	if (ethValue.Cmp(WarningThreshold) == -1) {
		warnBool = false
	} else {
		warnBool = true
	}
	return errBool, warnBool
}

func InsertReceipt(ctx context.Context, now time.Time, filename string, rcpt *merkle.Chainpoint) error {
	jsonData, err := json.Marshal(rcpt)
	if err != nil {
		return err
	}

	var tx_hash string

	if len(rcpt.Anchors) > 0 {
		tx_hash = rcpt.Anchors[0].SourceID
	}
	_rcpt := Receipt{
		TargetHash:      rcpt.TargetHash,
		JSONData:        jsonData,
		TransactionHash: tx_hash,
		Date:            now,
		Filename:        filename,
	}

	db, ok := DBFromContext(ctx)
	if !ok {
		return fmt.Errorf("Could not obtain DB from Context")
	}
	if err := db.Create(&_rcpt).Error; err != nil {
		return err
	}

	log.Printf("DEBUG: InsertReceipt: %v", _rcpt)

	return nil
}

func GetReceiptByHash(ctx context.Context, hash string) (*Receipt, bool, error) {
	var rcpt Receipt
	db, ok := DBFromContext(ctx)
	if !ok {
		return nil, false, fmt.Errorf("Could not obtain DB from Context")
	}
	cursor := db.Where(Receipt{TargetHash: hash})
	if cursor.Error != nil {
		return nil, false, fmt.Errorf("Error for TargetHash (%v): %v", hash, cursor.Error)
	}
	if cursor.Last(&rcpt).RecordNotFound() {
		return nil, false, nil
	}
	return &rcpt, true, nil
}

func DelReceiptByHash(ctx context.Context, hash string) error {
	db, ok := DBFromContext(ctx)
	if !ok {
		return fmt.Errorf("Could not obtain DB from Context")
	}
	cursor := db.Where(Receipt{TargetHash: hash}).Delete(Receipt{})
	if cursor.Error != nil {
		return fmt.Errorf("Error deleting for TargetHash (%v): %v", hash, cursor.Error)
	}
	return nil
}

func GetAllReceipts(ctx context.Context) ([]Receipt, error) {
	var receipts []Receipt

	db, ok := DBFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("Could not obtain DB from Context")
	}
	if db.Find(&receipts).RecordNotFound() {
		return nil, fmt.Errorf("RecordNotFound: No receipts in database")
	}

	return receipts, nil
}

func InitDatabase(dbDsn string, nodeAddress string, lockedAddress string, errorThreshold big.Float, warningThreshold big.Float) (*gorm.DB, error) {
	var err error
	var db *gorm.DB

	NodeAddress = nodeAddress
	LockedAddress = lockedAddress
	ErrorThreshold = &errorThreshold
	WarningThreshold = &warningThreshold

	for i := 1; i < 10; i++ {
		db, err = gorm.Open("postgres", dbDsn)
		if err == nil || i == 10 {
			break
		}
		sleep := (2 << uint(i)) * time.Second
		log.Printf("Could not connect to DB: %v", err)
		log.Printf("Waiting %v before retry", sleep)
		time.Sleep(sleep)
	}
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&Receipt{}).Error; err != nil {
		db.Close()
		return nil, err
	}
	DB = db
	return db, nil
}