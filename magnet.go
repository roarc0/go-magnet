package magnet

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

var (
	magnetPrefix = "magnet:?"
)

// Magnet represents a parsed magnet link.
type Magnet struct {
	ExactTopics       []string // xt
	DisplayNames      []string // dn
	ExactLength       uint64   // xl
	AddressTrackers   []string // tr
	WebSeeds          []string // ws
	AcceptableSources []string // as
	ExactSources      []string // xs
	KeywordTopics     []string // kt
	ManifestTopics    []string // mt
	SelectOnly        []string // so
	PEer              []string // x.pe
}

// Parse parses a magnet URI and returns a Magnet.
func Parse(magnetURI string) (*Magnet, error) {
	if !strings.HasPrefix(magnetURI, magnetPrefix) {
		return nil, &ErrorInvalidMagnet{msg: "missing magnet URI prefix"}
	}

	values, err := url.ParseQuery(magnetURI[8:])
	if err != nil {
		return nil, &ErrorInvalidMagnet{msg: "failed to parse magnet URI", err: err}
	}

	magnet := &Magnet{}

	if dn, ok := values["dn"]; ok {
		magnet.DisplayNames = dn
	}

	if xl, err := parseExactLength(values["xl"]); err != nil {
		return nil, &ErrorInvalidMagnet{msg: "failed to parse magnet XL field", err: err}
	} else {
		magnet.ExactLength = xl
	}

	if tr, ok := values["tr"]; ok {
		magnet.AddressTrackers = tr
	}

	// TODO support xt.1, xt.2, xt.3, ...
	if xt, ok := values["xt"]; ok {
		magnet.ExactTopics = xt
	}

	if ws, ok := values["ws"]; ok {
		magnet.WebSeeds = ws
	}

	if as, ok := values["as"]; ok {
		magnet.AcceptableSources = as
	}

	if xs, ok := values["xs"]; ok {
		magnet.ExactSources = xs
	}

	if kt, ok := values["kt"]; ok {
		magnet.KeywordTopics = kt
	}

	if mt, ok := values["mt"]; ok {
		magnet.ManifestTopics = mt
	}

	if so, ok := values["so"]; ok {
		magnet.SelectOnly = so
	}

	if pe, ok := values["x.pe"]; ok {
		magnet.PEer = pe
	}

	return magnet, nil
}

func parseExactLength(value []string) (uint64, error) {
	if len(value) == 0 {
		return 0, nil
	} else if len(value) > 1 {
		return 0, fmt.Errorf("got %d", len(value))
	}

	return strconv.ParseUint(value[0], 10, 64)
}

// String returns the magnet URI.
func (m *Magnet) String() string {
	values := url.Values{}

	for _, xt := range m.ExactTopics {
		values.Add("xt", xt)
	}

	for _, dn := range m.DisplayNames {
		values.Add("dn", dn)
	}

	if m.ExactLength != 0 {
		values.Add("xl", strconv.FormatUint(m.ExactLength, 10))
	}

	for _, tr := range m.AddressTrackers {
		values.Add("tr", tr)
	}

	for _, ws := range m.WebSeeds {
		values.Add("ws", ws)
	}

	for _, as := range m.AcceptableSources {
		values.Add("as", as)
	}

	for _, xs := range m.ExactSources {
		values.Add("xs", xs)
	}

	for _, kt := range m.KeywordTopics {
		values.Add("kt", kt)
	}

	for _, mt := range m.ManifestTopics {
		values.Add("mt", mt)
	}

	for _, so := range m.SelectOnly {
		values.Add("so", so)
	}

	for _, pe := range m.PEer {
		values.Add("x.pe", pe)
	}

	return magnetPrefix + values.Encode()
}
