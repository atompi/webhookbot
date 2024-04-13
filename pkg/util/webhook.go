package util

import (
	"bytes"
	"html/template"

	"go.uber.org/zap"
)

type AlertLabelsStruct struct {
	Alertname any `json:"alertname"`
	Endpoint  any `json:"endpoint"`
	Instance  any `json:"instance"`
	Pod       any `json:"pod"`
	Container any `json:"container"`
	Severity  any `json:"severity"`
}

type AlertAnnotationsStruct struct {
	Summary            any `json:"summary"`
	Description        any `json:"description"`
	RunbookUrl         any `json:"runbook_url"`
	GrafanaValueString any `json:"__value_string__"`
}

type AlertStruct struct {
	Status       any                    `json:"status"`
	Labels       AlertLabelsStruct      `json:"labels"`
	Annotations  AlertAnnotationsStruct `json:"annotations"`
	StartsAt     any                    `json:"startsAt"`
	EndsAt       any                    `json:"endsAt"`
	Fingerprint  any                    `json:"fingerprint"`
	GeneratorURL any                    `json:"generatorURL"`
}

type CommonLabelsStruct struct {
	Namespace any `json:"namespace"`
}

type AlertsGroupStruct struct {
	Status       any                `json:"status"`
	Alerts       []AlertStruct      `json:"alerts"`
	CommonLabels CommonLabelsStruct `json:"commonLabels"`
}

func GenPostJsonData(structData AlertsGroupStruct, tmplFilePath string) (jsonData string, err error) {
	t, err := template.ParseFiles(tmplFilePath)
	if err != nil {
		zap.L().Sugar().Errorf("parse %s failed: %v", tmplFilePath, err)
		return
	}
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, structData)
	if err != nil {
		zap.L().Sugar().Errorf("generate json data failed: %v", err)
		return
	}
	jsonData = buffer.String()
	return
}
