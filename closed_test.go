package chanclosed

import "testing"

func TestChannel(t *testing.T) {
	testChan[int8](t)
	testChan[uint8](t)
	testChan[int16](t)
	testChan[uint16](t)
	testChan[int32](t)
	testChan[uint32](t)
	testChan[int64](t)
	testChan[uint64](t)
	testChan[int](t)
	testChan[uint](t)
	testChan[float32](t)
	testChan[float64](t)
	testChan[complex64](t)
	testChan[complex128](t)
	testChan[struct{}](t)
	testChan[string](t)
	testChan[[]byte](t)

	type myStruct struct {
		a int
		b string
	}

	testChan[myStruct](t)
}

func testChan[B any](t *testing.T) {
	ch := make(chan B)
	if Closed(ch) {
		t.Fatalf("chan T is not closed")
	}
	close(ch)

	if !Closed(ch) {
		t.Fatalf("chan T is closed")
	}

	chPtr := make(chan *B)
	if Closed(chPtr) {
		t.Fatalf("chan T is not closed")
	}
	close(chPtr)

	if !Closed(ch) {
		t.Fatalf("chan T is closed")
	}
}
