package youtubelength

import (
	"context"
	"testing"
)

func TestClient_Get(t *testing.T) {
	client := Client{}

	for _, v := range videos {
		t.Run(v.id, func(t *testing.T) {
			length, err := client.Get(context.Background(), v.id)
			if err != nil {
				t.Fatalf("got error %v", err)
			}
			if length == v.length {
				t.Logf("got %v, want %v", length, v.length)
			} else {
				t.Errorf("got %v, want %v", length, v.length)
			}
		})
	}
}
