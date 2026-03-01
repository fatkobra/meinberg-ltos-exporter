package main

import (
	"github.com/prometheus/common/version"
)

var (
	Version   = "snapshot"
	Revision  = "unknown"
	Branch    = "unknown"
	BuildUser = "unknown"
	BuildDate = "unknown"
)

func init() {
	version.Version = Version
	version.Revision = Revision
	version.Branch = Branch
	version.BuildUser = BuildUser
	version.BuildDate = BuildDate
}
