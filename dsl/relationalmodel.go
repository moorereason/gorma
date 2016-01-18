package dsl

import "github.com/bketelsen/gorma"

// RelationalModel is the DSL that represents a Relational Model
// Examples and more docs here later
func RelationalModel(name string, dsl func()) {
	// We can't rely on this being run first, any of the top level DSL could run
	// in any order. The top level DSLs are API, Version, Resource, MediaType and Type.
	// The first one to be called executes InitDesign.
	checkInit()
	if s, ok := relationalStoreDefinition(true); ok {
		if s.RelationalModels == nil {
			s.RelationalModels = make(map[string]*gorma.RelationalModelDefinition)
		}
		store, ok := s.RelationalModels[name]
		if !ok {
			store := &gorma.RelationalModelDefinition{
				Name: name,
				DSL:  dsl,
			}
		}

		s.RelationalModels[name] = store
	}

}