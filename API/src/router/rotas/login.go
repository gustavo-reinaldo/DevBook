package rotas

import (
	"api/src/controllers"
	"net/http"
)

// Receber requisição com email e senha, buscar no banco de dados para verificar se não está cadastrado
var rotaLogin = Rota{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
