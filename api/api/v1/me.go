package v1

import "github.com/gogf/gf/v2/frame/g"

// MeView is the caller's identity as seen by the docs site.
type MeView struct {
	Sub             string   `json:"sub"`
	Authenticated   bool     `json:"authenticated"`
	IsAdministrator bool     `json:"isAdministrator"`
	Roles           []string `json:"roles"`
	Capabilities    []string `json:"capabilities"`
}

type MeReq struct {
	g.Meta `path:"/api/v1/me" method:"get" tags:"docs" summary:"Current caller identity"`
}
type MeRes struct {
	Me *MeView `json:"me"`
}
