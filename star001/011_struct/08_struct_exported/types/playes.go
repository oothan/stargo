// Package athletes ... (your comments to describe the exported package)
package athletes

import "strings"

type Info struct {
	Country  string
	HairCode string
}

type Player struct {
	Name  string
	Sport string
	Age   int
	Info
}

func (p *Player) ToLowercase() *Player {
	p.Name = strings.ToLower(p.Name)
	p.Sport = strings.ToLower(p.Sport)
	p.Country = strings.ToLower(p.Country)
	p.HairCode = strings.ToLower(p.HairCode)
	return p
}
