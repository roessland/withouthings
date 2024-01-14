package subscription

import (
	"fmt"
	"github.com/google/uuid"
	"net/url"
	"strconv"
	"time"
)

// RawNotification represents a notification received from the Withings API.
type RawNotification struct {

	// rawNotificationUUID is the randomly created UUID of the RawNotification.
	rawNotificationUUID uuid.UUID

	// source is the IP address of the Withings API server that sent the notification.
	source string

	// data is the HTTP body of the POST request from Withings.
	// For webhook tests, it is an empty string.
	// For example: `userid=133337&startdate=1682809920&enddate=1682837220&appli=44`
	data string

	// status is the status of the RawNotification.
	status RawNotificationStatus

	// receivedAt is the time the notification was received.
	receivedAt time.Time

	// processedAt is the time the notification was processed.
	// If the notification has not been processed yet, it is nil.
	processedAt *time.Time
}

type RawNotificationData struct {
	WithingsUserID string
	StartDate      time.Time
	EndDate        time.Time
	Appli          int
}

// RawNotificationStatus represents the status of a RawNotification.
type RawNotificationStatus string

const (
	// RawNotificationStatusPending means the RawNotification has not been processed yet,
	// and a Notification has not been created yet.
	RawNotificationStatusPending RawNotificationStatus = "pending"

	// RawNotificationStatusProcessed means the RawNotification has been processed,
	// and a Notification has been created.
	RawNotificationStatusProcessed RawNotificationStatus = "processed"
)

// isValid returns true if the RawNotificationStatus is a valid status.
func (rns RawNotificationStatus) isValid() bool {
	return rns == RawNotificationStatusPending || rns == RawNotificationStatusProcessed
}

// MustRawNotificationStatusFromString returns a RawNotificationStatus from a string,
// or panics if the string is not a valid status.
func MustRawNotificationStatusFromString(s string) RawNotificationStatus {
	if RawNotificationStatus(s).isValid() {
		return RawNotificationStatus(s)
	}
	panic(fmt.Sprintf("invalid RawNotificationStatus: %s", s))
}

// NewRawNotification validates input and returns a new RawNotification
func NewRawNotification(rawNotificationUUID uuid.UUID, source string, data string, status RawNotificationStatus, receivedAt time.Time, processedAt *time.Time) RawNotification {
	if rawNotificationUUID == uuid.Nil {
		panic("rawNotificationUUID cannot be nil")
	}
	if status != RawNotificationStatusPending && status != RawNotificationStatusProcessed {
		panic(fmt.Sprintf("invalid status: %s", status))
	}

	if status == RawNotificationStatusPending && processedAt != nil {
		msg := fmt.Sprintf("processedAt must be nil if status is pending, got %s / %v", status, processedAt)
		panic(msg)
	}

	if status == RawNotificationStatusProcessed && processedAt == nil {
		panic("processedAt cannot be nil if status is processed")
	}

	return RawNotification{
		rawNotificationUUID: rawNotificationUUID,
		source:              source,
		data:                data,
		status:              status,
		receivedAt:          receivedAt,
		processedAt:         processedAt,
	}
}

// UUID returns the UUID.
func (r RawNotification) UUID() uuid.UUID {
	return r.rawNotificationUUID
}

// Source returns the source.
func (r RawNotification) Source() string {
	return r.source
}

// Data returns the data.
func (r RawNotification) Data() string {
	return r.data
}

func (r RawNotification) ParsedData() (RawNotificationData, error) {
	queryStr := r.data
	params, err := url.ParseQuery(queryStr)
	if err != nil {
		return RawNotificationData{}, fmt.Errorf("failed to parse data: %w", err)
	}

	startUnix, err := strconv.ParseInt(params.Get("startdate"), 10, 64)
	if err != nil {
		return RawNotificationData{}, fmt.Errorf("failed to parse startdate: %w", err)
	}

	endUnix, err := strconv.ParseInt(params.Get("enddate"), 10, 64)
	if err != nil {
		return RawNotificationData{}, fmt.Errorf("failed to parse enddate: %w", err)
	}

	appli, err := strconv.Atoi(params.Get("appli"))
	if err != nil {
		return RawNotificationData{}, fmt.Errorf("failed to parse appli: %w", err)
	}

	return RawNotificationData{
		WithingsUserID: params.Get("userid"),
		StartDate:      time.Unix(startUnix, 0),
		EndDate:        time.Unix(endUnix, 0),
		Appli:          appli,
	}, nil
}

// Status returns the status.
func (r RawNotification) Status() RawNotificationStatus {
	return r.status
}

func (r RawNotification) ReceivedAt() time.Time {
	return r.receivedAt
}

func (r RawNotification) ProcessedAt() *time.Time {
	return r.processedAt
}

func (r *RawNotification) MarkAsProcessed() {
	now := time.Now()
	r.status = RawNotificationStatusProcessed
	r.processedAt = &now
}
