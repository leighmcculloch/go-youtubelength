package youtubelength_test

import (
	"context"
	"fmt"

	"4d63.com/youtubelength"
)

func ExampleGet() {
	// https://www.youtube.com/watch?v=G_OlRWGLdnw
	length, err := youtubelength.Get(context.Background(), "G_OlRWGLdnw")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(length)
	// Output: 6m51s
}
