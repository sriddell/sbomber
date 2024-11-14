package json

import (
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/devops-kung-fu/bomber/models"
)

func TestRenderer_Render(t *testing.T) {
	afs := &afero.Afero{Fs: afero.NewMemMapFs()}
	err := writeTemplate(afs, "test.json", models.NewResults([]models.Package{}, models.Summary{}, []models.ScannedFile{}, []string{"GPL"}, "0.0.0", "test", ""))
	assert.NoError(t, err)
	
	b, err := afs.ReadFile("test.json")
	assert.NotNil(t, b)
	assert.NoError(t, err)

	info, err := afs.Stat("test.json")
	assert.NoError(t, err)
	assert.Equal(t, os.FileMode(0644), info.Mode().Perm())

}
