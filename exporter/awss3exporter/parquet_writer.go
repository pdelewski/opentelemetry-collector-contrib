// Copyright 2021, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package awss3exporter

import (
	"context"
	"io/ioutil"

	"github.com/xitongsys/parquet-go-source/s3"
	"github.com/xitongsys/parquet-go/writer"
	"go.uber.org/zap"
)

// read input schema file and parse it to json structure
func (e S3Exporter) parseParquetInputSchema() (string, error) {
	content, err := ioutil.ReadFile("./schema/parquet_input_schema")
	if err != nil {
		e.logger.Error("Can't read parquet input schema", zap.Error(err))
	}

	return string(content), nil
}

// read output schema file
func (e S3Exporter) parseParquetOutputSchema() (string, error) {
	content, err := ioutil.ReadFile("./schema/parquet_output_schema")
	if err != nil {
		e.logger.Error("Can't read parquet output schema", zap.Error(err))
	}

	return string(content), nil
}

// pdata.Metrics needs to translate into parquetMetrics structure
// and it needs to match parquetOutputSchema
func (e S3Exporter) writeParquet(metrics []*ParquetMetric, ctx context.Context,
	bucket string, keyPrefix string, partition string,
	filePrefix string, fileformat string, batchWriterNum int64) {

	key := e.getS3Key(bucket, keyPrefix, partition, filePrefix, fileformat)

	// create new S3 file writer
	fw, err := s3.NewS3FileWriter(ctx, bucket, key, "", nil, nil)
	if err != nil {
		e.logger.Error("Can't create parquet file writer", zap.Error(err))
		return
	}

	// create new parquet file writer
	parquetOutputSchema, err := e.parseParquetOutputSchema()
	if err != nil {
		e.logger.Error("Can't parse parquet output schema", zap.Error(err))
		return
	}

	pw, err := writer.NewParquetWriter(fw, parquetOutputSchema, batchWriterNum)
	if err != nil {
		e.logger.Error("Can't create parquet writer", zap.Error(err))
		return
	}

	for _, v := range metrics {
		if err = pw.Write(v); err != nil {
			e.logger.Error("Error write parquet output")
		}
	}

	if err = pw.WriteStop(); err != nil {
		e.logger.Error("Parquet write stop error", zap.Error(err))
	}

	err = fw.Close()
	if err != nil {
		e.logger.Error("Error closing S3 file writer")
	}
	e.logger.Error("Write Finished")
}
