/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta3 "k8s.io/api/flowcontrol/v1beta3"
	flowcontrolv1beta3 "k8s.io/client-go/applyconfigurations/flowcontrol/v1beta3"
	gentype "k8s.io/client-go/gentype"
	typedflowcontrolv1beta3 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta3"
)

// fakeFlowSchemas implements FlowSchemaInterface
type fakeFlowSchemas struct {
	*gentype.FakeClientWithListAndApply[*v1beta3.FlowSchema, *v1beta3.FlowSchemaList, *flowcontrolv1beta3.FlowSchemaApplyConfiguration]
	Fake *FakeFlowcontrolV1beta3
}

func newFakeFlowSchemas(fake *FakeFlowcontrolV1beta3) typedflowcontrolv1beta3.FlowSchemaInterface {
	return &fakeFlowSchemas{
		gentype.NewFakeClientWithListAndApply[*v1beta3.FlowSchema, *v1beta3.FlowSchemaList, *flowcontrolv1beta3.FlowSchemaApplyConfiguration](
			fake.Fake,
			"",
			v1beta3.SchemeGroupVersion.WithResource("flowschemas"),
			v1beta3.SchemeGroupVersion.WithKind("FlowSchema"),
			func() *v1beta3.FlowSchema { return &v1beta3.FlowSchema{} },
			func() *v1beta3.FlowSchemaList { return &v1beta3.FlowSchemaList{} },
			func(dst, src *v1beta3.FlowSchemaList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta3.FlowSchemaList) []*v1beta3.FlowSchema { return gentype.ToPointerSlice(list.Items) },
			func(list *v1beta3.FlowSchemaList, items []*v1beta3.FlowSchema) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}