package main

import (
	"maps"
	"slices"
)

func removeSubfolders(folder []string) []string {
	roots := map[string]bool{}
	slices.Sort(folder)
outer:
	for _, f := range folder {
		for i, c := range f {
			if c == '/' && roots[f[:i]] {
				continue outer
			}
		}
		roots[f] = true
	}
	return slices.Collect(maps.Keys(roots))
}
