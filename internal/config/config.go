package config

var Version string = "v1.0.0"

type MainConfigStruct struct {
	ListenAddress string          `yaml:"listenAddress"`
	GinMode       string          `yaml:"ginMode"`
	WebhookPath   string          `yaml:"webhookPath"`
	Log           LogConfigStruct `yaml:"log"`
}

type FeishuConfigStruct struct {
	WebhookUrl      string `yaml:"webhookUrl"`
	AlertMsgTmpl    string `yaml:"alertMsgTmpl"`
	ResolvedMsgTmpl string `yaml:"resolvedMsgTmpl"`
}

type LogConfigStruct struct {
	Path  string `yaml:"path"`
	Level string `yaml:"level"`
}

type RootConfigStruct struct {
	Main   MainConfigStruct   `yaml:"main"`
	Feishu FeishuConfigStruct `yaml:"feishu"`
}

type AlertStruct struct {
	Status      any `json:"status"`
	StartsAt    any `json:"startsAt"`
	EndsAt      any `json:"endsAt"`
	Fingerprint any `json:"fingerprint"`
}

type CommonLabelsStruct struct {
	Alertname any `json:"alertname"`
	Namespace any `json:"namespace"`
	Pod       any `json:"pod"`
	Severity  any `json:"severity"`
}

type CommonAnnotationsStruct struct {
	Description any `json:"description"`
	Summary     any `json:"summary"`
	RunbookUrl  any `json:"runbook_url"`
}

type AlertsGroupStruct struct {
	Status            any                     `json:"status"`
	Alerts            []AlertStruct           `json:"alerts"`
	CommonLabels      CommonLabelsStruct      `json:"commonLabels"`
	CommonAnnotations CommonAnnotationsStruct `json:"commonAnnotations"`
}
