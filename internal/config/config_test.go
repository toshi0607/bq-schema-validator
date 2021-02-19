package config

import (
	"flag"
	"testing"
)

func TestConfig_Init(t *testing.T) {
	testing.Init()
	tests := map[string]struct {
		args          map[string]string
		wantProjectID string
		wantDatasetID string
		wantTableID   string
		wantFile      string
		wantTarget    string
		wantIgnore    string
		wantError     bool
	}{
		"invalid args": {
			args: map[string]string{
				"file":    "test",
				"project": "test",
			},
			wantError: true,
		},
		"valid args": {
			args: map[string]string{
				"file":    "test1",
				"project": "test2",
				"dataset": "test3",
				"table":   "test4",
			},
			wantProjectID: "test2",
			wantDatasetID: "test3",
			wantTableID:   "test4",
			wantFile:      "test1",
			wantTarget:    "",
			wantIgnore:    "",
			wantError:     false,
		},
	}

	for name, te := range tests {
		te := te

		flag.CommandLine.Set("file", te.args["file"])
		flag.CommandLine.Set("project", te.args["project"])
		flag.CommandLine.Set("table", te.args["table"])
		flag.CommandLine.Set("dataset", te.args["dataset"])
		flag.CommandLine.Set("ignore", te.args["ignore"])
		flag.CommandLine.Set("target", te.args["target"])

		got, err := New().Init()
		if !te.wantError && err != nil {
			t.Error(err)
		}
		if te.wantError {
			continue
		}

		if got.GCPProjectID != te.wantProjectID {
			t.Errorf("[%s] got: %s, want: %s", name, got, te.wantProjectID)
		}
		if got.DatasetID != te.wantDatasetID {
			t.Errorf("[%s] got: %s, want: %s", name, got, te.wantDatasetID)
		}
		if got.TableID != te.wantTableID {
			t.Errorf("[%s] got: %s, want: %s", name, got, te.wantTableID)
		}
		if got.Target != te.wantTarget {
			t.Errorf("[%s] got: %s, want: %s", name, got, te.wantTarget)
		}
		if got.Target != te.wantIgnore {
			t.Errorf("[%s] got: %s, want: %s", name, got, te.wantIgnore)
		}
	}
}
