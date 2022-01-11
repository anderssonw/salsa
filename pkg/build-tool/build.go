package build_tool

import (
	"encoding/json"
	"fmt"
	"github.com/nais/salsa/pkg/intoto"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

type BuildTool interface {
	Build(project string) error
	BuildTool(pattern string) bool
	BuildFiles() []string
}

func Scan(workingDir, project string) error {
	gradle := NewGradle(workingDir)
	mvn := NewMaven(workingDir)
	golang := NewGolang(workingDir)

	supportedBuildFiles := sumSupported(
		gradle.BuildFiles(),
		mvn.BuildFiles(),
		golang.BuildFiles(),
	)

	totalSupported := len(supportedBuildFiles)
	for index, pattern := range supportedBuildFiles {
		log.Printf("search for build type '%s'", pattern)
		foundBuildType := findBuildType(workingDir, pattern)

		if index < totalSupported {
			log.Printf("searching..")
			if foundBuildType != "" {
				log.Printf("found build type %s", foundBuildType)
				switch true {
				case gradle.BuildTool(foundBuildType):
					err := gradle.Build(project)
					if err != nil {
						return err
					}
				case mvn.BuildTool(foundBuildType):
					err := mvn.Build(project)
					if err != nil {
						return err
					}
				case golang.BuildTool(foundBuildType):
					err := golang.Build(project)
					if err != nil {
						return err
					}

					// add more cases
				}
				// found break out!
				break
			}

		} else {
			return fmt.Errorf("unknown build type")
		}
	}
	return nil
}

func findBuildType(root, pattern string) (result string) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name() == pattern {
			result = file.Name()
			break
		}
	}
	return result
}

func GenerateProvenance(project string, deps map[string]string) error {
	app := createApp(project, deps)
	s := intoto.GenerateSlsaPredicate(app)

	statement, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("marshal: %v\n", err)
	}

	log.Println(string(statement))
	provenanceName := fmt.Sprintf("%s.provenance", project)

	err = os.WriteFile(provenanceName, statement, 0644)
	if err != nil {
		return fmt.Errorf("write to file: %v\n", err)
	}
	return nil
}

func createApp(name string, deps map[string]string) intoto.App {
	return intoto.App{
		Name:         name,
		BuilderId:    "todoId",
		BuildType:    "todoType",
		Dependencies: deps,
	}
}
