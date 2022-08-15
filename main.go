package main

import (
	"net/http"

	"github.com/Alura/loja/routes"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
