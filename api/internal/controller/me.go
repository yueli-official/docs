package controller

import (
	"context"

	v1 "platform/products/docs/api/api/v1"
)

// Me reports the caller's identity for the front-end manage gate. It lives on a
// JWT-optional group: anonymous callers get authenticated=false (NOT a 401), so
// the SPA can decide whether to send them to login vs render the console.
type Me struct{}

func NewMe() *Me { return &Me{} }

func (c *Me) Me(ctx context.Context, _ *v1.MeReq) (*v1.MeRes, error) {
	sub, _ := subject(ctx) // empty when anonymous; subject() returns Forbidden err we ignore
	return &v1.MeRes{Me: &v1.MeView{
		Sub:           sub,
		Authenticated: sub != "",
		IsOwner:       isAdmin(ctx),
	}}, nil
}
