package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type AnchorPoint struct {
	SourceID string `json:"sourceId"`
	Type     string `json:"type"`
}

type ChainpointLeafString struct {
	Left  string `json:"left,omitempty"`  // Hashes from leaf's sibling to a root's child.
	Right string `json:"right,omitempty"` // Hashes from leaf's sibling to a root's child.
}

type Chainpoint struct {
	Context    string                 `json:"@context"`
	Anchors    []AnchorPoint          `json:"anchors"`
	MerkleRoot string                 `json:"merkleRoot"`
	Proof      []ChainpointLeafString `json:"proof"`
	TargetHash string                 `json:"targetHash"`
	Type       string                 `json:"type"`
}

type ChainpointTex struct {
	Chainpoint

	JsonData string
}

func main() {
	tmplFile := "template.tex"
	rcptFile := "rcpt.json"
	outFile := "output.tex"

	// Open Json file
	raw, err := ioutil.ReadFile(rcptFile)
	if err != nil {
		log.Fatal(err)
	}

	var rcpt ChainpointTex
	err = json.Unmarshal(raw, &rcpt)
	if err != nil {
		log.Fatal(err)
	}
	jsonData := strings.Trim(string(raw), "\n")
	jsonData = strings.Replace(jsonData, "{", "\\{", -1)
	jsonData = strings.Replace(jsonData, "}", "\\}", -1)
	rcpt.JsonData = fmt.Sprintf(":_JsOn_begin:%v:_JsOn_end:", jsonData)
	tmpl := template.Must(template.ParseFiles(tmplFile))
	f, err := os.Create(outFile)
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(f, rcpt)
}
