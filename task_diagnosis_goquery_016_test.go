package goquery

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisGoquery016SourceContract(t *testing.T) {
    source, err := os.ReadFile("array.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return s.Slice(index, index+1)") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "return s.Slice(index, index- 1)") {
        t.Fatalf("mutated source contract is still present")
    }
}
