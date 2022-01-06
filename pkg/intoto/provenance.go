package intoto

import (
    "github.com/in-toto/in-toto-golang/in_toto"
    slsa "github.com/in-toto/in-toto-golang/in_toto/slsa_provenance/v0.2"
	"strings"
	"time"
)

const (
	SlsaPredicateType = "https://slsa.dev/provenance/v0.2"
	StatementType     = "https://in-toto.io/Statement/v0.1"
)

func GenerateStatement(app App) in_toto.Statement {
	return in_toto.Statement{
		StatementHeader: in_toto.StatementHeader{
			Type:          StatementType,
			PredicateType: SlsaPredicateType,
			Subject:       nil,
		},
		Predicate: withPredicate(app),
	}
}

func GenerateSlsaPredicate(app App) slsa.ProvenancePredicate {
    return withPredicate(app)
}

func withPredicate(app App) slsa.ProvenancePredicate {
	return slsa.ProvenancePredicate{
        Builder:     slsa.ProvenanceBuilder{},
        BuildType:   "yolo",
        Invocation:  slsa.ProvenanceInvocation{
            ConfigSource: slsa.ConfigSource{},
            Parameters:   nil,
            Environment:  nil,
        },
        BuildConfig: nil,
        Metadata:    withMetadata(false, time.Now(), time.Now()),
        Materials:   withMaterials(app),
    }
}

func FindMaterials(predicate slsa.ProvenancePredicate, value string) []slsa.ProvenanceMaterial {
	var found []slsa.ProvenanceMaterial
	for _, m := range predicate.Materials {
		if find(m, value) {
			found = append(found, m)
		}
	}
	return found
}

func find(material slsa.ProvenanceMaterial, value string) bool {
	return strings.Contains(material.URI, value)
}

func withMetadata(rp bool, buildStarted, buildFinished time.Time) *slsa.ProvenanceMetadata {
	return &slsa.ProvenanceMetadata{
		BuildStartedOn:  &buildStarted,
		BuildFinishedOn: &buildFinished,
		Completeness:    withCompleteness(false, false, false),
		Reproducible:    rp,
	}
}

func withCompleteness(arguments, environment, materials bool) slsa.ProvenanceComplete {
	return slsa.ProvenanceComplete{
		Environment: environment,
		Materials:   materials,
	}
}

// TODO: use other type of materials aswell, e.g. github actions run in the build
func withMaterials(app App) []slsa.ProvenanceMaterial {
	materials := make([]slsa.ProvenanceMaterial, 0)
	for k, v := range app.Dependencies {
		m := slsa.ProvenanceMaterial{
			URI:    k + ":" + v,
			Digest: nil,
		}
		materials = append(materials, m)
	}
	return materials
}
