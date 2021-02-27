package main

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

	music, exists  := artistList[artistName]
	if !exists {
		return []string{}
	}

	return music
}
