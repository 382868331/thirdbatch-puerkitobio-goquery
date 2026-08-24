package goquery

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGoquery013SourceContract(t *testing.T) {
    source, err := os.ReadFile("manipulation.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if ss.Nodes[0].Data != \"body\" {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if ss.Nodes[1].Data != \"body\" {") {
        t.Fatalf("mutated source contract is still present")
    }
}
