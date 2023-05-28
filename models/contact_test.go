package models

import (
	"testing"
)

func TestNormalizeVarejao(t *testing.T) {
	type fields struct {
		Name      string
		CellPhone string
		Company   string
	}
	tests := []struct {
		name     string
		fields   fields
		expected fields
	}{
		{
			name: "Case 1",
			fields: fields{
				Company:   "macapa",
				Name:      "Marina Rodrigues",
				CellPhone: "5541996941919",
			},
			expected: fields{
				Company:   "macapa",
				Name:      "MARINA RODRIGUES",
				CellPhone: "5541996941919",
			},
		},
		{
			name: "Case 2",
			fields: fields{
				Company:   "varejao",
				Name:      "Nicolas Rodrigues",
				CellPhone: "5541954122723",
			},
			expected: fields{
				Company:   "varejao",
				Name:      "NICOLAS RODRIGUES",
				CellPhone: "5541954122723",
			},
		}, {
			name: "Case 3",
			fields: fields{
				Company:   "macapa",
				Name:      "Srta. Isabelly Castro",
				CellPhone: "5541959365078",
			},
			expected: fields{
				Company:   "macapa",
				Name:      "SRTA. ISABELLY CASTRO",
				CellPhone: "5541959365078",
			},
		},
		{
			name: "Case 4",
			fields: fields{
				Company:   "varejao",
				Name:      "Ana Julia da Rocha",
				CellPhone: "5541954122723",
			},
			expected: fields{
				Company:   "varejao",
				Name:      "Ana Julia da Rocha",
				CellPhone: "5541954122723",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			contact := &Contacts{
				Name:      tt.fields.Name,
				CellPhone: tt.fields.CellPhone,
				Company:   tt.fields.Company,
			}
			contact.NormalizeContacts()
		})
	}
}
