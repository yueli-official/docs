package docscomments

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestDocumentCommentsKeepTwoLevelThreadsAndModeration(t *testing.T) {
	ctx := context.Background()
	store := NewMemory()
	store.SeedDocument(Document{
		ID: "doc-1", Title: "安装", CollectionSlug: "quickstart",
		VersionKey: "default", SlugPath: "install", Locale: "zh-CN",
	})
	module := New(store)
	now := time.Date(2026, 8, 24, 12, 0, 0, 0, time.UTC)
	module.clock = func() time.Time {
		now = now.Add(time.Second)
		return now
	}

	top, err := module.Create(ctx, "doc-1", "member-123456", "测试用户", "很好用", "")
	if err != nil {
		t.Fatal(err)
	}
	reply, err := module.Create(ctx, "doc-1", "member-2", "回复者", "我也这样觉得", top.ID)
	if err != nil {
		t.Fatal(err)
	}
	nested, err := module.Create(ctx, "doc-1", "member-3", "第三人", "补充一点", reply.ID)
	if err != nil {
		t.Fatal(err)
	}
	if nested.ParentID != top.ID {
		t.Fatalf("nested parent = %q, want top-level %q", nested.ParentID, top.ID)
	}

	threads, total, _, _, err := module.List(ctx, "doc-1", 1, 20)
	if err != nil {
		t.Fatal(err)
	}
	if total != 1 || len(threads) != 1 || len(threads[0].Replies) != 2 {
		t.Fatalf("threads=%d total=%d replies=%d", len(threads), total, len(threads[0].Replies))
	}

	if _, err := module.Moderate(ctx, reply.ID, StatusSpam); err != nil {
		t.Fatal(err)
	}
	threads, _, _, _, err = module.List(ctx, "doc-1", 1, 20)
	if err != nil {
		t.Fatal(err)
	}
	if len(threads[0].Replies) != 1 {
		t.Fatalf("approved replies = %d, want 1", len(threads[0].Replies))
	}

	managed, err := module.Manage(ctx, AdminQuery{Status: StatusSpam, Page: 1, Size: 20})
	if err != nil {
		t.Fatal(err)
	}
	if managed.Total != 1 || managed.Items[0].Document.Title != "安装" {
		t.Fatalf("managed = %+v", managed)
	}

	if err := module.Delete(ctx, top.ID); err != nil {
		t.Fatal(err)
	}
	threads, total, _, _, err = module.List(ctx, "doc-1", 1, 20)
	if err != nil {
		t.Fatal(err)
	}
	if total != 0 || len(threads) != 0 {
		t.Fatalf("after delete threads=%d total=%d", len(threads), total)
	}
}

func TestDocumentCommentsRejectInvalidOrMissingTargets(t *testing.T) {
	module := New(NewMemory())
	if _, err := module.Create(context.Background(), "missing", "user", "", "comment", ""); !errors.Is(err, ErrDocumentNotFound) {
		t.Fatalf("missing document error = %v", err)
	}
	if _, err := module.Create(context.Background(), "missing", "user", "", "", ""); !errors.Is(err, ErrInvalidInput) {
		t.Fatalf("empty content error = %v", err)
	}
	if _, err := module.Manage(context.Background(), AdminQuery{Status: "unknown"}); !errors.Is(err, ErrInvalidInput) {
		t.Fatalf("invalid status error = %v", err)
	}
}
