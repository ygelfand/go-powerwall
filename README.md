# go-powerwall

## Overview

`go-powerwall` is a utility designed to act as an api proxy to the tesla energy device protobuf private api.

## Starting Powerwall Proxy Server

### Usage
```bash
go-powerwall proxy [flags]
```

### Flags
- `-l, --listen string`: Specifies the host:port to listen on. Default is `:8080`.
- `-o, --ondemand`: Disables periodic refresh.
- `-r, --refresh uint32`: Specifies the periodic refresh frequency in seconds. Default is `30`.

### Global Flags
Global flags can be set using environment variables starting with `POWERWALL_`, e.g `POWERWALL_PASSWORD`
- `-e, --endpoint string`: Specifies the Powerwall endpoint URL. Default is `https://192.168.91.1/tedapi`.
- `-p, --password string`: Specifies the Powerwall installer password.

## Examples
- Start the Powerwall proxy server on the default port:
  ```bash
  go-powerwall proxy -p ABCDEFGHIJ
  ```

- Start the Powerwall proxy server on a custom port and disable periodic refresh:
  ```bash
  go-powerwall proxy :8888 -o -p ABCDEFGHIJ
  ```

## Notes
- Ensure that you have proper connectivity to the tedapi system when starting the proxy
