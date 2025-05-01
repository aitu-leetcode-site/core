package nullable

import (
	"fmt"
	"time"
)

type Time struct {
	mayEmptyT[time.Time]
}

func NewNullableTime(t time.Time) Time {
	return Time{
		mayEmptyT: mayEmptyT[time.Time]{
			t:     t,
			valid: true,
		},
	}
}

func (n *Time) Time() time.Time {
	if !n.valid {
		return time.Time{}
	}
	return n.t
}

func (n *Time) Scan(src interface{}) error {
	if src == nil {
		n.valid = false
		n.t = time.Time{}
		return nil
	}

	switch typedSrc := src.(type) {
	case time.Time:
		n.t = typedSrc
	case []byte:
		parsed, err := time.Parse(time.RFC3339, string(typedSrc))
		if err != nil {
			return fmt.Errorf("types: cannot scan value %q into NullTime: %w", typedSrc, err)
		}
		n.t = parsed
	case string:
		parsed, err := time.Parse(time.RFC3339, typedSrc)
		if err != nil {
			return fmt.Errorf("types: cannot scan value %q into NullTime: %w", typedSrc, err)
		}
		n.t = parsed
	default:
		return fmt.Errorf("types: cannot scan value of type %T into NullTime", src)
	}

	n.valid = true
	return nil
}
