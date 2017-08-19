package internal

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	models "github.com/Magicking/rc-ge-ch-pdf/models"
	op "github.com/Magicking/rc-ge-ch-pdf/restapi/operations"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	tmpl "github.com/Magicking/rc-ge-ch-pdf/template"
)

func newOctetStream(r io.Reader, fn string) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, _ runtime.Producer) {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", fn))
		io.Copy(w, r)
	})
}

func GetreceiptHandler(ctx context.Context, params op.GetreceiptParams) middleware.Responder {
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
	rcptPdf, err := tmpl.MakeTemplate(rcpt.JSONData, time.Now())
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
