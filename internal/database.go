package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Magicking/rc-ge-ch-pdf/merkle"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"math"
	"math/big"
	"time"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

type Receipt struct {
	gorm.Model
	TargetHash      string
	TransactionHash string
	Filename        string
	Date            time.Time
	JSONData        []byte
}

type Sonde struct {
	ethereumActive						bool
	//balanceErrorThresholdExceeded		bool
	//balanceWarningThresholdExceeded 	bool
	//persistenceActive					bool
}

var NodeAddress string
var LockedAddress string
//Maybe error is warning and vice-versa
var ErrorThreshold int
var WarningThreshold int

func GetNodeSignal(ctx context.Context) (bool){
	fmt.Println(NodeAddress, " URI")
	ctxt := NewCCToContext(ctx, NodeAddress)
	_, ok := CCFromContext(ctxt)
	if !ok {
		log.Fatalf("Could not obtain ClientConnector from context\n")
	}
	return ok
}

func GetEthereumBalance() (*big.Float, int, int) {
	client, err := ethclient.Dial(NodeAddress)
	account := common.HexToAddress(LockedAddress)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)
	return ethValue, ErrorThreshold, WarningThreshold
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

func InitDatabase(dbDsn string, nodeAddress string, lockedAddress string, errorThreshold int, warningThreshold int) (*gorm.DB, error) {
	var err error
	var db *gorm.DB

	NodeAddress = nodeAddress
	LockedAddress = lockedAddress
	ErrorThreshold = errorThreshold
	WarningThreshold = warningThreshold

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
	return db, nil
}