package goquery

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGoquery005SourceContract(t *testing.T) {
    source, err := os.ReadFile("property.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "for c := s.Nodes[0].FirstChild; c != nil; c = c.NextSibling {") {
        t.Fatalf("expected source contract is missing")
    }
}
