package handlers

import (
	"time"

	"github.com/kprbc/kprbc/google"
	"github.com/sirupsen/logrus"
	calendar "google.golang.org/api/calendar/v3"
)

// GetUpcomingEvents retrieves the next few events coming up
func GetUpcomingEvents() (*calendar.Events, error) {
	service, err := calendar.New(google.Client)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	events, err := service.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(time.Now().Format(time.RFC3339)).MaxResults(5).OrderBy("startTime").Do()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return events, nil
}
