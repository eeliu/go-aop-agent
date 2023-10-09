FROM ubuntu:20.04
USER root 
## install developmental env
ENV TZ=Asia/Shanghai
ENV LC_CTYPE=en_US.UTF-8
ENV LC_ALL=en_US.UTF-8
# referr from https://code.visualstudio.com/remote/advancedcontainers/add-nonroot-user
ARG USERNAME=pinpoint
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Create the user
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    #
    # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# ********************************************************
# * Anything else you want to do like clean up goes here *
# ********************************************************

# [Optional] Set the default user. Omit if you want to keep the default as root.
# USER $USERNAME


RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt update && apt install -y gcc g++ gdb  pkg-config  libxml2-dev libsqlite3-dev libcurl4-openssl-dev zlib1g-dev make cmake wget autoconf git clangd python3-pip clang-format protobuf-compiler

# install golang 1.18

RUN wget https://go.dev/dl/go1.18.10.linux-amd64.tar.gz && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.10.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin
ENV GO_PATH=~/go
ENV PATH=$PATH:/$GO_PATH/bin
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
