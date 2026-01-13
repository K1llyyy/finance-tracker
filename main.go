package main

import (
	"flag"
	"fmt"
	"main/functions"
	"os"
	"time"
)

func main() {
	filename := "file.json"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0664)
	if err != nil {
		fmt.Println("Ошибка при открытии файла: ", err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Ошибка при закрытии файла", err)
		}
	}()
	data, err := functions.Unmarshal(filename)
	if err != nil {
		return
	}
	defer func() {
		err := functions.Marshal(data, filename)
		if err != nil {
			return
		}
	}()

	var id int
	userinfo := os.Args[1:]
	if len(userinfo) == 0 {
		fmt.Println("Проверьте список  команд")
		return
	}
	switch userinfo[0] {
	case "add":
		addcmd := flag.NewFlagSet("add", flag.ExitOnError)
		amount := addcmd.Int("amount", -1, "цена покупки")
		description := addcmd.String("description", "", "description")
		err = addcmd.Parse(userinfo[1:])
		if err != nil {
			fmt.Println("Ошибка в флагах:", err)
		}
		if *amount == -1 || *description == "" {
			fmt.Println("Вы не ввели аргументы")
			return
		}
		if len(data) == 0 {
			id = 0
		} else {
			id = data[len(data)-1].Id + 1
		}
		data = functions.Add(data, *amount, *description, id)
		fmt.Println("Вы успешно создали операцию")
	case "list":
		for _, v := range data {
			fmt.Printf("Номер операции: %d\nДата операции: %s\nОписание операции: %s\nСумма операции: %d\n", v.Id, v.CreateADT.Format(time.ANSIC), v.Description, v.Amount)
		}

	case "delete":
		deletecmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deletecmd.Int("id", -1, "Номер операции")
		err = deletecmd.Parse(userinfo[1:])
		if err != nil {
			fmt.Println("Ошибка в флагах:", err)
		}
		if *id == -1 {
			fmt.Println("Вы не ввели аргумент")
			return
		}
		i := functions.Inlist(data, *id)
		if i == -1 {
			return
		}
		data = append(data[:i], data[i+1:]...)
		fmt.Println("Ваша операция успешно удалена")
	case "summary":
		sum := 0
		summarycmd := flag.NewFlagSet("summary", flag.ExitOnError)
		month := summarycmd.Int("month", -1, "Месяц операций")
		err = summarycmd.Parse(userinfo[1:])
		if err != nil {
			fmt.Println("Ошибка в флагах:", err)
		}
		if *month == -1 {
			if len(data) == 0 {
				fmt.Println("У вас нет операций")
				return
			}
			for _, v := range data {
				sum += v.Amount
			}
			fmt.Printf("ОБщая сумма всех операций %d", sum)
		} else if *month > 12 || (*month < 1) {
			fmt.Println("Вы ввели невалидный месяц")
		} else {
			if len(data) == 0 {
				fmt.Println("У вас нет операций")
				return
			}
			curyear := time.Now().Year()
			for _, v := range data {
				if int(v.CreateADT.Month()) == *month && v.CreateADT.Year() == curyear {
					sum += v.Amount
				}
			}
			fmt.Printf("ОБщая сумма всех операций %d, за месяц %d текущего года", sum, *month)
		}
	case "update":
		updatecmd := flag.NewFlagSet("update", flag.ExitOnError)
		id := updatecmd.Int("id", -1, "id")
		description := updatecmd.String("description", "", "new description")
		amount := updatecmd.Int("amount", -1, "new amount")
		err = updatecmd.Parse(userinfo[1:])
		if err != nil {
			fmt.Println("Ошибка в флагах:", err)
		}
		if *id == -1 {
			fmt.Println("Вы не ввели номер операции")
			return
		}
		if *description == "" && *amount == -1 {
			fmt.Println("Вы не ввели значения для изменения")
			return
		} else {
			i := functions.Inlist(data, *id)
			if i == -1 {
				return
			}
			data[i] = functions.Update(data[i], *description, *amount)
		}
		fmt.Println("Операция успешно обновлена")
	default:
		fmt.Println("Неизвестная команда")
	}
}
