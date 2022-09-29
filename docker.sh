#! /usr/bin/env sh

wd=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
git pull --all
## AUTH_KEY variable is in this file, edit it
. "${wd}/env.sh"
if [ "${AUTH_KEY}" = "0" ]; then
    if [ ! -d ../../cblgh/cerca ]; then
        git clone https://github.com/eyedeekay/cerca ../../cblgh/cerca
    fi
    if [ ! -f "${wd}/about.i2p" ]; then
        go build
    fi
    export AUTH_KEY=$(./about.i2p -genauthkey)
    echo "$AUTH_KEY"
    exit 0
fi
docker build -t eyedeekay/about.i2p .
docker rm -f about.i2p
mkdir -p "${HOME}/abouti2p"
docker run -d --net=host --restart=always --user=user --name=about.i2p --volume="${HOME}/abouti2p":/home/user/about.i2p eyedeekay/about.i2p about.i2p --authkey="${AUTH_KEY}"
docker logs -f about.i2p