package base_type

import (
	"database/sql"
	"encoding/json"
)

type NullableBool struct {
	sql.NullBool
}

func (n NullableBool) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Bool)
	}
	return json.Marshal(nil)
}

func (n *NullableBool) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Valid = false
		return nil
	}

	var value bool
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	n.Bool = value
	n.Valid = true
	return nil
}

type NullableString struct {
	sql.NullString
}

func (n NullableString) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.String)
	}
	return json.Marshal(nil)
}

func (n *NullableString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Valid = false
		return nil
	}

	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	n.String = value
	n.Valid = true
	return nil
}
