package backend

import (
	"context"
	"fmt"
	"strings"

	github "github.com/google/go-github/v44/github"
	"github.com/spf13/viper"
)

// Format the repository config source.
func FmtRepo() (string, string) {
	repo := viper.GetString("app.proton_source")
	parts := strings.Split(repo, "/")
	return parts[0], parts[1]
}

// Fetch all of the releases from the proton source.
func GetReleases() ([]*github.RepositoryRelease, error) {
	owner, repo := FmtRepo()
	client := github.NewClient(nil)

	releases, _, err := client.Repositories.ListReleases(context.Background(), owner, repo, nil)

	if err != nil {
		return nil, err
	}

	return releases, nil
}

// Fetch all of the data for a specified tag from the proton source.
func GetReleaseData(tag string) (*github.RepositoryRelease, error) {
	owner, repo := FmtRepo()
	client := github.NewClient(nil)

	release, _, err := client.Repositories.GetReleaseByTag(context.Background(), owner, repo, tag)

	if err != nil {
		return nil, err
	}

	return release, nil
}

// Sorts through a release to find valid assets for downloading a release.
func GetValidAssets(release *github.RepositoryRelease) (*github.ReleaseAsset, *github.ReleaseAsset, error) {
	var protonTar *github.ReleaseAsset
	var protonSum *github.ReleaseAsset

	for _, asset := range release.Assets {

		// Once we have both assets, we don't need to keep looking.
		if protonTar != nil && protonSum != nil {
			break
		}

		// Find the needed files for an installation
		if strings.HasSuffix(asset.GetName(), ".tar.gz") {
			protonTar = asset
		} else if strings.HasSuffix(asset.GetName(), ".sha512sum") {
			protonSum = asset
		}
	}

	// No proton tar found, cannot install.
	if protonTar == nil {
		return nil, nil, fmt.Errorf("unable to find a proton tarball")
	}

	// A proton tar was found, but there is no sha512sum file to verify it with.
	if protonSum == nil {
		return protonTar, nil, nil
	}

	// We have everything we need to do a proper installation.
	return protonTar, protonSum, nil
}
