package html

import "embed"

// Templates contain the raw HTML of all of our templates.
//
//go:embed *.html assets/*
var Templates embed.FS
