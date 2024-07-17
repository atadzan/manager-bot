package processor

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/atadzan/bv-manager-bot/messages"
)

func readFromFile() (proxies []messages.Proxy) {
	proxies = []messages.Proxy{}
	file, err := os.Open(storageFilePath)
	if err != nil {
		log.Printf("Failed to open file: %s\n", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		log.Printf("Failed to open file: %s\n", err)
		return
	}

	records, err := reader.ReadAll()
	if err != nil {
		log.Printf("Failed to open file: %s\n", err)
		return
	}

	for _, record := range records {
		proxy := messages.Proxy{}
		if err := mapToStruct(header, record, &proxy); err != nil {
			log.Printf("Failed to open file: %s\n", err)
			return
		}
		proxies = append(proxies, proxy)
	}

	return
}

func mapToStruct(header []string, record []string, v interface{}) error {
	val := reflect.ValueOf(v).Elem()
	for i, field := range header {
		structField := val.FieldByName(field)
		if !structField.IsValid() {
			continue
		}

		value := record[i]
		if err := setFieldValue(structField, value); err != nil {
			return err
		}
	}
	return nil
}

func setFieldValue(field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	default:
		return fmt.Errorf("unsupported kind: %s", field.Kind())
	}
	return nil
}

func saveToFile(proxies []messages.Proxy) error {
	removeStorageFile()
	file, err := os.OpenFile(storageFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := getStructFields(proxies[0])
	if err = writer.Write(header); err != nil {
		log.Fatalf("Failed to write header to CSV file: %s", err)
	}

	for _, proxy := range proxies {
		record := structToSlice(proxy)
		if err = writer.Write(record); err != nil {
			log.Fatalf("Failed to write record to CSV file: %s", err)
		}
	}
	return nil
}

func getStructFields(v interface{}) []string {
	val := reflect.ValueOf(v)
	fields := make([]string, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		fields[i] = val.Type().Field(i).Name
	}
	return fields
}

func structToSlice(v interface{}) []string {
	val := reflect.ValueOf(v)
	values := make([]string, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		values[i] = val.Field(i).String()
	}
	return values
}

func removeStorageFile() {
	if err := os.Remove(storageFilePath); err != nil {
		log.Printf("Can't remove storage file. Err: %v", err)
	}
}
