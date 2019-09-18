package internal

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	models "github.com/geneva_horodateur/models"
	op "github.com/geneva_horodateur/restapi/operations"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/geneva_horodateur/template"
)

func newOctetStream(r io.Reader, fn string) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, _ runtime.Producer) {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", fn))
		io.Copy(w, r)
	})
}

func GetreceiptHandler(ctx context.Context, params op.GetreceiptParams) middleware.Responder {
	var lang string
	if params.Lang != nil {
		lang = strings.ToLower(*params.Lang)
	}
	switch lang {
	case "fr", "de", "it", "en":
	default:
		lang = "fr"
	}
	rcpt, ok, err := GetReceiptByHash(ctx, params.Hash)
	if err != nil {
		err_str := fmt.Sprintf("Failed to call %s: %v", "GetReceiptByHash", err)
		log.Printf(err_str)
		return op.NewGetreceiptDefault(500).WithPayload(&models.Error{Message: &err_str})
	}
	if !ok {
		err_str := fmt.Sprintf("No receipt found for %s", params.Hash)
		log.Printf(err_str)
		return op.NewGetreceiptDefault(500).WithPayload(&models.Error{Message: &err_str})
	}
	if !ok {
		err_str := fmt.Sprintf("No receipt found for %s", params.Hash)
		log.Printf(err_str)
		return op.NewGetreceiptDefault(500).WithPayload(&models.Error{Message: &err_str})
	}

	rcptPdf, err := template.MakeTemplate(rcpt.JSONData, lang, time.Now())
	if err != nil {
		err_str := fmt.Sprintf("Failed to call %s: %v", "MakeTemplate", err)
		log.Printf(err_str)
		return op.NewGetreceiptDefault(500).WithPayload(&models.Error{Message: &err_str})
	}
	reader := bytes.NewReader(rcptPdf)
	filename := fmt.Sprintf("%s_recu.pdf", rcpt.Filename)
	return newOctetStream(reader, filename)
}

func DelreceiptsHandler(ctx context.Context, params op.DelreceiptsParams) middleware.Responder {
	for _, hash := range params.Hashes {
		if(len(hash) == 0) {
			err_str := fmt.Sprintf("Please make sure that hash \"%s\" is valid", hash)
			log.Printf(err_str)
			return op.NewDelreceiptsDefault(400).WithPayload(&models.Error{Message: &err_str})
		}

		_, ok, err := GetReceiptByHash(ctx, hash)
		if err != nil {
			err_str := fmt.Sprintf("Failed to call %s: %v", "GetReceiptByHash", err)
			log.Printf(err_str)
			return op.NewGetreceiptDefault(500).WithPayload(&models.Error{Message: &err_str})
		}
		if !ok {
			err_str := fmt.Sprintf("No receipt found for %s", hash)
			log.Printf(err_str)
			return op.NewGetreceiptDefault(500).WithPayload(&models.Error{Message: &err_str})
		}
	}

	for _, hash := range params.Hashes {
		err := DelReceiptByHash(ctx, hash)
		if err != nil {
			log.Println(err)
		}
	}
	return op.NewDelreceiptsOK()
}

func ListtimestampedHandler(ctx context.Context, params op.ListtimestampedParams) middleware.Responder {
	rcpts, err := GetAllReceipts(ctx)
	if err != nil {
		err_str := fmt.Sprintf("Failed to call %s: %v", "GetAllReceipts", err)
		log.Printf(err_str)
		return op.NewListtimestampedDefault(500).WithPayload(&models.Error{Message: &err_str})
	}
	var ret []*models.ReceiptFile
	for _, rcpt := range rcpts {
		ret_rcpt := models.ReceiptFile{Filename: rcpt.Filename,
			Hash:              rcpt.TargetHash,
			Horodatingaddress: "",
			Date:              rcpt.Date.Unix(),
			Transactionhash:   rcpt.TransactionHash,
		}
		ret = append(ret, &ret_rcpt)
	}
	return op.NewListtimestampedOK().WithPayload(ret)
}


//func MonitoringHandler(ctx context.Context, params op.MonitoringParams, w http.ResponseWriter, r *http.Request) (middleware.Responder){
func MonitoringHandler(ctx context.Context, params op.MonitoringParams) middleware.Responder {
	nodeOk := GetNodeSignal(ctx)
	persistence, err := GetDBTests()
	if err != nil {
		persistence = false
	}

	if nodeOk {
		errorThreshold, warningThreshold := GetEthereumBalance()
		var sondeResp []*models.Sonde
		sondeResp_rcpt := models.Sonde{EthereumActive: true,
			BalanceErrorThresholdExceeded: errorThreshold,
			BalanceWarningThresholdExceeded: warningThreshold,
			PersistenceActive: persistence,
		}
		sondeResp = append(sondeResp, &sondeResp_rcpt)
		return op.NewMonitoringOK().WithPayload(sondeResp)
	} else {
		var sondeResp []*models.Sonde
		sondeResp_rcpt := models.Sonde{EthereumActive: false,
			BalanceErrorThresholdExceeded: false,
			BalanceWarningThresholdExceeded: false,
			PersistenceActive: persistence,
		}
		sondeResp = append(sondeResp, &sondeResp_rcpt)
		return op.NewMonitoringOK().WithPayload(sondeResp)
	}
}

func ConfigureSAMLHandler(ctx context.Context, params op.ConfigureSAMLParams) middleware.Responder {
	//configureSAML()
	//samlSP.requireAccount()
	return op.NewConfigureSAMLOK()
}