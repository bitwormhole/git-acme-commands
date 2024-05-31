package ls

import (
	"github.com/starter-go/afs"
	"github.com/starter-go/application/properties"
)

// LoadProperties ...
func LoadProperties(file afs.Path) (properties.Table, error) {
	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}
	return properties.Parse(text, nil)
}

// SaveProperties ...
func SaveProperties(p properties.Table, file afs.Path) error {
	opt := afs.ToWriteFile()
	if !file.Exists() {
		opt = afs.ToCreateFile()
	}
	text := properties.Format(p, properties.FormatWithGroups)
	return file.GetIO().WriteText(text, opt)
}
