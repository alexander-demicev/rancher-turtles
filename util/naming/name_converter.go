/*
Copyright 2023 SUSE.

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

package naming

import (
	"fmt"
	"strings"
)

var rancherCAPISuffix = "-capi"

// Name is a wrapper around CAPI/Rancher cluster names to simplify convertation between the two.
type Name string

// ToRancherName converts a CAPI cluster name to Rancher cluster name.
func (n Name) ToRancherName(managementClusterName string) string {
	if managementClusterName == "" {
		return fmt.Sprintf("%s%s", n.ToCapiName(managementClusterName), rancherCAPISuffix)

	}
	return fmt.Sprintf("%s-%s%s", managementClusterName, n.ToCapiName(managementClusterName), rancherCAPISuffix)
}

// ToCapiName converts a Rancher cluster name to CAPI cluster name.
func (n Name) ToCapiName(managementClusterName string) string {
	trimManagementClusterPrefix := strings.TrimPrefix(string(n), fmt.Sprintf("%s-", managementClusterName))
	return strings.TrimSuffix(string(trimManagementClusterPrefix), rancherCAPISuffix)
}
