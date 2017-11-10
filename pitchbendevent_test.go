package midi

import "testing"

func TestPitchBendEvent_DeltaTime(t *testing.T) {
	event := &PitchBendEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("DeltaTime() don't return nil")
	}
}

func TestPitchBendEvent_String(t *testing.T) {
	event, err := NewPitchBendEvent(nil, 1, 50)
	if err != nil {
		t.Fatal(err)
	}

	expected := "&PitchBendEvent{channel: 1, pitch: 50}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestPitchBendEvent_Serialize(t *testing.T) {
	event, err := NewPitchBendEvent(nil, 1, 50)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x00, 0xe1, 0x32, 0x00}
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

func TestPitchBendEvent_SetChannel(t *testing.T) {
	event := &PitchBendEvent{}

	err := event.SetChannel(0x10)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetChannel(0x0f)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPitchBendEvent_Channel(t *testing.T) {
	event := &PitchBendEvent{channel: 1}

	expected := uint8(1)
	actual := event.Channel()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestPitchBendEvent_SetPitch(t *testing.T) {
	event := &PitchBendEvent{}

	err := event.SetPitch(0x4fff)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetPitch(0x3fff)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPitchBendEvent_Pitch(t *testing.T) {
	event := &PitchBendEvent{pitch: 1234}

	expected := uint16(1234)
	actual := event.Pitch()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNewPitchBendEvent(t *testing.T) {
	event, err := NewPitchBendEvent(nil, 1, 1234)
	if err != nil {
		t.Fatal(err)
	}
	if event.channel != 1 {
		t.Fatalf("expected: 1 actual: %v", event.channel)
	}
	if event.pitch != 1234 {
		t.Fatalf("expected: 1234 actual: %v", event.pitch)
	}
}