package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	t.Run("if the file doesn't exist", func(t *testing.T) {
		conf, err := loadConfig("non-existing-file")
		assert.Nil(t, conf, "should return nil if the file doesn't exist")
		assert.Error(t, err, "should return the error if the file doesn't exist")
	})
	t.Run("if the json can't be decoded", func(t *testing.T) {
		conf, err := loadConfig("config_invalid_test.json")
		assert.Nil(t, conf, "should return nil if the json can't be decoded")
		assert.Error(t, err, "should return the error if the json can't be decoded")
	})
	t.Run("if the config can be loaded", func(t *testing.T) {
		conf, err := loadConfig("config_test.json")
		require.Nil(t, err, "should return no error if the config can be loaded")
		assert.Equal(t, Config{
			Csv: CsvConfig{
				Delimiter: ",",
			},
			Sampling: SamplingConfig{
				Mod:      1,
				IDColumn: 0,
			},
			Actions: []ActionConfig{
				ActionConfig{
					Name: "hash",
				},
				ActionConfig{
					Name: "outcode",
				},
				ActionConfig{
					Name: "year",
					DateConfig: DateConfig{
						Format: "20060102",
					},
				},
				ActionConfig{
					Name: "nothing",
				},
			},
		}, *conf, "should return the config properly decoded")
	})
}
