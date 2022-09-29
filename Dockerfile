FROM golang
ADD . /usr/src/about.i2p
WORKDIR /usr/src/about.i2p
RUN adduser --home /home/user --gecos "user,,,," --disabled-password user
RUN git clone https://github.com/eyedeekay/cerca ../../cblgh/cerca && \
    go build && \
    cp about.i2p /usr/bin/about.i2p
WORKDIR /home/user/about.i2p
CMD about.i2p