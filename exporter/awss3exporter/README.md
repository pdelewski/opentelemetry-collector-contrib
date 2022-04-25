# AWS S3 Exporter for OpenTelemetry Collector
This exporter converts OpenTelemetry metrics to supported format and upload to S3

## Schema supported
This exporter targets to support parquet/ORC/json format

## Exporter Configuration

The following exporter configuration parameters are supported. 

| Name                   | Description                                                                        | Default |
| :--------------------- | :--------------------------------------------------------------------------------- | ------- |
| `region`               | AWS region.                                                                        |         |
| `s3_bucket`            | S3 bucket                                                                          |         |
| `s3_prefix`            | prefix for the S3 key.                                                             |         |
| `s3_partition`         | time granularity of S3 key: hour or minute                                         |"minute" |
| `uploader_buffer_size` | S3 upload buffer size                                                              |  1MB    |
| `uploader_concurrency` | S3 upload buffer size concurrency                                                  |   1     |
| `max_retries`          | max retries for upload                                                             |   1     |
| `max_backoff_seconds`  | max backoff seconds before next retry                                              |   10    |
| `batch_size`           | max backoff seconds before next retry                                              |  500KB  |
| `batch_count`          | max backoff seconds before next retry                                              |  1000   |

## AWS Credential Configuration

This exporter follows default credential resolution for the
[aws-sdk-go](https://docs.aws.amazon.com/sdk-for-go/api/index.html).

Follow the [guidelines](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html) for the
credential configuration.
