package goquery

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisGoquery004SourceContract(t *testing.T) {
    source, err := os.ReadFile("type.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if res == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if res != nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
