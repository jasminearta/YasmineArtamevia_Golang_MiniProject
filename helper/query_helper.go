// helper/query_helper.go
package helper

import "fmt"

// GenerateProductQuery membuat query deskripsi produk berdasarkan nama, material, dan apakah produk mengandung plastik atau tidak.
func GenerateProductQuery(productName, material string, isPlastic bool) string {
	var query string

	if isPlastic {
		query = fmt.Sprintf("Jelaskan produk '%s', yang terbuat dari %s. Karena mengandung plastik, jelaskan dampak lingkungannya dan penggunaannya.", productName, material)
	} else {
		query = fmt.Sprintf("Jelaskan produk '%s', yang terbuat dari %s. Jelaskan dampak lingkungannya dan penggunaannya.", productName, material)
	}

	return query
}
