# webhookbot

Integration of Lark-robot for Prometheus Alertmanager via webhook.

## webhook raw json

- firing

```json
{
  "receiver": "WebhookAlert",
  "status": "firing",
  "alerts": [
    {
      "status": "firing",
      "labels": {
        "alertname": "etcdDatabaseHighFragmentationRatio",
        "endpoint": "http-metrics",
        "instance": "172.22.0.2:2381",
        "job": "kube-etcd",
        "namespace": "kube-system",
        "pod": "etcd-kind-control-plane",
        "prometheus": "monitoring/kube-prometheus-stack-prometheus",
        "service": "kube-prometheus-stack-kube-etcd",
        "severity": "warning"
      },
      "annotations": {
        "description": "etcd cluster \"kube-etcd\": database size in use on instance 172.22.0.2:2381 is 21.03% of the actual allocated disk space, please run defragmentation (e.g. etcdctl defrag) to retrieve the unused fragmented disk space.",
        "runbook_url": "https://etcd.io/docs/v3.5/op-guide/maintenance/#defragmentation",
        "summary": "etcd database size in use is less than 50% of the actual allocated storage."
      },
      "startsAt": "2022-12-14T01:32:49.278Z",
      "endsAt": "0001-01-01T00:00:00Z",
      "generatorURL": "http://kube-prometheus-stack-prometheus.monitoring:9090/graph?g0.expr=%28last_over_time%28etcd_mvcc_db_total_size_in_use_in_bytes%5B5m%5D%29+%2F+last_over_time%28etcd_mvcc_db_total_size_in_bytes%5B5m%5D%29%29+%3C+0.5&g0.tab=1",
      "fingerprint": "bc6e3ddcbe61b398"
    }
  ],
  "groupLabels": {
    "namespace": "kube-system"
  },
  "commonLabels": {
    "alertname": "etcdDatabaseHighFragmentationRatio",
    "endpoint": "http-metrics",
    "instance": "172.22.0.2:2381",
    "job": "kube-etcd",
    "namespace": "kube-system",
    "pod": "etcd-kind-control-plane",
    "prometheus": "monitoring/kube-prometheus-stack-prometheus",
    "service": "kube-prometheus-stack-kube-etcd",
    "severity": "warning"
  },
  "commonAnnotations": {
    "description": "etcd cluster \"kube-etcd\": database size in use on instance 172.22.0.2:2381 is 21.03% of the actual allocated disk space, please run defragmentation (e.g. etcdctl defrag) to retrieve the unused fragmented disk space.",
    "runbook_url": "https://etcd.io/docs/v3.5/op-guide/maintenance/#defragmentation",
    "summary": "etcd database size in use is less than 50% of the actual allocated storage."
  },
  "externalURL": "http://kube-prometheus-stack-alertmanager.monitoring:9093",
  "version": "4",
  "groupKey": "{}:{namespace=\"kube-system\"}",
  "truncatedAlerts": 0
}
```

- resolved

```json
{
  "receiver": "WebhookAlert",
  "status": "resolved",
  "alerts": [
    {
      "status": "resolved",
      "labels": {
        "alertname": "etcdDatabaseHighFragmentationRatio",
        "endpoint": "http-metrics",
        "instance": "172.22.0.2:2381",
        "job": "kube-etcd",
        "namespace": "kube-system",
        "pod": "etcd-kind-control-plane",
        "prometheus": "monitoring/kube-prometheus-stack-prometheus",
        "service": "kube-prometheus-stack-kube-etcd",
        "severity": "warning"
      },
      "annotations": {
        "description": "etcd cluster \"kube-etcd\": database size in use on instance 172.22.0.2:2381 is 20.99% of the actual allocated disk space, please run defragmentation (e.g. etcdctl defrag) to retrieve the unused fragmented disk space.",
        "runbook_url": "https://etcd.io/docs/v3.5/op-guide/maintenance/#defragmentation",
        "summary": "etcd database size in use is less than 50% of the actual allocated storage."
      },
      "startsAt": "2022-12-14T01:32:49.278Z",
      "endsAt": "2022-12-14T02:13:49.278Z",
      "generatorURL": "http://kube-prometheus-stack-prometheus.monitoring:9090/graph?g0.expr=%28last_over_time%28etcd_mvcc_db_total_size_in_use_in_bytes%5B5m%5D%29+%2F+last_over_time%28etcd_mvcc_db_total_size_in_bytes%5B5m%5D%29%29+%3C+0.5&g0.tab=1",
      "fingerprint": "bc6e3ddcbe61b398"
    }
  ],
  "groupLabels": {
    "namespace": "kube-system"
  },
  "commonLabels": {
    "alertname": "etcdDatabaseHighFragmentationRatio",
    "endpoint": "http-metrics",
    "instance": "172.22.0.2:2381",
    "job": "kube-etcd",
    "namespace": "kube-system",
    "pod": "etcd-kind-control-plane",
    "prometheus": "monitoring/kube-prometheus-stack-prometheus",
    "service": "kube-prometheus-stack-kube-etcd",
    "severity": "warning"
  },
  "commonAnnotations": {
    "description": "etcd cluster \"kube-etcd\": database size in use on instance 172.22.0.2:2381 is 20.99% of the actual allocated disk space, please run defragmentation (e.g. etcdctl defrag) to retrieve the unused fragmented disk space.",
    "runbook_url": "https://etcd.io/docs/v3.5/op-guide/maintenance/#defragmentation",
    "summary": "etcd database size in use is less than 50% of the actual allocated storage."
  },
  "externalURL": "http://kube-prometheus-stack-alertmanager.monitoring:9093",
  "version": "4",
  "groupKey": "{}:{namespace=\"kube-system\"}",
  "truncatedAlerts": 0
}
```

## Build

```
docker build -t hub.github.com/atompi/webhookbot:v1.1.0 .
```

## Deploy

```
mkdir -p conf/tmpl
# prepare conf/webhookbot.yaml and conf/tmpl/ (copy from examples/tmpl) folder
docker-compose up -d
```
