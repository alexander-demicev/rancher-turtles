/*
Copyright © 2025 SUSE LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package examples

import _ "embed"

var (
	//go:embed applications/azure/clusterresourceset-cloud-provider.yaml
	CAPIAzureCPI []byte

	//go:embed clusterclasses/azure/clusterclass-aks-example.yaml
	CAPIAzureAKSClusterclass []byte

	//go:embed clusterclasses/azure/clusterclass-rke2-example.yaml
	CAPIAzureRKE2Clusterclass []byte

	//go:embed applications/ccm/azure/helm-chart.yaml
	CAAPFAzureCCMHelmApp []byte

	//go:embed applications/cni/calico/helm-chart.yaml
	CAAPFCalicoCNIHelmApp []byte

	//go:embed clusterclasses/vsphere/clusterclass-rke2-example.yaml
	CAPIVSphereRKE2Clusterclass []byte
)
