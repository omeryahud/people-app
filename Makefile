# Current Operator version
VERSION ?= 0.0.1
# Default bundle image tag
BUNDLE_IMG ?= people-operator:$(VERSION)
# Options for 'bundle-build'
ifneq ($(origin CHANNELS), undefined)
BUNDLE_CHANNELS := --channels=$(CHANNELS)
endif
ifneq ($(origin DEFAULT_CHANNEL), undefined)
BUNDLE_DEFAULT_CHANNEL := --default-channel=$(DEFAULT_CHANNEL)
endif
BUNDLE_METADATA_OPTS ?= $(BUNDLE_CHANNELS) $(BUNDLE_DEFAULT_CHANNEL)

# Image URL to use all building/pushing image targets
REGISTRY     ?= quay.io/omeryahud
IMG_TAG      ?= devel
IMG          ?= ${REGISTRY}/people-app-operator:${IMG_TAG}
FRONTEND_IMG ?= ${REGISTRY}/people-app-frontend:${IMG_TAG}
BACKEND_IMG  ?= ${REGISTRY}/people-app-backend:${IMG_TAG}
DATABASE_IMG ?= ${REGISTRY}/people-app-database:${IMG_TAG}

# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true"

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

all: manager frontend backend database

# Run tests
test: generate fmt vet manifests
	go test ./... -coverprofile cover.out

test-frontend:
	echo Testing frontend

test-backend:
	echo Testing backend

test-database:
	echo Testing database

test-all: test test-frontend test-backend test-database

# Build manager binary
manager: generate fmt vet
	go mod tidy
	go mod vendor
	go build --mod=vendor -o bin/manager main.go

frontend: fmt vet
	go mod tidy
	go mod vendor
	go build --mod=vendor -o bin/frontend cmd/frontend/main.go

backend: fmt vet
	go mod tidy
	go mod vendor
	go build --mod=vendor -o bin/backend cmd/backend/main.go

database: fmt vet
	go mod tidy
	go mod vendor
	go build --mod=vendor -o bin/databse cmd/database/main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet manifests
	go run ./main.go

# Install CRDs into a cluster
install: manifests kustomize
	$(KUSTOMIZE) build config/crd | kubectl apply -f -

# Uninstall CRDs from a cluster
uninstall: manifests kustomize
	$(KUSTOMIZE) build config/crd | kubectl delete -f -

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests kustomize
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	$(KUSTOMIZE) build config/default | kubectl apply -f -

undeploy: manifests kustomize
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	$(KUSTOMIZE) build config/default | kubectl delete -f -

deploy-install : deploy install

undeploy-install: uninstall undeploy

delete: manifests kustomize
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	$(KUSTOMIZE) build config/default | kubectl delete -f -

# Generate manifests e.g. CRD, RBAC etc.
manifests: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Generate code
generate: controller-gen
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

# Build the docker image
docker-build: test
	docker build . -t ${IMG}

# Build the frontend docker image
docker-build-frontend: test-frontend
	docker build . -t ${FRONTEND_IMG} -f cmd/frontend/Dockerfile

# Build the backend docker image
docker-build-backend: test-backend
	docker build . -t ${BACKEND_IMG} -f cmd/backend/Dockerfile

# Build the database docker image
docker-build-database: test-database
	docker build . -t ${DATABASE_IMG} -f cmd/database/Dockerfile

# Build all docker images
docker-build-all: docker-build docker-build-frontend docker-build-backend docker-build-database

# Push the docker image
docker-push:
	docker push ${IMG}

docker-push-frontend:
	docker push ${FRONTEND_IMG}

docker-push-backend:
	docker push ${BACKEND_IMG}

docker-push-database:
	docker push ${DATABASE_IMG}

docker-push-all: docker-push docker-push-frontend docker-push-backend docker-push-database

# find or download controller-gen
# download controller-gen if necessary
controller-gen:
ifeq (, $(shell which controller-gen))
	@{ \
	set -e ;\
	CONTROLLER_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$CONTROLLER_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.3.0 ;\
	rm -rf $$CONTROLLER_GEN_TMP_DIR ;\
	}
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

kustomize:
ifeq (, $(shell which kustomize))
	@{ \
	set -e ;\
	KUSTOMIZE_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$KUSTOMIZE_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/kustomize/kustomize/v3@v3.5.4 ;\
	rm -rf $$KUSTOMIZE_GEN_TMP_DIR ;\
	}
KUSTOMIZE=$(GOBIN)/kustomize
else
KUSTOMIZE=$(shell which kustomize)
endif

# Generate bundle manifests and metadata, then validate generated files.
bundle: manifests
	operator-sdk generate kustomize manifests -q
	kustomize build config/manifests | operator-sdk generate bundle -q --overwrite --version $(VERSION) $(BUNDLE_METADATA_OPTS)
	operator-sdk bundle validate ./bundle

# Build the bundle image.
bundle-build:
	docker build -f bundle.Dockerfile -t $(BUNDLE_IMG) .
