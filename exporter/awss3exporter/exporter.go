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
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type S3Exporter struct {
	config config.Exporter
	logger *zap.Logger
}

func New(
	config config.Exporter,
	params component.ExporterCreateSettings,
) (component.MetricsExporter, error) {
	fmt.Println("S3Exporter")
	if config == nil {
		return nil, errors.New("s3 exporter config is nil")
	}

	logger := params.Logger

	s3Exporter := S3Exporter{
		config: config,
		logger: logger,
	}

	return s3Exporter, nil
}

func NewS3Exporter(
	config config.Exporter,
	params component.ExporterCreateSettings,
) (component.MetricsExporter, error) {
	exp, err := New(config, params)
	if err != nil {
		return nil, err
	}

	return exporterhelper.NewMetricsExporter(
		config,
		params,
		exp.(*S3Exporter).pushMetricsData,
		nil,
		exporterhelper.WithShutdown(exp.(*S3Exporter).Shutdown),
	)
}

func (e S3Exporter) pushMetricsData(ctx context.Context, md pmetric.Metrics) error {
	labels := map[string]string{}

	e.logger.Info("Processing resource metrics", zap.Any("labels", labels))

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
