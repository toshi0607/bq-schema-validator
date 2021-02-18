package main

import (
	"bufio"
	"bytes"
	"cloud.google.com/go/bigquery"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"toshi0607/bq-schema-validator/v2/internal/config"
	"toshi0607/bq-schema-validator/v2/pkg/cmp"
)

const (
	exitOK    = 0
	exitError = 1
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("recovered, error: %v", err)
			os.Exit(1)
		}
	}()
	os.Exit(realMain(os.Args))
}

func realMain(_ []string) int {
	ctx := context.Background()
	c, err := config.New().Init()
	if err != nil {
		flag.Usage()
		return exitError
	}

	client, err := bigquery.NewClient(ctx, c.GCPProjectID)
	if err != nil {
		fmt.Println(fmt.Errorf("bigquery.NewClient: %v", err))
		os.Exit(1)
	}
	defer client.Close()

	meta, err := client.Dataset(c.DatasetID).Table(c.TableID).Metadata(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	w := os.Stdout
	fmt.Fprintf(w, "Schema has %d top-level fields\n", len(meta.Schema))

	bqschema := make(map[string]string, len(meta.Schema))
	fs := []*bigquery.FieldSchema(meta.Schema)
	for _, v := range fs {
		if v.Name == "jsonPayload" {
			fss := []*bigquery.FieldSchema(v.Schema)
			for _, v := range fss {
				bqschema[v.Name] = string(v.Type)
			}
		}
	}

	var r io.Reader
	if c.File != "" {
		j, err := ioutil.ReadFile(c.File)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		r = bytes.NewReader(j)
	} else {
		r = bufio.NewReader(os.Stdin)
		bufio.NewReader(os.Stdin)
	}
	d := json.NewDecoder(r)

	fmt.Println("-----------------------------")
	fmt.Printf("current schema for %s dataset: %s\n", c.DatasetID, bqschema)
	fmt.Println("-----------------------------")
	for {
		var result map[string]interface{}
		if err := d.Decode(&result); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Skip the non structured log")
			continue
		}
		ds, ok := cmp.Diff(result, bqschema, c.Ignore, c.Target)
		if ok {
			fmt.Println(ds)
			fmt.Println("-----------------------------")
			fmt.Printf("Log line: %s\n", result)
			fmt.Println("-----------------------------")
		}
	}

	return exitOK
}
