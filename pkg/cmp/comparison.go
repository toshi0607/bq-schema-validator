package cmp

import (
	"fmt"
	"reflect"
	"strings"
)

func Diff(logLine interface{}, bqschema map[string]string, ignore, target string) ([]string, bool) {
	var diffs []string
	exist := false
	m, ok := logLine.(map[string]interface{})
	if !ok {
		return []string{fmt.Sprintf("-------- orig %s", logLine)}, exist
	}

	for k, v := range m {
		if k == ignore || k != target && target != "" {
			continue
		}
		t := reflect.TypeOf(v).Name()
		if t == "" {
			t = "map[string]interface{}"
		}
		if !isSameType(t, bqschema[strings.ToLower(k)]) {
			exist = true
			str := fmt.Sprintf("！！！！type diff exists field name: %s, log value: %s, log type: %s, schema type: %s\n", k, v, t, bqschema[k])
			diffs = append(diffs, str)
		}
	}
	return diffs, exist
}

func isSameType(log, schema string) bool {
	// https://tour.golang.org/basics/11
	m := map[string]string{
		"string":                 "STRING",
		"map[string]interface{}": "RECORD",
		"int":                    "INTEGER",
		"int8":                   "INTEGER",
		"int16":                  "INTEGER",
		"int32":                  "INTEGER",
		"int64":                  "INTEGER",
		"uint":                   "INTEGER",
		"uint8":                  "INTEGER",
		"uint16":                 "INTEGER",
		"uint32":                 "INTEGER",
		"uint64":                 "INTEGER",
		"uintptr":                "INTEGER",
		"float64":                "FLOAT",
		"float32":                "FLOAT",
		"bool":                   "BOOLEAN",
	}

	if m[log] == schema {
		return true
	}
	return false
}
