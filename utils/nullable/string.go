package nullable

import "fmt"

type String struct {
	mayEmptyT[string]
}

func NewNullableString(s string) String {
	return String{
		mayEmptyT: mayEmptyT[string]{
			t:     s,
			valid: true,
		},
	}
}

func (s *String) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	switch typedSrc := src.(type) {
	case string:
		s.t = typedSrc
	case []byte:
		s.t = string(typedSrc)
	default:
		return fmt.Errorf("nullable: cannot scan value of type %T into a NullString", src)
	}
	s.valid = true
	return nil
}
