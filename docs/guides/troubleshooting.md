# Troubleshooting

This page describes how to troubleshoot and gather information from the ExternalDNS Hetzner webhook.

## Version

We recommend running the latest version of the ExternalDNS and the Hetzner webhook.

To get the version of your ExternalDNS deployment and the webhook, you can run the following command:

```shell
kubectl get deployments -l app.kubernetes.io/name=external-dns -o yaml | grep 'image:'
```

## Logs

The webhook runs as a sidecar to ExternalDNS. To extract the logs only for the webhook you can run the following command:

```shell
kubectl logs -l app.kubernetes.io/name=external-dns -c webhook
```

To get a combined output for ExternalDNS and the webhook you can run the following command:

```shell
kubectl logs -l app.kubernetes.io/name=external-dns --all-containers --prefix
```

You may increase the log level by setting the `LOG_LEVEL` environment variable to `debug` via the Helm value `provider.webhook.env`. See the [configuration reference for more details](../reference/configuration.md).
