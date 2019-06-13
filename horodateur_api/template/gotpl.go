package template

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/geneva_horodateur/merkle"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"
)

type ChainpointTex struct {
	merkle.Chainpoint
	Date     string
	JsonData string
	Address  string
}

func MakeTemplate(rcpt []byte, lang string, now time.Time) ([]byte, error) {
	tmplFile := fmt.Sprintf("./template/latex/template.%s.tex", lang)
	_uuid := uuid.NewV4().String()
	outFile := fmt.Sprintf("/tmp/%s.tex", _uuid)
	pdfFile := fmt.Sprintf("/tmp/%s.pdf", _uuid)

	var _rcpt ChainpointTex
	err := json.Unmarshal(rcpt, &_rcpt)
	if err != nil {
		return nil, err
	}
	jsonData := strings.Trim(string(rcpt), "\n")
	jsonData = strings.Replace(jsonData, "{", "\\{", -1)
	jsonData = strings.Replace(jsonData, "}", "\\}", -1)
	_rcpt.JsonData = fmt.Sprintf(":JsOnbegin:%v:JsOnend:", jsonData)
	tmpl := template.Must(template.ParseFiles(tmplFile))
	f, err := os.Create(outFile)
	if err != nil {
		return nil, err
	}
	_rcpt.Date = fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())

	privateKeyHex := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Error on private key conversion: %s", err.Error())
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	_rcpt.Address = crypto.PubkeyToAddress(*publicKeyECDSA).String()

	tmpl.Execute(f, _rcpt)
	cmd := exec.Command("./template/make_receipt.sh", _uuid)

	/*
		go func() { // Watchdog function
			time.Sleep(10 * time.Second)
			cmd.Process.Kill()
		}()*/
	if err = cmd.Start(); err != nil {
		return nil, err
	}
	if err = cmd.Wait(); err != nil {
		return nil, err
	}

	pdf_receipt, err := ioutil.ReadFile(pdfFile)

	os.Remove(pdfFile)
	return pdf_receipt, nil
}
