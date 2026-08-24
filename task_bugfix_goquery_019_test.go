package goquery

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixGoquery019SourceContract(t *testing.T) {
    source, err := os.ReadFile("expand.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if sel == nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
