package sass

import "testing"

func TestCompile(t *testing.T) {
	_, err := Compile(".sass{.inner{}}")
	if err != nil {
		t.Fatal("Failed to compile sass:", err)
	}
}
