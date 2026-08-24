package goquery

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGoquery006SourceContract(t *testing.T) {
    source, err := os.ReadFile("array.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "for n := s.Nodes[0].PrevSibling; n != nil; n = n.PrevSibling {") {
        t.Fatalf("expected source contract is missing")
    }
}
