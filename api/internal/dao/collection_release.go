package dao

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/yueli-official/docs/api/internal/model"
)

type CloneCollectionReleaseInput struct {
	Source                *model.Collection
	SourceInternalVersion string
	SourceSemanticVersion string
	Target                *model.Collection
	TargetInternalVersion *model.CollectionVersion
	FamilyID              string
	FamilyName            string
}

func (p *PG) InitializeCollectionRelease(
	ctx context.Context,
	collectionID, familyID, familyName, semanticVersion string,
	hook TransactionHook,
) error {
	return p.db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if _, err := tx.Model("collection_release_families").Ctx(ctx).Data(g.Map{
			"id": familyID, "name": familyName,
		}).Insert(); err != nil {
			return err
		}
		result, err := tx.Model(tCollections).Ctx(ctx).
			Where("id", collectionID).
			WhereNull("release_family_id").
			Data(g.Map{
				"release_family_id": familyID,
				"semantic_version":  semanticVersion,
				"updated_at":        gtime.Now(),
			}).Update()
		if err != nil {
			return err
		}
		if affected, err := result.RowsAffected(); err != nil {
			return err
		} else if affected != 1 {
			return fmt.Errorf("collection release was initialized concurrently")
		}
		return runTransactionHook(ctx, tx, hook)
	})
}

// CloneCollectionAsRelease creates an independently editable collection from
// source. The old collection_versions row remains a hidden storage partition;
// the public release identity belongs to the collection itself.
func (p *PG) CloneCollectionAsRelease(
	ctx context.Context,
	in CloneCollectionReleaseInput,
	hook TransactionHook,
) ([]*model.Doc, error) {
	cloned := []*model.Doc{}
	err := p.db.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if in.Source.ReleaseFamilyID == "" {
			if _, err := tx.Model("collection_release_families").Ctx(ctx).Data(g.Map{
				"id": in.FamilyID, "name": in.FamilyName,
			}).Insert(); err != nil {
				return err
			}
			if _, err := tx.Model(tCollections).Ctx(ctx).Where("id", in.Source.ID).Data(g.Map{
				"release_family_id": in.FamilyID,
				"semantic_version":  in.SourceSemanticVersion,
				"updated_at":        gtime.Now(),
			}).Update(); err != nil {
				return err
			}
		}

		if _, err := tx.Model(tCollections).Ctx(ctx).Data(g.Map{
			"id": in.Target.ID, "slug": in.Target.Slug, "title": in.Target.Title,
			"description": in.Target.Description, "cover_asset_id": in.Target.CoverAssetID,
			"cover_url": in.Target.CoverURL, "icon": in.Target.Icon,
			"sort_order": in.Target.SortOrder, "author_sub": in.Target.AuthorSub,
			"release_family_id":          in.Target.ReleaseFamilyID,
			"semantic_version":           in.Target.SemanticVersion,
			"derived_from_collection_id": in.Source.ID,
		}).Insert(); err != nil {
			return err
		}
		if _, err := tx.Model(tCollectionVersions).Ctx(ctx).Data(g.Map{
			"id": in.TargetInternalVersion.ID, "collection_id": in.Target.ID,
			"key": "default", "label": "默认版本", "status": "published",
			"is_default": true, "sort_order": 0,
		}).Insert(); err != nil {
			return err
		}
		if _, err := tx.Ctx(ctx).Exec(`
INSERT INTO collection_locales (
    collection_id, locale, label, title, html_lang, direction,
    is_default, enabled, sort_order, created_at, updated_at
)
SELECT ?::uuid, locale, label, CASE WHEN is_default THEN ? ELSE title END, html_lang, direction,
       is_default, enabled, sort_order, NOW(), NOW()
FROM collection_locales
WHERE collection_id = ?::uuid`, in.Target.ID, in.Target.Title, in.Source.ID); err != nil {
			return err
		}

		var sourceDocs []*model.Doc
		if err := tx.Model(tDocs).Ctx(ctx).
			Where("collection_id", in.Source.ID).
			Where("version_id", in.SourceInternalVersion).
			WhereNull("deleted_at").
			OrderAsc("created_at").
			Scan(&sourceDocs); err != nil {
			return err
		}
		var err error
		cloned, err = cloneDocumentSet(ctx, tx, sourceDocs, documentCloneTarget{
			CollectionID: in.Target.ID, VersionID: in.TargetInternalVersion.ID,
			AuthorSub: in.Target.AuthorSub, Status: "draft",
		})
		if err != nil {
			return err
		}
		return runTransactionHook(ctx, tx, hook)
	})
	return cloned, err
}
