package helloworld

import "testing"
import "github.com/bmizerany/assert"

func TestHelloWorld(t *testing.T) {
	got := HelloWorld()
	want := "HelloWorld"
	assert.Equal(t, want, got, "")
}
