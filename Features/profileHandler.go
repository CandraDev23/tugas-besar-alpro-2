package Features

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Profile struct {
	Id        int
	Name      string
	Interests []string
	Skills    []string
}

var profile *Profile

func CreateProfile() {
	if profile != nil {
		fmt.Println("Data profile sudah ada.")
		ToMenu()
		return
	}

	reader := bufio.NewReader(os.Stdin)

	tempProfile := Profile{}
	tempProfile.Id = 1

	fmt.Println("\n---------------------------------------------")
	fmt.Println("                 Buat Profil")
	fmt.Println("---------------------------------------------")
	fmt.Println("Buat profil kamu dengan memasukkan data dibawah ini.")
	fmt.Print("Masukkan Nama: ")
	os.Stdout.Sync()
	nameInput, _ := reader.ReadString('\n')
	tempProfile.Name = strings.TrimSpace(nameInput)

	fmt.Print("Masukkan Minat (pisahkan dengan koma): ")
	os.Stdout.Sync()
	interestInput, _ := reader.ReadString('\n')
	interestInput = strings.TrimSpace(interestInput)
	if interestInput != "" {
		tempProfile.Interests = strings.Split(interestInput, ",")
	}

	fmt.Print("Masukkan Keterampilan (pisahkan dengan koma): ")
	os.Stdout.Sync()
	skillInput, _ := reader.ReadString('\n')
	skillInput = strings.TrimSpace(skillInput)
	if skillInput != "" {
		tempProfile.Skills = strings.Split(skillInput, ",")
	}
	profile = &tempProfile

	fmt.Println("\nData Profile berhasil ditambahkan:")
	fmt.Printf("ID: %d\n", profile.Id)
	fmt.Printf("Nama: %s\n", profile.Name)
	fmt.Println("Minat:")
	for i, interest := range profile.Interests {
		fmt.Printf("  %d. %s\n", i+1, strings.TrimSpace(interest))
	}
	fmt.Println("Keterampilan:")
	for i, skill := range profile.Skills {
		fmt.Printf("  %d. %s\n", i+1, strings.TrimSpace(skill))
	}
	fmt.Println("---------------------------------------------")
	ToMenu()
}

func ShowProfile() {
	if profile == nil {
		fmt.Println("Data profile belum ada.")
		ToMenu()
		return
	}

	fmt.Println("\n---------------------------------------------")
	fmt.Println("                 Lihat Profil")
	fmt.Println("---------------------------------------------")
	fmt.Printf("Nama: %s\n", profile.Name)

	fmt.Println("Minat:")
	for i, interest := range profile.Interests {
		fmt.Printf("  %d. %s\n", i+1, strings.TrimSpace(interest))
	}

	fmt.Println("Keterampilan:")
	for i, skill := range profile.Skills {
		fmt.Printf("  %d. %s\n", i+1, strings.TrimSpace(skill))
	}
	fmt.Println("---------------------------------------------")

	ToMenu()
}

func EditProfile() {
	fmt.Println("\n---------------------------------------------")
	fmt.Println("                 Edit Profil")
	fmt.Println("---------------------------------------------")
	fmt.Println("Data saat ini.")
	if profile != nil {
		fmt.Println("Nama:", profile.Name)
		fmt.Println("Minat:", profile.Interests)
		fmt.Println("Keterampilan:", profile.Skills)
	} else if profile == nil {
		fmt.Println("Belum ada data profile.")
		ToMenu()
		return
	}
	fmt.Println("---------------------------------------------")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nama:")
	newName, _ := reader.ReadString('\n')
	newName = strings.TrimSpace(newName)
	if newName != "" {
		profile.Name = newName
	}

	fmt.Print("Ubah minat (pisahkan dengan koma): ")
	inputInterest, _ := reader.ReadString('\n')
	inputInterest = strings.TrimSpace(inputInterest)
	if inputInterest != "" {
		profile.Interests = strings.Split(inputInterest, ",")
	}

	fmt.Print("Ubah keterampilan (pisahkan dengan koma): ")
	inputSkill, _ := reader.ReadString('\n')
	inputSkill = strings.TrimSpace(inputSkill)
	if inputSkill != "" {
		profile.Skills = strings.Split(inputSkill, ",")
	}
	fmt.Println("Profil berhasil diperbarui.")
	fmt.Println("---------------------------------------------")

	ToMenu()
}

func DeleteProfile() {
	fmt.Println("\n---------------------------------------------")
	fmt.Println("                 Hapus Profil")
	fmt.Println("---------------------------------------------")
	if profile != nil {
		fmt.Println("Data saat ini.")
		fmt.Println("Name:", profile.Name)
		fmt.Println("Minat:", profile.Interests)
		fmt.Println("Keterampilan:", profile.Skills)
	} else if profile == nil {
		fmt.Println("Belum ada data profile.")
		ToMenu()
		return
	}

	fmt.Println("\nApakah kamu yakin ingin menghapus data profil? (y/n)")
	var input string
	fmt.Scanln(&input)
	if input == "y" {
		profile = nil
		fmt.Println("Data profile berhasil dihapus.")
		fmt.Println("---------------------------------------------")
		ToMenu()
		return
	} else {
		fmt.Println("Data profile tidak jadi dihapus.")
		fmt.Println("---------------------------------------------")
		ToMenu()
		return
	}
}

func InsertInterestandSkill() {
	fmt.Println("\n---------------------------------------------")
	fmt.Println("           Tambah Minat dan Keterampilan")
	fmt.Println("---------------------------------------------")
	fmt.Println("Data saat ini.")
	if profile != nil {
		fmt.Println("Minat:", profile.Interests)
		fmt.Println("Keterampilan:", profile.Skills)
	} else if profile == nil {
		fmt.Println("Belum ada data profile.")
		ToMenu()
		return
	}
	fmt.Println("---------------------------------------------")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Tambah minat (pisahkan dengan koma): ")
	newInterest, _ := reader.ReadString('\n')
	newInterest = strings.TrimSpace(newInterest)
	if newInterest != "" {
		profile.Interests = append(profile.Interests, strings.Split(newInterest, ",")...)
	}

	fmt.Print("Tambah keterampilan (pisahkan dengan koma): ")
	newSkill, _ := reader.ReadString('\n')
	newSkill = strings.TrimSpace(newSkill)
	if newSkill != "" {
		profile.Skills = append(profile.Skills, strings.Split(newSkill, ",")...)
	}
	fmt.Println("Data minat dan keterampilan berhasil ditambahkan.")
	fmt.Println("---------------------------------------------")

	ToMenu()
}

func EditInterestandSkill() {
	fmt.Println("\n---------------------------------------------")
	fmt.Println("           Ubah Minat dan Keterampilan")
	fmt.Println("---------------------------------------------")
	fmt.Println("Data saat ini.")
	if profile != nil {
		fmt.Println("Minat:", profile.Interests)
		fmt.Println("Keterampilan:", profile.Skills)
	} else {
		fmt.Println("Belum ada data profile.")
		ToMenu()
		return
	}
	fmt.Println("---------------------------------------------")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ubah minat (pisahkan dengan koma): ")
	inputInterest, _ := reader.ReadString('\n')
	inputInterest = strings.TrimSpace(inputInterest)
	if inputInterest != "" {
		profile.Interests = strings.Split(inputInterest, ",")
	}

	fmt.Print("Ubah keterampilan (pisahkan dengan koma): ")
	inputSkill, _ := reader.ReadString('\n')
	inputSkill = strings.TrimSpace(inputSkill)
	if inputSkill != "" {
		profile.Skills = strings.Split(inputSkill, ",")
	}

	fmt.Println("Data minat dan keterampilan berhasil diperbarui.")
	fmt.Println("---------------------------------------------")
	ToMenu()
}

func DeleteInterestandSkill() {
	fmt.Println("\n---------------------------------------------")
	fmt.Println("           Hapus Minat dan Bakat")
	fmt.Println("---------------------------------------------")
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

	fmt.Println("---------------------------------------------")

	fmt.Println("\nApakah anda yakin ingin menghapus data minat? (y/n)")
	var deleteInterestConfirm string
	fmt.Scanln(&deleteInterestConfirm)
	if deleteInterestConfirm == "y" {
		profile.Interests = nil
		fmt.Println("Data minat berhasil dihapus.")
	}
	fmt.Println("\nApakah anda yakin ingin menghapus data keterampilan? (y/n)")
	var deleteSkillConfirm string
	fmt.Scanln(&deleteSkillConfirm)
	if deleteSkillConfirm == "y" {
		profile.Skills = nil
		fmt.Println("Data keterampilan berhasil dihapus.")
	}
	fmt.Println("---------------------------------------------")
	ToMenu()
}
