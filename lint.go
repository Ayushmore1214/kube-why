package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// lintDocument is one YAML document's syntax-check result within a
// (possibly multi-document) file.
type lintDocument struct {
	Index int    `json:"document"`
	Valid bool   `json:"valid"`
	Error string `json:"error,omitempty"`
}

// lintFile decodes every document in a YAML file and reports whether each
// one is syntactically valid. It doesn't know anything about Kubernetes
// specifically yet, it just catches the kind of mistake (bad indentation, a
// stray tab, an unclosed quote, a duplicate key) that would otherwise show
// up later as a confusing kubectl error.
func lintFile(data []byte) []lintDocument {
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	var docs []lintDocument
	docIndex := 0

	for {
		var doc interface{}
		err := decoder.Decode(&doc)
		if err == io.EOF {
			break
		}
		docIndex++

		if err != nil {
			docs = append(docs, lintDocument{Index: docIndex, Valid: false, Error: err.Error()})
			// A syntax error leaves the decoder unable to find the next
			// document boundary, re-calling Decode() on it just returns the
			// same error forever instead of reaching EOF. Stop here rather
			// than loop indefinitely.
			break
		}
		if doc == nil {
			// an empty document, e.g. a trailing "---" with nothing after it
			docIndex--
			continue
		}
		docs = append(docs, lintDocument{Index: docIndex, Valid: true})
	}
	return docs
}

func runLint(path string, jsonMode bool) {
	data, err := os.ReadFile(path)
	if err != nil {
		if jsonMode {
			printJSONError(fmt.Sprintf("couldn't read %s: %v", path, err), nil)
		} else {
			fmt.Fprintf(os.Stderr, "kube-why: couldn't read %s: %v\n", path, err)
		}
		os.Exit(1)
	}

	docs := lintFile(data)

	if jsonMode {
		printJSON(struct {
			File      string         `json:"file"`
			Documents []lintDocument `json:"documents"`
		}{File: path, Documents: docs})
		for _, d := range docs {
			if !d.Valid {
				os.Exit(1)
			}
		}
		if len(docs) == 0 {
			os.Exit(1)
		}
		return
	}

	hadError := false
	for _, d := range docs {
		if d.Valid {
			fmt.Printf("%sdocument %d: valid YAML%s\n", colorGreen, d.Index, colorReset)
			continue
		}
		hadError = true
		fmt.Printf("%sdocument %d: %s%s\n", colorRed, d.Index, d.Error, colorReset)
	}

	if len(docs) == 0 {
		fmt.Printf("kube-why: %s contains no YAML documents\n", path)
		os.Exit(1)
	}

	fmt.Println()
	if hadError {
		fmt.Printf("%s failed syntax check.\n", path)
		os.Exit(1)
	}
	fmt.Printf("%s: all %d document(s) are syntactically valid.\n", path, len(docs))
}
