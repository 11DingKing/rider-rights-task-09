package domain

import (
	"testing"
	"time"
)

func TestEscalatedCaseRemainsOverdue(t *testing.T) {
	deadline := time.Date(2026, 8, 19, 12, 0, 0, 0, time.UTC)
	item := &RightsCase{Status: StatusEscalated, Deadline: deadline}
	if !item.IsOverdue(deadline.Add(time.Minute)) {
		t.Fatal("escalated case was incorrectly treated as terminal")
	}
}
