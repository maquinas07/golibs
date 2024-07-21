package getopt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortFlags(t *testing.T) {
	var s bool
	AddFlag('s', "set", &s)
	parse([]string{"test", "-s"})
	assert.True(t, s)
}

func TestMultipleShortFlags(t *testing.T) {
	var s, u, v bool
	AddFlag('s', "set", &s)
	AddFlag('u', "uet", &u)
	AddFlag('v', "vet", &v)
	parse([]string{"test", "-suv"})
	assert.True(t, s)
	assert.True(t, u)
	assert.True(t, v)
}

func TestLongFlags(t *testing.T) {
	var s bool
	AddFlag('s', "set", &s)
	parse([]string{"test", "--set"})
	assert.True(t, s)
}

func TestShortOptions(t *testing.T) {
	var s string
	AddOption('s', "set", &s, false, nil)
	parse([]string{"test", "-s", "success"})
	assert.Equal(t, "success", s)
}

func TestLongOptions(t *testing.T) {
	var s string
	AddOption('s', "set", &s, false, nil)
	parse([]string{"test", "--set", "success"})
	assert.Equal(t, "success", s)
	s = ""
	parse([]string{"test", "--set=success"})
	assert.Equal(t, "success", s)
}

func TestShortOptionsFailsWhenRequiredAndNotPresent(t *testing.T) {
	var s string
	AddOption('s', "set", &s, false, nil)
	_, err := parse([]string{"test", "-s"})
	assert.NotEqual(t, "success", s)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("missing argument value"), err)
}

func TestLongOptionsFailsWhenRequiredAndNotPresent(t *testing.T) {
	var s string
	AddOption('s', "set", &s, false, nil)
	_, err := parse([]string{"test", "--set"})
	assert.NotEqual(t, "success", s)
	assert.NotNil(t, err)
	assert.Equal(t, fmt.Errorf("missing argument value"), err)
}
