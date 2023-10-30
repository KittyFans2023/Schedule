package main

import (
	"net/http"
	"schedule/GO/schedule/update"

	"github.com/gorilla/mux"
)

func main() {
	// router := mux.NewRouter()
	update.Get_info_about()

	// router.HandleFunc("/schedule", homeHendler)
	// router.HandleFunc("/schedule/update", updateHendler)
	// router.HandleFunc("/schedule/get_info/{group_number}", groupHendler)

	// http.ListenAndServe(":8080", router)
}

func homeHendler(rw http.ResponseWriter, _ *http.Request) {
	rw.Write([]byte(`Пустая`))
}

func groupHendler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	number := vars["number"]
	update.Get_info_about(number)
	rw.Write([]byte(`выполнено`))
}
func updateHendler(rw http.ResponseWriter, req *http.Request) {
	update.Update("https://cfuv.ru/raspisanie-fakultativov-fiziko-tekhnicheskogo-instituta")
}
