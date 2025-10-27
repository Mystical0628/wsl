package wsl

import (
	"fmt"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("wsl", Register)
}

func Register(settings any) (register.LinterPlugin, error) {
	// The configuration type will be map[string]any or []interface, it depends on your configuration.
	// You can use https://github.com/go-viper/mapstructure to convert map to struct.

	if settings != nil {
		s, err := register.DecodeSettings[PluginSettings](settings)
		if err != nil {
			return nil, err
		}

		return &PluginWSL{settings: &s}, nil
	}

	return &PluginWSL{}, nil
}

type PluginSettings struct {
	AllowFirstInBlock                bool     `mapstructure:"allow-first-in-block"`
	AllowWholeBlock                  bool     `mapstructure:"allow-whole-block"`
	MaxAllowedStatementsAboveBlock   int      `json:"max-allowed-statements-above-block"`
	MaxAllowedStatementsAboveIfBlock int      `json:"max-allowed-statements-above-if-block"`
	BranchMaxLines                   int      `mapstructure:"branch-max-lines"`
	CaseMaxLines                     int      `mapstructure:"case-max-lines"`
	Default                          string   `mapstructure:"default"`
	Enable                           []string `mapstructure:"enable"`
	Disable                          []string `mapstructure:"disable"`
}

func NewPluginSettings() PluginSettings {
	return PluginSettings{
		AllowFirstInBlock:                true,
		AllowWholeBlock:                  false,
		MaxAllowedStatementsAboveBlock:   1,
		MaxAllowedStatementsAboveIfBlock: 1,
		CaseMaxLines:                     0,
		BranchMaxLines:                   2,
	}
}

type PluginWSL struct {
	settings *PluginSettings
}

func (p *PluginWSL) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	var conf *Configuration

	if p.settings != nil {
		checkSet, err := NewCheckSet(p.settings.Default, p.settings.Enable, p.settings.Disable)
		if err != nil {
			return nil, fmt.Errorf("wsl: invalid check: %w", err)
		}

		conf = &Configuration{
			IncludeGenerated:                 true, // force to true because golangci-lint already has a way to filter generated files.
			AllowFirstInBlock:                p.settings.AllowFirstInBlock,
			AllowWholeBlock:                  p.settings.AllowWholeBlock,
			MaxAllowedStatementsAboveBlock:   p.settings.MaxAllowedStatementsAboveBlock,
			MaxAllowedStatementsAboveIfBlock: p.settings.MaxAllowedStatementsAboveIfBlock,
			BranchMaxLines:                   p.settings.BranchMaxLines,
			CaseMaxLines:                     p.settings.CaseMaxLines,
			Checks:                           checkSet,
		}
	}

	return []*analysis.Analyzer{
		NewAnalyzer(conf),
	}, nil
}

func (p *PluginWSL) GetLoadMode() string {
	// NOTE: the mode can be `register.LoadModeSyntax` or `register.LoadModeTypesInfo`.
	// - `register.LoadModeSyntax`: if the linter doesn't use types information.
	// - `register.LoadModeTypesInfo`: if the linter uses types information.

	return register.LoadModeTypesInfo
}
