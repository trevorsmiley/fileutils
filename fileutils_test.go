package fileutils

import "testing"

func TestGetFileNameWithoutExtension(t *testing.T) {
	testFile := "foo.bar"
	expected := "foo"
	actual := GetFileNameWithoutExtension(testFile)
	if actual != expected {
		t.Errorf("Actual: %s | Expected: %s\n", actual, expected)
	}
}

func BenchmarkGetFileNameWithoutExtension(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetFileNameWithoutExtension("foo.bar")
	}
}
