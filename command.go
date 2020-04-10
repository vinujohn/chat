package chat

import "fmt"

type Command interface {
	Format() []byte
}

type MSG struct {
	To      string
	Payload []byte
}

// Format - MSG <entity-id> <payload-length>\n<payload>
func (c MSG) Format() []byte {
	header := fmt.Sprintf("MSG %s %d\n", c.To, len(c.Payload))
	return append([]byte(header), c.Payload...)
}

type OK struct{}

// Format - OK
func (c OK) Format() []byte {
	return []byte("OK")
}
