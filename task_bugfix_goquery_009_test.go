package goquery

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGoquery009SourceContract(t *testing.T) {
    source, err := os.ReadFile("property.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if attr == nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
