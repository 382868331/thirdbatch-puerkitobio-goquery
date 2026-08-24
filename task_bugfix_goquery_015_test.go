package goquery

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGoquery015SourceContract(t *testing.T) {
    source, err := os.ReadFile("array.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return s.Eq(-1)") {
        t.Fatalf("expected source contract is missing")
    }
}
