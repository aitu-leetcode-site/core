package nullable

import "fmt"

type Bool struct {
	mayEmptyT[bool]
}

func NewNullableBool(b bool) Bool {
	return Bool{
		mayEmptyT: mayEmptyT[bool]{
			t:     b,
			valid: true,
		},
	}
}

func (n *Bool) Bool() bool {
	if !n.valid {
		return false
	}
	return n.t
}

func (n *Bool) Scan(src interface{}) error {
	if src == nil {
		n.valid = false
		n.t = false
		return nil
	}

	switch typedSrc := src.(type) {
	case bool:
		n.t = typedSrc
	case int64:
		n.t = typedSrc != 0
	case []byte:
		s := string(typedSrc)
		if s == "true" || s == "1" {
			n.t = true
		} else if s == "false" || s == "0" {
			n.t = false
		} else {
			return fmt.Errorf("types: cannot scan value %q into a NullBool", s)
		}
	default:
		return fmt.Errorf("types: cannot scan value of type %T into a NullBool", src)
	}

	n.valid = true
	return nil
}
