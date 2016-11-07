package rabbitsmpp

import (
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/mock"
)

// MockConsumer is an autogenerated mock type for the Consumer type
type MockConsumer struct {
	mock.Mock
}

// Consume provides a mock function with given fields:
func (_m *MockConsumer) Consume() (<-chan Job, <-chan error, error) {
	ret := _m.Called()

	var r0 <-chan Job
	if rf, ok := ret.Get(0).(func() <-chan Job); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan Job)
		}
	}

	var r1 <-chan error
	if rf, ok := ret.Get(1).(func() <-chan error); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(chan error)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ID provides a mock function with given fields:
func (_m *MockConsumer) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *MockConsumer) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockChannel is an autogenerated mock type for the Channel type
type MockChannel struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *MockChannel) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Consume provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5, _a6
func (_m *MockChannel) Consume(_a0 string, _a1 string, _a2 bool, _a3 bool, _a4 bool, _a5 bool, _a6 amqp.Table) (<-chan amqp.Delivery, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5, _a6)

	var r0 <-chan amqp.Delivery
	if rf, ok := ret.Get(0).(func(string, string, bool, bool, bool, bool, amqp.Table) <-chan amqp.Delivery); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan amqp.Delivery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, bool, bool, bool, bool, amqp.Table) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExchangeDeclare provides a mock function with given fields: name, kind, durable, autoDelete, internal, noWait, args
func (_m *MockChannel) ExchangeDeclare(name string, kind string, durable bool, autoDelete bool, internal bool, noWait bool, args amqp.Table) error {
	ret := _m.Called(name, kind, durable, autoDelete, internal, noWait, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, bool, bool, bool, bool, amqp.Table) error); ok {
		r0 = rf(name, kind, durable, autoDelete, internal, noWait, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Publish provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *MockChannel) Publish(_a0 string, _a1 string, _a2 bool, _a3 bool, _a4 amqp.Publishing) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, bool, bool, amqp.Publishing) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Qos provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockChannel) Qos(_a0 int, _a1 int, _a2 bool) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, bool) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueueBind provides a mock function with given fields: name, key, exchange, noWait, args
func (_m *MockChannel) QueueBind(name string, key string, exchange string, noWait bool, args amqp.Table) error {
	ret := _m.Called(name, key, exchange, noWait, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, bool, amqp.Table) error); ok {
		r0 = rf(name, key, exchange, noWait, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueueDeclare provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *MockChannel) QueueDeclare(_a0 string, _a1 bool, _a2 bool, _a3 bool, _a4 bool, _a5 amqp.Table) (amqp.Queue, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5)

	var r0 amqp.Queue
	if rf, ok := ret.Get(0).(func(string, bool, bool, bool, bool, amqp.Table) amqp.Queue); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r0 = ret.Get(0).(amqp.Queue)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool, bool, bool, bool, amqp.Table) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueueInspect provides a mock function with given fields: _a0
func (_m *MockChannel) QueueInspect(_a0 string) (amqp.Queue, error) {
	ret := _m.Called(_a0)

	var r0 amqp.Queue
	if rf, ok := ret.Get(0).(func(string) amqp.Queue); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(amqp.Queue)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPublisherClient is an autogenerated mock type for the PublisherClient type
type MockPublisherClient struct {
	mock.Mock
}

// Bind provides a mock function with given fields:
func (_m *MockPublisherClient) Bind() (chan *amqp.Error, error) {
	ret := _m.Called()

	var r0 chan *amqp.Error
	if rf, ok := ret.Get(0).(func() chan *amqp.Error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *amqp.Error)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Channel provides a mock function with given fields:
func (_m *MockPublisherClient) Channel() (Channel, error) {
	ret := _m.Called()

	var r0 Channel
	if rf, ok := ret.Get(0).(func() Channel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *MockPublisherClient) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueueName provides a mock function with given fields:
func (_m *MockPublisherClient) QueueName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ReBind provides a mock function with given fields:
func (_m *MockPublisherClient) ReBind() (PublisherClient, chan *amqp.Error, error) {
	ret := _m.Called()

	var r0 PublisherClient
	if rf, ok := ret.Get(0).(func() PublisherClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PublisherClient)
		}
	}

	var r1 chan *amqp.Error
	if rf, ok := ret.Get(1).(func() chan *amqp.Error); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(chan *amqp.Error)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

// Bind provides a mock function with given fields:
func (_m *MockClient) Bind() (chan *amqp.Error, error) {
	ret := _m.Called()

	var r0 chan *amqp.Error
	if rf, ok := ret.Get(0).(func() chan *amqp.Error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *amqp.Error)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Channel provides a mock function with given fields:
func (_m *MockClient) Channel() (Channel, error) {
	ret := _m.Called()

	var r0 Channel
	if rf, ok := ret.Get(0).(func() Channel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *MockClient) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueueName provides a mock function with given fields:
func (_m *MockClient) QueueName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockConsumerContainer is an autogenerated mock type for the ConsumerContainer type
type MockConsumerContainer struct {
	mock.Mock
}

// AddFromConfig provides a mock function with given fields: _a0
func (_m *MockConsumerContainer) AddFromConfig(_a0 Config) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(Config) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Config) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddRun provides a mock function with given fields: _a0
func (_m *MockConsumerContainer) AddRun(_a0 Consumer) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(Consumer) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Consume provides a mock function with given fields:
func (_m *MockConsumerContainer) Consume() (<-chan Job, <-chan error, error) {
	ret := _m.Called()

	var r0 <-chan Job
	if rf, ok := ret.Get(0).(func() <-chan Job); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan Job)
		}
	}

	var r1 <-chan error
	if rf, ok := ret.Get(1).(func() <-chan error); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(<-chan error)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ID provides a mock function with given fields:
func (_m *MockConsumerContainer) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RemoveStop provides a mock function with given fields: _a0
func (_m *MockConsumerContainer) RemoveStop(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *MockConsumerContainer) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
