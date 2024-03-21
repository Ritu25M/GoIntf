package species

import "goIntf/tag"

// Species handles the species-related functionalities.
type species struct {
}

// speciesHandler is an instance of Species used for accessing species-related methods.
var speciesHandler species

const unknown = "unknown"

// catToSpeciesMap maps species to their respective categories.
var catToSpeciesMap = map[string][]string{"herbivore": {"cow", "monkey"}}

// GenerateTag generates a tag for the given species.
// It is called by the main routine.
func GenerateTag(speciesName string) string {
	sTagger := tag.InitTagger(speciesHandler)
	return string(sTagger.CreateTagPrefix(speciesName))
}

// GetCategory retrieves the category of the given species.
// It is called by CreateTagPrefix method in the "tag" package.
func (s species) GetCategory(species string) string {
	for cat, speciesSlice := range catToSpeciesMap {
		for _, r := range speciesSlice {
			if r == species {
				return cat
			}
		}
	}
	return unknown
}

// dummyfunction to include other methods that species might handle
func (s species) dummyfunction(species string) {
	return
}
