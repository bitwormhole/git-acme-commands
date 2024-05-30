package core

import "strings"

// DomainListItem ...
type DomainListItem struct {
	Raw     string
	Domain  string
	Enabled bool
}

func parseDomainList(str string) ([]*DomainListItem, error) {

	const (
		ch1 = "\r"
		ch2 = "\n"
	)

	str = strings.ReplaceAll(str, ch1, ch2)
	src := strings.Split(str, ch2)
	dst := make([]*DomainListItem, 0)

	for _, row := range src {
		item := new(DomainListItem)
		item.Raw = row
		text := strings.TrimSpace(row)
		if text == "" {
		} else if strings.HasPrefix(text, "#") {
		} else if strings.HasPrefix(text, "//") {
		} else {
			item.Domain = strings.ToLower(text)
			item.Enabled = true
		}
		dst = append(dst, item)
	}

	return dst, nil
}
