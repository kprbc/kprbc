package google

import (
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	calendar "google.golang.org/api/calendar/v3"
)

var conf = &jwt.Config{
	Email:      os.Getenv("GCAL_USER"),
	PrivateKey: []byte(os.Getenv("GCAL_PRIVATE_KEY")),
	Scopes: []string{
		calendar.CalendarEventsScope,
	},
	TokenURL: google.JWTTokenURL,
}

// Client is a Google authenticated client with the relevant authentication set
var Client = conf.Client(oauth2.NoContext)

// ListUpcomingEvents returns a list of upcoming events
func ListUpcomingEvents(calendarID string, maxResults int64) (*calendar.Events, error) {
	srv, err := calendar.New(Client)
	if err != nil {
		return nil, err
	}
	return srv.Events.
		List(calendarID).
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(time.Now().Format(time.RFC3339)).
		MaxResults(maxResults).
		OrderBy("startTime").
		Do()
}
