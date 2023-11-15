package main

import (
	"net/http"
	"schedule/GO/schedule/API"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/schedule", homeHendler)
	router.HandleFunc("/schedule/update", updateHendler)
	router.HandleFunc("/schedule/next/{group_number}", NextHendler)
	router.HandleFunc("/schedule/get_info/{year}/{month}/{day}/{group_number}", groupHendler)
	http.ListenAndServe(":8080", router)
}

func homeHendler(rw http.ResponseWriter, _ *http.Request) {
	rw.Write([]byte(`Пустая`))
}
func NextHendler(rw http.ResponseWriter, req *http.Request) {
	value := mux.Vars(req)
	group := value["group_number"]

	json_data := API.Next(group)
	rw.Header().Set("Content-Type", "application/json") //устанавливаем какой будет контент страницы
	rw.Write([]byte(json_data))
}

//	func teacherHendler(rw http.ResponseWriter, req *http.Request){
//		value := mux.Vars(req)
//		name := value{"name"}
//		json_data := API
//	}
func groupHendler(rw http.ResponseWriter, req *http.Request) {
	value := mux.Vars(req)
	group := value["group_number"]
	year, _ := strconv.Atoi(value["year"])
	month, _ := strconv.Atoi(value["month"])
	day, _ := strconv.Atoi(value["day"])
	json_data := API.Get_info_about(group, year, month, day) //получаем наш документ формата json
	rw.Header().Set("Content-Type", "application/json")      //устанавливаем какой будет контент страницы
	rw.Write([]byte(json_data))                              //переводим в массив байтов

}
func updateHendler(rw http.ResponseWriter, req *http.Request) {
	API.Update("https://cfuv.ru/raspisanie-fakultativov-fiziko-tekhnicheskogo-instituta")
}
