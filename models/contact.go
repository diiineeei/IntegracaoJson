package models

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type Contacts struct {
	gorm.Model
	Name      string `json:"name"`
	CellPhone string `json:"cellphone"`
	Company   string `json:"company"`
}

const MACAPA = "macapa"
const VAREJAO = "varejao"

func (contact *Contacts) NormalizeContacts() {
	switch contact.Company {
	case VAREJAO:
		contact.CellPhone = strings.TrimSpace(contact.CellPhone)
	case MACAPA:
		cell := strings.TrimSpace(contact.CellPhone)
		if len(cell) == 13 {
			contact.CellPhone = fmt.Sprintf("+%s (%s) %s-%s", cell[:2], cell[2:4], cell[4:9], cell[9:13])
		}
		contact.Name = strings.ToUpper(contact.Name)
	}
}
