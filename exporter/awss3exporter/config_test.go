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
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configcheck"
	"go.opentelemetry.io/collector/config/configtest"
)

func TestDefaultConfig(t *testing.T) {
	factories, err := componenttest.NopFactories()
	assert.Nil(t, err)

	factory := NewFactory()
	factories.Exporters[factory.Type()] = factory
	cfg, err := configtest.LoadConfigFile(
		t, path.Join(".", "testdata", "default.yaml"), factories,
	)
	require.NoError(t, err)
	require.NotNil(t, cfg)

	e := cfg.Exporters["awss3"]

	assert.Equal(t, e,
		&Config{
			ExporterSettings: config.NewExporterSettings(typeStr),

			S3Uploader: S3UploaderConfig{
				UploaderBufferSize:  1000000,
				UploaderConcurrency: 1,
				MaxRetries:          1,
				MaxBackoffSeconds:   10,
			},
			BatchSize:  500000,
			BatchCount: 1000,
		},
	)
}

func TestConfig(t *testing.T) {
	factories, err := componenttest.NopFactories()
	assert.Nil(t, err)

	factory := NewFactory()
	factories.Exporters[factory.Type()] = factory
	cfg, err := configtest.LoadConfigFile(
		t, path.Join(".", "testdata", "config.yaml"), factories,
	)

	require.NoError(t, err)
	require.NotNil(t, cfg)

	e := cfg.Exporters["awss3"]

	assert.Equal(t, e,
		&Config{
			ExporterSettings: config.NewExporterSettings(typeStr),

			S3Uploader: S3UploaderConfig{
				Region:              "us-east-1",
				S3Bucket:            "foo",
				S3Prefix:            "bar",
				UploaderBufferSize:  1000000,
				UploaderConcurrency: 1,
				MaxRetries:          1,
				MaxBackoffSeconds:   10,
			},
			BatchSize:  500000,
			BatchCount: 1000,
		},
	)
}

func TestConfigCheck(t *testing.T) {
	cfg := (NewFactory()).CreateDefaultConfig()
	assert.NoError(t, configcheck.ValidateConfig(cfg))
}
