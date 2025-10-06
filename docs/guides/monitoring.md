# Monitoring

This page describes how to set up monitoring for the ExternalDNS Hetzner webhook.

## Metrics

By default, the ExternalDNS Helm chart creates a Kubernetes service exposing the metrics on port `8080`.

The ExternalDNS Helm chart supports configuring a [ServiceMonitor](https://prometheus-operator.dev/docs/api-reference/api/#monitoring.coreos.com/v1.ServiceMonitor) object via the Helm
value `provider.webhook.serviceMonitor`. You can find information about this in the Helm charts default [`values.yaml`](https://github.com/kubernetes-sigs/external-dns/blob/master/charts/external-dns/values.yaml#L167).

In the following snippet an additional `ServiceMonitor` resource with default values is created:

```yaml
# values.yml
# [...]
provider:
  webhook:
    serviceMonitor:
      enable: true
# [...]
```
