package main

import (
	"net/http"
	"schedule/GO/schedule/API"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// API.Update("https://cfuv.ru/raspisanie-fakultativov-fiziko-tekhnicheskogo-instituta")
	// API.Get_info_about("233-1")

	router.HandleFunc("/schedule", homeHendler)
	router.HandleFunc("/schedule/update", updateHendler)
	router.HandleFunc("/schedule/get_info/{group_number}", groupHendler)

	http.ListenAndServe(":8080", router)
}

func homeHendler(rw http.ResponseWriter, _ *http.Request) {
	rw.Write([]byte(`Пустая`))
}

func groupHendler(rw http.ResponseWriter, req *http.Request) {
	value := mux.Vars(req)
	group := value["group_number"]
	json_data := API.Get_info_about(group)              //получаем наш документ формата json
	rw.Header().Set("Content-Type", "application/json") //устанавливаем какой будет контент страницы
	rw.Write([]byte(json_data))                         //переводим в массив байтов

}
func updateHendler(rw http.ResponseWriter, req *http.Request) {
	API.Update("https://cfuv.ru/raspisanie-fakultativov-fiziko-tekhnicheskogo-instituta")
}
