// Code generated by mockery v2.20.0. DO NOT EDIT.

package subscription

import (
	context "context"

	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

// MockRepo is an autogenerated mock type for the Repo type
type MockRepo struct {
	mock.Mock
}

type MockRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepo) EXPECT() *MockRepo_Expecter {
	return &MockRepo_Expecter{mock: &_m.Mock}
}

// AllNotificationCategories provides a mock function with given fields: ctx
func (_m *MockRepo) AllNotificationCategories(ctx context.Context) ([]NotificationCategory, error) {
	ret := _m.Called(ctx)

	var r0 []NotificationCategory
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]NotificationCategory, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []NotificationCategory); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]NotificationCategory)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepo_AllNotificationCategories_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllNotificationCategories'
type MockRepo_AllNotificationCategories_Call struct {
	*mock.Call
}

// AllNotificationCategories is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockRepo_Expecter) AllNotificationCategories(ctx interface{}) *MockRepo_AllNotificationCategories_Call {
	return &MockRepo_AllNotificationCategories_Call{Call: _e.mock.On("AllNotificationCategories", ctx)}
}

func (_c *MockRepo_AllNotificationCategories_Call) Run(run func(ctx context.Context)) *MockRepo_AllNotificationCategories_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockRepo_AllNotificationCategories_Call) Return(_a0 []NotificationCategory, _a1 error) *MockRepo_AllNotificationCategories_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepo_AllNotificationCategories_Call) RunAndReturn(run func(context.Context) ([]NotificationCategory, error)) *MockRepo_AllNotificationCategories_Call {
	_c.Call.Return(run)
	return _c
}

// CreateNotification provides a mock function with given fields: ctx, notification
func (_m *MockRepo) CreateNotification(ctx context.Context, notification Notification) error {
	ret := _m.Called(ctx, notification)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, Notification) error); ok {
		r0 = rf(ctx, notification)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepo_CreateNotification_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateNotification'
type MockRepo_CreateNotification_Call struct {
	*mock.Call
}

// CreateNotification is a helper method to define mock.On call
//   - ctx context.Context
//   - notification Notification
func (_e *MockRepo_Expecter) CreateNotification(ctx interface{}, notification interface{}) *MockRepo_CreateNotification_Call {
	return &MockRepo_CreateNotification_Call{Call: _e.mock.On("CreateNotification", ctx, notification)}
}

func (_c *MockRepo_CreateNotification_Call) Run(run func(ctx context.Context, notification Notification)) *MockRepo_CreateNotification_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(Notification))
	})
	return _c
}

func (_c *MockRepo_CreateNotification_Call) Return(_a0 error) *MockRepo_CreateNotification_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepo_CreateNotification_Call) RunAndReturn(run func(context.Context, Notification) error) *MockRepo_CreateNotification_Call {
	_c.Call.Return(run)
	return _c
}

// CreateRawNotification provides a mock function with given fields: ctx, rawNotification
func (_m *MockRepo) CreateRawNotification(ctx context.Context, rawNotification RawNotification) error {
	ret := _m.Called(ctx, rawNotification)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, RawNotification) error); ok {
		r0 = rf(ctx, rawNotification)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepo_CreateRawNotification_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRawNotification'
type MockRepo_CreateRawNotification_Call struct {
	*mock.Call
}

// CreateRawNotification is a helper method to define mock.On call
//   - ctx context.Context
//   - rawNotification RawNotification
func (_e *MockRepo_Expecter) CreateRawNotification(ctx interface{}, rawNotification interface{}) *MockRepo_CreateRawNotification_Call {
	return &MockRepo_CreateRawNotification_Call{Call: _e.mock.On("CreateRawNotification", ctx, rawNotification)}
}

func (_c *MockRepo_CreateRawNotification_Call) Run(run func(ctx context.Context, rawNotification RawNotification)) *MockRepo_CreateRawNotification_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(RawNotification))
	})
	return _c
}

func (_c *MockRepo_CreateRawNotification_Call) Return(_a0 error) *MockRepo_CreateRawNotification_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepo_CreateRawNotification_Call) RunAndReturn(run func(context.Context, RawNotification) error) *MockRepo_CreateRawNotification_Call {
	_c.Call.Return(run)
	return _c
}

// CreateSubscriptionIfNotExists provides a mock function with given fields: ctx, subscription
func (_m *MockRepo) CreateSubscriptionIfNotExists(ctx context.Context, subscription *Subscription) error {
	ret := _m.Called(ctx, subscription)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Subscription) error); ok {
		r0 = rf(ctx, subscription)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepo_CreateSubscriptionIfNotExists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSubscriptionIfNotExists'
type MockRepo_CreateSubscriptionIfNotExists_Call struct {
	*mock.Call
}

// CreateSubscriptionIfNotExists is a helper method to define mock.On call
//   - ctx context.Context
//   - subscription *Subscription
func (_e *MockRepo_Expecter) CreateSubscriptionIfNotExists(ctx interface{}, subscription interface{}) *MockRepo_CreateSubscriptionIfNotExists_Call {
	return &MockRepo_CreateSubscriptionIfNotExists_Call{Call: _e.mock.On("CreateSubscriptionIfNotExists", ctx, subscription)}
}

func (_c *MockRepo_CreateSubscriptionIfNotExists_Call) Run(run func(ctx context.Context, subscription *Subscription)) *MockRepo_CreateSubscriptionIfNotExists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Subscription))
	})
	return _c
}

func (_c *MockRepo_CreateSubscriptionIfNotExists_Call) Return(_a0 error) *MockRepo_CreateSubscriptionIfNotExists_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepo_CreateSubscriptionIfNotExists_Call) RunAndReturn(run func(context.Context, *Subscription) error) *MockRepo_CreateSubscriptionIfNotExists_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteRawNotification provides a mock function with given fields: ctx, rawNotification
func (_m *MockRepo) DeleteRawNotification(ctx context.Context, rawNotification RawNotification) error {
	ret := _m.Called(ctx, rawNotification)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, RawNotification) error); ok {
		r0 = rf(ctx, rawNotification)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepo_DeleteRawNotification_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRawNotification'
type MockRepo_DeleteRawNotification_Call struct {
	*mock.Call
}

// DeleteRawNotification is a helper method to define mock.On call
//   - ctx context.Context
//   - rawNotification RawNotification
func (_e *MockRepo_Expecter) DeleteRawNotification(ctx interface{}, rawNotification interface{}) *MockRepo_DeleteRawNotification_Call {
	return &MockRepo_DeleteRawNotification_Call{Call: _e.mock.On("DeleteRawNotification", ctx, rawNotification)}
}

func (_c *MockRepo_DeleteRawNotification_Call) Run(run func(ctx context.Context, rawNotification RawNotification)) *MockRepo_DeleteRawNotification_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(RawNotification))
	})
	return _c
}

func (_c *MockRepo_DeleteRawNotification_Call) Return(_a0 error) *MockRepo_DeleteRawNotification_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepo_DeleteRawNotification_Call) RunAndReturn(run func(context.Context, RawNotification) error) *MockRepo_DeleteRawNotification_Call {
	_c.Call.Return(run)
	return _c
}

// GetPendingRawNotifications provides a mock function with given fields: ctx
func (_m *MockRepo) GetPendingRawNotifications(ctx context.Context) ([]RawNotification, error) {
	ret := _m.Called(ctx)

	var r0 []RawNotification
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]RawNotification, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []RawNotification); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]RawNotification)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepo_GetPendingRawNotifications_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPendingRawNotifications'
type MockRepo_GetPendingRawNotifications_Call struct {
	*mock.Call
}

// GetPendingRawNotifications is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockRepo_Expecter) GetPendingRawNotifications(ctx interface{}) *MockRepo_GetPendingRawNotifications_Call {
	return &MockRepo_GetPendingRawNotifications_Call{Call: _e.mock.On("GetPendingRawNotifications", ctx)}
}

func (_c *MockRepo_GetPendingRawNotifications_Call) Run(run func(ctx context.Context)) *MockRepo_GetPendingRawNotifications_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockRepo_GetPendingRawNotifications_Call) Return(_a0 []RawNotification, _a1 error) *MockRepo_GetPendingRawNotifications_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepo_GetPendingRawNotifications_Call) RunAndReturn(run func(context.Context) ([]RawNotification, error)) *MockRepo_GetPendingRawNotifications_Call {
	_c.Call.Return(run)
	return _c
}

// GetPendingSubscriptions provides a mock function with given fields: ctx
func (_m *MockRepo) GetPendingSubscriptions(ctx context.Context) ([]*Subscription, error) {
	ret := _m.Called(ctx)

	var r0 []*Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*Subscription, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*Subscription); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepo_GetPendingSubscriptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPendingSubscriptions'
type MockRepo_GetPendingSubscriptions_Call struct {
	*mock.Call
}

// GetPendingSubscriptions is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockRepo_Expecter) GetPendingSubscriptions(ctx interface{}) *MockRepo_GetPendingSubscriptions_Call {
	return &MockRepo_GetPendingSubscriptions_Call{Call: _e.mock.On("GetPendingSubscriptions", ctx)}
}

func (_c *MockRepo_GetPendingSubscriptions_Call) Run(run func(ctx context.Context)) *MockRepo_GetPendingSubscriptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockRepo_GetPendingSubscriptions_Call) Return(_a0 []*Subscription, _a1 error) *MockRepo_GetPendingSubscriptions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepo_GetPendingSubscriptions_Call) RunAndReturn(run func(context.Context) ([]*Subscription, error)) *MockRepo_GetPendingSubscriptions_Call {
	_c.Call.Return(run)
	return _c
}

// GetRawNotificationByUUID provides a mock function with given fields: ctx, rawNotificationUUID
func (_m *MockRepo) GetRawNotificationByUUID(ctx context.Context, rawNotificationUUID uuid.UUID) (RawNotification, error) {
	ret := _m.Called(ctx, rawNotificationUUID)

	var r0 RawNotification
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (RawNotification, error)); ok {
		return rf(ctx, rawNotificationUUID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) RawNotification); ok {
		r0 = rf(ctx, rawNotificationUUID)
	} else {
		r0 = ret.Get(0).(RawNotification)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, rawNotificationUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepo_GetRawNotificationByUUID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRawNotificationByUUID'
type MockRepo_GetRawNotificationByUUID_Call struct {
	*mock.Call
}

// GetRawNotificationByUUID is a helper method to define mock.On call
//   - ctx context.Context
//   - rawNotificationUUID uuid.UUID
func (_e *MockRepo_Expecter) GetRawNotificationByUUID(ctx interface{}, rawNotificationUUID interface{}) *MockRepo_GetRawNotificationByUUID_Call {
	return &MockRepo_GetRawNotificationByUUID_Call{Call: _e.mock.On("GetRawNotificationByUUID", ctx, rawNotificationUUID)}
}

func (_c *MockRepo_GetRawNotificationByUUID_Call) Run(run func(ctx context.Context, rawNotificationUUID uuid.UUID)) *MockRepo_GetRawNotificationByUUID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockRepo_GetRawNotificationByUUID_Call) Return(_a0 RawNotification, _a1 error) *MockRepo_GetRawNotificationByUUID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepo_GetRawNotificationByUUID_Call) RunAndReturn(run func(context.Context, uuid.UUID) (RawNotification, error)) *MockRepo_GetRawNotificationByUUID_Call {
	_c.Call.Return(run)
	return _c
}

// GetSubscriptionByUUID provides a mock function with given fields: ctx, subscriptionUUID
func (_m *MockRepo) GetSubscriptionByUUID(ctx context.Context, subscriptionUUID uuid.UUID) (*Subscription, error) {
	ret := _m.Called(ctx, subscriptionUUID)

	var r0 *Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*Subscription, error)); ok {
		return rf(ctx, subscriptionUUID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *Subscription); ok {
		r0 = rf(ctx, subscriptionUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, subscriptionUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepo_GetSubscriptionByUUID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubscriptionByUUID'
type MockRepo_GetSubscriptionByUUID_Call struct {
	*mock.Call
}

// GetSubscriptionByUUID is a helper method to define mock.On call
//   - ctx context.Context
//   - subscriptionUUID uuid.UUID
func (_e *MockRepo_Expecter) GetSubscriptionByUUID(ctx interface{}, subscriptionUUID interface{}) *MockRepo_GetSubscriptionByUUID_Call {
	return &MockRepo_GetSubscriptionByUUID_Call{Call: _e.mock.On("GetSubscriptionByUUID", ctx, subscriptionUUID)}
}

func (_c *MockRepo_GetSubscriptionByUUID_Call) Run(run func(ctx context.Context, subscriptionUUID uuid.UUID)) *MockRepo_GetSubscriptionByUUID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockRepo_GetSubscriptionByUUID_Call) Return(_a0 *Subscription, _a1 error) *MockRepo_GetSubscriptionByUUID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepo_GetSubscriptionByUUID_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*Subscription, error)) *MockRepo_GetSubscriptionByUUID_Call {
	_c.Call.Return(run)
	return _c
}

// GetSubscriptionByWebhookSecret provides a mock function with given fields: ctx, webhookSecret
func (_m *MockRepo) GetSubscriptionByWebhookSecret(ctx context.Context, webhookSecret string) (*Subscription, error) {
	ret := _m.Called(ctx, webhookSecret)

	var r0 *Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*Subscription, error)); ok {
		return rf(ctx, webhookSecret)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *Subscription); ok {
		r0 = rf(ctx, webhookSecret)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, webhookSecret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepo_GetSubscriptionByWebhookSecret_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubscriptionByWebhookSecret'
type MockRepo_GetSubscriptionByWebhookSecret_Call struct {
	*mock.Call
}

// GetSubscriptionByWebhookSecret is a helper method to define mock.On call
//   - ctx context.Context
//   - webhookSecret string
func (_e *MockRepo_Expecter) GetSubscriptionByWebhookSecret(ctx interface{}, webhookSecret interface{}) *MockRepo_GetSubscriptionByWebhookSecret_Call {
	return &MockRepo_GetSubscriptionByWebhookSecret_Call{Call: _e.mock.On("GetSubscriptionByWebhookSecret", ctx, webhookSecret)}
}

func (_c *MockRepo_GetSubscriptionByWebhookSecret_Call) Run(run func(ctx context.Context, webhookSecret string)) *MockRepo_GetSubscriptionByWebhookSecret_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRepo_GetSubscriptionByWebhookSecret_Call) Return(_a0 *Subscription, _a1 error) *MockRepo_GetSubscriptionByWebhookSecret_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepo_GetSubscriptionByWebhookSecret_Call) RunAndReturn(run func(context.Context, string) (*Subscription, error)) *MockRepo_GetSubscriptionByWebhookSecret_Call {
	_c.Call.Return(run)
	return _c
}

// GetSubscriptionsByAccountUUID provides a mock function with given fields: ctx, accountUUID
func (_m *MockRepo) GetSubscriptionsByAccountUUID(ctx context.Context, accountUUID uuid.UUID) ([]*Subscription, error) {
	ret := _m.Called(ctx, accountUUID)

	var r0 []*Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]*Subscription, error)); ok {
		return rf(ctx, accountUUID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []*Subscription); ok {
		r0 = rf(ctx, accountUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, accountUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepo_GetSubscriptionsByAccountUUID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubscriptionsByAccountUUID'
type MockRepo_GetSubscriptionsByAccountUUID_Call struct {
	*mock.Call
}

// GetSubscriptionsByAccountUUID is a helper method to define mock.On call
//   - ctx context.Context
//   - accountUUID uuid.UUID
func (_e *MockRepo_Expecter) GetSubscriptionsByAccountUUID(ctx interface{}, accountUUID interface{}) *MockRepo_GetSubscriptionsByAccountUUID_Call {
	return &MockRepo_GetSubscriptionsByAccountUUID_Call{Call: _e.mock.On("GetSubscriptionsByAccountUUID", ctx, accountUUID)}
}

func (_c *MockRepo_GetSubscriptionsByAccountUUID_Call) Run(run func(ctx context.Context, accountUUID uuid.UUID)) *MockRepo_GetSubscriptionsByAccountUUID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockRepo_GetSubscriptionsByAccountUUID_Call) Return(_a0 []*Subscription, _a1 error) *MockRepo_GetSubscriptionsByAccountUUID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepo_GetSubscriptionsByAccountUUID_Call) RunAndReturn(run func(context.Context, uuid.UUID) ([]*Subscription, error)) *MockRepo_GetSubscriptionsByAccountUUID_Call {
	_c.Call.Return(run)
	return _c
}

// ListSubscriptions provides a mock function with given fields: ctx
func (_m *MockRepo) ListSubscriptions(ctx context.Context) ([]*Subscription, error) {
	ret := _m.Called(ctx)

	var r0 []*Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*Subscription, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*Subscription); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepo_ListSubscriptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSubscriptions'
type MockRepo_ListSubscriptions_Call struct {
	*mock.Call
}

// ListSubscriptions is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockRepo_Expecter) ListSubscriptions(ctx interface{}) *MockRepo_ListSubscriptions_Call {
	return &MockRepo_ListSubscriptions_Call{Call: _e.mock.On("ListSubscriptions", ctx)}
}

func (_c *MockRepo_ListSubscriptions_Call) Run(run func(ctx context.Context)) *MockRepo_ListSubscriptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockRepo_ListSubscriptions_Call) Return(_a0 []*Subscription, _a1 error) *MockRepo_ListSubscriptions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepo_ListSubscriptions_Call) RunAndReturn(run func(context.Context) ([]*Subscription, error)) *MockRepo_ListSubscriptions_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, subscriptionUUID, updateFn
func (_m *MockRepo) Update(ctx context.Context, subscriptionUUID uuid.UUID, updateFn func(context.Context, *Subscription) (*Subscription, error)) error {
	ret := _m.Called(ctx, subscriptionUUID, updateFn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, func(context.Context, *Subscription) (*Subscription, error)) error); ok {
		r0 = rf(ctx, subscriptionUUID, updateFn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepo_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockRepo_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - subscriptionUUID uuid.UUID
//   - updateFn func(context.Context , *Subscription)(*Subscription , error)
func (_e *MockRepo_Expecter) Update(ctx interface{}, subscriptionUUID interface{}, updateFn interface{}) *MockRepo_Update_Call {
	return &MockRepo_Update_Call{Call: _e.mock.On("Update", ctx, subscriptionUUID, updateFn)}
}

func (_c *MockRepo_Update_Call) Run(run func(ctx context.Context, subscriptionUUID uuid.UUID, updateFn func(context.Context, *Subscription) (*Subscription, error))) *MockRepo_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(func(context.Context, *Subscription) (*Subscription, error)))
	})
	return _c
}

func (_c *MockRepo_Update_Call) Return(_a0 error) *MockRepo_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepo_Update_Call) RunAndReturn(run func(context.Context, uuid.UUID, func(context.Context, *Subscription) (*Subscription, error)) error) *MockRepo_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateRawNotification provides a mock function with given fields: ctx, rawNotificationUUID, updateFn
func (_m *MockRepo) UpdateRawNotification(ctx context.Context, rawNotificationUUID uuid.UUID, updateFn func(context.Context, *RawNotification) (*RawNotification, error)) error {
	ret := _m.Called(ctx, rawNotificationUUID, updateFn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, func(context.Context, *RawNotification) (*RawNotification, error)) error); ok {
		r0 = rf(ctx, rawNotificationUUID, updateFn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepo_UpdateRawNotification_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateRawNotification'
type MockRepo_UpdateRawNotification_Call struct {
	*mock.Call
}

// UpdateRawNotification is a helper method to define mock.On call
//   - ctx context.Context
//   - rawNotificationUUID uuid.UUID
//   - updateFn func(context.Context , *RawNotification)(*RawNotification , error)
func (_e *MockRepo_Expecter) UpdateRawNotification(ctx interface{}, rawNotificationUUID interface{}, updateFn interface{}) *MockRepo_UpdateRawNotification_Call {
	return &MockRepo_UpdateRawNotification_Call{Call: _e.mock.On("UpdateRawNotification", ctx, rawNotificationUUID, updateFn)}
}

func (_c *MockRepo_UpdateRawNotification_Call) Run(run func(ctx context.Context, rawNotificationUUID uuid.UUID, updateFn func(context.Context, *RawNotification) (*RawNotification, error))) *MockRepo_UpdateRawNotification_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(func(context.Context, *RawNotification) (*RawNotification, error)))
	})
	return _c
}

func (_c *MockRepo_UpdateRawNotification_Call) Return(_a0 error) *MockRepo_UpdateRawNotification_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepo_UpdateRawNotification_Call) RunAndReturn(run func(context.Context, uuid.UUID, func(context.Context, *RawNotification) (*RawNotification, error)) error) *MockRepo_UpdateRawNotification_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRepo creates a new instance of MockRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRepo(t mockConstructorTestingTNewMockRepo) *MockRepo {
	mock := &MockRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
