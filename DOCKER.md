# Docker usage

If you want to run it in docker, just clone the repository

```sh
git clone https://github.com/eyedeekay/about.i2p
```

Edit the `env.sh` file to contain a cookie `authkey`

```sh
export AUTH_KEY="secure key"
```

and run `./docker.sh`

```sh
./docker.sh
```