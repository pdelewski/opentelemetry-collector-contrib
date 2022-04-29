// Copyright 2021 OpenTelemetry Authors
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
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type S3Exporter struct {
	config           config.Exporter
	metricTranslator metricTranslator
	logger           *zap.Logger
}

func NewS3MetricsExporter(
	config config.Exporter,
	params component.ExporterCreateSettings,
) (component.MetricsExporter, error) {
	if config == nil {
		return nil, errors.New("s3 exporter config is nil")
	}

	logger := params.Logger
	expConfig := config.(*Config)
	expConfig.logger = logger

	expConfig.Validate()

	s3Exporter := &S3Exporter{
		config:           config,
		metricTranslator: newMetricTranslator(*expConfig),
		logger:           logger,
	}

	return exporterhelper.NewMetricsExporter(
		config,
		params,
		s3Exporter.ConsumeMetrics)
}

func NewS3LogsExporter(
	config config.Exporter,
	params component.ExporterCreateSettings,
) (component.LogsExporter, error) {
	if config == nil {
		return nil, errors.New("s3 exporter config is nil")
	}

	logger := params.Logger
	expConfig := config.(*Config)
	expConfig.logger = logger

	expConfig.Validate()

	s3Exporter := &S3Exporter{
		config:           config,
		metricTranslator: newMetricTranslator(*expConfig),
		logger:           logger,
	}

	return exporterhelper.NewLogsExporter(
		config,
		params,
		s3Exporter.ConsumeLogs)
}

func NewS3TracesExporter(
	config config.Exporter,
	params component.ExporterCreateSettings,
) (component.TracesExporter, error) {
	if config == nil {
		return nil, errors.New("s3 exporter config is nil")
	}

	logger := params.Logger
	expConfig := config.(*Config)
	expConfig.logger = logger

	expConfig.Validate()

	s3Exporter := &S3Exporter{
		config:           config,
		metricTranslator: newMetricTranslator(*expConfig),
		logger:           logger,
	}

	return exporterhelper.NewTracesExporter(
		config,
		params,
		s3Exporter.ConsumeTraces)
}

func (e *S3Exporter) pushMetricsData(ctx context.Context, md pmetric.Metrics) error {
	rms := md.ResourceMetrics()
	labels := map[string]string{}

	for i := 0; i < rms.Len(); i++ {
		rm := rms.At(i)
		am := rm.Resource().Attributes()
		if am.Len() > 0 {
			am.Range(func(k string, v pcommon.Value) bool {
				labels[k] = v.StringVal()
				return true
			})
		}
	}

	e.logger.Info("Processing resource metrics", zap.Any("labels", labels))

	expConfig := e.config.(*Config)
	var parquetMetrics []*ParquetMetric

	for i := 0; i < rms.Len(); i++ {
		rm := rms.At(i)
		e.metricTranslator.translateOTelToParquetMetric(&rm, &parquetMetrics, expConfig)
	}
	e.writeParquet(parquetMetrics, ctx, expConfig.S3Uploader.S3Bucket,
		expConfig.S3Uploader.S3Prefix, expConfig.S3Uploader.S3Partition,
		expConfig.S3Uploader.FilePrefix, expConfig.FileFormat, expConfig.BatchCount)
	return nil

}

func (e S3Exporter) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func (e S3Exporter) Start(ctx context.Context, host component.Host) error {
	return nil
}

func (e S3Exporter) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	return e.pushMetricsData(ctx, md)
}

func (e S3Exporter) ConsumeLogs(ctx context.Context, logs plog.Logs) error {
	return nil
}

func (e S3Exporter) ConsumeTraces(ctx context.Context, traces ptrace.Traces) error {
	return nil
}

func (e S3Exporter) Shutdown(context.Context) error {
	return nil
}

// generate the s3 time key based on partition configuration
func (e S3Exporter) getTimeKey(partition string) string {
	var timeKey string
	t := time.Now()
	year, month, day := t.Date()
	hour, minute, _ := t.Clock()

	rand.Int()

	if partition == "hour" {
		timeKey = fmt.Sprintf("year=%d/month=%02d/day=%02d/hour=%02d/", year, month, day, hour)
	} else {
		timeKey = fmt.Sprintf("year=%d/month=%02d/day=%02d/hour=%02d/minute=%02d/", year, month, day, hour, minute)
	}

	e.logger.Info("Start processing resource metrics", zap.Any("timeKey", timeKey))
	return timeKey
}

func (e S3Exporter) getS3Key(bucket string, keyPrefix string, partition string, filePrefix string, fileformat string) string {
	timeKey := e.getTimeKey(partition)
	randomID := rand.Int()

	s3Key := bucket + "/" + keyPrefix + "/" + timeKey + "/" + filePrefix + "_" + strconv.Itoa(randomID) + "." + fileformat

	e.logger.Info("Start processing resource metrics", zap.Any("s3Key", s3Key))

	return s3Key
}
