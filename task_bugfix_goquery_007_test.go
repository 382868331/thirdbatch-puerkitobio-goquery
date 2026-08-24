package goquery

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGoquery007SourceContract(t *testing.T) {
    source, err := os.ReadFile("array.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if sel != nil && len(sel.Nodes) > 0 {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if sel == nil && len(sel.Nodes) > 0 {") {
        t.Fatalf("mutated source contract is still present")
    }
}
