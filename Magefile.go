//go:build mage
// +build mage

package main

import (
	"fmt"

	// mage:import
	build "github.com/grafana/opensearch-datasource/build"
)

// Hello prints a message (shows that you can define custom Mage targets).
func Hello() {
	fmt.Println("hello plugin developer!")
}

// Default configures the default target.
var Default = build.BuildAll
