package jvm

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/nais/salsa/pkg/digest"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"

	"github.com/nais/salsa/pkg/build"
	"github.com/nais/salsa/pkg/utils"
)

const gradleBuildFileName = "build.gradle.kts"

type Gradle struct {
	BuildFilePatterns []string
}

func (g Gradle) ResolveDeps(workDir string) (*build.ArtifactDependencies, error) {
	cmd := exec.Command(
		"gradle",
		"-q", "dependencies", "--configuration", "runtimeClasspath", "-M", "sha256",
	)
	cmd.Dir = workDir

	err := utils.RequireCommand("gradle")
	if err != nil {
		return nil, fmt.Errorf("exec: %v\n", err)
	}

	depsOutput, err := utils.Exec(cmd)
	if err != nil {
		return nil, fmt.Errorf("exec: %v\n", err)
	}

	xmlData, err := ioutil.ReadFile(workDir + "/gradle/verification-metadata.xml")
	if err != nil {
		return nil, fmt.Errorf("readfile: %v\n", err)
	}

	deps, err := GradleDeps(depsOutput, xmlData)
	if err != nil {
		return nil, fmt.Errorf("could not get gradle deps: %w", err)
	}

	return &build.ArtifactDependencies{
		Cmd: build.Cmd{
			Path:     cmd.Path,
			CmdFlags: strings.Join(cmd.Args, " "),
		},
		RuntimeDeps: deps,
	}, nil
}

func NewGradle() build.BuildTool {
	return &Gradle{
		BuildFilePatterns: []string{gradleBuildFileName},
	}
}

func (g Gradle) BuildFiles() []string {
	return g.BuildFilePatterns
}

func GradleDeps(depsOutput string, checksumXml []byte) (map[string]build.Dependency, error) {
	pattern := regexp.MustCompile(`(?m)---\s[a-zA-Z0-9.]+:.*$`)
	matches := pattern.FindAllString(depsOutput, -1)
	if matches == nil {
		return nil, errors.New("unable to find any dependencies")
	}

	deps := make(map[string]build.Dependency, 0)

	sum := GradleChecksum{}
	err := xml.Unmarshal(checksumXml, &sum)
	if err != nil {
		return nil, fmt.Errorf("xml parsing: %v", err)
	}

	for _, match := range matches {
		elements := strings.Split(strings.Replace(match, "--- ", "", -1), ":")
		groupId := elements[0]
		artifactId := elements[1]
		version := filterVersion(elements[2])
		coordinates := fmt.Sprintf("%s:%s", groupId, artifactId)
		checksum := sum.checksum(groupId, artifactId, version)
		if noneEmpty(checksum) {
			deps[coordinates] = build.Dependency{
				Coordinates: coordinates,
				Version:     version,
				CheckSum: build.CheckSum{
					Algorithm: digest.AlgorithmSHA256,
					Digest:    checksum,
				},
			}
		}
	}

	return deps, nil
}

func noneEmpty(checksum string) bool {
	return checksum != ""
}

func filterVersion(rawVersion string) string {
	// 1.6.0 -> 1.6.10 (*)
	// 1.6.0 (*)
	// 1.6.10
	// 1.5.2-native-mt (*)
	filteredSuffix := filterSuffixes(rawVersion, " (*)", " (c)")
	useLatest := strings.Split(filteredSuffix, " -> ")
	if len(useLatest) > 1 {
		return useLatest[1]
	}
	return useLatest[0]
}

func filterSuffixes(orgString string, suffixes ...string) string {
	result := orgString
	for _, suffix := range suffixes {
		result = strings.TrimSuffix(result, suffix)
	}
	return result
}

func (g GradleChecksum) checksum(groupId, artifactId, version string) string {
	for _, c := range g.Components.Components {
		if c.Group == groupId && c.Name == artifactId && c.Version == version {
			for _, a := range c.Artifacts {
				if strings.HasSuffix(a.Name, ".jar") {
					return a.Sha256.Value
				}
			}
		}
	}

	return ""
}

type GradleChecksum struct {
	XMLName       xml.Name      `xml:"verification-metadata"`
	Configuration Configuration `xml:"configuration"`
	Components    Components    `xml:"components"`
}

type Configuration struct {
	XMLName          xml.Name `xml:"configuration"`
	VerifyMetadata   bool     `xml:"verify-metadata"`
	VerifySignatures bool     `xml:"verify-signatures"`
}

type Components struct {
	XMLName    xml.Name    `xml:"components"`
	Components []Component `xml:"component"`
}

type Component struct {
	XMLName   xml.Name   `xml:"component"`
	Group     string     `xml:"group,attr"`
	Name      string     `xml:"name,attr"`
	Version   string     `xml:"version,attr"`
	Artifacts []Artifact `xml:"artifact"`
}

type Artifact struct {
	XMLName xml.Name `xml:"artifact"`
	Name    string   `xml:"name,attr"`
	Sha256  Sha256   `xml:"sha256"`
}

type Sha256 struct {
	XMLName xml.Name `xml:"sha256"`
	Value   string   `xml:"value,attr"`
}
