package htmlsafe

import (
	"golang.org/x/net/html/atom"
)

var (
	// voidMap is a set of atom.Atom, elements which the HTML spec lists as self-closing tags.
	voidMap = make(map[atom.Atom]struct{})
)

func init() {
	voidMap[atom.Area] = struct{}{}
	voidMap[atom.Base] = struct{}{}
	voidMap[atom.Br] = struct{}{}
	voidMap[atom.Col] = struct{}{}
	voidMap[atom.Embed] = struct{}{}
	voidMap[atom.Hr] = struct{}{}
	voidMap[atom.Img] = struct{}{}
	voidMap[atom.Input] = struct{}{}
	voidMap[atom.Keygen] = struct{}{}
	voidMap[atom.Link] = struct{}{}
	voidMap[atom.Menuitem] = struct{}{}
	voidMap[atom.Meta] = struct{}{}
	voidMap[atom.Param] = struct{}{}
	voidMap[atom.Source] = struct{}{}
	voidMap[atom.Track] = struct{}{}
	voidMap[atom.Wbr] = struct{}{}
}
