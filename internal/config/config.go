package config

import (
	"errors"
	"flag"

	_ "testing" // for arranging import order for test
)

type Config struct {
	GCPProjectID string
	DatasetID    string
	TableID      string
	Ignore       string
	Target       string
	File         string
}

var (
	filePath  string
	projectID string
	datasetID string
	tableID   string
	target    string
	ignore    string
)

func init() {
	flag.StringVar(&filePath, "file", "", "[Optional] A file path")
	flag.StringVar(&projectID, "project", "", "[Required] A name of GCP project")
	flag.StringVar(&datasetID, "dataset", "", "[Required] A name of BigQuery dataset")
	flag.StringVar(&tableID, "table", "", "[Required] A name of BigQuery table")
	flag.StringVar(&ignore, "ignore", "", "[Optional] Ignore field when comparing log and schema")
	flag.StringVar(&target, "target", "", "[Optional] Target field when comparing log and schema")
}

func New() *Config {
	flag.Parse()
	return &Config{
		GCPProjectID: projectID,
		DatasetID:    datasetID,
		TableID:      tableID,
		File:         filePath,
		Target:       target,
		Ignore:       ignore,
	}
}

func (c *Config) Init() (*Config, error) {
	if c.GCPProjectID == "" || c.TableID == "" || c.DatasetID == "" {
		return nil, errors.New("missing args")
	}

	return c, nil
}
