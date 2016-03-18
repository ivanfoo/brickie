# brickie
Dependencies free docker CLI tool. Useful when you can't install the docker client locally but you have access to a remote docker host (CI systems, for instance)

## Usage

brickie will use DOCKER_* environment variables to change its behaviour. To use a remote Docker host securized with TLS:

```
export DOCKER_HOST="tcp://remote_host:port
export DOCKER_VERIFY_TLS=1
export DOCKER_CERT_PATH=<CERTS_PATH>
```

You can use the your local Docker dameon instead if you don't set up any of this.

**LOGIN**

`brickie login --username=foo --password=XXXX --email=foo@bar.com --registry=baz.io`

**BUILD**

`brickie build --name=baz.io/rab/foo:v1`

**PUSH**

`brickie push --image=baz.io/rab/foo:v1 --registry=baz.io`
