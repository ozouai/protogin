FROM gitpod/workspace-base
LABEL dazzle/layer=tool-docker
LABEL dazzle/test=tests/tool-docker.yaml
USER root
ENV TRIGGER_REBUILD=2
# https://docs.docker.com/engine/install/ubuntu/
RUN curl -o /var/lib/apt/dazzle-marks/docker.gpg -fsSL https://download.docker.com/linux/ubuntu/gpg \
    && apt-key add /var/lib/apt/dazzle-marks/docker.gpg \
    && add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" \
    && install-packages docker-ce=5:19.03.15~3-0~ubuntu-focal docker-ce-cli=5:19.03.15~3-0~ubuntu-focal containerd.io

RUN curl -o /usr/bin/slirp4netns -fsSL https://github.com/rootless-containers/slirp4netns/releases/download/v1.1.11/slirp4netns-$(uname -m) \
    && chmod +x /usr/bin/slirp4netns

RUN curl -o /usr/local/bin/docker-compose -fsSL https://github.com/docker/compose/releases/download/1.29.2/docker-compose-Linux-x86_64 \
    && chmod +x /usr/local/bin/docker-compose

# https://github.com/wagoodman/dive
RUN curl -o /tmp/dive.deb -fsSL https://github.com/wagoodman/dive/releases/download/v0.10.0/dive_0.10.0_linux_amd64.deb \
    && apt install /tmp/dive.deb \
    && rm /tmp/dive.deb

WORKDIR /
ENV HOME="/root"
RUN apt-get update
RUN wget -q -c https://golang.org/dl/go1.18.2.linux-amd64.tar.gz -O- | tar -C /usr/local -xzvf -
ENV PATH="/usr/local/go/bin:${PATH}"
RUN go version
RUN curl -L -o /tmp/protoc.zip "https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip" && (cd /tmp && unzip protoc.zip -d /usr/local) && rm /tmp/protoc.zip
RUN protoc --version
ENV HOME="/home/gitpod"
WORKDIR /tmp/
USER gitpod
COPY tools /tmp/tools
WORKDIR /tmp/tools
RUN sudo chmod -R 777 /tmp/tools
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN go install golang.org/x/tools/gopls
RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs
RUN go install github.com/ramya-rao-a/go-outline
RUN go install github.com/go-delve/delve/cmd/dlv
RUN go install github.com/go-delve/delve/cmd/dlv@master
RUN go install honnef.co/go/tools/cmd/staticcheck

ENV PATH="/home/gitpod/go/bin:${PATH}"