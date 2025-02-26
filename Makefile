# vars
IMAGE_REPO ?= local/local
NS ?= default

# const
VERSION = $(shell git rev-parse --short HEAD)
IMAGE_NAME = $(IMAGE_REPO)/sample-controller:$(VERSION)

build:
	go build -o _output/bin/sample-controller ./cmd

build-image:
	docker build -f ./artifacts/images/Dockerfile -t $(IMAGE_NAME) .

deploy-kind:
	kind load docker-image $(IMAGE_NAME)
	kubectl apply -f ./artifacts/crds/samplecontroller.k8s.io_foos.yaml
	sed -E "s#\{CONTROLLER_IMAGE}#$(IMAGE_NAME)#; s#\{NS\}#$(NS)#" ./artifacts/examples/deployment.yaml | kubectl apply -n $(NS) -f -

controller-gen:
	go build -o _output/bin/controller-gen ./vendor/sigs.k8s.io/controller-tools/cmd/controller-gen

update-crds: controller-gen
	./_output/bin/controller-gen crd paths="./api/..." output:crd:artifacts:config="artifacts/crds/"

update: update-crds
	./hack/update-codegen.sh
	./hack/verify-codegen.sh

clean:
	rm -rf _output/

