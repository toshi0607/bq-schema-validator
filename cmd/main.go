package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"toshi0607/bq-schema-validator/v2/internal/config"
	"toshi0607/bq-schema-validator/v2/internal/shcema"
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

	s, err := schema.Schema(ctx, c.GCPProjectID, c.DatasetID, c.TableID)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to get the schema, error: %v", err))
		return exitError
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
	fmt.Printf("current schema for %s dataset: %s\n", c.DatasetID, s)
	fmt.Println("-----------------------------")
	for {
		var result map[string]interface{}
		if err := d.Decode(&result); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Skip the non structured log")
			continue
		}
		ds, ok := cmp.Diff(result, s, c.Ignore, c.Target)
		if ok {
			fmt.Println(ds)
			fmt.Println("-----------------------------")
			fmt.Printf("Log line: %s\n", result)
			fmt.Println("-----------------------------")
		}
	}

	return exitOK
}
