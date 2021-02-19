package cmp

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"testing"
)

func TestDiff(t *testing.T) {
	t.Helper()
	tests := map[string]struct {
		file     string
		schema   map[string]string
		ignore   string
		target   string
		wantDiff bool
	}{
		"no diff": {
			file: "../../testdata/no_diff.json",
			schema: map[string]string{
				"level":   "STRING",
				"ts":      "FLOAT",
				"call_id": "STRING",
				"ff":      "RECORD",
			},
			wantDiff: false,
		},
		"unexpected log field": {
			file: "../../testdata/unexpected_log_field.json",
			schema: map[string]string{
				"level":   "STRING",
				"ts":      "FLOAT",
				"call_id": "STRING",
				"ff":      "RECORD",
			},
			wantDiff: true,
		},
		"ignore the specific field": {
			file: "../../testdata/unexpected_log_field.json",
			schema: map[string]string{
				"level":   "STRING",
				"ts":      "FLOAT",
				"call_id": "STRING",
				"ff":      "RECORD",
			},
			ignore:   "no_schema",
			wantDiff: false,
		},
		"focus on the specific field": {
			file: "../../testdata/unexpected_log_field.json",
			schema: map[string]string{
				"level":   "STRING",
				"ts":      "FLOAT",
				"call_id": "STRING",
				"ff":      "RECORD",
			},
			target:   "level",
			wantDiff: false,
		},
	}
	for name, te := range tests {
		te := te

		logr := Setup(t, te.file)
		for {
			var result map[string]interface{}
			if err := json.NewDecoder(logr).Decode(&result); err == io.EOF {
				break
			} else if err != nil {
				continue
			}
			d, ok := Diff(result, te.schema, te.ignore, te.target)
			if ok != te.wantDiff {
				t.Fatalf("[%s] unexpected diff: %s", name, d)
			}
		}
	}

}

func Setup(t *testing.T, filePath string) io.Reader {
	j, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read the file, path: %s, error: %v", filePath, err)
	}
	return bytes.NewReader(j)
}
