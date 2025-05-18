package Features

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"TugasBesar/Helpers"
)

type Career struct {
	Id               int
	Name             string
	Interests        []string
	Skills           []string
	Desc             string
	Salary           int
	IndustryCategory string // Tambahkan kategori industri
}

var careers = []Career{
	{
		Id:               1,
		Name:             "Software Engineer",
		Interests:        []string{"Teknologi", "Komputer"},
		Skills:           []string{"Problem Solving", "Pemecahan Masalah", "Kreatif", "Analisis", "Logika"},
		Desc:             "Merancang, mengembangkan, dan memelihara perangkat lunak dan aplikasi",
		Salary:           8000000,
		IndustryCategory: "Teknologi",
	},
	{
		Id:               2,
		Name:             "Data Scientist",
		Interests:        []string{"Analisis Data", "Statistik", "Machine Learning"},
		Skills:           []string{"Matematika", "Programing", "Statistik", "Visualisasi Data"},
		Desc:             "Menganalisis dan menginterpretasi data kompleks untuk mendukung pengambilan keputusan",
		Salary:           10000000,
		IndustryCategory: "Teknologi",
	},
	{
		Id:               3,
		Name:             "UI/UX Designer",
		Interests:        []string{"Desain", "Pengalaman Pengguna", "Visual Arts"},
		Skills:           []string{"Kreatif", "Desain Grafis", "Empati", "Prototyping"},
		Desc:             "Merancang antarmuka dan pengalaman pengguna untuk produk digital",
		Salary:           7500000,
		IndustryCategory: "Kreatif",
	},
	{
		Id:               4,
		Name:             "Front-end Engineer",
		Interests:        []string{"Web Development", "User Interface", "Teknologi"},
		Skills:           []string{"JavaScript", "HTML", "CSS", "React", "Responsive Design"},
		Desc:             "Membangun antarmuka web yang interaktif dan responsif",
		Salary:           7800000,
		IndustryCategory: "Teknologi",
	},
	{
		Id:               5,
		Name:             "Back-end Engineer",
		Interests:        []string{"Arsitektur Sistem", "Database", "API"},
		Skills:           []string{"Programing", "Database Design", "API Development", "Server Management"},
		Desc:             "Membangun dan mengelola infrastruktur server dan database aplikasi",
		Salary:           8500000,
		IndustryCategory: "Teknologi",
	},
	{
		Id:               6,
		Name:             "Mobile App Developer",
		Interests:        []string{"Mobile Technology", "UX Design", "Teknologi"},
		Skills:           []string{"Android/iOS Development", "UI Design", "Cross-Platform Development"},
		Desc:             "Mengembangkan aplikasi untuk platform mobile seperti Android dan iOS",
		Salary:           8200000,
		IndustryCategory: "Teknologi",
	},
	{
		Id:               7,
		Name:             "Marketing Manager",
		Interests:        []string{"Pemasaran", "Strategi Bisnis", "Media Sosial"},
		Skills:           []string{"Komunikasi", "Analisis Pasar", "Strategi Kampanye", "Branding"},
		Desc:             "Mengembangkan dan mengelola strategi pemasaran untuk produk atau layanan",
		Salary:           9000000,
		IndustryCategory: "Pemasaran",
	},
	{
		Id:               8,
		Name:             "Financial Analyst",
		Interests:        []string{"Keuangan", "Investasi", "Ekonomi"},
		Skills:           []string{"Analisis Finansial", "Excel", "Pemodelan Keuangan", "Forecasting"},
		Desc:             "Menganalisis data keuangan dan membuat rekomendasi untuk keputusan bisnis",
		Salary:           8800000,
		IndustryCategory: "Keuangan",
	},
	{
		Id:               9,
		Name:             "Content Writer",
		Interests:        []string{"Menulis", "Storytelling", "Komunikasi"},
		Skills:           []string{"Penulisan Kreatif", "SEO", "Riset", "Copywriting"},
		Desc:             "Membuat konten tertulis untuk berbagai platform media",
		Salary:           6500000,
		IndustryCategory: "Kreatif",
	},
	{
		Id:               10,
		Name:             "HR Manager",
		Interests:        []string{"Sumber Daya Manusia", "Pengembangan Organisasi", "Training"},
		Skills:           []string{"Rekrutmen", "Employee Relations", "Manajemen Konflik", "Leadership"},
		Desc:             "Mengelola aspek sumber daya manusia dalam organisasi",
		Salary:           9500000,
		IndustryCategory: "Sumber Daya Manusia",
	},
	{
		Id:               11,
		Name:             "Project Manager",
		Interests:        []string{"Manajemen Proyek", "Kepemimpinan", "Organisasi"},
		Skills:           []string{"Perencanaan", "Komunikasi", "Manajemen Tim", "Problem Solving"},
		Desc:             "Merencanakan, mengeksekusi, dan menyelesaikan proyek sesuai deadline dan budget",
		Salary:           9200000,
		IndustryCategory: "Manajemen",
	},
	{
		Id:               12,
		Name:             "Graphic Designer",
		Interests:        []string{"Desain", "Seni Visual", "Kreativitas"},
		Skills:           []string{"Adobe Creative Suite", "Desain Visual", "Ilustrasi", "Typography"},
		Desc:             "Menciptakan elemen visual untuk media cetak dan digital",
		Salary:           7000000,
		IndustryCategory: "Kreatif",
	},
}


func Careers() {
	fmt.Println("\n---------------------------------------------")
	fmt.Println("              DAFTAR KARIER")
	fmt.Println("---------------------------------------------")

	for _, career := range careers {
		fmt.Println("\n---------------------------------------------")
		fmt.Printf("      %s\n", career.Name)
		fmt.Println("---------------------------------------------")
		fmt.Printf("Kategori Industri: %s\n", career.IndustryCategory)
		fmt.Printf("Gaji: %s\n", Helpers.FormatSalary(career.Salary))
		fmt.Printf("Deskripsi: %s\n", career.Desc)

		fmt.Println("\nMinat:")
		for i, interest := range career.Interests {
			fmt.Printf("  %d. %s\n", i+1, strings.TrimSpace(interest))
		}

		fmt.Println("\nKeterampilan:")
		for i, talent := range career.Skills {
			fmt.Printf("  %d. %s\n", i+1, strings.TrimSpace(talent))
		}
		fmt.Println("---------------------------------------------")
	}
	ToMenu()
}

func SequentialSearchOptions() {
	fmt.Println("\n---------------------------------------------")
	fmt.Println("           PILIH JENIS PENCARIAN")
	fmt.Println("---------------------------------------------")
	fmt.Println("1. Berdasarkan Nama")
	fmt.Println("2. Berdasarkan Kategori Industri")
	fmt.Println("0. Exit")

	var ByData string
	fmt.Println("Masukan pilihan (angka): ")
	fmt.Scanln(&ByData)

	if ByData == "1" {
		SequantialSearchByName()
	} else if ByData == "2" {
		SequentialSearchByIndustry()
	}
}

func BinarySearchOptions() {
	fmt.Println("\n---------------------------------------------")
	fmt.Println("           PILIH JENIS PENCARIAN")
	fmt.Println("---------------------------------------------")
	fmt.Println("1. Berdasarkan Nama")
	fmt.Println("2. Berdasarkan Kategori Industri")
	fmt.Println("0. Exit")

	var ByData string
	fmt.Println("Masukan pilihan (angka): ")
	fmt.Scanln(&ByData)

	if ByData == "1" {
		BinarySearchByName()
	} else if ByData == "2" {
		BinarySearchByIndustry()
	}
}

func RecommendationCareers() {
	fmt.Println("\n---------------------------------------------------")
	fmt.Println("REKOMENDASI KARIER BERDASARKAN KECOCOKAN ATAU GAJI")
	fmt.Println("---------------------------------------------------")

	if profile == nil {
		fmt.Println("Belum ada data profile.")
		ToMenu()
		return
	} else if profile.Interests != nil && profile.Skills != nil {
		fmt.Println("Minat:", profile.Interests)
		fmt.Println("Keterampilan:", profile.Skills)
	} else if profile.Interests == nil && profile.Skills == nil {
		fmt.Println("Belum ada data minat dan keterampilan.")
		ToMenu()
		return
	}

	recommendations := GenerateCareerRecommendations()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nPilih metode sorting:")
	fmt.Println("1. Selection Sort berdasarkan kecocokan")
	fmt.Println("2. Selection Sort berdasarkan gaji")
	fmt.Println("3. Insertion Sort berdasarkan kecocokan")
	fmt.Println("4. Insertion Sort berdasarkan gaji")
	fmt.Print("Masukkan pilihan (1-4): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fmt.Print("\nUrutan ascending (true/false): ")
	orderInput, _ := reader.ReadString('\n')
	ascending := strings.TrimSpace(strings.ToLower(orderInput)) == "true"

	switch input {
	case "1":
		SelectionSort(recommendations, false, ascending)
	case "2":
		SelectionSort(recommendations, true, ascending)
	case "3":
		InsertionSort(recommendations, false, ascending)
	case "4":
		InsertionSort(recommendations, true, ascending)
	default:
		fmt.Println("Pilihan tidak valid. Menampilkan default (tanpa sort).")
	}

	fmt.Println("\nDaftar Rekomendasi Karier:")
	DisplayRecommendations(recommendations)
	fmt.Println("---------------------------------------------------")

	ToMenu()
}
