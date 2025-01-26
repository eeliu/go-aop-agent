FROM golang:1.23.1
USER root 
## install deve`lopmental env
# ENV TZ=Asia/Shanghai
# ENV LC_CTYPE=en_US.UTF-8
# ENV LC_ALL=en_US.UTF-8
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
ENV GOPATH=/home/pinpoint/go
ENV GOROOT=/usr/local/go
ENV PATH=$PATH:$GOPATH/bin:$GOROOT/bin

