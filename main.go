package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

var (
	tableName      = ""
	fields         = ""
	quote          = `'`
	separator      = `,`
	querys         = ""
	fieldsArray    = []string{}
	defaultsArray  = []string{}
	dataTypesArray = []string{}
)

func main() {
	file, err := os.Open("./input/sample.csv")
	if err != nil {
		log.Fatal("Error leyendo el archivo", err)
	}
	defer file.Close()
	reader := csv.NewReader(charmap.ISO8859_15.NewDecoder().Reader(file))

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error leyendo records")
	}

	//  queries creation
	for indice, record := range records {
		if indice == 0 {
			tableName = strings.Split(record[0], ",")[0]
		}
		if indice == 1 {
			fields = record[0]
			fieldsArray = strings.Split(record[0], ",")
		}
		if indice == 2 {
			dataTypesArray = strings.Split(record[0], ",")
		}
		if indice == 3 {
			defaultsArray = strings.Split(record[0], ",")
			addCreateSentenceToQuery()
		}
		if indice > 3 {
			addInsertsToQuery(record[0])
		}
	}
	createFile()
}

func addCreateSentenceToQuery() {
	querys += fmt.Sprintf("CREATE TABLE %s (\n\tid serial NOT NULL,\n", tableName)
	for indice, field := range fieldsArray {
		separatorForCreate := ",\n"
		if indice == len(fieldsArray)-1 {
			separatorForCreate = "\n"
		}
		querys += fmt.Sprintf("\t%s %s %s"+separatorForCreate, field, dataTypesArray[indice], defaultsArray[indice])
	}
	querys += ");\n\n"
}

func addInsertsToQuery(record string) {
	stringRow := ""
	row := strings.Split(record, ",")
	for ind, field := range row {
		if ind == len(row)-1 {
			stringRow += quote + field + quote
		} else {
			stringRow += quote + field + quote + separator
		}
	}
	querys += fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);\n", tableName, fields, stringRow)
}

func createFile() {
	archivo, err := os.Create("./output/create_and_inserts_table_" + tableName)
	if err != nil {
		fmt.Println("Error al crear el achivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, querys)
	defer archivo.Close()
	println("Queries created, check in output folder.")
}
