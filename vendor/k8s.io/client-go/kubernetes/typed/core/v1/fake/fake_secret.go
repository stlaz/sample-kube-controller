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
	v1 "k8s.io/api/core/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	gentype "k8s.io/client-go/gentype"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// fakeSecrets implements SecretInterface
type fakeSecrets struct {
	*gentype.FakeClientWithListAndApply[*v1.Secret, *v1.SecretList, *corev1.SecretApplyConfiguration]
	Fake *FakeCoreV1
}

func newFakeSecrets(fake *FakeCoreV1, namespace string) typedcorev1.SecretInterface {
	return &fakeSecrets{
		gentype.NewFakeClientWithListAndApply[*v1.Secret, *v1.SecretList, *corev1.SecretApplyConfiguration](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("secrets"),
			v1.SchemeGroupVersion.WithKind("Secret"),
			func() *v1.Secret { return &v1.Secret{} },
			func() *v1.SecretList { return &v1.SecretList{} },
			func(dst, src *v1.SecretList) { dst.ListMeta = src.ListMeta },
			func(list *v1.SecretList) []*v1.Secret { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.SecretList, items []*v1.Secret) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}