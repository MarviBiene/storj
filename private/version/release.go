// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package version

import _ "unsafe" // needed for go:linkname

//go:linkname buildTimestamp storj.io/private/version.buildTimestamp
var buildTimestamp string = "1611966133"

//go:linkname buildCommitHash storj.io/private/version.buildCommitHash
var buildCommitHash string = "ae4b97023c6e604c34bf59b007306ac5303a01a7"

//go:linkname buildVersion storj.io/private/version.buildVersion
var buildVersion string = "v1.21.2"

//go:linkname buildRelease storj.io/private/version.buildRelease
var buildRelease string = "true"

// ensure that linter understands that the variables are being used.
func init() { use(buildTimestamp, buildCommitHash, buildVersion, buildRelease) }

func use(...interface{}) {}
