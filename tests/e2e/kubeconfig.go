package e2e

import (
	"os"

	"k8s.io/client-go/rest"
	"k8s.io/kops/pkg/kubeconfig"
	"sigs.k8s.io/yaml"
)

func CreateKubeConfig(restConfig *rest.Config) (string, error) {
	c := &kubeconfig.KubectlConfig{
		Kind:       "Config",
		ApiVersion: "v1",
		Clusters: []*kubeconfig.KubectlClusterWithName{
			{
				Name: "default",
				Cluster: kubeconfig.KubectlCluster{
					Server:                   restConfig.Host,
					CertificateAuthorityData: restConfig.CAData,
				},
			},
		},
		Users: []*kubeconfig.KubectlUserWithName{
			{
				Name: "default",
				User: kubeconfig.KubectlUser{
					ClientCertificateData: restConfig.CertData,
					ClientKeyData:         restConfig.KeyData,
				},
			},
		},
		Contexts: []*kubeconfig.KubectlContextWithName{
			{
				Name: "default",
				Context: kubeconfig.KubectlContext{
					Cluster: "default",
					User:    "default",
				},
			},
		},
		CurrentContext: "default",
	}

	file, err := os.CreateTemp("/tmp", "kubeconfig")
	if err != nil {
		return "", err
	}
	defer file.Close()

	bytes, err := yaml.Marshal(c)
	if err != nil {
		return "", err
	}

	_, err = file.Write(bytes)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}
