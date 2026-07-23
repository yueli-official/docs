package docsabuse

import (
	"context"
	"testing"
	"time"

	"github.com/yueli-official/foundation/go/abuse"
)

func TestAuthorApplicationUsesActorAndNetworkBudgets(t *testing.T) {
	module, err := abuse.NewMemory(
		abuse.MustCompile(Definition(Policy{
			ActorCapacity: 2, NetworkCapacity: 10, Window: time.Hour,
		})),
		abuse.MemoryOptions{Secret: []byte("docs-abuse-test-secret-at-least-32-bytes")},
	)
	if err != nil {
		t.Fatal(err)
	}
	actions, err := Bind(module)
	if err != nil {
		t.Fatal(err)
	}
	network, err := NetworkPrefix("192.0.2.20")
	if err != nil {
		t.Fatal(err)
	}
	for index, want := range []abuse.Disposition{
		abuse.DispositionAllow,
		abuse.DispositionAllow,
		abuse.DispositionReject,
	} {
		got, err := actions.AuthorApplication.Admit(context.Background(), abuse.Input{
			ID:      abuse.AttemptID("application-" + string(rune('a'+index))),
			Signals: abuse.Signals{Network: network, Actor: "user-1"},
		})
		if err != nil {
			t.Fatal(err)
		}
		if got.Disposition != want {
			t.Fatalf("attempt %d: got %q, want %q", index+1, got.Disposition, want)
		}
	}
}
