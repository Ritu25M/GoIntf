package tag

// notice that package species is not imported here

type tagStruct struct {
	speciesManager infSpecies
}

type infSpecies interface {
	GetCategory(string) string
}

var categorytagmap = map[string]tagtype{
	"herbivore": "HR",
	"carnivore": "CR",
	"omnivore":  "OR",
}

type tagtype string

// InitStructB initializes structB with an instance of interfaceA and returns structB.
// serves as a function called by FunctionA1 from packageA.
func InitTagger(iSpecies infSpecies) tagStruct {
	return tagStruct{iSpecies}
}

// CreateTagPrefix is a method on StructB that calls GetCategory without using the "species" package directly..
func (s tagStruct) CreateTagPrefix(species string) tagtype {
	if tag, ok := categorytagmap[s.speciesManager.GetCategory(species)]; ok {
		return tag
	} else {
		return "Unknown"
	}
	// import cycle will not arise here.
}
