// Package secretsource registers all the internal secret sources when imported
package secretsource

import (
	_ "github.com/shiqinzhi/k6/internal/secretsource/file" // import them for init
	_ "github.com/shiqinzhi/k6/internal/secretsource/mock" // import them for init
)
