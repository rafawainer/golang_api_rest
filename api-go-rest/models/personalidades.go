package models

type Personalidade struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personalidades []Personalidade

func RetornaUmaPersonalidade(p []Personalidade, id int) (Personalidade, string) {

	var pessoaEncontrada Personalidade
	message := "Personalidade Não Encontrada"

	for _, v := range p {
		if v.Id == id {
			pessoaEncontrada = v
			message = "Personalidade Encontrada com Sucesso"
			break
		}
	}
	return pessoaEncontrada, message
}
