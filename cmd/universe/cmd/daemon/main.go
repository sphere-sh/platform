package main

import (
	"sphere_platform_universe"
	sphere_platform_universe_internal "sphere_platform_universe/internal"
)

var (
	// Universe - the universe.
	Universe *sphere_platform_universe.Universe
)

func init() {

	// Configure the logging
	//
	sphere_platform_universe_internal.SetupLogging()

	Universe = &sphere_platform_universe.Universe{}
}

func main() {

}
