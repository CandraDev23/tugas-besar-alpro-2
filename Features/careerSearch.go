package Features

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"TugasBesar/Helpers"
)

func ProcessSequantialSearchByName(nama string) []Career {
	nama = strings.TrimSpace(nama)
	fmt.Println("\nMencari:", nama)

	var hasilPencarian []Career

	for i := 0; i < len(careers); i++ {
		if strings.EqualFold(careers[i].Name, nama) {
			hasilPencarian = append(hasilPencarian, careers[i])
			continue
		}

		if strings.Contains(strings.ToLower(careers[i].Name), strings.ToLower(nama)) {
			hasilPencarian = append(hasilPencarian, careers[i])
		}
	}

	return hasilPencarian
}

func SequantialSearchByName() {
	fmt.Println("\n---------------------------------------------")
	fmt.Println("  CARI KARIER BERDASARKAN NAMA PEKERJAAN")
	fmt.Println("   	  SEQUENTIAL SEARCH")
	fmt.Println("---------------------------------------------")

	fmt.Println("Cari Karier (Sequential Search)")
	fmt.Print("Masukan nama karir/pekerjaan: ")
	var nama string

	scanner := bufio.NewReader(os.Stdin)
	nama, _ = scanner.ReadString('\n')
	nama = strings.TrimSpace(nama)

	hasilPencarian := ProcessSequantialSearchByName(nama)

	if len(hasilPencarian) > 0 {
		fmt.Printf("Ditemukan %d hasil pencarian:\n", len(hasilPencarian))

		for i, hasil := range hasilPencarian {
			fmt.Printf("\n--- Hasil #%d ---\n", i+1)
			fmt.Println("Nama : ", hasil.Name)
			fmt.Println("Deskripsi : ", hasil.Desc)
			fmt.Println("Gaji : ", Helpers.FormatSalary(hasil.Salary))

			fmt.Print("Minat : ")
			for j, interest := range hasil.Interests {
				if j > 0 {
					fmt.Print(", ")
				}
				fmt.Print(interest)
			}
			fmt.Println()

			fmt.Print("Keterampilan : ")
			for j, skill := range hasil.Skills {
				if j > 0 {
					fmt.Print(", ")
				}
				fmt.Print(skill)
			}
			fmt.Println()
		}
	} else {
		fmt.Println("\nKarir tidak ditemukan")
		fmt.Println("Daftar karir yang tersedia:")
		for i, career := range careers {
			fmt.Printf("%d. %s\n", i+1, career.Name)
		}
	}
	fmt.Println("---------------------------------------------")
	ToMenu()
}

// sortName urutkan karir berdasarkan nama
func sortName() {
	// Menggunakan bubble sort
	for i := 0; i < len(careers)-1; i++ {
		for j := 0; j < len(careers)-i-1; j++ {
			if careers[j].Name > careers[j+1].Name {
				careers[j], careers[j+1] = careers[j+1], careers[j]
			}
		}
	}

	// Debug: print sorted careers
	fmt.Println("\nDaftar karir setelah diurutkan:")
	for i, career := range careers {
		fmt.Printf("%d. %s\n", i+1, career.Name)
	}
}

// BinarySearchName mencari karier berdasarkan nama dengan binary search
// CariKarierBinary mencari karier dengan binary search dan menampilkan informasi
func BinarySearchByName() {
	fmt.Println("\n---------------------------------------------")
	fmt.Println("   CARI KARIER BERDASARKAN NAMA PEKERJAAN")
	fmt.Println("   		BINARY SEARCH")
	fmt.Println("---------------------------------------------")
	fmt.Println("Cari Karier (Binary Search)")
	fmt.Print("Masukan nama karir : ")
	var nama string

	// Menangani input dengan lebih baik
	scanner := bufio.NewReader(os.Stdin)
	nama, _ = scanner.ReadString('\n')
	nama = strings.TrimSpace(nama)

	hasilPencarian := ProcessBinarySearchNameMultiple(nama)

	if len(hasilPencarian) > 0 {
		fmt.Printf("Ditemukan %d hasil pencarian:\n", len(hasilPencarian))

		for i, hasil := range hasilPencarian {
			fmt.Printf("\n--- Hasil #%d ---\n", i+1)
			fmt.Println("Nama : ", hasil.Name)
			fmt.Println("Deskripsi : ", hasil.Desc)
			fmt.Println("Gaji : ", Helpers.FormatSalary(hasil.Salary))

			// Menampilkan interests
			fmt.Print("Minat : ")
			for j, interest := range hasil.Interests {
				if j > 0 {
					fmt.Print(", ")
				}
				fmt.Print(interest)
			}
			fmt.Println()

			// Menampilkan skills
			fmt.Print("Keterampilan : ")
			for j, skill := range hasil.Skills {
				if j > 0 {
					fmt.Print(", ")
				}
				fmt.Print(skill)
			}
			fmt.Println()
		}
	} else {
		fmt.Println("\nKarir tidak ditemukan")
		fmt.Println("Daftar karir yang tersedia:")
		for i, career := range careers {
			fmt.Printf("%d. %s\n", i+1, career.Name)
		}
	}
	fmt.Println("---------------------------------------------")

	ToMenu()
}

// BinarySearchNameMultiple mencari semua karier yang sesuai dengan nama
func ProcessBinarySearchNameMultiple(nama string) []Career {
	// Binary search bekerja untuk exact match
	exactMatch := ProcessBinarySearchByName(nama)

	// Untuk multiple matches, kita tetap perlu linear search
	// karena binary search hanya efektif untuk exact match
	var hasilPencarian []Career

	// Jika ada exact match, tambahkan ke hasil
	if exactMatch != nil {
		hasilPencarian = append(hasilPencarian, *exactMatch)
	}

	// Tambahkan partial matches dengan linear search
	nama = strings.ToLower(strings.TrimSpace(nama))
	for i := 0; i < len(careers); i++ {
		currentName := strings.ToLower(careers[i].Name)

		// Skip jika sudah ada di hasil (exact match)
		if exactMatch != nil && currentName == strings.ToLower(exactMatch.Name) {
			continue
		}

		// Cek apakah nama mengandung keyword
		if strings.Contains(currentName, nama) {
			hasilPencarian = append(hasilPencarian, careers[i])
		}
	}

	return hasilPencarian
}

// BinarySearchName mencari karier berdasarkan nama dengan binary search
func ProcessBinarySearchByName(nama string) *Career {
	// Urutkan dahulu
	sortName()

	awal := 0
	akhir := len(careers) - 1
	nama = strings.ToLower(strings.TrimSpace(nama))

	fmt.Println("\nMencari dengan binary search:", nama)

	for awal <= akhir {
		tengah := (awal + akhir) / 2
		mid_name := strings.ToLower(careers[tengah].Name)

		if mid_name == nama {
			return &careers[tengah]
		} else if mid_name < nama {
			awal = tengah + 1
		} else {
			akhir = tengah - 1
		}
	}

	// Binary search hanya menemukan exact match, jadi cek partial match setelahnya
	return nil
}

// CariKarierByIndustry mencari karier berdasarkan kategori industri (sequential search)
func SequentialSearchByIndustry() {
	fmt.Println("\n---------------------------------------------")
	fmt.Println(" CARI KARIER BERDASARKAN KATEGORI INDUSTRI")
	fmt.Println("   	   SEQUENTIAL SEARCH")
	fmt.Println("---------------------------------------------")
	fmt.Println("Cari Karier Berdasarkan Kategori Industri (Sequential Search)")
	fmt.Print("Masukan kategori industri : ")
	var category string

	// Menangani input dengan lebih baik
	scanner := bufio.NewReader(os.Stdin)
	category, _ = scanner.ReadString('\n')
	category = strings.TrimSpace(category)

	hasilPencarian := ProcessSequentialSearchByIndustry(category)

	if len(hasilPencarian) > 0 {
		fmt.Printf("Ditemukan %d hasil pencarian:\n", len(hasilPencarian))

		for i, hasil := range hasilPencarian {
			fmt.Printf("\n--- Hasil #%d ---\n", i+1)
			fmt.Println("Nama : ", hasil.Name)
			fmt.Println("Kategori Industri : ", hasil.IndustryCategory)
			fmt.Println("Deskripsi : ", hasil.Desc)
			fmt.Println("Gaji : ", Helpers.FormatSalary(hasil.Salary))

			// Menampilkan interests
			fmt.Print("Interests : ")
			for j, interest := range hasil.Interests {
				if j > 0 {
					fmt.Print(", ")
				}
				fmt.Print(interest)
			}
			fmt.Println()

			// Menampilkan skills
			fmt.Print("Skills : ")
			for j, skill := range hasil.Skills {
				if j > 0 {
					fmt.Print(", ")
				}
				fmt.Print(skill)
			}
			fmt.Println()
		}
	} else {
		fmt.Println("\nKarir dengan kategori tersebut tidak ditemukan")
		fmt.Println("Daftar kategori industri yang tersedia:")

		// Membuat map untuk menghindari duplikasi kategori
		industryMap := make(map[string]bool)
		for _, career := range careers {
			industryMap[career.IndustryCategory] = true
		}

		// Menampilkan daftar kategori yang unik
		i := 1
		for industry := range industryMap {
			fmt.Printf("%d. %s\n", i, industry)
			i++
		}
	}
	fmt.Println("---------------------------------------------")

	ToMenu()
}

// SearchCareerByIndustry mencari karier berdasarkan kategori industri (sequential search)
func ProcessSequentialSearchByIndustry(category string) []Career {
	// Trim input untuk menghilangkan whitespace di awal dan akhir
	category = strings.TrimSpace(category)
	fmt.Println("\nMencari:", category)

	var hasilPencarian []Career

	for i := 0; i < len(careers); i++ {

		// Metode 1: Exact match case insensitive
		if strings.EqualFold(careers[i].IndustryCategory, category) {
			hasilPencarian = append(hasilPencarian, careers[i])
			continue // Lanjut ke iterasi berikutnya setelah menambahkan ke hasil
		}

		// Metode 2: Contains check (untuk pencarian parsial)
		if strings.Contains(
			strings.ToLower(careers[i].IndustryCategory),
			strings.ToLower(category)) {
			hasilPencarian = append(hasilPencarian, careers[i])
		}
	}

	return hasilPencarian
}

// sortIndustry urutkan karir berdasarkan kategori industri
func sortIndustry() {
	// Menggunakan bubble sort
	for i := 0; i < len(careers)-1; i++ {
		for j := 0; j < len(careers)-i-1; j++ {
			if careers[j].IndustryCategory > careers[j+1].IndustryCategory {
				careers[j], careers[j+1] = careers[j+1], careers[j]
			}
		}
	}

	// Debug: print sorted careers by industry
	fmt.Println("\nDaftar karir setelah diurutkan berdasarkan industri:")
	for i, career := range careers {
		fmt.Printf("%d. %s (Industri: %s)\n", i+1, career.Name, career.IndustryCategory)
	}
}

// CariKarierByIndustryBinary mencari karier berdasarkan kategori industri (binary search)
func BinarySearchByIndustry() {
	fmt.Println("\n---------------------------------------------")
	fmt.Println("  CARI KARIER BERDASARKAN KATEGORI INDUSTRI")
	fmt.Println("   		BINARY SEARCH")
	fmt.Println("---------------------------------------------")
	fmt.Println("Cari Karier Berdasarkan Kategori Industri (Binary Search)")
	fmt.Print("Masukan kategori industri : ")
	var category string

	// Menangani input dengan lebih baik
	scanner := bufio.NewReader(os.Stdin)
	category, _ = scanner.ReadString('\n')
	category = strings.TrimSpace(category)

	hasilPencarian := ProcessBinarySearchByIndustryMultiple(category)

	if len(hasilPencarian) > 0 {
		fmt.Printf("Ditemukan %d hasil pencarian:", len(hasilPencarian))

		for i, hasil := range hasilPencarian {
			fmt.Printf("\n--- Hasil #%d ---\n", i+1)
			fmt.Println("Nama : ", hasil.Name)
			fmt.Println("Kategori Industri : ", hasil.IndustryCategory)
			fmt.Println("Deskripsi : ", hasil.Desc)
			fmt.Println("Gaji : ", Helpers.FormatSalary(hasil.Salary))

			// Menampilkan interests
			fmt.Print("Minat : ")
			for j, interest := range hasil.Interests {
				if j > 0 {
					fmt.Print(", ")
				}
				fmt.Print(interest)
			}
			fmt.Println()

			// Menampilkan skills
			fmt.Print("Keterampilan : ")
			for j, skill := range hasil.Skills {
				if j > 0 {
					fmt.Print(", ")
				}
				fmt.Print(skill)
			}
			fmt.Println()
		}
	} else {
		fmt.Println("\nKarir dengan kategori tersebut tidak ditemukan")
		fmt.Println("Daftar kategori industri yang tersedia:")

		// Membuat map untuk menghindari duplikasi kategori
		industryMap := make(map[string]bool)
		for _, career := range careers {
			industryMap[career.IndustryCategory] = true
		}

		// Menampilkan daftar kategori yang unik
		i := 1
		for industry := range industryMap {
			fmt.Printf("%d. %s\n", i, industry)
			i++
		}
	}
	fmt.Println("---------------------------------------------")

	ToMenu()
}

// BinarySearchIndustry mencari karier berdasarkan kategori industri dengan binary search
func ProcessBinarySearchByIndustry(category string) *Career {
	// Urutkan dahulu
	sortIndustry()

	awal := 0
	akhir := len(careers) - 1
	category = strings.ToLower(strings.TrimSpace(category))

	fmt.Println("\nMencari :", category)

	for awal <= akhir {
		tengah := (awal + akhir) / 2
		mid_category := strings.ToLower(careers[tengah].IndustryCategory)

		if mid_category == category {
			return &careers[tengah]
		} else if mid_category < category {
			awal = tengah + 1
		} else {
			akhir = tengah - 1
		}
	}

	return nil
}

// BinarySearchIndustryMultiple mencari semua karier berdasarkan kategori industri
func ProcessBinarySearchByIndustryMultiple(category string) []Career {
	// Binary search bekerja untuk exact match
	exactMatch := ProcessBinarySearchByIndustry(category)

	// Untuk multiple matches, kita tetap perlu linear search
	// karena binary search hanya efektif untuk exact match
	var hasilPencarian []Career

	// Jika ada exact match, tambahkan ke hasil dan cari yang lain dengan kategori sama
	if exactMatch != nil {
		// Mencari semua dengan kategori sama secara linear
		category = strings.ToLower(strings.TrimSpace(category))
		for i := 0; i < len(careers); i++ {
			currentCategory := strings.ToLower(careers[i].IndustryCategory)

			// Cek exact match kategori
			if currentCategory == category {
				hasilPencarian = append(hasilPencarian, careers[i])
			}
		}
	} else {
		// Jika tidak ada exact match, cari partial match
		category = strings.ToLower(strings.TrimSpace(category))
		for i := 0; i < len(careers); i++ {
			currentCategory := strings.ToLower(careers[i].IndustryCategory)

			// Cek apakah kategori mengandung keyword
			if strings.Contains(currentCategory, category) {
				hasilPencarian = append(hasilPencarian, careers[i])
			}
		}
	}

	return hasilPencarian
}