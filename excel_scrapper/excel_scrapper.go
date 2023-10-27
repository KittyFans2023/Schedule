package excel_scrapper

import (
	"fmt"
	"slices"
	"strings"

	"github.com/xuri/excelize/v2"
)

func ERROR(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func read_file(PATH string) [][]string {
	f, err := excelize.OpenFile(PATH)
	ERROR(err)
	defer func() {
		err := f.Close()
		ERROR(err)
	}()
	rows, err := f.GetRows("курс 1 ПИ ")
	ERROR(err)
	return rows
}
func to_pretty(elem []string) []string {
	for i := 0; i < 20; i++ {
		elem = append(elem, "")
	}
	elem = elem[0:22]
	elem = slices.Delete(elem, 11, 14) //"сращиваем обе части таблицы(четную и нечетную)
	lenght := len(elem)
	for i := 0; i < 20-lenght; i++ {
		elem = append(elem, "") //дополняем массив, чтобы в случае чего были обнаружены "окна" в парах
	}
	return elem
}

func read_schedule() map[string]map[int][][]string {
	PATH := "excel_scrapper/PI.xlsx"
	rows := read_file(PATH)
	week := []string{"понедельник", "вторник", "среда", "четверг", "пятница", "суббота"}
	lessons_type := []string{"лк", "пз"} //тип пары
	//день недели
	lt := ""                                                                                                                                                                                                                       // Тип пары_2
	together := []string{"высшая математика", "основы российской государственности", "иностранный язык(немецкий,французкий)", "физическая культура", "история россии/с 13.10", "история россии/с 09.10", "история россии/с 16.10"} //предметы для которых практические вместе

	var all_info map[string]map[int][][]string //вся информация. КЛЮЧ_1 - День. Ключ_2-номер пары, значение-массив pair
	all_info = make(map[string]map[int][][]string)

	comment := ""   //комментарий преподавателя
	day := ""       //день недели
	number := 1     //номер пары
	subject_1 := "" //предмет первой подгруппы
	teacher_1 := "" //преподаватель первой подгруппы
	address_1 := "" //адрес первой подгруппы

	subject_2 := "" //предмет второй подгруппы
	teacher_2 := "" //преподаватель второй подгруппы
	address_2 := "" //адрес второй подгруппы

	var pair map[int][][]string //массив всех предметов за number пару у каждой группы/подгруппы
	pair = make(map[int][][]string)

	var new_day_index []int //при смене дня недели записывает индекс строки, где он меняется
	for index, elem := range rows {

		if index%5 != 4 { //ускоряет работы, минуя просмотр ненужных строк
			continue
		}
		if len(elem) == 0 { //если длина строки столбца равна 0, то мы ее скипаем
			continue
		}

		teach := rows[index+1]
		addr := rows[index+2]

		teach = to_pretty(teach)
		addr = to_pretty(addr)
		elem = to_pretty(elem)

		if slices.Contains(week, strings.ToLower(elem[0])) {
			new_day_index = append(new_day_index, index) //если первый элемент строки-день неделю, то его индекс записываем
		}
		if len(new_day_index) != 0 { //если длинна не равна 0, то значит, что как минимум понедельник уже был
			d := strings.ToLower(rows[new_day_index[len(new_day_index)-1]][0])
			if day != d { //сравниваем нынешний день с последним записанным

				if day != "" { //если день не пуст, то значит предыдущий день закончился
					all_info[strings.ToLower(day)] = pair //записываем данные
					pair = make(map[int][][]string)       //создаем новый массив для этого дня

				}
				number = 1                                                          //начинаем с 1 пары
				day = strings.ToLower(rows[new_day_index[len(new_day_index)-1]][0]) //начинаем новый день
				if day == "суббота" {
					return all_info
				}
			} else if strings.Contains("1234567890", elem[1]) {

				number++ //если в элементе есть числа-элемент содержит номер пары -> увеличиваем на единицу
			}

			for i := 2; i < len(elem); i++ { //начинаем сразу со второго индекса, чтобы начать с "типа лекции"
				nlt := strings.ToLower(elem[i]) //нынешний элемент

				if slices.Contains(lessons_type, nlt) { //если он показывает тип, то выясняем какой именно
					if strings.Contains(" лк ", nlt) {
						lt = "лк"
						continue
					} else if strings.Contains(" пз ", nlt) {
						lt = "пз"
						continue
					}
				}
				if lt == "лк" || (lt == "пз" && slices.Contains(together, nlt)) { //если тип лекции или ПЗ,но те,которые проходят всей группой,то
					subject_1 = elem[i] //записываем их для обеих подгрупп
					teacher_1 = teach[i]
					address_1 = addr[i]

					subject_2 = elem[i]
					teacher_2 = rows[index+1][i]
					address_2 = addr[i]

				} else if lt == "пз" { //если обычное пз, то взависимости от того, кому принадлежат
					teach = append(teach, "") //чтобы избежать ошибки при попытке доступа к элементу
					addr = append(addr, "")   //добавляем пустую строку
					subject_1 = elem[i]
					teacher_1 = teach[i]
					address_1 = addr[i]

					subject_2 = elem[i+1]
					teacher_2 = teach[i+1]
					address_2 = addr[i+1]

				} else if lt == "" && (i == 2 || i == 5 || i == 8 || i == 11 || i == 14 || i == 17) { // если данная ячейка-это ячейка, в которой хранится тип
					//в таком случае, если она пуста, то пар на эту пару нет
					subject_1 = "" //предмет первой подгруппы
					teacher_1 = "" //преподаватель первой подгруппы
					address_1 = "" //адрес первой подгруппы

					subject_2 = "" //предмет второй подгруппы
					teacher_2 = "" //преподаватель второй подгруппы
					address_2 = "" //адрес второй подгруппы
				} else if lt == "" {
					continue
				}
				pair[number] = append(pair[number], []string{subject_1, teacher_1, address_1, lt, comment})
				pair[number] = append(pair[number], []string{subject_2, teacher_2, address_2, lt, comment})
				lt = ""
			} //Информация за number пару

		}
	}
	return all_info

}

func Get_information() map[string]map[string]map[int][][]string {
	all_info := read_schedule()
	var re_info map[string]map[string]map[int][][]string
	re_info = make(map[string]map[string]map[int][][]string)
	groups := []string{"231-1", "231-2", "232-1", "232-2", "233-1", "233-2"}

	for i := 0; i < len(groups); i++ {
		elem := groups[i]
		re_info[elem] = map[string]map[int][][]string{}
		for day := range all_info {
			re_info[elem][day] = map[int][][]string{}
			for number, subject := range all_info[day] {
				re_info[elem][day][number] = make([][]string, 0)
				re_info[elem][day][number] = [][]string{subject[i], subject[i+6]}
			}
		}
	}
	return re_info
}
