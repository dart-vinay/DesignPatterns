package Structural

import "fmt"

// Important: Composite represent "has a relationship" while inheritance represents "is a relationship"
// Used when the components form a hierarchy. When we have to implement a tree like structure then we prefer this structural pattern.
// Eg. Directory structure, Representing the positional hierarchy in companies, graphical objects on a screen (one of the best example)
// It helps to treat the object and the composition of the objects uniformly.

type Graphic interface {
	Create()
	GetIntro()
}

// Leaf graphical Components
type Line struct {
	Intro string
}

type Dot struct {
	Intro string
}

// Composite Graphic
type Picture struct {
	ChildGraphics []Graphic
	Intro string
}

func (leaf *Line) Create() {
	leaf.Intro = "This is a line"
}

func (line *Line) GetIntro() {
	fmt.Print(line.Intro)
}

func (dot *Dot) Create() {
	dot.Intro = "This is a dot"
}

func (dot *Dot) GetIntro() {
	fmt.Print(dot.Intro)
}

func (picture *Picture) Create() {
	picture.Intro = "This is a picture"
}

func (picture *Picture) GetIntro(){
	fmt.Print(picture.Intro)
}

func (picture *Picture) AddComponent(component Graphic) {
	picture.ChildGraphics = append(picture.ChildGraphics, component)
}
func (picture *Picture) ShowAllComponents() {
	for _, component := range picture.ChildGraphics {
		component.GetIntro()
		fmt.Print("\n")
		if newComponent, ok := component.(*Picture); ok {
			newComponent.ShowAllComponents()
		}
	}
}
