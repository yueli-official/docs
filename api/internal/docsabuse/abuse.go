package docsabuse

import (
	"net/netip"
	"strings"
	"time"

	"github.com/yueli-official/foundation/go/abuse"
)

const ActionAuthorApplication abuse.ActionKey = "docs.author_application.create"

type Policy struct {
	ActorCapacity   int64
	NetworkCapacity int64
	Window          time.Duration
	Challenge       *abuse.ChallengeDefinition
}

func Definition(policy Policy) abuse.Definition {
	if policy.ActorCapacity <= 0 {
		policy.ActorCapacity = 3
	}
	if policy.NetworkCapacity <= 0 {
		policy.NetworkCapacity = 20
	}
	if policy.Window <= 0 {
		policy.Window = 24 * time.Hour
	}
	challengeAt := int64(0)
	if policy.Challenge != nil {
		challengeAt = policy.ActorCapacity
	}
	return abuse.Definition{
		Version:  1,
		Consumer: "docs",
		Actions: []abuse.ActionDefinition{{
			Key: ActionAuthorApplication,
			Required: abuse.SignalRequirements{
				Network: abuse.Required,
				Actor:   abuse.Required,
			},
			Meters: []abuse.MeterDefinition{
				{
					ID:        "docs.author_application.network",
					Slot:      abuse.SlotNetwork,
					Algorithm: abuse.TokenBucket(policy.NetworkCapacity, policy.NetworkCapacity, policy.Window),
				},
				{
					ID:          "docs.author_application.actor",
					Slot:        abuse.SlotActor,
					Algorithm:   abuse.FixedWindow(policy.ActorCapacity, policy.Window),
					ChallengeAt: challengeAt,
				},
			},
			Challenge: policy.Challenge,
		}},
	}
}

type Actions struct {
	AuthorApplication abuse.Action
}

func Bind(module abuse.Module) (Actions, error) {
	action, err := module.Action(ActionAuthorApplication)
	if err != nil {
		return Actions{}, err
	}
	return Actions{AuthorApplication: action}, nil
}

func NetworkPrefix(value string) (netip.Prefix, error) {
	address, err := netip.ParseAddr(strings.TrimSpace(value))
	if err != nil {
		return netip.Prefix{}, err
	}
	address = address.Unmap()
	bits := address.BitLen()
	if address.Is6() {
		bits = 64
	}
	return netip.PrefixFrom(address, bits).Masked(), nil
}
