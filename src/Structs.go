package main


type Associated struct {
	Image string `json:"image"`
	Name string `json:"name"`
	Description string `json:"description"`
	Contact []Contact `json:"contact"`
	Address []Address `json:"address"`
}

type Contact struct {
	Email string `json:"email"`
	Number string `json:"number"`
	Pix string `json:"pix"`
}

type Address struct {
	Street string `json:"street"`
	DescriptionAddr string `json:"descriptionAddress"`
}