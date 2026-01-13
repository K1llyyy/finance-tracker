package functions

import (
	"encoding/json"
	"fmt"
	"main/struc"
	"os"
	"time"
)

func Unmarshal(filename string) ([]struc.Struc, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка при чтении файла", err)
		return []struc.Struc{}, err
	}
	if len(data) == 0 {
		return []struc.Struc{}, nil
	}
	var structure []struc.Struc
	err = json.Unmarshal(data, &structure)
	if err != nil {
		fmt.Println("Ошибка при десирелизации", err)
		return []struc.Struc{}, err
	}
	return structure, nil
}

func Marshal(exp []struc.Struc, filename string) error {
	file, err := json.MarshalIndent(exp, "", " ")
	if err != nil {
		fmt.Println("Ошибка при сирелизации", err)
		return err
	}
	err = os.WriteFile(filename, file, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи файла", err)
		return err
	}
	return nil
}

func Add(data []struc.Struc, amount int, descp string, id int) []struc.Struc {
	data = append(data, struc.Struc{id, descp, amount, time.Now(), time.Now()})
	return data
}

func Inlist(data []struc.Struc, id int) int {
	for i, v := range data {
		if v.Id == id {
			return i
		}
	}
	fmt.Println("Нет операции с таким номером")
	return -1
}

func Update(exp struc.Struc, desc string, amount int) struc.Struc {
	if desc != "" && amount == -1 {
		return struc.Struc{exp.Id, desc, exp.Amount, exp.CreateADT, time.Now()}
	} else if desc == "" && amount != -1 {
		return struc.Struc{exp.Id, exp.Description, amount, exp.CreateADT, time.Now()}
	} else {
		return struc.Struc{exp.Id, desc, amount, exp.CreateADT, time.Now()}
	}
}
