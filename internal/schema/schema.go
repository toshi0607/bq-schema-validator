package schema

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
)

func Schema(ctx context.Context, projectID, datasetID, tableID string) (map[string]string, error) {
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := client.Close(); err != nil {
			fmt.Println("failed to close the bq client")
		}
	}()

	meta, err := client.Dataset(datasetID).Table(tableID).Metadata(ctx)
	if err != nil {
		return nil, err
	}

	s := make(map[string]string, len(meta.Schema))
	fs := []*bigquery.FieldSchema(meta.Schema)
	for _, v := range fs {
		if v.Name == "jsonPayload" {
			fss := []*bigquery.FieldSchema(v.Schema)
			for _, v := range fss {
				s[v.Name] = string(v.Type)
			}
		}
	}
	return s, nil
}
