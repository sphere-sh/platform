package main

import (
	"os"
	"sphere_platform_universe"
	sphere_platform_universe_internal "sphere_platform_universe/internal"
)

var (
	// Universe - the universe.
	Universe *sphere_platform_universe.Universe

	// RootDirectory - the root directory of the universe.
	RootDirectory = "/etc/sphere/platform/universe"
)

func init() {

	// First, we initialize the universe's root directory.
	//
	if _, err := os.Stat(RootDirectory); os.IsNotExist(err) {
		if err := os.MkdirAll(RootDirectory, os.ModePerm); err != nil {
			panic("Failed to create root directory: " + err.Error())
		}
	}

	// Next, we set up the universes logging.
	//
	sphere_platform_universe_internal.SetupLogging(RootDirectory)

	// Finally, we turn on the lights of the universe.
	//
	Universe = &sphere_platform_universe.Universe{}
}

func main() {

}
