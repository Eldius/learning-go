/*
Package specs groups the setup definition code
*/
package specs

import (
	"fmt"
	"path/filepath"
	"io/ioutil"
	yml "gopkg.in/yaml.v3"
)

/*
SetupSpecsList is a list of specs
*/
type SetupSpecsList struct {
	Specs map[string]SetupSpec `yaml:"specs"`
}

/*
SetupSpec is an abstraction of a install spec
*/
type SetupSpec struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
    URL string `yaml:"url"`
    Version string `yaml:"version"`
    VersionCmd string `yaml:"versionCmd"`
}

/*
ImportSpecs imports specs files
*/
func ImportSpecs(files []string) {
	fmt.Println("Importing files: ", files)
	ParseSpecsFromFile(files)
}

/*
ParseSpecsFromFile parse spec file to an object slice
*/
func ParseSpecsFromFile(files []string) []SetupSpec {
	parsedSpecs := make([]SetupSpec, 0)

	for _, f := range files {
		fileName, err := filepath.Abs(f)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Parsing file: ", fileName)
		if data, err := ioutil.ReadFile(fileName); err != nil {
			panic(err.Error())
		} else {
			var specsMap SetupSpecsList
			if errr := yml.Unmarshal(data, &specsMap); err != nil {
				panic(errr.Error())
			} else {
				for k, v := range specsMap.Specs {
					v.Name = k
					parsedSpecs = append(parsedSpecs, v)
				}
			}
		}
	}
	return parsedSpecs
}