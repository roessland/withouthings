// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	AccountID                 int64
	WithingsUserID            string
	WithingsAccessToken       string
	WithingsRefreshToken      string
	WithingsAccessTokenExpiry time.Time
	WithingsScopes            string
	AccountUuid               uuid.UUID
}

type Notification struct {
	NotificationUuid    uuid.UUID
	AccountUuid         uuid.UUID
	ReceivedAt          time.Time
	ParamsJson          []byte
	Data                []byte
	FetchedAt           time.Time
	RawNotificationUuid uuid.UUID
	Source              string
	Params              string
}

type NotificationCategory struct {
	Appli       int32
	Scope       string
	Description string
}

type RawNotification struct {
	RawNotificationID   int64
	Source              string
	Status              string
	Data                string
	RawNotificationUuid uuid.UUID
	ReceivedAt          time.Time
	ProcessedAt         *time.Time
}

type Session struct {
	Token  string
	Data   []byte
	Expiry time.Time
}

type Subscription struct {
	SubscriptionID      int64
	Appli               int32
	Callbackurl         string
	WebhookSecret       string
	Status              string
	Comment             string
	SubscriptionUuid    uuid.UUID
	AccountUuid         uuid.UUID
	StatusLastCheckedAt time.Time
}
