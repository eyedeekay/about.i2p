FROM golang
ADD . /usr/src/about.i2p
WORKDIR /usr/src/about.i2p
RUN adduser --home /home/user --gecos "user,,,," --disabled-password user && \
    mkdir -p /home/user/about.i2p && \
    chown -R user /home/user/about.i2p
RUN git clone https://github.com/eyedeekay/cerca ../../cblgh/cerca && \
    go build && \
    cp about.i2p /usr/bin/about.i2p
USER user
WORKDIR /home/user/about.i2p
CMD about.i2p