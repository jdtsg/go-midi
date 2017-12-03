package midi

import "fmt"

// Quantity represents variable length quantity in MIDI.
type Quantity struct {
	value []byte
}

// SetUint32 sets value.
func (q *Quantity) SetUint32(u32 uint32) error {
	if u32 > 0x0fffffff {
		return fmt.Errorf("midi: 0x%x is larger than 0xffffff", u32)
	}

	q.value = []byte{}
	mask := uint32(0xfe00000)

	for i := uint32(21); i >= 7; i -= 7 {
		b := byte((u32&mask)>>i) + 0x80
		mask = mask >> 7
		if b > 0x80 {
			q.value = append(q.value, byte(b))
		}
	}

	q.value = append(q.value, byte(u32&0x7f))

	return nil
}

// Uint32 returns value as uint32.
func (q *Quantity) Uint32() uint32 {
	var u32 uint32

	for i, b := range q.Value() {
		u := uint32(b) & 0x7f
		j := len(q.value) - i - 1
		shift := (uint32(j * 8)) - uint32(j)
		u = u << shift
		u32 += u
	}

	return u32
}

// SetValue sets value.
func (q *Quantity) SetValue(value []byte) error {
	if len(value) > 4 {
		return fmt.Errorf("midi: maximum length of byte slice is 4, but len(value) = %v", len(value))
	}
	q.value = value

	return nil
}

// Value returns value.
func (q *Quantity) Value() []byte {
	if q.value == nil {
		q.value = make([]byte, 1)
	}
	return q.value
}

// Serialize serializes value of variable length quantity.
func (q *Quantity) Serialize() []byte {
	return q.Value()
}
