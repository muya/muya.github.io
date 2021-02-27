---
layout: post
title: Go Testify - Force Test Exit on Failed Assertion
date: '2021-02-27'
author: Fred Muya
excerpt: How to force a test to exit in GoLang when using Testify
tags:
- go
- golang
- testing
- test
- testify
---

**TL;DR:** `require.*` in `"testify"` is the equivalent of `t.Fatal/t.Fatalf` in `"testing"`

When writing test cases, sometimes you may want the current test case to exit in case of a failed assertion.

Take the method below (in which you also get some insight into some of the music I enjoy 😜 ) as an example:

{% highlight go %}

func FetchArtistMusic(artistName string) []string {
    artistList := map[string][]string{
        "Terrace Martin": {
            "Butterfly",
            "Valdez Off Crenshaw",
            "Almond Butter",
        },
        "Fuzzy Logic": {
            "In The Morning",
        },
        "Fall Out Boy": {
            "Young and Menace",
            "HOLD ME TIGHT OR DON'T",
            "Sunshine Riptide (feat. Burna Boy)",
        },
        "Rapsody": {
            "Never Fail",
            "Sojourner",
            "Maya",
            "Serena",
        },
        "Smino": {
            "Amphetamine",
            "Glass Flows (feat. Ravyn Lenae)",
        },
        "Seba Kaapstad": {
            "Our People",
            "You Better",
        },
    }

    music, exists := artistList[artistName]
    if !exists {
        return []string{}
    }

    return music
}

{% endhighlight %}

The `FetchArtistMusic` method returns an array of songs that we have for a given artist's name.

Let's say we want to have a test where we log the artist's 2nd song.

For example, for "Terrace Martin", we'd want to log "Valdez Off Crenshaw".


##### Using Default Go "testing" Package


Using Go's default [`"testing"` package](https://golang.org/pkg/testing/), the test could look something like this:

{% highlight go linenos %}

 import "testing"

 func TestFetchArtistMusic(t *testing.T) {
     // Using "testing" package
     artistWithMultipleSongs := FetchArtistMusic("Fuzzy Logic")
     if len(artistWithMultipleSongs) < 2 {
         t.Errorf("Expected artist to have at least 2 songs, found %d.", len(artistWithMultipleSongs))
     }

     t.Logf("Artist 2nd song is: %s", artistWithMultipleSongs[1])
 }

{% endhighlight %}

The test above has been written intentionally to make sure that we're fetching an artist with one song listed ("Fuzzy Logic").

Executing the test above results in a panic. Why?

This is because we're attempting to log the 2nd song on line 10, despite our assertion on line 6 failing.

Ideally, we'd want this test to exit upon the failed assertion.

When using the `"testing"` package, we can force the test to exit by using the [`Fatal`](https://golang.org/pkg/testing/#T.Fatal)/[`Fatalf`](https://golang.org/pkg/testing/#T.Fatalf) methods instead of [`Error`](https://golang.org/pkg/testing/#T.Error)/[`Errorf`](https://golang.org/pkg/testing/#T.Errorf):

{% highlight go linenos %}

 import "testing"

 func TestFetchArtistMusic_Fatal(t *testing.T) {
     // Using "testing" package
     artistWithMultipleSongs := FetchArtistMusic("Fuzzy Logic")
     if len(artistWithMultipleSongs) < 2 {
         t.Fatalf("Expected artist to have at least 2 songs, found %d.", len(artistWithMultipleSongs))
     }

     t.Logf("Artist 2nd song is: %s", artistWithMultipleSongs[1])
 }

{% endhighlight %}

Using `Fatalf`, the test exits on the failed assertion, and there is no panic.


##### Using testify
As you may have noticed, the test code above is a little verbose, because we have to validate the length using an `if` statement, then take the appropriate action.

The [`testify` package](https://github.com/stretchr/testify) provides some additional tools that ease writing of tests in Go.

One of these tools is "Assert", which makes assertions a breeze.

Using our example, compare the 2 code snippets below:

{% highlight go linenos %}

 // Using "testing" package
 artistWithMultipleSongs := FetchArtistMusic("Fuzzy Logic")
 if len(artistWithMultipleSongs) < 2 {
     t.Fatalf("Expected artist to have at least 2 songs, found %d.", len(artistWithMultipleSongs))
 }

 // Using "testify.assert"
 assert.Len(2, artistWithMultipleSongs, 2, "Expected artist to have multiple songs")


{% endhighlight %}


Using testify, the equivalent of our panic-inducing test will look as follows:

{% highlight go linenos %}

 import (
     "github.com/stretchr/testify/assert"
     "testing"
 )

 func TestFetchArtistMusic_withAssert(t *testing.T) {
     // Using "testing" package
     artistWithMultipleSongs := FetchArtistMusic("Fuzzy Logic")

     assert.Len(t, artistWithMultipleSongs, 2, "Expected artist to have at least 2 songs.")

     t.Logf("Artist 2nd song is: %s", artistWithMultipleSongs[1])
 }

{% endhighlight %}

The assertion on line 10 will fail, and attempting to log the value on line 12 will result in a panic.

To solve our problem while using the `"testify"` package, we'll need to use `require` instead:


{% highlight go linenos %}

 import (
     "github.com/stretchr/testify/require"
     "testing"
 )

 func TestFetchArtistMusic_withAssert(t *testing.T) {
     // Using "testing" package
     artistWithMultipleSongs := FetchArtistMusic("Fuzzy Logic")

     require.Len(t, artistWithMultipleSongs, 2, "Expected artist to have at least 2 songs.")

     t.Logf("Artist 2nd song is: %s", artistWithMultipleSongs[1])
 }

{% endhighlight %}

In this case, the test will exit upon the failed assertion, ensuring that there is no panic being induced!


In this blog post, we demonstrated how to force a test written in Go using the `"testify"` package to exit in case of a failed assertion.

Happy coding, and stay safe!

_**Note**: All the code snippets used in this post are [available on GitHub](https://github.com/muya/muya.github.io/blob/master/snippets/2021-02-27)_


_**Disclaimer**: All songs used in this article are the property of the respective amazing artists! Music was thoroughly enjoyed while writing this article._


_You may find a playlist of the songs on [Apple Music](https://music.apple.com/ke/playlist/go-testify/pl.u-BWVJhexlXvP) & [Spotify](https://open.spotify.com/playlist/26w7TL5iRVVTDzCW8Ws1C2?si=48bd73ca49984a62)._

