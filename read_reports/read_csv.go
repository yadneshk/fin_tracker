package read_reports

import (
	"encoding/csv"
	"fmt"
	"os"
	"github.com/yadneshk/fin_tracker/database"
)

func ReadCsv(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file %s", path)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file %s", path)
	}
	for _, record := range records {
		fmt.Println(record)
	}
	database.InsertJSONData()
}