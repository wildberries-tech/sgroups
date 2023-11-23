ARG GO_VER
ARG ALPINE_VER

# first build stage
# by default uses alpine-based image
FROM golang:${GO_VER}-alpine${ALPINE_VER} as builder

# copy all source code for build
WORKDIR /src
COPY . .

ARG ALPINE_VER
ARG VERSION
ARG WB_MIRROR

# install updates and build executable
RUN apk upgrade --no-cache -X $WB_MIRROR/alpine/v${ALPINE_VER}/main && \
    apk add make && \
    make sgroups-tf-v2

# second release stage
ARG APP_VER \
    HOME=/root \
    TERRAFORM_VER=1.6.3

FROM $REGISTRY/$WB_DOCKER_PROXY/hashicorp/terraform:$TERRAFORM_VER

RUN apk add --no-cache vault libcap && \
    setcap cap_ipc_lock= /usr/sbin/vault

ARG PLUGIN_PATH=$HOME/.terraform.d/plugins/registry.terraform.io/wb/sgroups/$APP_VER/linux_amd64
COPY --from=builder /app/bin/terraform-provider-sgroups $PLUGIN_PATH/terraform-provider-sgroups_v2$APP_VER
WORKDIR $HOME
COPY deploy/terraform-provider-sgroups-v2/.terraformrc .
