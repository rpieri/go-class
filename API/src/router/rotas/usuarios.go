package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:              "/usuarios",
		Metodo:           http.MethodPost,
		Funcao:           controllers.CriarUsuario,
		RequerAutencicao: false,
	},
	{
		URI:              "/usuarios",
		Metodo:           http.MethodGet,
		Funcao:           controllers.BuscarUsuarios,
		RequerAutencicao: false,
	},
	{
		URI:              "/usuarios/{usuarioId}",
		Metodo:           http.MethodGet,
		Funcao:           controllers.BuscarUsuario,
		RequerAutencicao: false,
	},
	{
		URI:              "/usuarios/{usuarioId}",
		Metodo:           http.MethodPut,
		Funcao:           controllers.AtualizarUsuario,
		RequerAutencicao: false,
	},
	{
		URI:              "/usuarios/{usuarioId}",
		Metodo:           http.MethodDelete,
		Funcao:           controllers.RemoverUsuario,
		RequerAutencicao: false,
	},
}
