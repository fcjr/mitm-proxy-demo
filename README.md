# mitm-proxy-demo

Proxies all network requests and logs the request urls to stdout.

## Building

```sh
make build
```

## Usage

Node: requires running with `sudo`

Flags:

1) --autoconfig

    Automatically configure system settings with the proxy

2) --port <int> (default 9000)

    Change listening port