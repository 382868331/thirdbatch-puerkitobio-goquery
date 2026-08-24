package goquery

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGoquery018SourceContract(t *testing.T) {
    source, err := os.ReadFile("manipulation.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if parent == nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
