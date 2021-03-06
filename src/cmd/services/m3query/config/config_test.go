// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package config

import (
	"testing"

	"github.com/m3db/m3/src/query/models"
	xconfig "github.com/m3db/m3x/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/validator.v2"
	yaml "gopkg.in/yaml.v2"
)

func TestTagOptionsFromEmptyConfig(t *testing.T) {
	cfg := TagOptionsConfiguration{}
	opts, err := TagOptionsFromConfig(cfg)
	require.NoError(t, err)
	require.NotNil(t, opts)
	assert.Equal(t, []byte("__name__"), opts.MetricName())
}

func TestTagOptionsFromConfig(t *testing.T) {
	name := "foobar"
	cfg := TagOptionsConfiguration{
		MetricName: name,
	}
	opts, err := TagOptionsFromConfig(cfg)
	require.NoError(t, err)
	require.NotNil(t, opts)
	assert.Equal(t, []byte(name), opts.MetricName())
}

func TestConfigLoading(t *testing.T) {
	var cfg Configuration
	require.NoError(t, xconfig.LoadFile(&cfg, "./testdata/sample_config.yml", xconfig.Options{}))

	assert.Equal(t, &LimitsConfiguration{
		MaxComputedDatapoints: 12000,
	}, &cfg.Limits)
	// TODO: assert on more fields here.
}

func TestConfigValidation(t *testing.T) {
	baseCfg := func(t *testing.T) *Configuration {
		var cfg Configuration
		require.NoError(t, xconfig.LoadFile(&cfg, "./testdata/sample_config.yml", xconfig.Options{}),
			"sample configuration is no longer valid or loadable. Fix it up to provide a base config here")

		return &cfg
	}

	// limits configuration
	limitsCfgCases := []struct {
		Name  string
		Limit int64
	}{{
		Name:  "empty LimitsConfiguration is valid (implies disabled)",
		Limit: 0,
	}, {
		Name:  "LimitsConfiguration with positive limit is valid",
		Limit: 5,
	}, {}, {
		Name:  "LimitsConfiguration with negative limit is valid (implies disabled)",
		Limit: -5,
	}}

	for _, tc := range limitsCfgCases {
		t.Run(tc.Name, func(t *testing.T) {
			cfg := baseCfg(t)
			cfg.Limits = LimitsConfiguration{
				MaxComputedDatapoints: 5,
			}

			assert.NoError(t, validator.Validate(cfg))
		})
	}
}

func TestDefaultTagOptionsConfig(t *testing.T) {
	var cfg TagOptionsConfiguration
	require.NoError(t, yaml.Unmarshal([]byte(""), &cfg))
	opts, err := TagOptionsFromConfig(cfg)
	require.NoError(t, err)
	assert.Equal(t, []byte("__name__"), opts.MetricName())
	assert.Equal(t, models.TypeLegacy, opts.IDSchemeType())
}

func TestTagOptionsConfig(t *testing.T) {
	var cfg TagOptionsConfiguration
	config := "metricName: abcdefg\nidScheme: prepend_meta"
	require.NoError(t, yaml.Unmarshal([]byte(config), &cfg))
	opts, err := TagOptionsFromConfig(cfg)
	require.NoError(t, err)
	assert.Equal(t, []byte("abcdefg"), opts.MetricName())
	assert.Equal(t, models.TypePrependMeta, opts.IDSchemeType())
}
