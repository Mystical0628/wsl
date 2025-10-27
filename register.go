package wsl

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("wsl", Register)
}

func Register(settings any) (register.LinterPlugin, error) {
	// The configuration type will be map[string]any or []interface, it depends on your configuration.
	// You can use https://github.com/go-viper/mapstructure to convert map to struct.

	s, err := register.DecodeSettings[Configuration](settings)
	if err != nil {
		return nil, err
	}

	return &PluginWSL{config: s}, nil
}

type PluginWSL struct {
	config Configuration
}

func (p *PluginWSL) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		NewAnalyzer(&p.config),
	}, nil
}

func (p *PluginWSL) GetLoadMode() string {
	// NOTE: the mode can be `register.LoadModeSyntax` or `register.LoadModeTypesInfo`.
	// - `register.LoadModeSyntax`: if the linter doesn't use types information.
	// - `register.LoadModeTypesInfo`: if the linter uses types information.

	return register.LoadModeTypesInfo
}
