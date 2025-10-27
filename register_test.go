package wsl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/golangci/plugin-module-register/register"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestPluginWSLDefaultConfig(t *testing.T) {
	newPlugin, err := register.GetPlugin("wsl")
	require.NoError(t, err)

	plugin, err := newPlugin(NewPluginSettings())
	require.NoError(t, err)

	analyzers, err := plugin.BuildAnalyzers()
	require.NoError(t, err)

	dirs, err := os.ReadDir("./testdata/src/default_config")
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range dirs {
		t.Run(tc.Name(), func(t *testing.T) {
			testdata := analysistest.TestData()
			analyzer := analyzers[0]

			analysistest.RunWithSuggestedFixes(t, testdata, analyzer, filepath.Join("default_config", tc.Name()))
		})
	}
}
