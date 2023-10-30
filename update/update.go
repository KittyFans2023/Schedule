package update

import (
	"schedule/GO/schedule/db"
	"schedule/GO/schedule/scrapper"
)

func Update(URL string) {
	data := scrapper.Parse(URL)
	db.Make_db(data)
}

func Get_info_about(group ...string) {
	if len(group) == 0 {
		db.Info_about()
	} else {
		db.Info_about(group[0])
	}

}
