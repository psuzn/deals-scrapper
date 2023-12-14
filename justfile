set dotenv-load

default:
    @just --list

namespace:= "deals-extractor"
deployment:= "deals-extractor"
deployerServiceAccountName:= "deals-extractor-deployer"
tag := "$(git rev-parse --short HEAD)"
imageRepo := "ghcr.io/psuzn/deals-scrapper"

build-push-image imageTag=tag:
    #!/usr/bin/env sh
    tags="{{tag}},latest"
    docker build -t {{imageRepo}}:$tags

    echo Pushed tags $tags

ns-create:
    kubectl create namespace {{ namespace }}

ns-delete:
    kubectl delete namespace {{ namespace }}

# Uninstalls the helm chart
helm-delete:
    helm uninstall -n {{namespace}} {{deployment}}

# Runs helm upgrade
helm-upgrade imageTag=tag:
    #!/usr/bin/env sh
    echo "Deploying {{imageRepo}}:{{imageTag}} to $(kubectl config current-context)"
    helm upgrade {{deployment}} --create-namespace \
        --install --namespace {{namespace}} ./deployment/helm/backend \
        --set image.tag={{imageTag}} \
        --set image.repository={{imageRepo}}
        --set serverUrl=$SERVER_URL
        --set urls=$URLS

# creates a service account and token for a deployer
helm-create-deployer:
    echo "Creating deployer service account to $(kubectl config current-context)"
    helm install {{deployment}}-deployer ./deployment/helm/service-account \
        --namespace {{namespace}}  --create-namespace \
        --set serviceAccountName={{deployerServiceAccountName}} \
        --set namespace={{namespace}}

# deletes the deployer service account
helm-delete-deployer:
    echo "Deleting deployer service account to $(kubectl config current-context)"
    helm uninstall {{deployment}}-deployer --namespace {{namespace}}

# gets the deployer kubeconfig
deployer-cubeconfig:
    #!/usr/bin/env sh
    CLUSTER_NAME=$(kubectl config current-context)
    SECRET_NAME="sa-{{ deployerServiceAccountName }}-token"
    SA_TOKEN=$(kubectl get secret $SECRET_NAME -n {{namespace}}  -o jsonpath='{.data.token}' | base64 -D)
    CA_DATA=$(kubectl get secret $SECRET_NAME -n {{namespace}} -o jsonpath='{.data.ca\.crt}')
    K8S_ENDPOINT=$(kubectl config view -o jsonpath="{.clusters[?(@.name=='${CLUSTER_NAME}')].cluster.server}")
    echo "
    apiVersion: v1
    kind: Config
    clusters:
      - name: default-cluster
        cluster:
          certificate-authority-data: ${CA_DATA}
          server: ${K8S_ENDPOINT}
    contexts:
    - name: default-context
      context:
        cluster: default-cluster
        namespace: {{namespace}}
        user: default-user
    current-context: default-context
    users:
    - name: default-user
      user:
        token: ${SA_TOKEN}
    "