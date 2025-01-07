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
	v1 "k8s.io/api/batch/v1"
	batchv1 "k8s.io/client-go/applyconfigurations/batch/v1"
	gentype "k8s.io/client-go/gentype"
	typedbatchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
)

// fakeJobs implements JobInterface
type fakeJobs struct {
	*gentype.FakeClientWithListAndApply[*v1.Job, *v1.JobList, *batchv1.JobApplyConfiguration]
	Fake *FakeBatchV1
}

func newFakeJobs(fake *FakeBatchV1, namespace string) typedbatchv1.JobInterface {
	return &fakeJobs{
		gentype.NewFakeClientWithListAndApply[*v1.Job, *v1.JobList, *batchv1.JobApplyConfiguration](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("jobs"),
			v1.SchemeGroupVersion.WithKind("Job"),
			func() *v1.Job { return &v1.Job{} },
			func() *v1.JobList { return &v1.JobList{} },
			func(dst, src *v1.JobList) { dst.ListMeta = src.ListMeta },
			func(list *v1.JobList) []*v1.Job { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.JobList, items []*v1.Job) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}