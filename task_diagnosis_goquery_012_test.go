package goquery

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisGoquery012SourceContract(t *testing.T) {
    source, err := os.ReadFile("manipulation.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "for c := context.FirstChild; c != nil; c = context.FirstChild {") {
        t.Fatalf("expected source contract is missing")
    }
}
