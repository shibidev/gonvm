package releases

import "shibidev.xyz/apps/gonvm/internal/semver"

type NodeRelease struct {
	Version     *semver.Version
	ReleaseDate string `json:"date"`

	Files []string

	NpmVer     *semver.Version `json:"npm"`
	V8Ver      *semver.Version `json:"v8"`
	UvVer      *semver.Version `json:"uv"`
	ZLibVer    *semver.Version `json:"zlib"`
	OpenSSLVer *semver.Version `json:"openssl"`

	ModuleCount string `json:"modules"`

	Lts      interface{} `json:"lts"`
	Security bool        `json:"security"`
}

func (r *NodeRelease) IsLTS() (bool, string) {
	if ltsCodename, isLts := r.Lts.(string); isLts {
		return true, ltsCodename
	}

	return false, ""
}

func (r *NodeRelease) IsSecurity() bool {
	return r.Security
}

type NodeReleaseLibraries struct {
	NPM, V8, UV, ZLib, OpenSSL *semver.Version
}

func (r *NodeRelease) LibraryVersions() NodeReleaseLibraries {
	return NodeReleaseLibraries{
		NPM:     r.NpmVer,
		V8:      r.V8Ver,
		UV:      r.UvVer,
		ZLib:    r.ZLibVer,
		OpenSSL: r.OpenSSLVer,
	}
}
