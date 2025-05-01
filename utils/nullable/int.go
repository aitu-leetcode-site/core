package nullable

import "fmt"

type Int struct {
	mayEmptyT[int]
}

func NewNullableInt(i int) Int {
	return Int{
		mayEmptyT: mayEmptyT[int]{
			t:     i,
			valid: true,
		},
	}
}

func (i *Int) Int() int {
	return i.t
}
func (i *Int) Scan(src interface{}) error {
	if src == nil {
		i.valid = false
		i.t = 0
		return nil
	}

	switch typedSrc := src.(type) {
	case int64:
		i.t = int(typedSrc)
	case int:
		i.t = typedSrc
	case float64:
		i.t = int(typedSrc)
	case []byte:
		var ii int
		_, err := fmt.Sscanf(string(typedSrc), "%d", &ii)
		if err != nil {
			return fmt.Errorf("nullable: cannot scan value %q into a NullInt: %v", typedSrc, err)
		}
		i.t = ii
	default:
		return fmt.Errorf("nullable: cannot scan value of type %T into a NullInt", src)
	}
	i.valid = true
	return nil
}
