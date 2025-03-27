package releases

import (
	"errors"
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/goccy/go-json"
	"io"
	"net/http"
	"strings"
)

var (
	ErrVersionNotFound     = errors.New("node version not found")
	ErrInvalidVersion      = errors.New("invalid version string")
	ErrInvalidReleaseIndex = errors.New("invalid release index")
)

func SearchVersion(ver string, releaseIndexURL string) (release *NodeRelease, err error) {

	if len(releaseIndexURL) == 0 { // Should I replace this with something else?
		releaseIndexURL = "https://nodejs.org/download/release/index.json"
	}

	getReleaseResp, err := http.Get(releaseIndexURL)
	if err != nil {
		return nil, err
	}
	defer getReleaseResp.Body.Close()

	decoder := json.NewDecoder(getReleaseResp.Body)

	var releaseIndex []NodeRelease
	err = decoder.Decode(&releaseIndex)
	if err != nil && err != io.EOF {
		return nil, err
	}

	ltsMap := make(map[string]NodeRelease)

	for _, nodeRelease := range releaseIndex {
		if isLTS, codename := nodeRelease.IsLTS(); isLTS {
			ltsMap[strings.ToLower(codename)] = nodeRelease
		}
	}

	latestLts := (*semver.Version)(nil)

	for codename, ltsRelease := range ltsMap {
		ltsVerStr, err := semver.NewVersion(fmt.Sprintf("%d-%s", ltsRelease.Version.Major(), codename))
		if err != nil {
			return nil, err
		}

		if latestLts != nil {
			if latestLts.GreaterThan(ltsVerStr) {
				latestLts = ltsVerStr
			}
		} else {
			latestLts = ltsVerStr
		}
	}

	if latestLts == nil {
		//TODO: Replace with proper alerts.
		panic("Could not find latest LTS, Things are broke")
	}

	ltsMap["lts"] = ltsMap[latestLts.Prerelease()]
	ver = strings.ToLower(ver)

	if release, doesLtsExist := ltsMap[ver]; doesLtsExist {
		return &release, nil
	}

	return nil, ErrVersionNotFound
}
