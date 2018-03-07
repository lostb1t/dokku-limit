FROM ubuntu:14.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get -y install gcc git build-essential wget ruby-dev ruby1.9.1 lintian rpm help2man man-db
RUN command -v fpm > /dev/null || sudo gem install fpm --no-ri --no-rdoc

ARG GOLANG_VERSION

RUN wget -qO /tmp/go${GOLANG_VERSION}.linux-amd64.tar.gz https://storage.googleapis.com/golang/go${GOLANG_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf /tmp/go${GOLANG_VERSION}.linux-amd64.tar.gz \
    && cp /usr/local/go/bin/* /usr/local/bin

ARG WORKDIR=/go/src/github.com/dokku/dokku

WORKDIR ${WORKDIR}

COPY Makefile ${WORKDIR}/
COPY *.mk ${WORKDIR}/

RUN make deb-setup rpm-setup sshcommand plugn

COPY . ${WORKDIR}

ARG PLUGIN_MAKE_TARGET
ARG DOKKU_VERSION=master
ARG DOKKU_GIT_REV
ARG IS_RELEASE=false

ENV GOPATH=/go

RUN PLUGIN_MAKE_TARGET=${PLUGIN_MAKE_TARGET} \
    DOKKU_VERSION=${DOKKU_VERSION} \
    DOKKU_GIT_REV=${DOKKU_GIT_REV} \
    IS_RELEASE=${IS_RELEASE} \
    SKIP_GO_CLEAN=true \
    make version copyfiles \
    && rm -rf plugins/common/*.go  plugins/common/glide*  plugins/common/vendor/ \
    && make deb-herokuish deb-dokku deb-plugn deb-sshcommand deb-sigil deb-dokku-update \
            rpm-herokuish rpm-dokku rpm-plugn rpm-sshcommand rpm-sigil rpm-dokku-update

RUN mkdir -p /data \
    && cp /tmp/*.deb /data \
    && cp /tmp/*.rpm /data \
    && ls -lha /data/
