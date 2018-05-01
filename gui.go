package main

// GUI ...
type GUI struct {
	buttons []Button
}

// Button ...
type Button struct {
	Element
	state int
}
