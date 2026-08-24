package goquery

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGoquery003SourceContract(t *testing.T) {
    source, err := os.ReadFile("traversal.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if sel == nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
