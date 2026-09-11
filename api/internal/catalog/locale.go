package catalog

import (
	"context"
	"regexp"
	"strings"

	"github.com/yueli-official/docs/api/internal/docserr"
	"github.com/yueli-official/docs/api/internal/model"
)

var localePattern = regexp.MustCompile(`^[A-Za-z]{2,3}(-[A-Za-z0-9]{2,8})*$`)

type UpsertLocaleInput struct {
	Title        *string
	CollectionID string
	Locale       string
	Label        string
	HTMLLang     string
	Direction    string
	IsDefault    bool
	Enabled      bool
	SortOrder    int
}

type CloneLocaleInput struct {
	CollectionID    string
	SourceLocale    string
	TargetLocale    string
	TargetLabel     string
	TargetHTMLLang  string
	TargetDirection string
	TargetSortOrder int
}

func (s *Service) ListLocales(ctx context.Context, collectionID string, enabledOnly bool) ([]*model.CollectionLocale, error) {
	return s.dao.ListCollectionLocales(ctx, collectionID, enabledOnly)
}

func (s *Service) DefaultLocale(ctx context.Context, collectionID string) (*model.CollectionLocale, error) {
	return s.dao.GetDefaultCollectionLocale(ctx, collectionID)
}

func (s *Service) ResolveLocale(ctx context.Context, collectionID, locale string, enabledOnly bool) (*model.CollectionLocale, error) {
	var (
		value *model.CollectionLocale
		err   error
	)
	if strings.TrimSpace(locale) == "" {
		value, err = s.dao.GetDefaultCollectionLocale(ctx, collectionID)
	} else {
		value, err = s.dao.GetCollectionLocale(ctx, collectionID, locale)
	}
	if err != nil {
		return nil, err
	}
	if value == nil || enabledOnly && !value.Enabled {
		return nil, docserr.NotFound(locale)
	}
	return value, nil
}

func (s *Service) UpsertLocale(ctx context.Context, in UpsertLocaleInput) (*model.CollectionLocale, error) {
	in.Locale = strings.TrimSpace(in.Locale)
	in.Label = strings.TrimSpace(in.Label)
	in.HTMLLang = strings.TrimSpace(in.HTMLLang)
	in.Direction = strings.TrimSpace(in.Direction)
	if !localePattern.MatchString(in.Locale) {
		return nil, docserr.InvalidInput("invalid documentation locale")
	}
	if in.Label == "" {
		return nil, docserr.InvalidInput("locale label is required")
	}
	if in.HTMLLang == "" {
		in.HTMLLang = in.Locale
	}
	if !localePattern.MatchString(in.HTMLLang) {
		return nil, docserr.InvalidInput("invalid locale htmlLang")
	}
	if in.Direction == "" {
		in.Direction = "ltr"
	}
	if in.Direction != "ltr" && in.Direction != "rtl" {
		return nil, docserr.InvalidInput("locale direction must be ltr or rtl")
	}
	if in.IsDefault && !in.Enabled {
		return nil, docserr.InvalidInput("default locale must be enabled")
	}
	collection, err := s.dao.GetCollectionByID(ctx, in.CollectionID)
	if err != nil {
		return nil, err
	}
	if collection == nil {
		return nil, docserr.NotFound(in.CollectionID)
	}
	existing, err := s.dao.GetCollectionLocale(ctx, in.CollectionID, in.Locale)
	if err != nil {
		return nil, err
	}
	if existing != nil && existing.IsDefault && (!in.IsDefault || !in.Enabled) {
		return nil, docserr.InvalidInput("choose another default locale before disabling this locale")
	}
	title := collection.Title
	if existing != nil {
		title = existing.Title
	}
	if in.Title != nil {
		title = strings.TrimSpace(*in.Title)
		if title == "" {
			return nil, docserr.InvalidInput("collection locale title is required")
		}
	}
	value := &model.CollectionLocale{
		Title:        title,
		CollectionID: in.CollectionID,
		Locale:       in.Locale,
		Label:        in.Label,
		HTMLLang:     in.HTMLLang,
		Direction:    in.Direction,
		IsDefault:    in.IsDefault,
		Enabled:      in.Enabled,
		SortOrder:    in.SortOrder,
	}
	if err := s.dao.UpsertCollectionLocale(ctx, value, s.urlReconcileHook(in.CollectionID, "docs locale updated")); err != nil {
		return nil, err
	}
	return s.dao.GetCollectionLocale(ctx, in.CollectionID, in.Locale)
}

func (s *Service) DeleteLocale(ctx context.Context, collectionID, locale string) error {
	value, err := s.dao.GetCollectionLocale(ctx, collectionID, locale)
	if err != nil {
		return err
	}
	if value == nil {
		return docserr.NotFound(locale)
	}
	if value.IsDefault {
		return docserr.InvalidInput("default locale cannot be deleted")
	}
	count, err := s.dao.CountCollectionLocaleDocs(ctx, collectionID, locale)
	if err != nil {
		return err
	}
	if count > 0 {
		return docserr.InvalidInput("locale with documents cannot be deleted")
	}
	return s.dao.DeleteCollectionLocale(ctx, collectionID, locale)
}

func (s *Service) CloneLocale(ctx context.Context, authorSub string, in CloneLocaleInput) (int, error) {
	in.SourceLocale = strings.TrimSpace(in.SourceLocale)
	in.TargetLocale = strings.TrimSpace(in.TargetLocale)
	in.TargetLabel = strings.TrimSpace(in.TargetLabel)
	in.TargetHTMLLang = strings.TrimSpace(in.TargetHTMLLang)
	in.TargetDirection = strings.TrimSpace(in.TargetDirection)
	if in.SourceLocale == in.TargetLocale {
		return 0, docserr.InvalidInput("source and target locales must differ")
	}
	if !localePattern.MatchString(in.TargetLocale) {
		return 0, docserr.InvalidInput("invalid target documentation locale")
	}
	if in.TargetLabel == "" {
		return 0, docserr.InvalidInput("target locale label is required")
	}
	if in.TargetHTMLLang == "" {
		in.TargetHTMLLang = in.TargetLocale
	}
	if !localePattern.MatchString(in.TargetHTMLLang) {
		return 0, docserr.InvalidInput("invalid target locale htmlLang")
	}
	if in.TargetDirection == "" {
		in.TargetDirection = "ltr"
	}
	if in.TargetDirection != "ltr" && in.TargetDirection != "rtl" {
		return 0, docserr.InvalidInput("target locale direction must be ltr or rtl")
	}
	source, err := s.dao.GetCollectionLocale(ctx, in.CollectionID, in.SourceLocale)
	if err != nil {
		return 0, err
	}
	if source == nil || !source.Enabled {
		return 0, docserr.InvalidInput("source locale must be enabled for this collection")
	}
	target, err := s.dao.GetCollectionLocale(ctx, in.CollectionID, in.TargetLocale)
	if err != nil {
		return 0, err
	}
	if target != nil && !target.Enabled {
		return 0, docserr.InvalidInput("target locale must be enabled for this collection")
	}
	sourceCount, err := s.dao.CountCollectionLocaleDocs(ctx, in.CollectionID, in.SourceLocale)
	if err != nil {
		return 0, err
	}
	if sourceCount == 0 {
		return 0, docserr.InvalidInput("source locale has no documents to copy")
	}
	targetCount, err := s.dao.CountCollectionLocaleDocs(ctx, in.CollectionID, in.TargetLocale)
	if err != nil {
		return 0, err
	}
	if targetCount > 0 {
		return 0, docserr.InvalidInput("target locale already contains documents")
	}
	version, err := s.dao.GetDefaultCollectionVersion(ctx, in.CollectionID)
	if err != nil {
		return 0, err
	}
	if version == nil {
		return 0, docserr.InvalidInput("collection has no storage partition")
	}
	locale := &model.CollectionLocale{
		CollectionID: in.CollectionID, Locale: in.TargetLocale,
		Label: in.TargetLabel, Title: source.Title, HTMLLang: in.TargetHTMLLang,
		Direction: in.TargetDirection, Enabled: true, SortOrder: in.TargetSortOrder,
	}
	return s.dao.CloneCollectionLocale(
		ctx, in.CollectionID, version.ID, in.SourceLocale, locale, authorSub,
		s.urlReconcileHook(in.CollectionID, "docs locale derived"),
	)
}
