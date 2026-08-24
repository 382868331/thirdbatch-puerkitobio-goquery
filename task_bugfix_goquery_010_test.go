package goquery

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGoquery010SourceContract(t *testing.T) {
    source, err := os.ReadFile("filter.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if sel == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if sel != nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
