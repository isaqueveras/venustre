package main

import (
	"log"
	"net/http"

	"github.com/isaqueveras/venustre"
)

func main() {
	w := venustre.Watcher("#001", "Atualiza horario")
	defer w.Wait()

	http.Handle("/venustre/", w.Handler())

	log.Println("Venustre Dashboard: http://localhost:9091")
	log.Fatal(http.ListenAndServe(":9091", nil))
}
