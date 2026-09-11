package docsauthz_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/yueli-official/docs/api/internal/docsauthz"
	"github.com/yueli-official/docs/api/internal/testidentity"
	foundationauth "github.com/yueli-official/foundation/go/auth"
	"github.com/yueli-official/foundation/go/authorization"
)

func TestInitialAdministratorRequiresExplicitUserClaim(t *testing.T) {
	module, err := authorization.NewMemory(authorization.MustCompile(docsauthz.Definition()), authorization.MemoryOptions{
		RootScopeID: docsauthz.RootScopeID, AllowUnclaimed: true,
		Constraints: docsauthz.ConstraintEvaluators(), Predicates: docsauthz.PredicateEvaluators(),
	})
	if err != nil {
		t.Fatal(err)
	}
	service := docsauthz.New(module)
	ctx := context.Background()
	user := foundationauth.NewContext(ctx, testidentity.User(t, "owner", []string{"administrator"}, nil))
	// Identity roles and ordinary access must not initialize this product.
	if _, err := service.EffectiveAccess(user); err != nil {
		t.Fatal(err)
	}
	status, err := service.AdministratorClaimStatus(ctx)
	if err != nil || status.Claimed {
		t.Fatalf("ordinary access claimed site: %#v %v", status, err)
	}
	for _, denied := range []context.Context{ctx, foundationauth.NewContext(ctx, testidentity.Client(t, "service", nil))} {
		if _, err := service.ClaimInitialAdministrator(denied); !authorization.Is(err, authorization.ErrorInvalidInput) {
			t.Fatalf("non-user claim: %v", err)
		}
	}
	result, err := service.ClaimInitialAdministrator(user)
	if err != nil || !result.Created || !result.Status.Claimed {
		t.Fatalf("claim: %#v %v", result, err)
	}
	retry, err := service.ClaimInitialAdministrator(user)
	if err != nil || retry.Created {
		t.Fatalf("winner retry: %#v %v", retry, err)
	}
	other := foundationauth.NewContext(ctx, testidentity.User(t, "other", nil, nil))
	if _, err := service.ClaimInitialAdministrator(other); !authorization.Is(err, authorization.ErrorConflict) {
		t.Fatalf("late claim: %v", err)
	}
}

func TestInitialAdministratorConcurrentClaimsHaveOneWinner(t *testing.T) {
	module, err := authorization.NewMemory(authorization.MustCompile(docsauthz.Definition()), authorization.MemoryOptions{
		RootScopeID: docsauthz.RootScopeID, AllowUnclaimed: true,
		Constraints: docsauthz.ConstraintEvaluators(), Predicates: docsauthz.PredicateEvaluators(),
	})
	if err != nil {
		t.Fatal(err)
	}
	service := docsauthz.New(module)
	outcomes := make(chan error, 20)
	start := make(chan struct{})
	for i := range 20 {
		ctx := foundationauth.NewContext(context.Background(), testidentity.User(t, fmt.Sprintf("candidate-%d", i), nil, nil))
		go func() { <-start; _, err := service.ClaimInitialAdministrator(ctx); outcomes <- err }()
	}
	close(start)
	winners := 0
	for range 20 {
		err := <-outcomes
		if err == nil {
			winners++
		} else if !authorization.Is(err, authorization.ErrorConflict) {
			t.Fatal(err)
		}
	}
	if winners != 1 {
		t.Fatalf("winners = %d", winners)
	}
}

func TestBootstrapClosesInitialAdministratorClaim(t *testing.T) {
	module, err := authorization.NewMemory(authorization.MustCompile(docsauthz.Definition()), authorization.MemoryOptions{
		RootScopeID:       docsauthz.RootScopeID,
		ProtectedSubjects: []authorization.SubjectRef{{Kind: authorization.SubjectUser, ID: "bootstrap"}},
		Constraints:       docsauthz.ConstraintEvaluators(), Predicates: docsauthz.PredicateEvaluators(),
	})
	if err != nil {
		t.Fatal(err)
	}
	service := docsauthz.New(module)
	ctx := foundationauth.NewContext(context.Background(), testidentity.User(t, "candidate", nil, nil))
	status, err := service.AdministratorClaimStatus(ctx)
	if err != nil || !status.Claimed {
		t.Fatalf("bootstrap status: %#v %v", status, err)
	}
	if _, err := service.ClaimInitialAdministrator(ctx); !authorization.Is(err, authorization.ErrorConflict) {
		t.Fatalf("bootstrap claim: %v", err)
	}
}
