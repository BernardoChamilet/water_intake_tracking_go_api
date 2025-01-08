package models

type RespostaLogin struct {
	Matricula int    `json:"matricula"`
	Token     string `json:"token"`
}
