package importkit

import (
	"encoding/json"
	"fmt"
)

func parseManifest(data []byte) (Manifest, error) {
	var m Manifest
	if err := json.Unmarshal(data, &m); err != nil {
		return Manifest{}, fmt.Errorf("manifest invalid json: %w", err)
	}
	if m.SchemaVersion != 1 {
		return Manifest{}, fmt.Errorf("manifest schemaVersion must be 1")
	}
	if m.Collection == "" {
		return Manifest{}, fmt.Errorf("manifest collection is required")
	}
	if m.Version == "" {
		return Manifest{}, fmt.Errorf("manifest version is required")
	}
	if m.DefaultLocale == "" {
		return Manifest{}, fmt.Errorf("manifest defaultLocale is required")
	}
	if len(m.Locales) == 0 {
		return Manifest{}, fmt.Errorf("manifest locales is required")
	}
	if m.Mode == "" {
		m.Mode = "upsert"
	}
	if m.Mode != "create-only" && m.Mode != "upsert" && m.Mode != "replace-version" {
		return Manifest{}, fmt.Errorf("manifest mode is invalid")
	}
	return m, nil
}
