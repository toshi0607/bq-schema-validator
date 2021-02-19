# bq-schema-validator

## Description

bq-schema-validator is a CLI tool to detect the log entry that causes a BigQuery table schema error via log sink.

## Usage

```console
$ kubectl kogs POD_NAME [CONTAINER_NAME] | bq-schema-validator -project GCP_PROJECT_ID -dataset BIGQUERY_DATASET_ID -table TABLE_ID

$ bq-schema-validator -project GCP_PROJECT_ID -dataset BIGQUERY_DATASET_ID -table TABLE_ID -file FILE_PATH_FOR_CONTAINER_LOGS
```

### Options

```
  -dataset string
        [Required] A name of BigQuery dataset
  -file string
        [Optional] A file path
  -ignore string
        [Optional] Ignore field when comparing log and schema
  -project string
        [Required] A name of GCP project
  -table string
        [Required] A name of BigQuery table
  -target string
        [Optional] Target field when comparing log and schema
```

## Reference

Data types in BigQuery schema

| Name | Data Type | Description |
| -----| --------- | ----------- |
| Integer	| INT64	| Numeric values without fractional components |
| Floating point	| FLOAT64	| Approximate numeric values with fractional components |
| Numeric	| NUMERIC	| Exact numeric values with fractional components |
| BigNumeric (Preview)	| BIGNUMERIC	| Exact numeric values with fractional components |
| Boolean	| BOOL	| TRUE or FALSE (case insensitive) |
| String	| STRING	| Variable-length character (Unicode) data |
| Bytes	| BYTES	| Variable-length binary data |
| Date	| DATE	| A logical calendar date |
| Date/Time	| DATETIME	| A year, month, day, hour, minute, second, and subsecond |
| Time	| TIME	| A time, independent of a specific date |
| Timestamp	| TIMESTAMP	| An absolute point in time, with microsecond precision |
| Struct (Record)	| STRUCT	| Container of ordered fields each with a type (required) and field name (optional) |
| Geography	| GEOGRAPHY	| A pointset on the Earth's surface (a set of points, lines and polygons on the WGS84 reference spheroid, with geodesic edges) |

* [Specifying a schema](https://cloud.google.com/bigquery/docs/schemas)
* [BigQuery table schema in Go](https://github.com/googleapis/google-cloud-go/blob/master/bigquery/schema.go#L158)
* [Overview of logs exports](https://cloud.google.com/logging/docs/export)