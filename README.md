# bq-schema-validator

[![ci](https://github.com/toshi0607/bq-schema-validator/actions/workflows/ci.yaml/badge.svg)](https://github.com/toshi0607/bq-schema-validator/actions/workflows/ci.yaml)

## Description

bq-schema-validator is a CLI tool to detect the log entry that causes a BigQuery table schema error via log sink.

BigQuery creates a schema (field, type, etc) in a table when a log router sends container logs to the table for the first time. If the type of field sent to the table is different from original one, a log sink error takes place. For example, a log including `log_id` field which is string type is sent and then type is changed to int in the application. When the log is sent, the error occurs due to a type mismatch, so we need to detect the change. But, we only can see the error [via an email and an activity stream](https://cloud.google.com/logging/docs/export/configure_export_v2?hl=en#troubleshooting) when it happens for the first time so that it's hard to find.

To solve this issue, this CLI detects the log of which field type is different from the existing type in the schema. As fsr as I know, the schema mismatch usually happens at a top level field, so nested fields aren't checked.


## Prerequisites

* Permission to view the container log
* Permission to view the table schema (`roles/bigquery.dataViewer` for the GCP project)
* `gcloud auth application-default login`

## Usage

```console
$ kubectl logs POD_NAME [CONTAINER_NAME] | bq-schema-validator -project GCP_PROJECT_ID -dataset BIGQUERY_DATASET_ID -table TABLE_ID

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

## Installation

Download the binary from [GitHub Releases][https://github.com/toshi0607/bq-schema-validator/releases] and drop it in your `$PATH`.

- [Darwin / Mac][https://github.com/toshi0607/bq-schema-validator/releases/latest]
- [Linux][https://github.com/toshi0607/bq-schema-validator/releases/latest]

## License

[MIT][./LICENSE]

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
