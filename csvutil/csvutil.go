package csvutil

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/kishore-tadapaneni/utils/model"
)

func ReadFile(filepath string) ([]model.Employee, error) {
	csvFile, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var employees []model.Employee
	for _, row := range records {
		if len(row) == 3 {
			name := row[0]
			ageInt := 0
			fmt.Sscanf(row[1], "%d", &ageInt)
			email := row[2]

			employee := model.Employee{
				Name:  name,
				Age:   ageInt,
				Email: email,
			}
			employees = append(employees, employee)
		}
	}

	return employees, nil
}
