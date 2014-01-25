package sass

import "testing"

func TestContext(t *testing.T) {
	c, err := NewContext()
	if err != nil {
		t.Fatal("Failed to create context:", err)
	}
	err = c.Free()
	if err != nil {
		t.Fatal("Failed to free context:", err)
	}
}

func TestFileContext(t *testing.T) {
	c, err := NewFileContext()
	if err != nil {
		t.Fatal("Failed to create file context:", err)
	}
	err = c.Free()
	if err != nil {
		t.Fatal("Failed to free file context:", err)
	}
}

func TestFolderContext(t *testing.T) {
	c, err := NewFolderContext()
	if err != nil {
		t.Fatal("Failed to create folder context:", err)
	}
	err = c.Free()
	if err != nil {
		t.Fatal("Failed to free folder context:", err)
	}
}
