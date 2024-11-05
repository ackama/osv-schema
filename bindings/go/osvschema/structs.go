package osvschema

import (
	"time"
)

// Package identifies the affected code library or command provided by the
// package.
//
// See: https://ossf.github.io/osv-schema/#affectedpackage-field
type Package struct {
	Ecosystem Ecosystem `json:"ecosystem"`
	Name      string    `json:"name"`
	Purl      string    `json:"purl"`
}

// Event describes a single version that either:
//
//   - Introduces a vulnerability: {"introduced": string}
//   - Fixes a vulnerability: {"fixed": string}
//   - Describes the last known affected version: {"last_affected": string}
//   - Sets an upper limit on the range being described: {"limit": string}
//
// Event instances form part of a “timeline” of status changes for the affected
// package described by the Affected struct.
//
// See: https://ossf.github.io/osv-schema/#affectedrangesevents-fields
type Event struct {
	Introduced   string `json:"introduced"`
	Fixed        string `json:"fixed"`
	LastAffected string `json:"last_affected"`
	Limit        string `json:"limit"`
}

// Range describes the affected range of given version for a specific package.
//
// See: https://ossf.github.io/osv-schema/#affectedranges-field
type Range struct {
	Type             RangeType              `json:"type"`
	Events           []Event                `json:"events"`
	Repo             string                 `json:"repo"`
	DatabaseSpecific map[string]interface{} `json:"database_specific"`
}

// Severity is used to describe the severity of a vulnerability for an affected
// package using one or more quantitative scoring methods.
//
// See: https://ossf.github.io/osv-schema/#severity-field
type Severity struct {
	Type  SeverityType `json:"type"`
	Score string       `json:"score"`
}

// Affected describes an affected package version, meaning one instance that
// contains the vulnerability.
//
// See: https://ossf.github.io/osv-schema/#affected-fields
type Affected struct {
	Package           Package                `json:"package"`
	Severity          []Severity             `json:"severity"`
	Ranges            []Range                `json:"ranges"`
	Versions          []string               `json:"versions"`
	DatabaseSpecific  map[string]interface{} `json:"database_specific"`
	EcosystemSpecific map[string]interface{} `json:"ecosystem_specific"`
}

// Reference links to additional information, advisories, issue tracker entries,
// and so on about the vulnerability itself.
//
// See: https://ossf.github.io/osv-schema/#references-field
type Reference struct {
	Type ReferenceType `json:"type"`
	URL  string        `json:"url"`
}

// Credit gives credit for the discovery, confirmation, patch, or other events
// in the life cycle of a vulnerability.
//
// See: https://ossf.github.io/osv-schema/#credits-fields
type Credit struct {
	Name    string     `json:"name"`
	Type    CreditType `json:"type"`
	Contact []string   `json:"contact"`
}

// Vulnerability is the core Open Source Vulnerability (OSV) data type.
//
// The full documentation for the schema is available at
// https://ossf.github.io/osv-schema.
type Vulnerability struct {
	SchemaVersion    string                 `json:"schema_version"`
	ID               string                 `json:"id"`
	Modified         time.Time              `json:"modified"`
	Published        time.Time              `json:"published"`
	Withdrawn        time.Time              `json:"withdrawn"`
	Aliases          []string               `json:"aliases"`
	Related          []string               `json:"related"`
	Details          string                 `json:"details"`
	Summary          string                 `json:"summary"`
	Severity         []Severity             `json:"severity"`
	Affected         []Affected             `json:"affected"`
	References       []Reference            `json:"references"`
	Credits          []Credit               `json:"credits"`
	DatabaseSpecific map[string]interface{} `json:"database_specific"`
}
