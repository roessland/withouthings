// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: subscription.sql

package db

import (
	"context"
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
INSERT INTO raw_notification (source, status, data)
VALUES ($1, $2, $3)
`

type CreateRawNotificationParams struct {
	Source string
	Status string
	Data   string
}

func (q *Queries) CreateRawNotification(ctx context.Context, arg CreateRawNotificationParams) error {
	_, err := q.db.Exec(ctx, createRawNotification, arg.Source, arg.Status, arg.Data)
	return err
}

const createSubscription = `-- name: CreateSubscription :exec
INSERT INTO subscription (account_id, appli, callbackurl, webhook_secret, status, comment)
VALUES ($1, $2, $3, $4, $5, $6)
`

type CreateSubscriptionParams struct {
	AccountID     int64
	Appli         int32
	Callbackurl   string
	WebhookSecret string
	Status        string
	Comment       string
}

func (q *Queries) CreateSubscription(ctx context.Context, arg CreateSubscriptionParams) error {
	_, err := q.db.Exec(ctx, createSubscription,
		arg.AccountID,
		arg.Appli,
		arg.Callbackurl,
		arg.WebhookSecret,
		arg.Status,
		arg.Comment,
	)
	return err
}

const deleteSubscription = `-- name: DeleteSubscription :exec
DELETE
FROM subscription
WHERE subscription_id = $1
`

func (q *Queries) DeleteSubscription(ctx context.Context, subscriptionID int64) error {
	_, err := q.db.Exec(ctx, deleteSubscription, subscriptionID)
	return err
}

const getPendingRawNotifications = `-- name: GetPendingRawNotifications :many
SELECT raw_notification_id, source, status, data
FROM raw_notification
WHERE status == 'pending'
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
SELECT subscription_id, account_id, appli, callbackurl, webhook_secret, status, comment
FROM subscription
WHERE status = 'pending'
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
			&i.AccountID,
			&i.Appli,
			&i.Callbackurl,
			&i.WebhookSecret,
			&i.Status,
			&i.Comment,
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

const getSubscriptionByID = `-- name: GetSubscriptionByID :one
SELECT subscription_id, account_id, appli, callbackurl, webhook_secret, status, comment
FROM subscription
WHERE subscription_id = $1
`

func (q *Queries) GetSubscriptionByID(ctx context.Context, subscriptionID int64) (Subscription, error) {
	row := q.db.QueryRow(ctx, getSubscriptionByID, subscriptionID)
	var i Subscription
	err := row.Scan(
		&i.SubscriptionID,
		&i.AccountID,
		&i.Appli,
		&i.Callbackurl,
		&i.WebhookSecret,
		&i.Status,
		&i.Comment,
	)
	return i, err
}

const getSubscriptionByWebhookSecret = `-- name: GetSubscriptionByWebhookSecret :one
SELECT subscription_id, account_id, appli, callbackurl, webhook_secret, status, comment
FROM subscription
WHERE webhook_secret = $1
`

func (q *Queries) GetSubscriptionByWebhookSecret(ctx context.Context, webhookSecret string) (Subscription, error) {
	row := q.db.QueryRow(ctx, getSubscriptionByWebhookSecret, webhookSecret)
	var i Subscription
	err := row.Scan(
		&i.SubscriptionID,
		&i.AccountID,
		&i.Appli,
		&i.Callbackurl,
		&i.WebhookSecret,
		&i.Status,
		&i.Comment,
	)
	return i, err
}

const getSubscriptionsByAccountID = `-- name: GetSubscriptionsByAccountID :many
SELECT subscription_id, account_id, appli, callbackurl, webhook_secret, status, comment
FROM subscription
WHERE account_id = $1
`

func (q *Queries) GetSubscriptionsByAccountID(ctx context.Context, accountID int64) ([]Subscription, error) {
	rows, err := q.db.Query(ctx, getSubscriptionsByAccountID, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Subscription
	for rows.Next() {
		var i Subscription
		if err := rows.Scan(
			&i.SubscriptionID,
			&i.AccountID,
			&i.Appli,
			&i.Callbackurl,
			&i.WebhookSecret,
			&i.Status,
			&i.Comment,
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
SELECT subscription_id, account_id, appli, callbackurl, webhook_secret, status, comment
FROM subscription
ORDER BY account_id
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
			&i.AccountID,
			&i.Appli,
			&i.Callbackurl,
			&i.WebhookSecret,
			&i.Status,
			&i.Comment,
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
