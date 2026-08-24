package goquery

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGoquery002SourceContract(t *testing.T) {
    source, err := os.ReadFile("query.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return false") {
        t.Fatalf("expected source contract is missing")
    }
}
