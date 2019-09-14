# Android Web Debug Bridge

AWDB is an HTTP API wrapping the [Android Debug Bridge utility](https://developer.android.com/studio/command-line/adb).

It aims to make it easier to use remote Android devices for testing and debugging. Developers and CI tools can interact with devices on remote hosts through a simple HTTP interface, without the need to configure SSH access and port-forwarding.

## Installation

AWDB can be installed by calling `go get github.com/rightmesh/awdb`. It requires ADB to be installed.

Once installed, run AWDB with `$GOPATH/bin/awdb`, or just `awdb` if `$GOBIN` is on your `$PATH`.

**Note:** AWDB does not have any built-in security or authentication. Ensure any servers running AWDB are behind a firewall and/or reverse proxy with sufficient security features to prevent unauthorized access to your devices.

TODO: Pre-compiled release binaries.

## API

AWDB's API is documented with the OpenAPI specification [here](api/openapi-spec/awdb.yml). You can browse the documentation online [here](https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/RightMesh/awdb/master/api/openapi-spec/awdb.yml).
