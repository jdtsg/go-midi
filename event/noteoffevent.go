package event

import (
	"fmt"

	"github.com/moutend/go-midi/constant"
	"github.com/moutend/go-midi/deltatime"
)

// NoteOffEvent corresponds to note off event.
type NoteOffEvent struct {
	deltaTime     *deltatime.DeltaTime
	runningStatus bool
	channel       uint8
	note          constant.Note
	velocity      uint8
}

// deltatime.DeltaTime returns delta time of note off event.
func (e *NoteOffEvent) DeltaTime() *deltatime.DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &deltatime.DeltaTime{}
	}
	return e.deltaTime
}

// Serialize serializes note off event.
func (e *NoteOffEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, constant.NoteOff+e.channel)
	bs = append(bs, byte(e.note), e.velocity)

	return bs
}

// SetRunningStatus sets running status.
func (e *NoteOffEvent) SetRunningStatus(status bool) {
	e.runningStatus = status
}

// RunningStatus returns running status.
func (e *NoteOffEvent) RunningStatus() bool {
	return e.runningStatus
}

// SetChannel sets channel.
func (e *NoteOffEvent) SetChannel(channel uint8) error {
	if channel > 0x0f {
		return fmt.Errorf("midi: maximum channel number is 15 (0x0f)")
	}
	e.channel = channel

	return nil
}

// Channel returns channel.
func (e *NoteOffEvent) Channel() uint8 {
	return e.channel
}

// SetNote sets note.
func (e *NoteOffEvent) SetNote(note constant.Note) error {
	if note > 0x7f {
		return fmt.Errorf("midi: maximum value of note is 127 (0x7f)")
	}
	e.note = note

	return nil
}

// Note returns note.
func (e *NoteOffEvent) Note() constant.Note {
	return e.note
}

// SetVelocity sets velocity.
func (e *NoteOffEvent) SetVelocity(velocity uint8) error {
	if velocity > 0x7f {
		return fmt.Errorf("midi: maximum value of velocity is 127 (0x7f)")
	}
	e.velocity = velocity

	return nil
}

// Velocity returns velocity.
func (e *NoteOffEvent) Velocity() uint8 {
	return e.velocity
}

// String returns string representation of note off event.
func (e *NoteOffEvent) String() string {
	return fmt.Sprintf("&NoteOffEvent{channel: %v, note: %v, velocity: %v}", e.channel, e.note, e.velocity)
}

// NewNoteOffEvent returns NoteOffEvent with the given parameter.
func NewNoteOffEvent(deltaTime *deltatime.DeltaTime, channel uint8, note constant.Note, velocity uint8) (*NoteOffEvent, error) {
	var err error

	event := &NoteOffEvent{}
	event.deltaTime = deltaTime

	err = event.SetChannel(channel)
	if err != nil {
		return nil, err
	}
	err = event.SetNote(note)
	if err != nil {
		return nil, err
	}
	err = event.SetVelocity(velocity)
	if err != nil {
		return nil, err
	}
	return event, nil
}
