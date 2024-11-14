// Package json contains functionality to render output in json format
package json

import (
	"encoding/json"
	"log"

	"github.com/devops-kung-fu/bomber/models"
	"github.com/devops-kung-fu/common/util"
	"github.com/spf13/afero"
)

// Renderer contains methods to render to JSON format
type Renderer struct{}

// Render outputs json to STDOUT
func (Renderer) Render(results models.Results) error {
	var afs *afero.Afero

	if results.Meta.Provider == "test" {
		afs = &afero.Afero{Fs: afero.NewMemMapFs()}
	} else {
		afs = &afero.Afero{Fs: afero.NewOsFs()}
	}

	filename := "bomber-results.json"

	err := writeTemplate(afs, filename, results)

	return err
}

func writeTemplate(afs *afero.Afero, filename string, results models.Results) error {
	b, _ := json.MarshalIndent(results, "", "\t")
	util.PrintInfo("Writing JSON output:", filename)
	file, err := afs.Create(filename)
	if err != nil {
		log.Println(err)
		return err
	}
	if _, err := file.Write(b); err != nil {
		log.Println(err)
		return err
	}
	err = afs.Fs.Chmod(filename, 0644)
	return err
}
