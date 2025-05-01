package nullable

import (
	"fmt"
	"strconv"
)

type NullFloat struct {
	mayEmptyT[float64]
}

func NewNullableFloat(f float64) NullFloat {
	return NullFloat{
		mayEmptyT: mayEmptyT[float64]{
			t:     f,
			valid: true,
		},
	}
}

func (n *NullFloat) Float() float64 {
	if !n.valid {
		return 0
	}
	return n.t
}

func (n *NullFloat) Scan(src interface{}) error {
	if src == nil {
		n.valid = false
		n.t = 0
		return nil
	}

	switch typedSrc := src.(type) {
	case float64:
		n.t = typedSrc
	case int64:
		n.t = float64(typedSrc)
	case []byte:
		f, err := strconv.ParseFloat(string(typedSrc), 64)
		if err != nil {
			return fmt.Errorf("nullable: cannot scan value %q into a NullFloat: %v", typedSrc, err)
		}
		n.t = f
	case string:
		f, err := strconv.ParseFloat(typedSrc, 64)
		if err != nil {
			return fmt.Errorf("nullable: cannot scan value %q into a NullFloat: %v", typedSrc, err)
		}
		n.t = f
	default:
		return fmt.Errorf("nullable: cannot scan value of type %T into a NullFloat", src)
	}

	n.valid = true
	return nil
}
