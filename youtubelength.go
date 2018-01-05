package youtubelength

import (
	"context"
	"time"
)

// DefaultClient is a default youtubelength.Client.
var DefaultClient = Client{}

// Get gets the length of the video using the DefaultClient.
func Get(c context.Context, videoID string) (time.Duration, error) {
	return DefaultClient.Get(c, videoID)
}
