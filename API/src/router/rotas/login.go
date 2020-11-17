package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotaLogin = Rota{
	URI:              "/login",
	Metodo:           http.MethodPost,
	RequerAutencicao: false,
	Funcao:           controllers.Login,
}
