package model

type Option struct {
	Text string
	Arc  string
}

type Arc struct {
	Title   string
	Story   []string
	Options []Option
}

type Story struct {
	Arcs map[string]Arc
}
