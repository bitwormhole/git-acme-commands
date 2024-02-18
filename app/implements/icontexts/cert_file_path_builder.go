package icontexts

import (
	"fmt"
	"strings"
	"time"

	"github.com/starter-go/afs"
)

type certFilePathBuilder struct {
	Dir    afs.Path
	Domain string
	Time   time.Time
}

func (inst *certFilePathBuilder) createForLatest() afs.Path {
	return inst.createWithSessionName("latest")
}

func (inst *certFilePathBuilder) createForCurrent() afs.Path {
	str := inst.Time.Format(time.RFC3339)
	session := str
	parts := strings.Split(str, "T")
	for i, part := range parts {
		if i == 0 {
			session = part
			break
		}
	}
	return inst.createWithSessionName(session)
}

func (inst *certFilePathBuilder) createWithSessionName(session string) afs.Path {
	const f = "%s-%s.cer"
	if session == "" {
		session = "latest"
	}
	domain := inst.Domain
	name := fmt.Sprintf(f, domain, session)
	return inst.Dir.GetChild(name)
}
