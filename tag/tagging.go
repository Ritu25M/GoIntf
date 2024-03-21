package tag

import "time"

// notice that package species is not imported here

// Tagger handles the tag-related methods.
type tagger struct {
	speciesManager infSpecies
}

// SpeciesHandler defines the interface for species-related used by tagger.
type infSpecies interface {
	GetCategory(string) string
}

// catToTagMap maps species categories to their respective tag types.
var catToTagMap = map[string]tagtype{
	"herbivore": "HR",
	"carnivore": "CR",
	"omnivore":  "OR",
	"unknown":   "U",
}

type tagtype string

// InitTagger initializes tagger struct with an instance of speciesManager and returns tagger.
func InitTagger(iSpecies infSpecies) tagger {
	return tagger{iSpecies}
}

// CreateTagPrefix is a method on Tagger that calls GetCategory without using the "species" package directly.
func (t tagger) CreateTagPrefix(species string) tagtype {
	// Retrieve category of the species and map it to the corresponding tag type
	cat := t.speciesManager.GetCategory(species)
	if tag, ok := catToTagMap[cat]; ok {
		return tag + time.DateTime
	} else {
		return "CatNotSupported"
	}
	// import cycle will not arise here.
}
