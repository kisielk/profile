package profile_test

import (
	"os"

	"github.com/davecheney/profile"
)

func ExampleStart() {
	// start a simple CPU profile and register
	// a defer to Stop (flush) the profiling data.
	defer profile.Start().Stop()
}

func ExampleCPUProfile() {
	// CPU profiling is the default profiling mode, but you can specify it
	// explicitly for completeness.
	defer profile.Start(profile.MemProfile).Stop()
}

func ExampleMemProfile() {
	// use memory profiling, rather than the default cpu profiling.
	defer profile.Start(profile.MemProfile).Stop()
}

func ExampleProfilePath() {
	// set the location that the profile will be written to
	defer profile.Start(profile.ProfilePath(os.Getenv("HOME")))
}

func ExampleNoShutdownHook() {
	// disable the automatic shutdown hook.
	defer profile.Start(profile.NoShutdownHook).Stop()
}
