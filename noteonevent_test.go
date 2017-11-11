package midi

import "testing"

func TestNoteOnEvent_DeltaTime(t *testing.T) {
	event := &NoteOnEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("DeltaTime() don't return nil")
	}
}

func TestNoteOnEvent_String(t *testing.T) {
	event, err := NewNoteOnEvent(nil, 1, C3, 50)
	if err != nil {
		t.Fatal(err)
	}

	expected := "&NoteOnEvent{channel: 1, note: C3, velocity: 50}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNoteOnEvent_Serialize(t *testing.T) {
	event, err := NewNoteOnEvent(nil, 0, C3, 50)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x00, 0x90, 0x3c, 0x32}
	actual := event.Serialize()

	if len(expected) != len(actual) {
		t.Fatalf("expected: %v bytes actual: %v bytes", len(expected), len(actual))
	}
	for i, e := range expected {
		a := actual[i]
		if e != a {
			t.Fatalf("expected[%v] = 0x%x actual[%v] = 0x%x", i, e, i, a)
		}
	}
}

func TestNoteOnEvent_SetChannel(t *testing.T) {
	event := &NoteOnEvent{}

	err := event.SetChannel(0x10)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetChannel(0x0f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNoteOnEvent_Channel(t *testing.T) {
	event := &NoteOnEvent{channel: 1}

	expected := uint8(1)
	actual := event.Channel()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNoteOnEvent_SetNote(t *testing.T) {
	event := &NoteOnEvent{}

	err := event.SetNote(0x80)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetNote(0x7f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNoteOnEvent_Note(t *testing.T) {
	event := &NoteOnEvent{note: C3}

	expected := C3
	actual := event.Note()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNoteOnEvent_SetVelocity(t *testing.T) {
	event := &NoteOnEvent{}

	err := event.SetVelocity(0x80)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetVelocity(0x7f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNoteOnEvent_Velocity(t *testing.T) {
	event := &NoteOnEvent{velocity: 1}

	expected := uint8(1)
	actual := event.Velocity()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNewNoteOnEvent(t *testing.T) {
	event, err := NewNoteOnEvent(nil, 1, C3, 123)
	if err != nil {
		t.Fatal(err)
	}
	if event.channel != 1 {
		t.Fatalf("expected: 1 actual: %v", event.channel)
	}
	if event.note != C3 {
		t.Fatalf("expected: C3 actual: %v", event.note)
	}
	if event.velocity != 123 {
		t.Fatalf("expected: 123 actual: %v", event.velocity)
	}

}
