package port

import (
	"fmt"
	"strconv"
)

type Port int

func (p *Port) String() string {
	return strconv.Itoa(p.Int())
}

func (p *Port) Addr(host string) string {
	port := p.Int()
	if port < 0 {
		port = 0
	}
	return fmt.Sprintf("%s:%d", host, port)
}

func (p *Port) Int() int {
	return int(*p)
}
