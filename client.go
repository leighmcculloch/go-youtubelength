package youtubelength

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Client is the client that can get the length of a Youtube video.
type Client struct {
	// HTTPClient is an optional http.Client that will be used to get the
	// length of the Youtube video. If it is not set a zero http.Client with a
	// timeout of 10s will be used.
	HTTPClient *http.Client
}

func (cl *Client) httpClient() *http.Client {
	if cl == nil || cl.HTTPClient == nil {
		return &http.Client{Timeout: 10 * time.Second}
	}
	return cl.HTTPClient
}

// Get gets the length of the video.
func (cl *Client) Get(c context.Context, videoID string) (time.Duration, error) {
	u := url.URL{
		Scheme: "https",
		Host:   "www.youtube.com",
		Path:   "/get_video_info",
		RawQuery: func() string {
			q := url.Values{}
			q.Add("video_id", videoID)
			return q.Encode()
		}(),
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return 0, fmt.Errorf("creating request to %v: %v", u, err)
	}
	req = req.WithContext(c)

	res, err := cl.httpClient().Do(req)
	if err != nil {
		return 0, fmt.Errorf("doing request to %v: %v", u, err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, fmt.Errorf("reading response from %v: %v", u, err)
	}

	values, err := url.ParseQuery(string(body))
	if err != nil {
		return 0, fmt.Errorf("parsing response from %v: %v", u, err)
	}

	lengthSecsValues := values["length_seconds"]
	if numValues := len(lengthSecsValues); numValues != 1 {
		return 0, fmt.Errorf("parsing length_seconds in response from %v: %d values available but should only be one", u, numValues)
	}

	lengthSecs, err := strconv.Atoi(lengthSecsValues[0])
	if err != nil {
		return 0, fmt.Errorf("parsing length_seconds %q in response from %v: %v", lengthSecsValues[0], u, err)
	}

	return time.Duration(lengthSecs) * time.Second, nil
}
