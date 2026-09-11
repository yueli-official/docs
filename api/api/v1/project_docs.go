package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/yueli-official/docs/api/internal/projectdocs"
)

type ListProjectDocSourcesReq struct {
	g.Meta `path:"/api/v1/import-sources" method:"get" tags:"docs" summary:"List project documentation sources"`
	Page   int `json:"page" d:"1" v:"min:1"`
	Size   int `json:"size" d:"20" v:"between:1,100"`
}
type ListProjectDocSourcesRes struct {
	Page    int                  `json:"page"`
	Size    int                  `json:"size"`
	Items   []projectdocs.Source `json:"items"`
	Total   int                  `json:"total"`
	Enabled bool                 `json:"enabled"`
}
type CreateProjectDocSourceReq struct {
	g.Meta          `path:"/api/v1/import-sources" method:"post" tags:"docs" summary:"Bind a public GitHub release documentation source"`
	Repository      string `json:"repository" v:"required|length:3,300"`
	AssetName       string `json:"assetName"`
	Collection      string `json:"collection" v:"required"`
	DefaultLocale   string `json:"defaultLocale" v:"required"`
	Mode            string `json:"mode"`
	PersonalToken   string `json:"personalToken" v:"required"`
	ExclusiveMirror bool   `json:"exclusiveMirror"`
	AutoCheck       bool   `json:"autoCheck"`
}
type CreateProjectDocSourceRes struct {
	Source projectdocs.Source `json:"source"`
}
type GetProjectDocSourceReq struct {
	g.Meta `path:"/api/v1/import-sources/{id}" method:"get" tags:"docs" summary:"Get documentation source and recent runs"`
	ID     string `json:"id" in:"path" v:"required|length:36,36"`
}
type GetProjectDocSourceRes struct {
	Source projectdocs.Source `json:"source"`
	Runs   []projectdocs.Run  `json:"runs"`
}
type UpdateProjectDocSourceReq struct {
	g.Meta        `path:"/api/v1/import-sources/{id}" method:"patch" tags:"docs" summary:"Pause or resume a source and optionally replace its credential"`
	ID            string `json:"id" in:"path" v:"required|length:36,36"`
	Enabled       *bool  `json:"enabled"`
	AutoCheck     *bool  `json:"autoCheck"`
	PersonalToken string `json:"personalToken"`
}
type UpdateProjectDocSourceRes struct {
	Source projectdocs.Source `json:"source"`
}
type CheckProjectDocSourceReq struct {
	g.Meta `path:"/api/v1/import-sources/{id}/check" method:"post" tags:"docs" summary:"Queue a documentation source check"`
	ID     string `json:"id" in:"path" v:"required|length:36,36"`
}
type CheckProjectDocSourceRes struct {
	Queued bool `json:"queued"`
}
type DeleteProjectDocSourceReq struct {
	g.Meta `path:"/api/v1/import-sources/{id}" method:"delete" tags:"docs" summary:"Remove source binding without deleting documents"`
	ID     string `json:"id" in:"path" v:"required|length:36,36"`
}
type DeleteProjectDocSourceRes struct{}
