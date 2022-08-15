package routes

import (
	"net/http"

	"github.com/Alura/loja/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)

}
