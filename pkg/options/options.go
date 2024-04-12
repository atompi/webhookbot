package options

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Version string = "v1.1.0"

type MetricsOptions struct {
	Enable bool   `yaml:"enable"`
	Path   string `yaml:"path"`
}

type LogOptions struct {
	Path  string `yaml:"path"`
	Level string `yaml:"level"`
}

type CoreOptions struct {
	Mode    string         `yaml:"mode"`
	Listen  string         `yaml:"listen"`
	Metrics MetricsOptions `yaml:"metrics"`
	Log     LogOptions     `yaml:"log"`
}

type TemplateOptions struct {
	Alerting string `yaml:"alerting"`
	Resolved string `yaml:"resolved"`
}

type BotOptions struct {
	Path     string          `yaml:"path"`
	Webhook  string          `yaml:"webhookUrl"`
	Template TemplateOptions `yaml:"template"`
}

type Options struct {
	Core CoreOptions  `yaml:"core"`
	Bots []BotOptions `yaml:"bots"`
}

func NewOptions() (opts Options) {
	optsSource := viper.AllSettings()
	err := createOptions(optsSource, &opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "create options failed:", err)
		os.Exit(1)
	}
	return
}
