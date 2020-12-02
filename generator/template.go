package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const CalendarPath string = "calendar/"
const TemplateFolderPath = CalendarPath + "template"
const DayFolderPattern = CalendarPath + "day_%d"

func main() {
	cmd := flag.Int("d", time.Now().Day(), "the day to generate")
	flag.Parse()

	err := copyTemplateForDay(*cmd)
	if err != nil {
		log.Fatal(err)
	}
}

func copyTemplateForDay(dayNumber int) error {
	dayFolderName := fmt.Sprintf(DayFolderPattern, dayNumber)

	_, err := os.Stat(dayFolderName)
	folderForDayExists := err == nil

	if folderForDayExists {
		log.Printf("Skipping generation of Folder for day %d because it already exists", dayNumber)
		return nil
	}

	err = os.Mkdir(dayFolderName, os.ModePerm)
	if err != nil {
		return err
	}

	templateFiles, err := ioutil.ReadDir(TemplateFolderPath)

	for _, fileInfo := range templateFiles {
		templateFileName := fileInfo.Name()
		dayFileName := strings.Replace(templateFileName, "X", strconv.Itoa(dayNumber), 1)

		templateFile, err := os.Open(getFilenameInFolder(TemplateFolderPath, fileInfo.Name()))
		if err != nil {
			return err
		}

		dayFile, err := os.Create(getFilenameInFolder(dayFolderName, dayFileName))
		if err != nil {
			return err
		}

		_, err = io.Copy(dayFile, templateFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func getFilenameInFolder(folder string, filename string) string {
	return fmt.Sprintf("%s/%s", folder, filename)
}
