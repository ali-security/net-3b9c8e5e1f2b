// Command create_mod builds a Go module zip in proxy.golang.org layout from a
// source directory, using golang.org/x/mod/zip.CreateFromDir.
//
// Usage: create_mod <module-path> <version> <source-dir> <output-zip>
package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "usage: create_mod <module-path> <version> <source-dir> <output-zip>")
		os.Exit(2)
	}
	modPath, version, srcDir, outPath := os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	f, err := os.Create(outPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create %s: %v\n", outPath, err)
		os.Exit(1)
	}
	defer f.Close()

	m := module.Version{Path: modPath, Version: version}
	if err := zip.CreateFromDir(f, m, srcDir); err != nil {
		fmt.Fprintf(os.Stderr, "CreateFromDir %s@%s from %s: %v\n", modPath, version, srcDir, err)
		os.Exit(1)
	}
	fmt.Printf("wrote %s for %s@%s\n", outPath, modPath, version)
}
