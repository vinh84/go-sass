package sass

import (
	"os"
	"strings"
	"testing"
)

func TestDefaultOptions(t *testing.T) {
	opts := NewOptions()
	switch {
	case opts.output_style != STYLE_NESTED:
		t.Fatal("Default output style should be STYLE_NESTED")
	case opts.source_comments != SOURCE_COMMENTS_NONE:
		t.Fatal("Default source comments should be SOURCE_COMMENTS_NONE")
	case opts.include_paths != "":
		t.Fatal("Default include paths should be empty")
	case opts.image_path == "":
		t.Fatal("Default image path should not be empty")
	}
}

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

func TestCompileFile(t *testing.T) {
	const FILE = "test.sass"
	fhandle, err := os.Create(FILE)
	defer func() {
		fhandle.Close()
		os.Remove(FILE)
	}()
	if err != nil {
		t.Fatal("Failed to open test file for writing, test could not be run", err)
	}
	_, err = fhandle.WriteString(".sass{.inner{color:red}}")
	if err != nil {
		t.Fatal("Failed to write test file, test could not be run", err)
	}
	err = fhandle.Sync()
	if err != nil {
		t.Fatal("Failed to sync test file, test could not be run", err)
	}
	out, err := CompileFile(FILE, NewOptions())
	switch {
	case err != nil:
		t.Fatal("Failed to compile sass:", err)
	case len(strings.Fields(out)) == 0:
		t.Fatal("Compilation resulted in empty output.")
	}
}
