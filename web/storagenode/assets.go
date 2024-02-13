// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package storagenodeweb

import (
	"embed"
	"fmt"
	"io/fs"

	"storj.io/storj/storagenode"
)

//go:embed dist/*
var assets embed.FS

// Assets contains either the built storagenode from web/storagenode/dist directory or it is empty.
func init() {
	dist, err := fs.Sub(assets, "dist")
	if err != nil {
		panic(fmt.Errorf("invalid embedding: %w", err))
	}

	storagenode.Assets = dist
}
