package main

type Commands struct {
	Commands []Command `json:"commands"`
}

type Command struct {
	Botcommand string `json:"botcommand"`
	Botmessage string `json:"botmessage"`
}
