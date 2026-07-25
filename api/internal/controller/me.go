package controller

import (
	"context"
	"slices"

	v1 "platform/products/docs/api/api/v1"
	"platform/products/docs/api/internal/docserr"
)

// Me reports the caller's identity for the front-end manage gate. It lives on a
// JWT-optional group: anonymous callers get authenticated=false (NOT a 401), so
// the SPA can decide whether to send them to login vs render the console.
type Me struct{}

func NewMe() *Me { return &Me{} }

func (c *Me) Me(ctx context.Context, _ *v1.MeReq) (*v1.MeRes, error) {
	sub, _ := subject(ctx) // empty when anonymous; subject() returns Forbidden err we ignore
	service := authorizationService(ctx)
	if service == nil {
		return nil, docserr.AuthorizationUnavailable()
	}
	access, err := service.EffectiveAccess(ctx)
	if err != nil {
		return nil, docserr.AuthorizationUnavailable()
	}
	roles := make([]string, 0, len(access.Grants))
	for _, grant := range access.Grants {
		role := string(grant.Role)
		if !slices.Contains(roles, role) {
			roles = append(roles, role)
		}
	}
	capabilities := make([]string, len(access.Capabilities))
	for index, capability := range access.Capabilities {
		capabilities[index] = string(capability)
	}
	administrator := isAdministrator(ctx)
	return &v1.MeRes{Me: &v1.MeView{
		Sub: sub, Authenticated: sub != "", IsAdministrator: administrator,
		Roles: roles, Capabilities: capabilities,
	}}, nil
}
