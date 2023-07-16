package types

import (
	"github.com/ichiban/prolog/engine"
)

type PrologQueryRequest struct {
	PrologExtensionManifest *PrologExtensionManifestRequest `json:"prolog_extension_manifest,omitempty"`
	RunPredicate            *RunPredicateRequest            `json:"run_predicate,omitempty"`
}

type PrologQueryResponse struct {
	PrologExtensionManifest *PrologExtensionManifestResponse `json:"prolog_extension_manifest,omitempty"`
	RunPredicate            *RunPredicateResponse            `json:"run_predicate,omitempty"`
}

type PrologExtensionManifestRequest struct{}

type PrologExtensionManifestResponse struct {
	Predicates []PredicateManifest `json:"predicates"`
}

type PredicateManifest struct {
	Address string  `json:"address"`
	Name    string         `json:"name"`
	Cost    uint64         `json:"cost"`
}

type RunPredicateRequest struct {
	Name string   `json:"name"`
	Args []WasmTerm `json:"args"`
}

type RunPredicateResponse struct {
	Commands []Command `json:"commands"`
}

type Command struct {
	Unify []WasmTerm `json:"unify"`
}

type WasmTerm struct {
	Var *int64 `json:"var,omitempty"`
	Atom *string `json:"atom,omitempty"` 
}

func (term WasmTerm) ToTerm() engine.Term {
	if term.Var != nil {
		return engine.Variable(*term.Var)
	} else if term.Atom != nil {
		return engine.NewAtom(*term.Atom)
	} else {
		panic("invalid wasm term")
	}
}