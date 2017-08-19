package template

import (
	"encoding/json"
	"fmt"
	"github.com/Magicking/rc-ge-ch-pdf/merkle"
	"github.com/satori/go.uuid"
	"io/ioutil"
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
}

func MakeTemplate(rcpt []byte, now time.Time) ([]byte, error) {
	tmplFile := "./template/latex/template.tex"
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
	_rcpt.JsonData = fmt.Sprintf(":_JsOn_begin:%v:_JsOn_end:", jsonData)
	tmpl := template.Must(template.ParseFiles(tmplFile))
	f, err := os.Create(outFile)
	if err != nil {
		return nil, err
	}
	_rcpt.Date = fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())
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
