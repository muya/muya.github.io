package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)


func TestFetchArtistMusic(t *testing.T) {
	// Using "testing" package
	artistWithMultipleSongs := FetchArtistMusic("Fuzzy Logic")
	if len(artistWithMultipleSongs) < 2 {
		t.Errorf("Expected artist to have at least 2 songs, found %d.", len(artistWithMultipleSongs))
	}

	t.Logf("Artist 2nd song is: %s", artistWithMultipleSongs[1])
}

func TestFetchArtistMusic_Fatal(t *testing.T) {
	// Using "testing" package
	artistWithMultipleSongs := FetchArtistMusic("Fuzzy Logic")
	if len(artistWithMultipleSongs) < 2 {
		t.Fatalf("Expected artist to have at least 2 songs, found %d.", len(artistWithMultipleSongs))
	}

	t.Logf("Artist 2nd song is: %s", artistWithMultipleSongs[1])
}

func TestFetchArtistMusic_withAssert(t *testing.T) {
	// Using "testing" package
	artistWithMultipleSongs := FetchArtistMusic("Fuzzy Logic")
	assert.Len(t, artistWithMultipleSongs, 2, "Expected artist to have at least 2 songs.")

	t.Logf("Artist 2nd song is: %s", artistWithMultipleSongs[1])
}

func TestFetchArtistMusic_withRequire(t *testing.T) {
	// Using "testing" package
	artistWithMultipleSongs := FetchArtistMusic("Fuzzy Logic")
	require.Len(t, artistWithMultipleSongs, 2, "Expected artist to have at least 2 songs.")

	t.Logf("Artist 2nd song is: %s", artistWithMultipleSongs[1])
}
