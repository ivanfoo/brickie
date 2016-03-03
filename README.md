# brickie
docker CLI tool to build and push images using a remote docker host

## Usage

To use a remote Docker host securized with TLS:

```
export DOCKER_HOST="tcp://remote_host:port
export DOCKER_VERIFY_TLS=1
export DOCKER_CERT_PATH=<CERTS_PATH>
```

Also, brickie will use the docker's config.json file for auth purpose (if needed)

You can use the your local Docker dameon instead if you don't set up any of this

**BUILD**

`brickie build --name=foo.bar/baz/rab`

**PUSH**

`brickie push --image=foo.bar/baz/rab --registry=foo.bar`
