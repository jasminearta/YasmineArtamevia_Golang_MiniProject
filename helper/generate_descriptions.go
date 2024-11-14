// helper/generate_descriptions.go
package helper

import (
	"context"
	"fmt"
	"jessie_miniproject/config"
	"jessie_miniproject/models"
	"log"
)

// GenerateDescriptions melakukan iterasi pada entri ProductLog dan menambahkan deskripsi yang dihasilkan AI.
func GenerateDescriptions(ctx context.Context, productLogs []models.ProductLog) error {
	for i := range productLogs {
		query := GenerateProductQuery(productLogs[i].ProductName, productLogs[i].Material, productLogs[i].IsPlastic)
		rekomendasi, err := ResponseAI(ctx, query)
		if err != nil {
			log.Printf("Terjadi kesalahan saat membuat rekomendasi AI untuk produk %s: %v", productLogs[i].ProductName, err)
			continue
		}

		productLogs[i].Rekomendasi = rekomendasi
		fmt.Printf("Buatkan deskripsi untuk data ini dan beri tahu mana yang plastik dan mana yang bukan. Tolong buat juga rekomendasi pengurangan sampah atau rekomendasi kesehatan dari data ini %s: %s\n", productLogs[i].ProductName, rekomendasi)

		// Simpan ke database (misalnya jika menggunakan GORM)
		if err := config.DB.Save(&productLogs[i]).Error; err != nil {
			log.Printf("Terjadi kesalahan saat menyimpan deskripsi untuk produk %s: %v", productLogs[i].ProductName, err)
		}
	}
	return nil
}
