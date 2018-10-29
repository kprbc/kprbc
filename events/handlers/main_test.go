package handlers

import (
	"reflect"
	"testing"

	calendar "google.golang.org/api/calendar/v3"
)

func TestGetUpcomingEvents(t *testing.T) {
	tests := []struct {
		name string
		want *calendar.Events
	}{
		// TODO: Add test cases.
		{name: "First test", want: &calendar.Events{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUpcomingEvents(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUpcomingEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}
