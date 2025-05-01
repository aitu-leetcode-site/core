package log

import "github.com/sirupsen/logrus"

type Field struct {
	field string
	value interface{}
}

func NewField(field string, value interface{}) Field {
	return Field{
		field: field,
		value: value,
	}
}

func convertFields(fields []Field) logrus.Fields {
	out := make(logrus.Fields)
	for _, field := range fields {
		out[field.field] = field.value
	}
	return out
}
