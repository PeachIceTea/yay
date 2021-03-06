package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"unicode"
)

type intRange struct {
	min int
	max int
}

func makeIntRange(min, max int) intRange {
	return intRange{
		min,
		max,
	}
}

func (r intRange) get(n int) bool {
	return n >= r.min && n <= r.max
}

type intRanges []intRange

func (rs intRanges) get(n int) bool {
	for _, r := range rs {
		if r.get(n) {
			return true
		}
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func addMapStringSet(h map[string]stringSet, n string, v string) {
	_, ok := h[n]
	if !ok {
		h[n] = make(stringSet)
	}
	h[n].set(v)
}

func addMapStringSlice(h map[string][]string, n string, v string) {
	_, ok := h[n]
	if !ok {
		h[n] = make([]string, 0, 1)
	}
	h[n] = append(h[n], v)
}

func completeFileName(dir, name string) (string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if strings.HasPrefix(file.Name(), name) {
			return filepath.Join(dir, file.Name()), nil
		}
	}

	return "", nil
}

func lessRunes(iRunes, jRunes []rune) bool {
	max := len(iRunes)
	if max > len(jRunes) {
		max = len(jRunes)
	}

	for idx := 0; idx < max; idx++ {
		ir := iRunes[idx]
		jr := jRunes[idx]

		lir := unicode.ToLower(ir)
		ljr := unicode.ToLower(jr)

		if lir != ljr {
			return lir < ljr
		}

		// the lowercase runes are the same, so compare the original
		if ir != jr {
			return ir < jr
		}
	}

	return len(iRunes) < len(jRunes)
}
