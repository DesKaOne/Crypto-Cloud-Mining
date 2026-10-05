package application

import (
	"context"
	"testing"
	"time"

	"github.com/DesKaOne/Crypto-Cloud-Mining/backend/internal/domain"
)

type fixedClock struct{ now time.Time }

func (c fixedClock) Now() time.Time { return c.now }

type planReader struct{ plan *domain.MiningPlan }

func (p planReader) GetByID(context.Context, domain.MiningPlanID) (*domain.MiningPlan, error) {
	return p.plan, nil
}

type sessionStore struct{ created *domain.MiningSession }

func (s *sessionStore) Create(_ context.Context, session domain.MiningSession) error {
	s.created = &session
	return nil
}

func (s *sessionStore) GetByID(context.Context, domain.MiningSessionID) (*domain.MiningSession, error) {
	return s.created, nil
}

func (s *sessionStore) Update(_ context.Context, session domain.MiningSession) error {
	s.created = &session
	return nil
}

func TestStartMiningSessionUsesServerTimeAndPlanVersion(t *testing.T) {
	now := time.Date(2026, 10, 4, 10, 20, 30, 123, time.FixedZone("WIB", 7*60*60))
	plan := &domain.MiningPlan{ID: "plan-1", AssetID: "asset-1", Version: 7, Status: "active"}
	store := &sessionStore{}

	uc := StartMiningSession{
		Plans:    planReader{plan: plan},
		Sessions: store,
		Clock:    fixedClock{now: now},
	}

	err := uc.Execute(context.Background(), "user-1", "session-1", "plan-1")
	if err != nil {
		t.Fatal(err)
	}

	if store.created.StartedAt == nil || !store.created.StartedAt.Equal(now.UTC()) {
		t.Fatalf("expected authoritative UTC time, got %v", store.created.StartedAt)
	}
	if store.created.PlanVersion != 7 {
		t.Fatalf("expected plan version 7, got %d", store.created.PlanVersion)
	}
	if store.created.Status != "active" {
		t.Fatalf("expected active status, got %s", store.created.Status)
	}
}

func TestChangeMiningSessionStatusRejectsTerminalTransition(t *testing.T) {
	session := &domain.MiningSession{Status: "completed"}
	if err := ChangeMiningSessionStatus(session, "active", time.Now()); err == nil {
		t.Fatal("expected invalid transition error")
	}
}
