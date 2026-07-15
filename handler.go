package venustre

import (
	"net/http"
)

func (c *Watch) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/venustre/dashboard", func(w http.ResponseWriter, r *http.Request) {
		dashboardPage().Render(r.Context(), w)
	})

	mux.HandleFunc("/venustre/tasks", func(w http.ResponseWriter, r *http.Request) {
		tasksPage().Render(r.Context(), w)
	})

	mux.HandleFunc("/venustre/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		tasksDetailsPage(id).Render(r.Context(), w)
	})

	return mux
}
