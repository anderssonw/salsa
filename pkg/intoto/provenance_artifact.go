package intoto

import (
	"github.com/nais/salsa/pkg/digest"
	"time"

	slsa "github.com/in-toto/in-toto-golang/in_toto/slsa_provenance/v0.2"
	"github.com/nais/salsa/pkg/build"
	"github.com/nais/salsa/pkg/vcs"
)

type ProvenanceArtifact struct {
	BuildConfig       string
	BuilderId         string
	BuilderRepoDigest *slsa.ProvenanceMaterial
	BuildInvocationId string
	BuildStartedOn    time.Time
	BuildType         string
	Dependencies      *build.ArtifactDependencies
	Invocation        *slsa.ProvenanceInvocation
	Name              string
}

func CreateProvenanceArtifact(name string, deps *build.ArtifactDependencies, env *vcs.Environment) *ProvenanceArtifact {
	pa := &ProvenanceArtifact{
		BuildStartedOn: time.Now().UTC(),
		Dependencies:   deps,
		Name:           name,
	}

	if env != nil {
		pa.BuildType = vcs.BuildType
		pa.BuildInvocationId = env.BuildInvocationId()
		pa.BuilderId = env.BuilderId()
		pa.withBuilderRepoDigest(env).withBuilderInvocation(env)
		return pa
	}

	pa.BuildConfig = "Some commands that made this build"
	pa.BuilderId = vcs.DefaultBuildId
	pa.BuildType = vcs.AdHocBuildType
	pa.Invocation = nil
	return pa
}

func (in *ProvenanceArtifact) withBuilderRepoDigest(env *vcs.Environment) *ProvenanceArtifact {
	in.BuilderRepoDigest = &slsa.ProvenanceMaterial{
		URI: "git+" + env.RepoUri(),
		Digest: slsa.DigestSet{
			digest.AlgorithmSHA1: env.GithubSha(),
		},
	}
	return in
}

func (in *ProvenanceArtifact) withBuilderInvocation(env *vcs.Environment) *ProvenanceArtifact {
	in.Invocation = &slsa.ProvenanceInvocation{
		ConfigSource: slsa.ConfigSource{
			URI: "git+" + env.RepoUri(),
			Digest: slsa.DigestSet{
				digest.AlgorithmSHA1: env.GithubSha(),
			},
			EntryPoint: env.Workflow,
		},
		Parameters: env.EventInputJson(),
		// Should contain the architecture of the runner.
		Environment: env.RunnerContext,
	}
	return in
}

func (in *ProvenanceArtifact) HasLegitBuilderRepoDigest() bool {
	if in.BuilderRepoDigest == nil {
		return false
	}

	return in.BuilderRepoDigest.Digest != nil && in.BuilderRepoDigest.URI != ""

}

func (in *ProvenanceArtifact) HasLegitDependencies() bool {
	if in.Dependencies == nil {
		return false
	}

	return len(in.Dependencies.RuntimeDeps) > 0
}

func (in *ProvenanceArtifact) HasLegitParameters() bool {
	if in.Invocation != nil {
		if in.Invocation.Parameters != nil {
			return true
		}
	}
	return false
}
