package controller

import (
	"context"
	"github.com/yueli-official/foundation/go/authorization"

	v1 "github.com/yueli-official/docs/api/api/v1"
	"github.com/yueli-official/docs/api/internal/docsauthz"
	"github.com/yueli-official/docs/api/internal/docserr"
)

// AuthorizationSetup exposes only the non-sensitive, read-only claim state.
// The mutating claim endpoint remains on the authenticated Authorization API.
type AuthorizationSetup struct {
	service *docsauthz.Service
}

func NewAuthorizationSetup(service *docsauthz.Service) *AuthorizationSetup {
	return &AuthorizationSetup{service: service}
}

func (controller *AuthorizationSetup) GetInitialAdministratorClaimStatus(
	ctx context.Context,
	_ *v1.GetInitialAdministratorClaimStatusReq,
) (*v1.GetInitialAdministratorClaimStatusRes, error) {
	if controller == nil || controller.service == nil {
		return nil, docserr.AuthorizationUnavailable()
	}
	status, err := controller.service.AdministratorClaimStatus(ctx)
	if err != nil {
		return nil, mapAuthorizationError(err)
	}
	return &v1.GetInitialAdministratorClaimStatusRes{
		Claimed: status.Claimed, CanClaim: !status.Claimed,
	}, nil
}

func (*Authorization) ClaimInitialAdministrator(
	ctx context.Context,
	_ *v1.ClaimInitialAdministratorReq,
) (*v1.ClaimInitialAdministratorRes, error) {
	service := authorizationService(ctx)
	if service == nil {
		return nil, docserr.AuthorizationUnavailable()
	}
	result, err := service.ClaimInitialAdministrator(ctx)
	if err != nil {
		if authorization.Is(err, authorization.ErrorConflict) {
			return nil, docserr.InitialAdministratorAlreadyClaimed()
		}
		return nil, mapAuthorizationError(err)
	}
	return &v1.ClaimInitialAdministratorRes{
		Claimed: result.Status.Claimed, Created: result.Created,
	}, nil
}
