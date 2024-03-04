// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package version

import _ "unsafe" // needed for go:linkname

//go:linkname buildTimestamp storj.io/common/version.buildTimestamp
var buildTimestamp string = "1709562045"

//go:linkname buildCommitHash storj.io/common/version.buildCommitHash
var buildCommitHash string = "9a7e61846ca99f0172f6ac37622515c3c333a0d5"

//go:linkname buildVersion storj.io/common/version.buildVersion
var buildVersion string = "v1.99.2"

//go:linkname buildRelease storj.io/common/version.buildRelease
var buildRelease string = "true"

// ensure that linter understands that the variables are being used.
func init() { use(buildTimestamp, buildCommitHash, buildVersion, buildRelease) }

func use(...interface{}) {}
