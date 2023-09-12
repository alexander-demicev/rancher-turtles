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
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cluster name mapping", func() {
	It("Should only suffix rancher cluster name with -capi if management cluster name is not provided", func() {
		name := Name("some-cluster").ToRancherName("")
		Expect(name).To(Equal("some-cluster-capi"))
	})

	It("Should prefix rancher cluster name with management cluster name -capi", func() {
		name := Name("some-cluster").ToRancherName("mgmt-cluster-1")
		Expect(name).To(Equal("mgmt-cluster-1-some-cluster-capi"))
	})

	It("Should only add prefix and suffix once", func() {
		name := Name("some-cluster").ToRancherName("")
		name = Name(name).ToRancherName("mgmt-cluster-1")
		Expect(string(name)).To(Equal("mgmt-cluster-1-some-cluster-capi"))
	})

	It("Should remove suffix from rancher cluster if management cluster name is not provided", func() {
		name := Name("some-cluster").ToRancherName("")
		name = Name(name).ToCapiName("")
		Expect(string(name)).To(Equal("some-cluster"))
	})

	It("Should remove prefix and suffix from rancher cluster", func() {
		name := Name("some-cluster").ToRancherName("mgmt-cluster-1")
		name = Name(name).ToCapiName("mgmt-cluster-1")
		Expect(string(name)).To(Equal("some-cluster"))
	})

	It("Should remove suffix from rancher cluster only if it is present", func() {
		name := Name("some-cluster").ToCapiName("")
		Expect(string(name)).To(Equal("some-cluster"))
	})
})

func TestNameConverter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test naming convention")
}
