package event

import "testing"

func TestCuePointEventDeltaTime(t *testing.T) {
	event := &CuePointEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("deltatime.DeltaTime() don't return nil")
	}
}

func TestCuePointEvent_String(t *testing.T) {
	event, err := NewCuePointEvent(nil, []byte("text"))
	if err != nil {
		t.Fatal(err)
	}

	expected := "&CuePointEvent{text: \"text\"}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestCuePointEvent_Serialize(t *testing.T) {
	event, err := NewCuePointEvent(nil, []byte("text"))
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0xff, 0x07, 0x04, 0x74, 0x65, 0x78, 0x74}
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

func TestCuePointEvent_SetText(t *testing.T) {
	event := &CuePointEvent{}

	err := event.SetText(bigdata)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	err = event.SetText(bigdata[1:])
	if err != nil {
		t.Fatal(err)
	}
}

func TestCuePointEvent_Text(t *testing.T) {
	event := &CuePointEvent{}

	expected := ""
	actual := string(event.Text())
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}

	event = &CuePointEvent{text: []byte("text")}

	expected = "text"
	actual = string(event.Text())

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNewCuePointEvent(t *testing.T) {
	_, err := NewCuePointEvent(nil, bigdata)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	event, err := NewCuePointEvent(nil, []byte("text"))
	if err != nil {
		t.Fatal(err)
	}
	expected := []byte("text")
	actual := event.text

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
