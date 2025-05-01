package nullable

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

type mayEmptyT[T any] struct {
	t     T
	valid bool
}

func (m *mayEmptyT[T]) Valid() bool {
	if m.valid {
		return true
	}
	isValid := !reflect.ValueOf(m.t).IsZero()
	if isValid {
		m.valid = true
	}
	return isValid
}

func (m *mayEmptyT[T]) String() string {
	return fmt.Sprint(m.t)
}

func (m *mayEmptyT[T]) MarshalJSON() ([]byte, error) {
	if !m.Valid() {
		return nil, nil
	}
	return json.Marshal(m.t)
}

func (m *mayEmptyT[T]) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	m.valid = true
	return json.Unmarshal(data, &m.t)
}

func (m *mayEmptyT[T]) Value() (driver.Value, error) {
	if !m.Valid() {
		return nil, nil
	}
	return m.t, nil
}

func (m *mayEmptyT[T]) RawValue() (T, bool) {
	return m.t, m.valid
}

func (m *mayEmptyT[T]) Scan(_ interface{}) error {
	panic("implement me")
}
