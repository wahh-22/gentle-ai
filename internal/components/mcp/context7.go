package mcp

import (
	"fmt"

	"github.com/gentleman-programming/gentle-ai/v2/internal/versions"
)

var defaultContext7ServerJSON = []byte(fmt.Sprintf("{\n  \"command\": \"npx\",\n  \"args\": [\n    \"-y\",\n    \"--package=@upstash/context7-mcp@%s\",\n    \"--\",\n    \"context7-mcp\"\n  ]\n}\n", versions.Context7MCP))

var defaultContext7OverlayJSON = []byte(fmt.Sprintf("{\n  \"mcpServers\": {\n    \"context7\": {\n      \"command\": \"npx\",\n      \"args\": [\n        \"-y\",\n        \"--package=@upstash/context7-mcp@%s\",\n        \"--\",\n        \"context7-mcp\"\n      ]\n    }\n  }\n}\n", versions.Context7MCP))

// openCodeContext7OverlayJSON is the opencode.json overlay using the new MCP format.
// Context7 is a remote MCP server — no npx needed.
// The context7 entry must replace atomically so legacy local keys do not survive
// deep merge into OpenCode/KiloCode's strict MCP schema.
var openCodeContext7OverlayJSON = []byte("{\n  \"mcp\": {\n    \"context7\": {\n      \"__replace__\": {\n        \"type\": \"remote\",\n        \"url\": \"https://mcp.context7.com/mcp\",\n        \"enabled\": true\n      }\n    }\n  }\n}\n")

// openCodeOpenPetsTermuxOverlayJSON configures OpenPets MCP for Termux.
// Uses the explicit global npm dist entry to avoid shebang/env edge cases in
// Termux wrappers while still routing notifications through Termux:API.
var openCodeOpenPetsTermuxOverlayJSON = []byte("{\n  \"mcp\": {\n    \"openpets\": {\n      \"type\": \"local\",\n      \"command\": [\"node\", \"/data/data/com.termux/files/usr/lib/node_modules/@open-pets/cli/dist/index.js\", \"mcp\", \"--backend\", \"termux\"],\n      \"enabled\": true\n    }\n  }\n}\n")

var openCodeOpenPetsInstructionMarkdown = []byte("<!-- OPENPETS:START -->\n## OpenPets\nOpenPets MCP tools may be available.\nUse OpenPets as a short visible status channel for meaningful coding progress:\n- Use openpets_say only for brief, user-facing, non-sensitive status updates.\n- Do not include code, logs, secrets, URLs, file paths, prompts, or private data.\n- Use openpets_react for small visual progress changes.\n- Use openpets_status only when checking availability or the targeted pet.\n- Do not spam every internal step.\n<!-- OPENPETS:END -->\n")

// openClawContext7OverlayJSON is the OpenClaw openclaw.json overlay.
// OpenClaw rejects top-level mcpServers and expects MCP entries under
// mcp.servers.
var openClawContext7OverlayJSON = []byte(fmt.Sprintf("{\n  \"mcp\": {\n    \"servers\": {\n      \"context7\": {\n        \"command\": \"npx\",\n        \"args\": [\n          \"-y\",\n          \"--package=@upstash/context7-mcp@%s\",\n          \"--\",\n          \"context7-mcp\"\n        ]\n      }\n    }\n  }\n}\n", versions.Context7MCP))

// vsCodeContext7OverlayJSON is the VS Code mcp.json overlay using the "servers" key.
var vsCodeContext7OverlayJSON = []byte("{\n  \"servers\": {\n    \"context7\": {\n      \"type\": \"http\",\n      \"url\": \"https://mcp.context7.com/mcp\"\n    }\n  }\n}\n")

// antigravityContext7OverlayJSON is the Antigravity mcp_config.json overlay.
// Uses mcpServers key (same schema as Claude Code) with serverUrl for HTTP remote.
// The context7 entry must replace atomically so legacy local keys do not survive
// deep merge into this managed MCP server entry.
var antigravityContext7OverlayJSON = []byte("{\n  \"mcpServers\": {\n    \"context7\": {\n      \"__replace__\": {\n        \"serverUrl\": \"https://mcp.context7.com/mcp\"\n      }\n    }\n  }\n}\n")

// kimiContext7OverlayJSON follows Kimi's documented mcp.json "well-known MCP
// config format", using mcpServers + explicit http transport for Context7.
// The context7 entry must replace atomically so legacy local keys do not survive
// deep merge into this managed MCP server entry.
var kimiContext7OverlayJSON = []byte("{\n  \"mcpServers\": {\n    \"context7\": {\n      \"__replace__\": {\n        \"transport\": \"http\",\n        \"url\": \"https://mcp.context7.com/mcp\"\n      }\n    }\n  }\n}\n")

func DefaultContext7ServerJSON() []byte {
	content := make([]byte, len(defaultContext7ServerJSON))
	copy(content, defaultContext7ServerJSON)
	return content
}

func DefaultContext7OverlayJSON() []byte {
	content := make([]byte, len(defaultContext7OverlayJSON))
	copy(content, defaultContext7OverlayJSON)
	return content
}

func OpenCodeContext7OverlayJSON() []byte {
	content := make([]byte, len(openCodeContext7OverlayJSON))
	copy(content, openCodeContext7OverlayJSON)
	return content
}

func OpenCodeOpenPetsTermuxOverlayJSON() []byte {
	content := make([]byte, len(openCodeOpenPetsTermuxOverlayJSON))
	copy(content, openCodeOpenPetsTermuxOverlayJSON)
	return content
}

func OpenCodeOpenPetsInstructionMarkdown() []byte {
	content := make([]byte, len(openCodeOpenPetsInstructionMarkdown))
	copy(content, openCodeOpenPetsInstructionMarkdown)
	return content
}

func OpenClawContext7OverlayJSON() []byte {
	content := make([]byte, len(openClawContext7OverlayJSON))
	copy(content, openClawContext7OverlayJSON)
	return content
}

func VSCodeContext7OverlayJSON() []byte {
	content := make([]byte, len(vsCodeContext7OverlayJSON))
	copy(content, vsCodeContext7OverlayJSON)
	return content
}

func AntigravityContext7OverlayJSON() []byte {
	content := make([]byte, len(antigravityContext7OverlayJSON))
	copy(content, antigravityContext7OverlayJSON)
	return content
}

func KimiContext7OverlayJSON() []byte {
	content := make([]byte, len(kimiContext7OverlayJSON))
	copy(content, kimiContext7OverlayJSON)
	return content
}
