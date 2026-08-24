package goquery

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGoquery011SourceContract(t *testing.T) {
    source, err := os.ReadFile("property.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(s.Nodes) != 0 {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if len(s.Nodes) == 0 {") {
        t.Fatalf("mutated source contract is still present")
    }
}
