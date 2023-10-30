package API

import (
	"schedule/GO/schedule/db"
	"schedule/GO/schedule/scrapper"
)

func Update(URL string) {
	data := scrapper.Parse(URL)
	db.Make_db(data)
}

func Get_info_about(group ...string) string {
	if len(group) == 0 {
		return db.Info_about()
	} else {
		return db.Info_about(group[0])
	}

}
