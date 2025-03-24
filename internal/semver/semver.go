package semver

import (
	semverMaster "github.com/Masterminds/semver/v3"
	"strings"
)

type Version struct {
	*semverMaster.Version
	originalInput string
}

func (v *Version) UnmarshalJSON(data []byte) error {
	dataStr := strings.Trim(string(data), "\"")
	v.originalInput = dataStr

	// TODO: Possible conforming of version string to fit specs for semver package.
	/*
		So some of the versions for the libraries that NodeJS uses, like V8 and ZLib, don't necessarily follow semver,
		but they are pretty close into fitting the specs for our package here.
	*/
	parsedVer, _ := semverMaster.NewVersion(dataStr)
	v.Version = parsedVer

	return nil
}
