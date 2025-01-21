package repository

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type companyType int

const (
	Corporations companyType = iota
	NonProfit
	Cooperative
	SoleProprietorship
	Undefined
)

// Value - Implementation of valuer for database/sql
func (ct companyType) Value() (driver.Value, error) {
	return ct.String(), nil
}

// Scan - Implement the database/sql scanner interface
func (ct *companyType) Scan(value interface{}) error {
	if value == nil {
		*ct = Undefined
		return nil
	}
	if str_type, err := driver.String.ConvertValue(value); err == nil {

		if v, ok := str_type.(string); ok {
			*ct, err = parseStringFromCT(v)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("failed to scan YesNoEnum")
}

func parseStringFromCT(input string) (companyType, error) {
	switch input {
	case "Corporations":
		return Corporations, nil
	case "NonProfit":
		return NonProfit, nil
	case "Cooperative":
		return Cooperative, nil
	case "Sole Proprietorship":
		return SoleProprietorship, nil
	default:
		return Undefined, errors.New("invalid company type")
	}
}

func (ct *companyType) String() string {
	switch *ct {
	case 0:
		return "Corporations"
	case 1:
		return "NonProfit"
	case 2:
		return "Cooperative"
	case 3:
		return "Sole Proprietorship"
	default:
		return "error"
	}
}

func (t *companyType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	// switch s {
	// case "Corporations":
	// 	*t = Corporations
	// case "NonProfit":
	// 	*t = NonProfit
	// case "Cooperative":
	// 	*t = Cooperative
	// case "Sole Proprietorship":
	// 	*t = SoleProprietorship
	// default:
	// 	return fmt.Errorf("invalid company type")
	// }

	//2nd way: use of parseStringFromCT, ,to test
	ct, err := parseStringFromCT(s)
	if err != nil {
		return err
	}
	*t = ct

	return nil
}

// MarshalJSON custom implementation
func (t companyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

type Company struct {
	Uuid              string      `json:"uuid"`
	Name              string      `json:"name"`
	Description       string      `json:"description"`
	AmountOfEmployees int         `json:"amount_of_employees"`
	Registered        bool        `json:"registered"`
	Type              companyType `json:"type"`
	SysCreationDate   time.Time   `json:"sys_creation_date"`
	SysUpdateDate     time.Time   `json:"sys_update_date"`
}
