package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	f, _ := os.Open("annual-enterprise-survey-2021-financial-year-provisional-csv.csv")
	reader := bufio.NewReader(f)
	survey := make([]map[string]string, 0)

	firstLine := readLine(reader)
	keys := strings.Split(string(firstLine), ",")

	for line := readLine(reader); line != nil; line = readLine(reader) {
		values := strings.Split(string(line), ",")

		tmp := make(map[string]string)
		for i, key := range keys {
			tmp[key] = values[i]
		}

		survey = append(survey, tmp)
	}

	newFile, _ := os.Create("copy.json")
	defer newFile.Close()

	newFile.WriteString("[\n")

	for _, elem := range survey {
		newFile.WriteString("[")

		for _, key := range keys {
			keyJson := strings.Replace(key, `"`, "'", -1)
			valueJson := strings.Replace(elem[key], `"`, "'", -1)
			line := fmt.Sprintf(`{"%s" : "%s" },`, keyJson, valueJson)
			newFile.WriteString(line)
		}
		newFile.WriteString("],")
	}

	newFile.WriteString("]")

}

func readLine(reader *bufio.Reader) (line []byte) {
	line, _, _ = reader.ReadLine()
	return
}
