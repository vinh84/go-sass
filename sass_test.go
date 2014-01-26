package sass

import (
	"strings"
	"testing"
)

func TestCompile(t *testing.T) {
	t.Parallel()
	out, err := Compile(".sass{.inner{color:red}}", NewOptions())
	switch {
	case err != nil:
		t.Fatal("Failed to compile sass:", err)
	case len(strings.Fields(out)) == 0:
		t.Fatal("Compilation resulted in empty output.")
	}
}
