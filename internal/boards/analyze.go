package boards

import (
	"context"
	"sort"
	"strings"
)

func (o Options) sortAndGatherMetrics(ctx context.Context, wrapper BoardWrapper) BoardWrapper {
	sort.Slice(wrapper.Boards, func(i, j int) bool {
		if wrapper.Boards[i].Vendor != wrapper.Boards[j].Vendor {
			return wrapper.Boards[i].Vendor < wrapper.Boards[j].Vendor
		}

		return wrapper.Boards[i].Name < wrapper.Boards[j].Name
	})

	vendors := make(map[string]int)
	for _, board := range wrapper.Boards {
		wrapper.Metadata.Totals.Boards++
		if board.HasWifi {
			wrapper.Metadata.Totals.WifiEnabled++
		}
		vendor := sanitizeVendor(board.Vendor)
		_, ok := vendors[vendor]
		if ok {
			vendors[vendor]++
		} else {
			vendors[vendor] = 1
		}
	}
	wrapper.Metadata.Totals.Vendors = len(vendors)
	return wrapper
}

func sanitizeVendor(vendor string) string {
	return strings.ToLower(vendor)
}
