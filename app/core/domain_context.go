package core

import (
	"fmt"
	"strings"

	"github.com/starter-go/afs"
)

// GetLatestCertificateFile ...
func (inst *DomainContext) GetLatestCertificateFile() (afs.Path, error) {

	refer := inst.LatestFile

	text, err := refer.GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}

	text = strings.TrimSpace(text)
	if text == "" {
		p := refer.GetPath()
		return nil, fmt.Errorf("latest file [%s] is empty", p)
	}

	dir := refer.GetParent()
	target := dir.GetChild(text)
	return target, nil
}

// GetCurrentCertificateFile ...
func (inst *DomainContext) GetCurrentCertificateFile() (afs.Path, error) {

	refer := inst.CurrentFile

	text, err := refer.GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}

	text = strings.TrimSpace(text)
	if text == "" {
		p := refer.GetPath()
		return nil, fmt.Errorf("current file [%s] is empty", p)
	}

	dir := refer.GetParent()
	target := dir.GetChild(text)
	return target, nil
}
