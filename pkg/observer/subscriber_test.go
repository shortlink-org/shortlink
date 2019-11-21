package observer

import (
	"reflect"
	"testing"
	"time"
)

type mockWriter struct {
	data []string
}

func (mw *mockWriter) Write(b []byte) (n int, err error) {
	mw.data = append(mw.data, string(b))
	return len(b), nil
}

func TestSubscriber_Notify(t *testing.T) {
	excepted := make([]string, 50)
	subs := make([]subscriber, 50)

	for i := 0; i < 50; i++ {
		store := mockWriter{}
		subs[i] = NewSubscriber(&store)
	}

	msg := "test msg"
	for i := 0; i < 50; i++ {
		excepted[i] = msg
		for _, sub := range subs {
			sub.Notify(msg)
		}
	}
	time.Sleep(1 * time.Second)

	for _, sub := range subs {
		if !reflect.DeepEqual(sub.store.(*mockWriter).data, excepted) {
			t.Errorf("excepted:%v got:%v", sub.store.(*mockWriter).data, excepted)
		}
	}
}
