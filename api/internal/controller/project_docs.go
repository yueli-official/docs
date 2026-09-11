package controller

import (
	"context"
	v1 "github.com/yueli-official/docs/api/api/v1"
	"github.com/yueli-official/docs/api/internal/docsauthz"
	"github.com/yueli-official/docs/api/internal/docserr"
	"github.com/yueli-official/docs/api/internal/projectdocs"
	"github.com/yueli-official/foundation/go/authorization"
	"github.com/yueli-official/foundation/go/identifier"
)

type ProjectDocs struct{ service *projectdocs.Service }

func NewProjectDocs(service *projectdocs.Service) *ProjectDocs { return &ProjectDocs{service: service} }
func (c *ProjectDocs) authorize(ctx context.Context, ids ...string) error {
	if err := requireCapability(ctx, docsauthz.CapabilityImportManage, docsauthz.RootScopeID, authorization.ResourceFacts{}); err != nil {
		return err
	}
	for _, id := range ids {
		if _, err := identifier.Parse(id); err != nil {
			return docserr.InvalidInput("source_id_invalid")
		}
	}
	if c.service == nil {
		return docserr.ProjectSyncUnavailable()
	}
	return nil
}
func (c *ProjectDocs) ListProjectDocSources(ctx context.Context, req *v1.ListProjectDocSourcesReq) (*v1.ListProjectDocSourcesRes, error) {
	if err := requireCapability(ctx, docsauthz.CapabilityImportManage, docsauthz.RootScopeID, authorization.ResourceFacts{}); err != nil {
		return nil, err
	}
	if c.service == nil {
		return &v1.ListProjectDocSourcesRes{Items: []projectdocs.Source{}, Page: req.Page, Size: req.Size}, nil
	}
	items, total, err := c.service.Store.List(ctx, req.Page, req.Size)
	return &v1.ListProjectDocSourcesRes{Items: items, Total: total, Enabled: true, Page: req.Page, Size: req.Size}, err
}
func (c *ProjectDocs) CreateProjectDocSource(ctx context.Context, req *v1.CreateProjectDocSourceReq) (*v1.CreateProjectDocSourceRes, error) {
	if err := c.authorize(ctx); err != nil {
		return nil, err
	}
	owner, err := subject(ctx)
	if err != nil {
		return nil, err
	}
	source, err := c.service.Create(ctx, owner, projectdocs.CreateInput{Repository: req.Repository, AssetName: req.AssetName, Collection: req.Collection, DefaultLocale: req.DefaultLocale, Mode: req.Mode, Token: req.PersonalToken, ExclusiveMirror: req.ExclusiveMirror, AutoCheck: req.AutoCheck})
	if err != nil {
		return nil, err
	}
	return &v1.CreateProjectDocSourceRes{Source: source}, nil
}
func (c *ProjectDocs) GetProjectDocSource(ctx context.Context, req *v1.GetProjectDocSourceReq) (*v1.GetProjectDocSourceRes, error) {
	if err := c.authorize(ctx, req.ID); err != nil {
		return nil, err
	}
	source, runs, err := c.service.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &v1.GetProjectDocSourceRes{Source: source, Runs: runs}, nil
}
func (c *ProjectDocs) UpdateProjectDocSource(ctx context.Context, req *v1.UpdateProjectDocSourceReq) (*v1.UpdateProjectDocSourceRes, error) {
	if err := c.authorize(ctx, req.ID); err != nil {
		return nil, err
	}
	owner, err := subject(ctx)
	if err != nil {
		return nil, err
	}
	if err = c.service.Update(ctx, req.ID, owner, req.PersonalToken, req.Enabled, req.AutoCheck); err != nil {
		return nil, err
	}
	source, _, err := c.service.Get(ctx, req.ID)
	return &v1.UpdateProjectDocSourceRes{Source: source}, err
}
func (c *ProjectDocs) CheckProjectDocSource(ctx context.Context, req *v1.CheckProjectDocSourceReq) (*v1.CheckProjectDocSourceRes, error) {
	if err := c.authorize(ctx, req.ID); err != nil {
		return nil, err
	}
	if err := c.service.Queue(ctx, req.ID); err != nil {
		return nil, err
	}
	return &v1.CheckProjectDocSourceRes{Queued: true}, nil
}
func (c *ProjectDocs) DeleteProjectDocSource(ctx context.Context, req *v1.DeleteProjectDocSourceReq) (*v1.DeleteProjectDocSourceRes, error) {
	if err := c.authorize(ctx, req.ID); err != nil {
		return nil, err
	}
	return &v1.DeleteProjectDocSourceRes{}, c.service.Delete(ctx, req.ID)
}
