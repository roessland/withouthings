// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: subscription.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const allNotificationCategories = `-- name: AllNotificationCategories :many
SELECT appli, scope, description
FROM notification_category
ORDER BY appli
`

func (q *Queries) AllNotificationCategories(ctx context.Context) ([]NotificationCategory, error) {
	rows, err := q.db.Query(ctx, allNotificationCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []NotificationCategory
	for rows.Next() {
		var i NotificationCategory
		if err := rows.Scan(&i.Appli, &i.Scope, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createRawNotification = `-- name: CreateRawNotification :exec
INSERT INTO raw_notification (raw_notification_uuid, source, status, data)
VALUES ($1, $2, $3, $4)
`

type CreateRawNotificationParams struct {
	RawNotificationUuid uuid.UUID
	Source              string
	Status              string
	Data                string
}

func (q *Queries) CreateRawNotification(ctx context.Context, arg CreateRawNotificationParams) error {
	_, err := q.db.Exec(ctx, createRawNotification,
		arg.RawNotificationUuid,
		arg.Source,
		arg.Status,
		arg.Data,
	)
	return err
}

const createSubscription = `-- name: CreateSubscription :exec
INSERT INTO subscription(subscription_uuid,
                         account_uuid,
                         appli,
                         callbackurl,
                         webhook_secret,
                         status,
                         comment,
                         status_last_checked_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

type CreateSubscriptionParams struct {
	SubscriptionUuid    uuid.UUID
	AccountUuid         uuid.UUID
	Appli               int32
	Callbackurl         string
	WebhookSecret       string
	Status              string
	Comment             string
	StatusLastCheckedAt time.Time
}

func (q *Queries) CreateSubscription(ctx context.Context, arg CreateSubscriptionParams) error {
	_, err := q.db.Exec(ctx, createSubscription,
		arg.SubscriptionUuid,
		arg.AccountUuid,
		arg.Appli,
		arg.Callbackurl,
		arg.WebhookSecret,
		arg.Status,
		arg.Comment,
		arg.StatusLastCheckedAt,
	)
	return err
}

const deleteRawNotification = `-- name: DeleteRawNotification :exec
DELETE
FROM raw_notification
WHERE raw_notification_uuid = $1
`

func (q *Queries) DeleteRawNotification(ctx context.Context, rawNotificationUuid uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteRawNotification, rawNotificationUuid)
	return err
}

const deleteSubscription = `-- name: DeleteSubscription :exec
DELETE
FROM subscription
WHERE subscription_uuid = $1
`

func (q *Queries) DeleteSubscription(ctx context.Context, subscriptionUuid uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteSubscription, subscriptionUuid)
	return err
}

const getPendingRawNotifications = `-- name: GetPendingRawNotifications :many
SELECT raw_notification_id, source, status, data, raw_notification_uuid
FROM raw_notification
WHERE status = 'pending'
ORDER BY raw_notification_id
`

func (q *Queries) GetPendingRawNotifications(ctx context.Context) ([]RawNotification, error) {
	rows, err := q.db.Query(ctx, getPendingRawNotifications)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RawNotification
	for rows.Next() {
		var i RawNotification
		if err := rows.Scan(
			&i.RawNotificationID,
			&i.Source,
			&i.Status,
			&i.Data,
			&i.RawNotificationUuid,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPendingSubscriptions = `-- name: GetPendingSubscriptions :many
SELECT subscription_id, appli, callbackurl, webhook_secret, status, comment, subscription_uuid, account_uuid, status_last_checked_at
FROM subscription
WHERE status = 'pending'
ORDER BY subscription.account_uuid
`

func (q *Queries) GetPendingSubscriptions(ctx context.Context) ([]Subscription, error) {
	rows, err := q.db.Query(ctx, getPendingSubscriptions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Subscription
	for rows.Next() {
		var i Subscription
		if err := rows.Scan(
			&i.SubscriptionID,
			&i.Appli,
			&i.Callbackurl,
			&i.WebhookSecret,
			&i.Status,
			&i.Comment,
			&i.SubscriptionUuid,
			&i.AccountUuid,
			&i.StatusLastCheckedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRawNotification = `-- name: GetRawNotification :one
SELECT raw_notification_id, source, status, data, raw_notification_uuid
FROM raw_notification
WHERE raw_notification_uuid = $1
`

func (q *Queries) GetRawNotification(ctx context.Context, rawNotificationUuid uuid.UUID) (RawNotification, error) {
	row := q.db.QueryRow(ctx, getRawNotification, rawNotificationUuid)
	var i RawNotification
	err := row.Scan(
		&i.RawNotificationID,
		&i.Source,
		&i.Status,
		&i.Data,
		&i.RawNotificationUuid,
	)
	return i, err
}

const getSubscriptionByAccountUUIDAndAppli = `-- name: GetSubscriptionByAccountUUIDAndAppli :one
SELECT subscription_id, appli, callbackurl, webhook_secret, status, comment, subscription_uuid, account_uuid, status_last_checked_at
FROM subscription
WHERE account_uuid = $1
  AND appli = $2
`

type GetSubscriptionByAccountUUIDAndAppliParams struct {
	AccountUuid uuid.UUID
	Appli       int32
}

func (q *Queries) GetSubscriptionByAccountUUIDAndAppli(ctx context.Context, arg GetSubscriptionByAccountUUIDAndAppliParams) (Subscription, error) {
	row := q.db.QueryRow(ctx, getSubscriptionByAccountUUIDAndAppli, arg.AccountUuid, arg.Appli)
	var i Subscription
	err := row.Scan(
		&i.SubscriptionID,
		&i.Appli,
		&i.Callbackurl,
		&i.WebhookSecret,
		&i.Status,
		&i.Comment,
		&i.SubscriptionUuid,
		&i.AccountUuid,
		&i.StatusLastCheckedAt,
	)
	return i, err
}

const getSubscriptionByUUID = `-- name: GetSubscriptionByUUID :one
SELECT subscription_id, appli, callbackurl, webhook_secret, status, comment, subscription_uuid, account_uuid, status_last_checked_at
FROM subscription
WHERE subscription_uuid = $1
`

func (q *Queries) GetSubscriptionByUUID(ctx context.Context, subscriptionUuid uuid.UUID) (Subscription, error) {
	row := q.db.QueryRow(ctx, getSubscriptionByUUID, subscriptionUuid)
	var i Subscription
	err := row.Scan(
		&i.SubscriptionID,
		&i.Appli,
		&i.Callbackurl,
		&i.WebhookSecret,
		&i.Status,
		&i.Comment,
		&i.SubscriptionUuid,
		&i.AccountUuid,
		&i.StatusLastCheckedAt,
	)
	return i, err
}

const getSubscriptionByWebhookSecret = `-- name: GetSubscriptionByWebhookSecret :one
SELECT subscription_id, appli, callbackurl, webhook_secret, status, comment, subscription_uuid, account_uuid, status_last_checked_at
FROM subscription
WHERE webhook_secret = $1
`

func (q *Queries) GetSubscriptionByWebhookSecret(ctx context.Context, webhookSecret string) (Subscription, error) {
	row := q.db.QueryRow(ctx, getSubscriptionByWebhookSecret, webhookSecret)
	var i Subscription
	err := row.Scan(
		&i.SubscriptionID,
		&i.Appli,
		&i.Callbackurl,
		&i.WebhookSecret,
		&i.Status,
		&i.Comment,
		&i.SubscriptionUuid,
		&i.AccountUuid,
		&i.StatusLastCheckedAt,
	)
	return i, err
}

const getSubscriptionsByAccountUUID = `-- name: GetSubscriptionsByAccountUUID :many
SELECT subscription_id, appli, callbackurl, webhook_secret, status, comment, subscription_uuid, account_uuid, status_last_checked_at
FROM subscription
WHERE account_uuid = $1
ORDER BY appli
`

func (q *Queries) GetSubscriptionsByAccountUUID(ctx context.Context, accountUuid uuid.UUID) ([]Subscription, error) {
	rows, err := q.db.Query(ctx, getSubscriptionsByAccountUUID, accountUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Subscription
	for rows.Next() {
		var i Subscription
		if err := rows.Scan(
			&i.SubscriptionID,
			&i.Appli,
			&i.Callbackurl,
			&i.WebhookSecret,
			&i.Status,
			&i.Comment,
			&i.SubscriptionUuid,
			&i.AccountUuid,
			&i.StatusLastCheckedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSubscriptions = `-- name: ListSubscriptions :many
SELECT subscription_id, appli, callbackurl, webhook_secret, status, comment, subscription_uuid, account_uuid, status_last_checked_at
FROM subscription
ORDER BY account_uuid
`

func (q *Queries) ListSubscriptions(ctx context.Context) ([]Subscription, error) {
	rows, err := q.db.Query(ctx, listSubscriptions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Subscription
	for rows.Next() {
		var i Subscription
		if err := rows.Scan(
			&i.SubscriptionID,
			&i.Appli,
			&i.Callbackurl,
			&i.WebhookSecret,
			&i.Status,
			&i.Comment,
			&i.SubscriptionUuid,
			&i.AccountUuid,
			&i.StatusLastCheckedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSubscription = `-- name: UpdateSubscription :exec
UPDATE subscription
SET account_uuid           = $1,
    appli                  = $2,
    callbackurl            = $3,
    webhook_secret         = $4,
    status                 = $5,
    comment                = $6,
    status_last_checked_at = $7
WHERE subscription_uuid = $8
`

type UpdateSubscriptionParams struct {
	AccountUuid         uuid.UUID
	Appli               int32
	Callbackurl         string
	WebhookSecret       string
	Status              string
	Comment             string
	StatusLastCheckedAt time.Time
	SubscriptionUuid    uuid.UUID
}

func (q *Queries) UpdateSubscription(ctx context.Context, arg UpdateSubscriptionParams) error {
	_, err := q.db.Exec(ctx, updateSubscription,
		arg.AccountUuid,
		arg.Appli,
		arg.Callbackurl,
		arg.WebhookSecret,
		arg.Status,
		arg.Comment,
		arg.StatusLastCheckedAt,
		arg.SubscriptionUuid,
	)
	return err
}
