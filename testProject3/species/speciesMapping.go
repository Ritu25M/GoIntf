package species

import "testProject3/tag"

type species struct{}

var speciesHandler species

// GetCategory retrieves the category of the given species.
// It is called by CreateTagPrefix method in the "tag" package.
func (s species) GetCategory(species string) string {
	for category, speciesSlice := range speciesToCategoryMap {
		for _, r := range speciesSlice {
			if r == species {
				return category
			}
		}
	}
	return "Unknown"
}

var speciesToCategoryMap = map[string][]string{"herbivore": {"cow", "chicken"}}

func GenerateTag(species string) string {
	sTagger := tag.InitTagger(speciesHandler)
	return string(sTagger.CreateTagPrefix(species))
}
