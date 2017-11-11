package midi

import "testing"

func TestEndOfTrackEvent_DeltaTime(t *testing.T) {
	event := &EndOfTrackEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("DeltaTime() don't return nil")
	}
}

func TestEndOfTrackEvent_String(t *testing.T) {
	event, err := NewEndOfTrackEvent(nil)
	if err != nil {
		t.Fatal(err)
	}

	expected := "&EndOfTrackEvent{}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestEndOfTrackEvent_Serialize(t *testing.T) {
	event := &EndOfTrackEvent{}
	expected := []byte{0x00, 0xff, 0x2f, 0x00}
	actual := event.Serialize()

	if len(expected) != len(actual) {
		t.Fatalf("expected: %v bytes actual: %v bytes", len(expected), len(actual))
	}
}

func TestNewEndOfTrackEvent(t *testing.T) {
	_, err := NewEndOfTrackEvent(nil)
	if err != nil {
		t.Fatalf("err must be always nil")
	}
}
