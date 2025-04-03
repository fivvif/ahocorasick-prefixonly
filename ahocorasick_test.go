// ahocorasick_test.go: test suite for ahocorasick
//
// Copyright (c) 2013 CloudFlare, Inc.

package ahocorasick

import (
	"testing"
)

func assert(t *testing.T, b bool) {
	if !b {
		t.Fail()
	}
}

func TestNoPatterns(t *testing.T) {
	m := NewStringMatcher([]string{})
	hits := m.Match([]byte("foo bar baz"))
	assert(t, len(hits) == 0)

	hits = m.MatchThreadSafe([]byte("foo bar baz"))
	assert(t, len(hits) == 0)
}

func TestNoData(t *testing.T) {
	m := NewStringMatcher([]string{"foo", "baz", "bar"})
	hits := m.Match([]byte(""))
	assert(t, len(hits) == 0)

	hits = m.MatchThreadSafe([]byte(""))
	assert(t, len(hits) == 0)
}

func TestBasic(t *testing.T) {
	m := NewStringMatcher([]string{"foo", "baz", "bar"})
	hits := m.Match([]byte("foo baz bar"))
	assert(t, len(hits) == 1)
	assert(t, hits[0] == 0)

	hits = m.MatchThreadSafe([]byte("foo baz bar"))
	assert(t, len(hits) == 1)
	assert(t, hits[0] == 0)
}

func TestSpaces(t *testing.T) {
	m := NewStringMatcher([]string{"foo bar", "sunshine cosmetic"})

	// 1 case
	hits := m.Match([]byte("foo bar sunshine cosmetic"))
	assert(t, len(hits) == 1)
	assert(t, hits[0] == 0)
	hits = m.MatchThreadSafe([]byte("foo bar sunshine cosmetic"))
	assert(t, len(hits) == 1)
	assert(t, hits[0] == 0)

	// 2 case
	hits = m.Match([]byte("women sunshine cosmetic"))
	assert(t, len(hits) == 0)
	hits = m.MatchThreadSafe([]byte("women sunshine cosmetic"))
	assert(t, len(hits) == 0)

	// 3 case
	hits = m.Match([]byte("sunshine cosmetic for womens"))
	assert(t, len(hits) == 1)
	assert(t, hits[0] == 1)
	hits = m.MatchThreadSafe([]byte("sunshine cosmetic for womens"))
	assert(t, len(hits) == 1)
	assert(t, hits[0] == 1)
}
