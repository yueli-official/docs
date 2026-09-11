package v1

import "github.com/gogf/gf/v2/frame/g"

type GetInitialAdministratorClaimStatusReq struct {
	g.Meta `path:"/api/v1/authorization/setup" method:"get" tags:"authorization" summary:"Initial administrator claim status"`
}
type GetInitialAdministratorClaimStatusRes struct {
	Claimed  bool `json:"claimed"`
	CanClaim bool `json:"canClaim"`
}

type ClaimInitialAdministratorReq struct {
	g.Meta `path:"/api/v1/authorization/setup/claim" method:"post" tags:"authorization" summary:"Claim initial administrator"`
}
type ClaimInitialAdministratorRes struct {
	Claimed bool `json:"claimed"`
	Created bool `json:"created"`
}
