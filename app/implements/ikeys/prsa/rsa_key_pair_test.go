package prsa

import (
	"testing"

	"github.com/starter-go/afs/files"
)

func TestRSA(t *testing.T) {

	dir := files.FS().NewPath(t.TempDir())
	file := dir.GetChild("rsa.key.pem")
	p := new(Provider)

	pair1, err := p.Generator().Generate()
	if err != nil {
		t.Error(err)
		return
	}

	err = p.Saver().Save(pair1, file)
	if err != nil {
		t.Error(err)
		return
	}

	pair2, err := p.Loader().Load(file)
	if err != nil {
		t.Error(err)
		return
	}

	pair2.Signer()
}
