ARG GO_VER
ARG ALPINE_VER
ARG REGISTRY=harbor.wildberries.ru \
    WB_DOCKER_PROXY=docker-hub-proxy

# first build stage
# by default uses alpine-based image
FROM $REGISTRY/$WB_DOCKER_PROXY/golang:${GO_VER}-alpine${ALPINE_VER} as builder

# copy all source code for build
WORKDIR /src
COPY . .

ARG ALPINE_VER
ARG WB_MIRROR

# install updates and build executable
RUN apk upgrade --no-cache -X $WB_MIRROR/alpine/v${ALPINE_VER}/main && \
    apk add make && \
    make sgroups-tf-v2

# second release stage
ARG TERRAFORM_VER \
    REGISTRY=harbor.wildberries.ru \
    WB_DOCKER_PROXY=docker-hub-proxy \
    HOME=/root

FROM $REGISTRY/$WB_DOCKER_PROXY/hashicorp/terraform:$TERRAFORM_VER

COPY --from=builder /src/bin/terraform-provider-sgroups /tmp/terraform-provider-sgroups
COPY deploy/terraform-provider-sgroups-v2/version_setup.sh version_setup.sh

ARG CI_TAG
ARG CI_BRANCH
ARG CI_SHORT_SHA

RUN apk add --no-cache vault libcap && \
    setcap cap_ipc_lock= /usr/sbin/vault && \
    export CI_TAG=${CI_TAG} && \
    export CI_BRANCH=${CI_BRANCH} && \
    export CI_SHORT_SHA=${CI_SHORT_SHA} && \
    env | grep CI_ && \
    source version_setup.sh && \
    mkdir -p $HOME/.terraform.d/plugins/registry.terraform.io/wb/sgroups/$CUSTOM_VERSION/linux_amd64 && \
    cp /tmp/terraform-provider-sgroups $HOME/.terraform.d/plugins/registry.terraform.io/wb/sgroups/$CUSTOM_VERSION/linux_amd64/terraform-provider-sgroups_v$CUSTOM_VERSION && \
    rm -f /tmp/terraform-provider-sgroups

WORKDIR $HOME
COPY deploy/terraform-provider-sgroups-v2/.terraformrc .
