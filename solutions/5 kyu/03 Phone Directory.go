package kata

import "regexp"
import (
	"fmt"
	"strings"
)

type phoneRecord struct {
	phone   string
	name    string
	address string
}

func (p phoneRecord) ToString() string {
	return fmt.Sprintf("Phone => %v, Name => %v, Address => %v", p.phone, p.name, p.address)
}

//Phone ...
func Phone(dir, num string) string {

	numCount := strings.Count(dir, "+"+num)

	if numCount == 0 {
		return "Error => Not found: " + num
	} else if numCount > 1 {
		return "Error => Too many people: " + num
	}

	record := findRecordInBook(dir, num)
	result := parseBookRecord(record)
	return result.ToString()
}

//findRecordInBook ...
func findRecordInBook(dir, num string) string {

	records := strings.Split(dir, "\n")

	for _, record := range records {
		if strings.Count(record, "+"+num) == 1 {
			return record
		}
	}

	return ""
}

func parseBookRecord(bookRecord string) phoneRecord {
	record := phoneRecord{
		phone:   "",
		name:    "",
		address: "",
	}

	//Phone
	regPhone, _ := regexp.Compile(`\+(\S){1,}`) /*\+[0-9_.-]**/
	record.phone = regPhone.FindString(bookRecord)
	bookRecord = strings.ReplaceAll(bookRecord, record.phone, "")
	record.phone = strings.TrimLeft(record.phone, "+")
	record.phone = string(record.phone[:strings.LastIndex(record.phone, "-")+5])

	//Name
	regName, _ := regexp.Compile(`\<(\S).+>`)
	record.name = regName.FindString(bookRecord)
	bookRecord = strings.ReplaceAll(bookRecord, record.name, "")
	record.name = strings.TrimLeft(record.name, "<")
	record.name = strings.TrimRight(record.name, ">")

	//Address
	record.address = strings.TrimSpace(bookRecord)
	regAddressTrash, _ := regexp.Compile(`([^A-Za-z0-9.\-]+)`)
	record.address = regAddressTrash.ReplaceAllString(record.address, " ")

	if strings.Index(record.address, "  ") > 0 {
		record.address = strings.ReplaceAll(record.address, "  ", " ")
	}
	record.address = strings.ReplaceAll(record.address, "  ", " ")
	record.address = strings.TrimSpace(record.address)

	return record
}
