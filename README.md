# hgate

This is a proxy for interaction with the TokenD based systems, which performs the signature of the request headers for authentication.

## Install

1. Install https://github.com/golang/go/wiki/Ubuntu
2. Run `go get gitlab.com/tokend/hgate`
3. Run `go install gitlab.com/tokend/hgate/...`

## Usage

Run the hgate executable with argument `serve` to start proxy:

```bash
./hgate serve
```

Additional option:

- `--config [path/to/config] or -c [path/to/config]` default: "./config.yaml"

## Config
Config is a file in [YAML](https://en.wikipedia.org/wiki/YAML) format.

```yaml
port: 8842   #port where hgate will listen
log_level: warn   #level of log output
horizon_url: http://128.199.229.140:8011    #url of Horizon API
seed: SADB...   #your SecretSeed
```