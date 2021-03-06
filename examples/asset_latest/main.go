package main

import (
	"fmt"
	"github.com/dhillondeep/go-getrelease"
)

func main()  {
	client := getrelease.NewGithubClient(nil)

	// for a directory, the content is extracted inside the folder
	// regex can be used for names
	if assetName, err := getrelease.GetTagAsset(client, "./wio",
		"wio", "wio", "wio_.*_linux_64bit.tar.gz",
		"v0.9.0", getrelease.WithChecksum("asset:wio_.*_checksums.txt")); err != nil {
			panic(err)
	} else {
		fmt.Printf("Name of the asset is: %s", assetName)
	}
}
