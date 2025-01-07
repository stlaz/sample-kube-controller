build:
	go build -o _output/bin/sample-controller ./cmd

controller-gen:
	go build -o _output/bin/controller-gen ./vendor/sigs.k8s.io/controller-tools/cmd/controller-gen

update-crds: controller-gen
	./_output/bin/controller-gen crd paths="./api/..." output:crd:artifacts:config="artifacts/crds/"

update: update-crds
	./hack/update-codegen.sh
	./hack/verify-codegen.sh

clean:
	rm -rf _output/

