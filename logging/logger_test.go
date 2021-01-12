package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Used to test the output of the logger
type testOutput struct {
	Text string
}

func (o *testOutput) Write(p []byte) (n int, err error) {
	o.Text += string(p)

	return len(p), nil
}

func TestNew(t *testing.T) {
	o := &testOutput{}

	l := New().SetOutput(o)
	assert.Equal(t, o, (l.(*logger)).output)
	assert.NotNil(t, l)
}

func TestDebug(t *testing.T) {
	o := &testOutput{}
	defaultLogger = New().SetOutput(o).SetLevel(LevelDebug)

	Debug("Hello %s", "World")

	assert.Contains(t, o.Text, debugColor)
	assert.Contains(t, o.Text, "Hello World")
}

func TestDebug_WithName(t *testing.T) {
	o := &testOutput{}
	defaultLogger = New().SetOutput(o).SetLevel(LevelDebug)

	SetName("TestLogger")
	Debug("Hello %s", "World")

	assert.Contains(t, o.Text, debugColor)
	assert.Contains(t, o.Text, "TestLogger")
	assert.Contains(t, o.Text, "Hello World")
}

func TestDebugf_WithGreaterLogLevel(t *testing.T) {
	o := &testOutput{}
	defaultLogger = New().SetOutput(o)

	SetLevel(LevelInfo)
	Debug("Hello %s", "World")

	assert.Equal(t, "", o.Text)
}

func TestInfo(t *testing.T) {
	o := &testOutput{}
	defaultLogger = New().SetOutput(o)

	Info("Hello %s", "World")

	assert.Contains(t, o.Text, infoColor)
	assert.Contains(t, o.Text, "Hello World")
}

func TestInfo_WithName(t *testing.T) {
	o := &testOutput{}
	defaultLogger = New().SetOutput(o)

	SetName("TestLogger")
	Info("Hello %s", "World")

	assert.Contains(t, o.Text, infoColor)
	assert.Contains(t, o.Text, "TestLogger")
	assert.Contains(t, o.Text, "Hello World")
}

func TestInfo_WithGreaterLogLevel(t *testing.T) {
	o := &testOutput{}
	defaultLogger = New().SetOutput(o)

	SetLevel(LevelWarning)
	Info("Hello %s", "World")

	assert.Equal(t, "", o.Text)
}

func TestWarning(t *testing.T) {
	o := &testOutput{}
	defaultLogger = New().SetOutput(o)

	Warning("Hello %s", "World")

	assert.Contains(t, o.Text, warningColor)
	assert.Contains(t, o.Text, "Hello World")
}

func TestWarning_WithName(t *testing.T) {
	o := &testOutput{}
	defaultLogger = New().SetOutput(o)

	SetName("TestLogger")
	Warning("Hello %s", "World")

	assert.Contains(t, o.Text, warningColor)
	assert.Contains(t, o.Text, "TestLogger")
	assert.Contains(t, o.Text, "Hello World")
}

func TestWarning_WithGreaterLogLevel(t *testing.T) {
	o := &testOutput{}
	defaultLogger = New().SetOutput(o)

	SetLevel(LevelError)
	Warning("Hello %s", "World")

	assert.Equal(t, "", o.Text)
}

func TestError(t *testing.T) {
	o := &testOutput{}
	defaultLogger = New().SetOutput(o)

	Error("Hello %s", "World")

	assert.Contains(t, o.Text, errorColor)
	assert.Contains(t, o.Text, "Hello World")
}

func TestError_WithName(t *testing.T) {
	o := &testOutput{}
	defaultLogger = New().SetOutput(o)

	SetName("TestLogger")
	Error("Hello %s", "World")

	assert.Contains(t, o.Text, errorColor)
	assert.Contains(t, o.Text, "TestLogger")
	assert.Contains(t, o.Text, "Hello World")
}

func TestSetOutput(t *testing.T) {
	o := &testOutput{}
	SetOutput(o)

	assert.Equal(t, o, (defaultLogger.(*logger)).output)

	Info("Hello World")
	assert.Contains(t, o.Text, "Hello World")
}
