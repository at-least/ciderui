package reference

import (
	"bytes"
	"fmt"
	"go/format"
	"strconv"
	"strings"
)

func goDigest(d [32]byte) string {
	parts := make([]string, len(d))
	for i, b := range d {
		parts[i] = strconv.Itoa(int(b))
	}
	return "[32]byte{" + strings.Join(parts, ", ") + "}"
}

// SignedInGoFile is reference/<site>/results/<p.Name>.go: p as a SignedInPage literal
// named by GoName, gofmt'd.
func SignedInGoFile(site string, p SignedInPage) ([]byte, error) {
	name, err := GoName(p.Name)
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	b.WriteString(header(site))
	emit(&b, "var %s = reference.SignedInPage{\nName: %q,\nPath: %q,\nBrowser: %q,\nDate: %q,\n", name, p.Name, p.Path, p.Browser, p.Date)
	emit(&b, "Window: reference.Window{Width: %s, Height: %s, PixelRatio: %s},\n", float(p.Window.Width), float(p.Window.Height), float(p.Window.PixelRatio))
	if p.Probes != nil {
		b.WriteString("Probes: []reference.SignedInProbe{\n")
		for _, pr := range p.Probes {
			emit(&b, "{\nName: %q,\nHTML: %q,\n", pr.Name, pr.HTML)
			goStyles(&b, "Styles", pr.Styles)
			if pr.Box != nil {
				emit(&b, "Box: &reference.Box{Width: %s, Height: %s},\n", float(pr.Box.Width), float(pr.Box.Height))
			}
			b.WriteString("},\n")
		}
		b.WriteString("},\n")
	}
	if p.Tokens != nil {
		b.WriteString("Tokens: []reference.Token{\n")
		for _, t := range p.Tokens {
			emit(&b, "{Name: %q, Value: %q},\n", t.Name, t.Value)
		}
		b.WriteString("},\n")
	}
	emit(&b, "ProbesDigest: %s,\nReaderDigest: %s,\nSHA256: %s,\n}\n", goDigest(p.ProbesDigest), goDigest(p.ReaderDigest), goDigest(p.SHA256))
	out, err := format.Source(b.Bytes())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", p.Name, err)
	}
	return out, nil
}

// WriteSignedIn writes p to dir/<p.Name>.go, and dir/pages.go from the records dir
// then holds; site is the command that writes them, reference/<site>.
func WriteSignedIn(dir, site string, p SignedInPage) error {
	src, err := SignedInGoFile(site, p)
	if err != nil {
		return err
	}
	return writeResult(dir, site, "reference.SignedInPage", p.Name, src)
}
