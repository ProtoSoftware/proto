/*
Copyright © 2022 BitsOfAByte

*/
package shared

import (
	"context"
	"fmt"
	"os"
	"strings"

	github "github.com/google/go-github/v44/github"
	"github.com/spf13/viper"
)

// Format the the given index source into a valid format.
func FmtRepo(entryIndex int) (string, string) {

	sources := viper.GetStringSlice("app.sources")

	if len(sources) == 0 {
		panic("No sources found.")
	}

	split := strings.Split(sources[entryIndex], "/")

	return split[0], split[1]

}

// Get the source index the user is trying to install from.
func GetSourceIndex() int {
	sources := viper.GetStringSlice("app.sources")
	if len(sources) > 1 {
		Debug("GetSourceIndex: Found " + fmt.Sprintf("%d", len(sources)) + " sources.")
		Debug("GetSourceIndex: Prompting user for source.")
		fmt.Println("\nMultiple sources found. Which one do you want to use?")
		for i, source := range sources {
			fmt.Printf("%d. %s\n", i+1, source)
		}
		fmt.Println("0. Cancel")
		fmt.Print("Choice: ")
		var source int
		fmt.Scanf("%d", &source)

		if source == 0 {
			os.Exit(0)
		}

		if source < 1 || source > len(sources) {
			Debug("GetSourceIndex: User chose an invalid source.")
			return GetSourceIndex()
		}

		fmt.Println("")

		Debug("GetSourceIndex: User chose source: " + sources[source-1])
		return source - 1
	}

	return 0
}

// Fetch all of the releases from the proton source.
func GetReleases(entryIndex int) ([]*github.RepositoryRelease, error) {
	owner, repo := FmtRepo(entryIndex)
	client := github.NewClient(nil)

	releases, _, err := client.Repositories.ListReleases(context.Background(), owner, repo, nil)

	Debug("GetReleases: Found " + fmt.Sprintf("%d", len(releases)) + " releases for " + owner + "/" + repo)

	if err != nil {
		return nil, err
	}

	return releases, nil
}

// Fetch all of the data for a specified tag from the proton source.
func GetReleaseData(entryIndex int, tag string) (*github.RepositoryRelease, error) {
	owner, repo := FmtRepo(entryIndex)
	client := github.NewClient(nil)

	release, _, err := client.Repositories.GetReleaseByTag(context.Background(), owner, repo, tag)

	Debug("GetReleaseData: Looking for: " + owner + "/" + repo + "/" + tag)

	if err != nil {
		return nil, err
	}

	Debug("GetReleaseData: Found release " + release.GetTagName())

	return release, nil
}

// Total Asset size in bytes.
func GetTotalAssetSize(assets []*github.ReleaseAsset) int64 {
	var size int

	for _, asset := range assets {
		if strings.HasSuffix(asset.GetName(), ".tar.gz") {
			size += asset.GetSize()
		}

		if strings.HasSuffix(asset.GetName(), ".sha512sum") {
			size += asset.GetSize()
		}

		if strings.HasSuffix(asset.GetName(), ".tar.xz") {
			size += asset.GetSize()
		}
	}

	return int64(size)
}

// Sorts through a release to find valid assets for downloading a release.
func GetValidAssets(release *github.RepositoryRelease) (*github.ReleaseAsset, *github.ReleaseAsset, error) {
	var protonTar *github.ReleaseAsset
	var protonSum *github.ReleaseAsset

	for _, asset := range release.Assets {

		Debug("GetValidAssets: Validating asset: " + asset.GetName())

		// Once we have both assets, we don't need to keep looking.
		if protonTar != nil && protonSum != nil {
			Debug("GetValidAssets: Found both assets, finishing search.")
			break
		}

		// Find the needed files for an installation
		if strings.HasSuffix(asset.GetName(), ".tar.gz") {
			Debug("GetValidAssets: Found a valid tar.gz asset.")
			protonTar = asset
		} else if strings.HasSuffix(asset.GetName(), ".tar.xz") {
			Debug("GetValidAssets: Found a valid tar.xz asset.")
			protonTar = asset
		} else if strings.HasSuffix(asset.GetName(), ".sha512sum") {
			Debug("GetValidAssets: Found a valid sha512sum asset.")
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
