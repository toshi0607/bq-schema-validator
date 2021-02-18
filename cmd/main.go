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
			fmt.Println(fmt.Errorf("failed to read the file, path: %s, error: %v", c.File, err))
			return exitError
		}
		r = bytes.NewReader(j)
	} else {
		r = bufio.NewReader(os.Stdin)
	}

	fmt.Println("-----------------------------")
	fmt.Printf("current schema for the %s dataset: %s\n", c.DatasetID, s)
	fmt.Println("-----------------------------")
	for {
		var result map[string]interface{}
		if err := json.NewDecoder(r).Decode(&result); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("skip the non structured log")
			continue
		}
		ds, ok := cmp.Diff(result, s, c.Ignore, c.Target)
		if ok {
			fmt.Println(ds)
			fmt.Println("-----------------------------")
			fmt.Printf("log line: %s\n", result)
			fmt.Println("-----------------------------")
		}
	}

	return exitOK
}
