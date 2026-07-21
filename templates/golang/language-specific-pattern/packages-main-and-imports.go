// Go packages, imports, init, and main entry-point patterns.
//
// This file remains in package languagepatterns so all templates compile
// together. Copy the marked entry-point shape into a package main program when
// writing a standalone executable.
package languagepatterns

import (
	format "fmt"
	"strings"
)

// Imports are file-scoped: every file lists exactly the packages it uses.
// `format` above demonstrates an alias, though the conventional `fmt` name is
// normally clearer. Imports are grouped and sorted by gofmt.

var packageInitTracePattern []string

// init has no parameters or results and is called automatically after package
// variables initialize. Libraries should avoid surprising init side effects.
func init() {
	packageInitTracePattern = append(packageInitTracePattern, "languagepatterns initialized")
}

// PackageInitTracePattern returns a copy so callers cannot mutate package state.
func PackageInitTracePattern() []string {
	return append([]string(nil), packageInitTracePattern...)
}

// FormatGreetingPattern uses names from two imported standard packages.
func FormatGreetingPattern(name string) string {
	cleaned := strings.TrimSpace(name)
	if cleaned == "" {
		cleaned = "interviewer"
	}
	return format.Sprintf("hello, %s", cleaned)
}

// SolvePackagePattern keeps algorithm logic independent of process I/O.
func SolvePackagePattern(nums []int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}

// MainPackagePattern contains the work an executable's main function would
// perform. Keeping it separate makes the logic directly testable.
func MainPackagePattern() string {
	answer := SolvePackagePattern([]int{1, 2, 3})
	return format.Sprint(answer)
}

/*
Standalone executable shape:

	package main

	import "fmt"

	func solve(nums []int) int {
		// Algorithm only.
		return 0
	}

	func main() {
		fmt.Println(solve([]int{1, 2, 3}))
	}

Package model:

- Every .go file starts with one package clause. Files in the same directory
  normally declare the same package and compile together.
- `package main` plus `func main()` defines an executable entry point. A main
  function is not used for an importable library package.
- Imported package names are referenced explicitly (`strings.TrimSpace`).
  Import aliases are useful for real name collisions, not routine shortening.
- `import . "path"` injects exported names into the file and obscures origin;
  avoid dot imports in interview and production code.
- `import _ "path"` runs a package's initialization only for side effects. It
  is common for driver registration but uncommon in algorithm code.
- Importing a package initializes its dependencies, then package-level
  variables, then init functions, once before dependent code runs.
- Packages cannot have import cycles.

Visibility and naming:

- Names beginning with an uppercase Unicode letter are exported from a package.
  Lowercase names are accessible only within the package.
- The import path identifies a package to the build system. The identifier used
  in source normally comes from its package clause.
- Keep input/output in main and algorithm logic in small functions. LeetCode
  supplies its own harness, so submissions usually need neither package main nor
  a main function beyond the platform's provided template.
*/
