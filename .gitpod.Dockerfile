FROM eldius/gitpod-go1.14-workspace:latest

RUN echo "deb https://deb.torproject.org/torproject.org bionic main" | sudo tee -a /etc/apt/sources.list.d/tor.list > /dev/null && \
    echo "deb-src https://deb.torproject.org/torproject.org bionic main" | sudo tee -a /etc/apt/sources.list.d/tor.list > /dev/null && \
    sudo curl https://deb.torproject.org/torproject.org/A3C4F0F979CAA22CDBA8F512EE8CBC9E886DDD89.asc | gpg --import && \
    sudo gpg --export A3C4F0F979CAA22CDBA8F512EE8CBC9E886DDD89 | apt-key add - && \
    sudo apt update && \
    sudo apt install -y tor deb.torproject.org-keyring && \
    sudo rm -rf /var/lib/apt/lists/*

USER gitpod

# Install custom tools, runtime, etc. using apt-get
# For example, the command below would install "bastet" - a command line tetris clone:
#
# RUN sudo apt-get -q update && #     sudo apt-get install -yq bastet && #     sudo rm -rf /var/lib/apt/lists/*
#
# More information: https://www.gitpod.io/docs/config-docker/
