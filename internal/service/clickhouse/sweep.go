//go:build sweep
// +build sweep

package clickhouse

import (
	"github.com/aiven/terraform-provider-aiven/internal/sweep"
)

func init() {
	sweep.AddServiceSweeper("clickhouse")
}
