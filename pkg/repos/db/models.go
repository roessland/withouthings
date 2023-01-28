// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"time"
)

type Account struct {
	AccountID                 int64
	WithingsUserID            string
	WithingsAccessToken       string
	WithingsRefreshToken      string
	WithingsAccessTokenExpiry time.Time
	WithingsScopes            string
}

type NotificationCategory struct {
	Appli       int32
	Scope       sql.NullString
	Description sql.NullString
}

type RawNotification struct {
	RawNotificationID int64
	Source            string
	Status            string
	Data              string
}

type Session struct {
	Token  string
	Data   []byte
	Expiry time.Time
}

type Subscription struct {
	SubscriptionID int64
	AccountID      int64
	Appli          int32
	Callbackurl    string
	WebhookSecret  string
	Status         string
	Comment        string
}
