package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseSections_WithMalformedFrontmatter(t *testing.T) {
	data := []byte(`--
a: b
---

# Heading
`)
	sections, err := ParseSections(data)
	require.NoError(t, err)

	assert.Equal(t, string(data), string(sections.Content))
	assert.Equal(t, "", string(sections.FrontMatter))
	assert.Equal(t, 0, sections.ContentOffset)
}

func TestParseSections_WithoutFrontMatter(t *testing.T) {
	data := []byte(`# Example

A paragraph
`)
	sections, err := ParseSections(data)
	require.NoError(t, err)
	assert.Equal(t, string(data), string(sections.Content))
}

func TestParseSections_WithFrontMatterYAML(t *testing.T) {
	data := []byte(`---
prop1: val1
prop2: val2
---

# Example

A paragraph
`)
	sections, err := ParseSections(data)
	require.NoError(t, err)
	assert.Equal(t, "---\nprop1: val1\nprop2: val2\n---", string(sections.FrontMatter))
	assert.Equal(t, "# Example\n\nA paragraph\n", string(sections.Content))
}

func TestParseSections_WithFrontMatterTOML(t *testing.T) {
	data := []byte(`+++
prop1 = "val1"
prop2 = "val2"
+++

# Example

A paragraph
`)
	sections, err := ParseSections(data)
	require.NoError(t, err)
	assert.Equal(t, "+++\nprop1 = \"val1\"\nprop2 = \"val2\"\n+++", string(sections.FrontMatter))
	assert.Equal(t, "# Example\n\nA paragraph\n", string(sections.Content))
}

func TestParseSections_WithFrontMatterJSON(t *testing.T) {
	data := []byte(`{
    "prop1": "val1",
    "prop2": "val2"
}

# Example

A paragraph
`)
	sections, err := ParseSections(data)
	require.NoError(t, err)
	assert.Equal(t, "{\n    \"prop1\": \"val1\",\n    \"prop2\": \"val2\"\n}", string(sections.FrontMatter))
	assert.Equal(t, "# Example\n\nA paragraph\n", string(sections.Content))
}
