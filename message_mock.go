package libemail

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
)

// MessageMock implements Message
type MessageMock struct {
	t minimock.Tester

	funcCompile          func() (ba1 []byte, err error)
	afterCompileCounter  uint64
	beforeCompileCounter uint64
	CompileMock          mMessageMockCompile

	funcRecipients          func() (sa1 []string)
	afterRecipientsCounter  uint64
	beforeRecipientsCounter uint64
	RecipientsMock          mMessageMockRecipients

	funcSender          func() (s1 string)
	afterSenderCounter  uint64
	beforeSenderCounter uint64
	SenderMock          mMessageMockSender
}

// NewMessageMock returns a mock for Message
func NewMessageMock(t minimock.Tester) *MessageMock {
	m := &MessageMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}
	m.CompileMock = mMessageMockCompile{mock: m}
	m.RecipientsMock = mMessageMockRecipients{mock: m}
	m.SenderMock = mMessageMockSender{mock: m}

	return m
}

type mMessageMockCompile struct {
	mock               *MessageMock
	defaultExpectation *MessageMockCompileExpectation
	expectations       []*MessageMockCompileExpectation
}

// MessageMockCompileExpectation specifies expectation struct of the Message.Compile
type MessageMockCompileExpectation struct {
	mock *MessageMock

	results *MessageMockCompileResults
	Counter uint64
}

// MessageMockCompileResults contains results of the Message.Compile
type MessageMockCompileResults struct {
	ba1 []byte
	err error
}

// Expect sets up expected params for Message.Compile
func (m *mMessageMockCompile) Expect() *mMessageMockCompile {
	if m.mock.funcCompile != nil {
		m.mock.t.Fatalf("MessageMock.Compile mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &MessageMockCompileExpectation{}
	}

	return m
}

// Return sets up results that will be returned by Message.Compile
func (m *mMessageMockCompile) Return(ba1 []byte, err error) *MessageMock {
	if m.mock.funcCompile != nil {
		m.mock.t.Fatalf("MessageMock.Compile mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &MessageMockCompileExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &MessageMockCompileResults{ba1, err}
	return m.mock
}

//Set uses given function f to mock the Message.Compile method
func (m *mMessageMockCompile) Set(f func() (ba1 []byte, err error)) *MessageMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Message.Compile method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Message.Compile method")
	}

	m.mock.funcCompile = f
	return m.mock
}

// Compile implements Message
func (m *MessageMock) Compile() (ba1 []byte, err error) {
	mm_atomic.AddUint64(&m.beforeCompileCounter, 1)
	defer mm_atomic.AddUint64(&m.afterCompileCounter, 1)

	if m.CompileMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.CompileMock.defaultExpectation.Counter, 1)

		results := m.CompileMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the MessageMock.Compile")
		}
		return (*results).ba1, (*results).err
	}
	if m.funcCompile != nil {
		return m.funcCompile()
	}
	m.t.Fatalf("Unexpected call to MessageMock.Compile.")
	return
}

// CompileAfterCounter returns a count of finished MessageMock.Compile invocations
func (m *MessageMock) CompileAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterCompileCounter)
}

// CompileBeforeCounter returns a count of MessageMock.Compile invocations
func (m *MessageMock) CompileBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforeCompileCounter)
}

// MinimockCompileDone returns true if the count of the Compile invocations corresponds
// the number of defined expectations
func (m *MessageMock) MinimockCompileDone() bool {
	for _, e := range m.CompileMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CompileMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCompileCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCompile != nil && mm_atomic.LoadUint64(&m.afterCompileCounter) < 1 {
		return false
	}
	return true
}

// MinimockCompileInspect logs each unmet expectation
func (m *MessageMock) MinimockCompileInspect() {
	for _, e := range m.CompileMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to MessageMock.Compile")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CompileMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCompileCounter) < 1 {
		m.t.Error("Expected call to MessageMock.Compile")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCompile != nil && mm_atomic.LoadUint64(&m.afterCompileCounter) < 1 {
		m.t.Error("Expected call to MessageMock.Compile")
	}
}

type mMessageMockRecipients struct {
	mock               *MessageMock
	defaultExpectation *MessageMockRecipientsExpectation
	expectations       []*MessageMockRecipientsExpectation
}

// MessageMockRecipientsExpectation specifies expectation struct of the Message.Recipients
type MessageMockRecipientsExpectation struct {
	mock *MessageMock

	results *MessageMockRecipientsResults
	Counter uint64
}

// MessageMockRecipientsResults contains results of the Message.Recipients
type MessageMockRecipientsResults struct {
	sa1 []string
}

// Expect sets up expected params for Message.Recipients
func (m *mMessageMockRecipients) Expect() *mMessageMockRecipients {
	if m.mock.funcRecipients != nil {
		m.mock.t.Fatalf("MessageMock.Recipients mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &MessageMockRecipientsExpectation{}
	}

	return m
}

// Return sets up results that will be returned by Message.Recipients
func (m *mMessageMockRecipients) Return(sa1 []string) *MessageMock {
	if m.mock.funcRecipients != nil {
		m.mock.t.Fatalf("MessageMock.Recipients mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &MessageMockRecipientsExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &MessageMockRecipientsResults{sa1}
	return m.mock
}

//Set uses given function f to mock the Message.Recipients method
func (m *mMessageMockRecipients) Set(f func() (sa1 []string)) *MessageMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Message.Recipients method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Message.Recipients method")
	}

	m.mock.funcRecipients = f
	return m.mock
}

// Recipients implements Message
func (m *MessageMock) Recipients() (sa1 []string) {
	mm_atomic.AddUint64(&m.beforeRecipientsCounter, 1)
	defer mm_atomic.AddUint64(&m.afterRecipientsCounter, 1)

	if m.RecipientsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.RecipientsMock.defaultExpectation.Counter, 1)

		results := m.RecipientsMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the MessageMock.Recipients")
		}
		return (*results).sa1
	}
	if m.funcRecipients != nil {
		return m.funcRecipients()
	}
	m.t.Fatalf("Unexpected call to MessageMock.Recipients.")
	return
}

// RecipientsAfterCounter returns a count of finished MessageMock.Recipients invocations
func (m *MessageMock) RecipientsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterRecipientsCounter)
}

// RecipientsBeforeCounter returns a count of MessageMock.Recipients invocations
func (m *MessageMock) RecipientsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforeRecipientsCounter)
}

// MinimockRecipientsDone returns true if the count of the Recipients invocations corresponds
// the number of defined expectations
func (m *MessageMock) MinimockRecipientsDone() bool {
	for _, e := range m.RecipientsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RecipientsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRecipientsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRecipients != nil && mm_atomic.LoadUint64(&m.afterRecipientsCounter) < 1 {
		return false
	}
	return true
}

// MinimockRecipientsInspect logs each unmet expectation
func (m *MessageMock) MinimockRecipientsInspect() {
	for _, e := range m.RecipientsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to MessageMock.Recipients")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RecipientsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRecipientsCounter) < 1 {
		m.t.Error("Expected call to MessageMock.Recipients")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRecipients != nil && mm_atomic.LoadUint64(&m.afterRecipientsCounter) < 1 {
		m.t.Error("Expected call to MessageMock.Recipients")
	}
}

type mMessageMockSender struct {
	mock               *MessageMock
	defaultExpectation *MessageMockSenderExpectation
	expectations       []*MessageMockSenderExpectation
}

// MessageMockSenderExpectation specifies expectation struct of the Message.Sender
type MessageMockSenderExpectation struct {
	mock *MessageMock

	results *MessageMockSenderResults
	Counter uint64
}

// MessageMockSenderResults contains results of the Message.Sender
type MessageMockSenderResults struct {
	s1 string
}

// Expect sets up expected params for Message.Sender
func (m *mMessageMockSender) Expect() *mMessageMockSender {
	if m.mock.funcSender != nil {
		m.mock.t.Fatalf("MessageMock.Sender mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &MessageMockSenderExpectation{}
	}

	return m
}

// Return sets up results that will be returned by Message.Sender
func (m *mMessageMockSender) Return(s1 string) *MessageMock {
	if m.mock.funcSender != nil {
		m.mock.t.Fatalf("MessageMock.Sender mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &MessageMockSenderExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &MessageMockSenderResults{s1}
	return m.mock
}

//Set uses given function f to mock the Message.Sender method
func (m *mMessageMockSender) Set(f func() (s1 string)) *MessageMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Message.Sender method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Message.Sender method")
	}

	m.mock.funcSender = f
	return m.mock
}

// Sender implements Message
func (m *MessageMock) Sender() (s1 string) {
	mm_atomic.AddUint64(&m.beforeSenderCounter, 1)
	defer mm_atomic.AddUint64(&m.afterSenderCounter, 1)

	if m.SenderMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.SenderMock.defaultExpectation.Counter, 1)

		results := m.SenderMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the MessageMock.Sender")
		}
		return (*results).s1
	}
	if m.funcSender != nil {
		return m.funcSender()
	}
	m.t.Fatalf("Unexpected call to MessageMock.Sender.")
	return
}

// SenderAfterCounter returns a count of finished MessageMock.Sender invocations
func (m *MessageMock) SenderAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterSenderCounter)
}

// SenderBeforeCounter returns a count of MessageMock.Sender invocations
func (m *MessageMock) SenderBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforeSenderCounter)
}

// MinimockSenderDone returns true if the count of the Sender invocations corresponds
// the number of defined expectations
func (m *MessageMock) MinimockSenderDone() bool {
	for _, e := range m.SenderMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SenderMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSenderCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSender != nil && mm_atomic.LoadUint64(&m.afterSenderCounter) < 1 {
		return false
	}
	return true
}

// MinimockSenderInspect logs each unmet expectation
func (m *MessageMock) MinimockSenderInspect() {
	for _, e := range m.SenderMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to MessageMock.Sender")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SenderMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSenderCounter) < 1 {
		m.t.Error("Expected call to MessageMock.Sender")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSender != nil && mm_atomic.LoadUint64(&m.afterSenderCounter) < 1 {
		m.t.Error("Expected call to MessageMock.Sender")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MessageMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCompileInspect()

		m.MinimockRecipientsInspect()

		m.MinimockSenderInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MessageMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *MessageMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCompileDone() &&
		m.MinimockRecipientsDone() &&
		m.MinimockSenderDone()
}