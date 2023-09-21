package main

import (
	_ "net/http/pprof"

	"github.com/nlandolfi/distribution/v3/registry"
	_ "github.com/nlandolfi/distribution/v3/registry/auth/htpasswd"
	_ "github.com/nlandolfi/distribution/v3/registry/auth/silly"
	_ "github.com/nlandolfi/distribution/v3/registry/auth/token"
	_ "github.com/nlandolfi/distribution/v3/registry/proxy"
	_ "github.com/nlandolfi/distribution/v3/registry/storage/driver/azure"
	_ "github.com/nlandolfi/distribution/v3/registry/storage/driver/filesystem"
	_ "github.com/nlandolfi/distribution/v3/registry/storage/driver/gcs"
	_ "github.com/nlandolfi/distribution/v3/registry/storage/driver/inmemory"
	_ "github.com/nlandolfi/distribution/v3/registry/storage/driver/middleware/cloudfront"
	_ "github.com/nlandolfi/distribution/v3/registry/storage/driver/middleware/redirect"
	_ "github.com/nlandolfi/distribution/v3/registry/storage/driver/s3-aws"
)

func main() {
	registry.RootCmd.Execute()
}
