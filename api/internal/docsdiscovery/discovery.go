package docsdiscovery

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/yueli-official/foundation/go/discovery"

	"github.com/yueli-official/docs/api/internal/dao"
	"github.com/yueli-official/docs/api/internal/model"
)

type Config struct {
	Origin        string
	Name          string
	Description   string
	DefaultLocale string
	TTL           time.Duration
	Clock         func() time.Time
}

func New(store *dao.PG, config Config) (*discovery.Module, *discovery.Cache, error) {
	module, err := discovery.Compile(discovery.Definition{
		ContractVersion: discovery.ContractVersion,
		Site: discovery.SiteProfile{
			Origin: config.Origin, Name: config.Name,
			Description: config.Description, DefaultLocale: config.DefaultLocale,
		},
		URLPolicy: discovery.URLPolicy{PreserveQuery: true},
	})
	if err != nil {
		return nil, nil, err
	}
	cache, err := discovery.NewCache(module, discovery.CacheOptions{
		TTL: config.TTL, Clock: config.Clock,
		Build: func(context.Context) (discovery.PublicationPlan, discovery.Sources, error) {
			return discovery.PublicationPlan{
					Sitemap: &discovery.SitemapPlan{Source: "pages"},
					Robots:  &discovery.RobotsPlan{},
				}, discovery.Sources{
					"pages": &source{module: module, store: store, config: config},
				}, nil
		},
	})
	if err != nil {
		return nil, nil, err
	}
	return module, cache, nil
}

func ProjectDoc(
	module *discovery.Module,
	doc *model.Doc,
	collection *model.Collection,
	version *model.CollectionVersion,
	pathValue string,
	defaultLocale string,
) (discovery.PageProjection, error) {
	if module == nil || doc == nil || collection == nil || version == nil || doc.Status != "published" {
		return discovery.PageProjection{}, fmt.Errorf("published doc, collection, version and Discovery module are required")
	}
	title := doc.Title
	if doc.SEOTitle != "" {
		title = doc.SEOTitle
	}
	description := doc.Excerpt
	if doc.SEODescription != "" {
		description = doc.SEODescription
	}
	pagePath := "/" + strings.Trim(collection.Slug+"/"+pathValue, "/")
	query := url.Values{}
	if doc.Locale != "" && doc.Locale != defaultLocale {
		query.Set("locale", doc.Locale)
	}
	if !version.IsDefault {
		query.Set("version", version.Key)
	}
	if encoded := query.Encode(); encoded != "" {
		pagePath += "?" + encoded
	}
	projection, _, err := module.Project(discovery.PageDescriptor{
		Key: "doc:" + doc.ID, Path: pagePath, Locale: doc.Locale,
		Subject: discovery.WebPageSubject(discovery.WebPage{
			Title: title, Description: description,
		}),
		Breadcrumbs: []discovery.Breadcrumb{{
			Name: collection.Title, Path: "/" + collection.Slug,
		}},
	})
	return projection, err
}

type source struct {
	module *discovery.Module
	store  *dao.PG
	config Config
}

func (value *source) Next(ctx context.Context, cursor discovery.Cursor, limit int) (discovery.Batch, error) {
	after := string(cursor)
	rowsLimit := limit + 1
	records := make([]discovery.Record, 0, limit)
	if after == "" {
		page := discovery.PageDescriptor{
			Key: "site:home", Path: "/", Locale: value.config.DefaultLocale,
			Subject: discovery.WebPageSubject(discovery.WebPage{
				Title: value.config.Name, Description: value.config.Description,
			}),
		}
		projection, _, err := value.module.Project(page)
		if err != nil {
			return discovery.Batch{}, err
		}
		records = append(records, discovery.Record{SortKey: projection.CanonicalURL, Page: page})
		after = projection.CanonicalURL
		rowsLimit = limit
	}
	rows, err := value.store.ListDiscoveryPages(
		ctx, value.config.Origin, value.config.DefaultLocale, after, rowsLimit,
	)
	if err != nil {
		return discovery.Batch{}, err
	}
	hasMore := len(rows) >= rowsLimit
	if hasMore {
		rows = rows[:rowsLimit-1]
	}
	for _, row := range rows {
		record, err := value.record(row)
		if err != nil {
			return discovery.Batch{}, err
		}
		records = append(records, record)
		if len(records) == limit {
			hasMore = true
			break
		}
	}
	var next discovery.Cursor
	if len(records) > 0 {
		next = discovery.Cursor(records[len(records)-1].SortKey)
	}
	return discovery.Batch{Records: records, NextCursor: next, Done: !hasMore}, nil
}

func (value *source) record(row dao.DiscoveryRow) (discovery.Record, error) {
	var modifiedAt *time.Time
	if row.UpdatedAt != nil {
		updated := row.UpdatedAt.Time.UTC()
		modifiedAt = &updated
	}
	var subject discovery.Subject
	if row.Kind == "collection" {
		var image *discovery.Image
		if row.ImageURL != "" {
			image = &discovery.Image{URL: row.ImageURL, Alt: row.Title}
		}
		subject = discovery.CollectionSubject(discovery.Collection{
			Title: row.Title, Description: row.Description, Image: image,
		})
	} else {
		subject = discovery.WebPageSubject(discovery.WebPage{
			Title: row.Title, Description: row.Description,
		})
	}
	page := discovery.PageDescriptor{
		Key: row.Key, Path: row.Path, Locale: row.Locale,
		ModifiedAt: modifiedAt, Subject: subject,
	}
	projection, _, err := value.module.Project(page)
	if err != nil {
		return discovery.Record{}, err
	}
	return discovery.Record{SortKey: projection.CanonicalURL, Page: page}, nil
}
