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
    make sgroups-tf

# second release stage
FROM alpine:${ALPINE_VER}

ARG ALPINE_VER
ARG WB_MIRROR

# create user other than root and install updated
RUN addgroup -g 101 app && \
    adduser -H -u 101 -G app -s /bin/sh -D app && \
    apk update --no-cache -X $WB_MIRROR/alpine/v${ALPINE_VER}/main && \
    apk upgrade --no-cache -X $WB_MIRROR/alpine/v${ALPINE_VER}/main

# place all necessary executables and other files into /app directory
WORKDIR /app/
COPY --from=builder --chown=app:app /src .

# run container as new non-root user
USER app

CMD ["/app/bin/terraform-provider-sgroups"]
