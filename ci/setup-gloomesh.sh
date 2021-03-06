#!/bin/bash -ex

#####################################
#
# Set up gloo mesh in the target kind cluster.
#
#####################################

cluster=$1
glooMeshChart=$2
agentChart=$3
agentImage=$4
apiServerAddress=$5

PROJECT_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )/.."

if [ "${cluster}" == "" ]; then
  cluster=mgmt-cluster
fi

K="kubectl --context kind-${cluster}"

echo "deploying gloo-mesh to ${cluster} from local images..."

## build and load GlooMesh docker images
MAKE="make -C $PROJECT_ROOT"
eval "${MAKE} manifest-gen package-helm build-all-images -B"

INSTALL_DIR="${PROJECT_ROOT}/install/"
DEFAULT_MANIFEST="${INSTALL_DIR}/gloo-mesh-default.yaml"

# load GlooMesh discovery and networking images
grep "image:" "${DEFAULT_MANIFEST}" \
  | awk '{print $3}' \
  | while read -r image; do
  kind load docker-image --name "${cluster}" "${image}"
done
# load cert-agent image
kind load docker-image --name "${cluster}" "${agentImage}"

## install to kube

go run "${PROJECT_ROOT}/cmd/meshctl/main.go" install \
  --kubecontext kind-"${cluster}" \
  --chart-file "${glooMeshChart}" \
  --namespace gloo-mesh \
  --register \
  --cluster-name "${cluster}" \
  --verbose  \
  --api-server-address "${apiServerAddress}" \
  --cert-agent-chart-file "${agentChart}"


${K} -n gloo-mesh rollout status deployment networking
${K} -n gloo-mesh rollout status deployment discovery

echo setup successfully set up gloo-mesh
