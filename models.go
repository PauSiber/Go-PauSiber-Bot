package main

type JSONData struct {
	Commands []struct {
		Botcommand string `json:"botcommand"`
		Botmessage string `json:"botmessage"`
	} `json:"commands"`
}
