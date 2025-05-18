package Features

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
)

func menu() {
	fmt.Println("\n---------------------------------------------")
	fmt.Println("                 Menu Utama")
	fmt.Println("---------------------------------------------")
	fmt.Println("1. Buat Profil")
	fmt.Println("2. Lihat Profil")
	fmt.Println("3. Edit Profil")
	fmt.Println("4. Hapus Profil")
	fmt.Println("5. Tambah Minat & Keterampilan")
	fmt.Println("6. Ubah Minat & Keterampilan")
	fmt.Println("7. Hapus Minat & Keterampilan")
	fmt.Println("8. Lihat Daftar Karier")
	fmt.Println("9. Cari Karier (Sequential Search)")
	fmt.Println("10. Cari Karier (Binary Search)")
	fmt.Println("11. Lihat Rekomendasi Karier")
	fmt.Println("0. Keluar")
	fmt.Println("---------------------------------------------")
}

func ToMenu() {
	reader := bufio.NewReader(os.Stdin)
    fmt.Println("Tekan Enter untuk melanjutkan...")
    _, _ = reader.ReadString('\n') 
    ShowAllMenu()
}

func ShowAllMenu() {
	menu()

	var input int
	fmt.Println("\nPilih Menu berdasarkan Angka pada Menunya.")
	fmt.Println("Masukkan pilihan: ")
	fmt.Scanln(&input)

	switch input {
	case 1:
		CreateProfile()
	case 2:
		ShowProfile()
	case 3:
		EditProfile()
	case 4:
		DeleteProfile()
	case 5:
		InsertInterestandSkill()
	case 6:
		EditInterestandSkill()
	case 7:
		DeleteInterestandSkill()
	case 8:
		Careers()
	case 9:
		SequentialSearchOptions()
	case 10:
		BinarySearchOptions()
	case 11:
		RecommendationCareers()
	case 0:
		os.Exit(0)
	default:
		fmt.Println("Input tidak valid, mohon lihat daftar menu yang ada")
	}
}
