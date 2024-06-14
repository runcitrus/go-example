package frontend

import (
	"embed"
)

//go:embed web/.output/public/*
var FS embed.FS
