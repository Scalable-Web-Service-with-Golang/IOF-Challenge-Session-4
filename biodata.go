package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	Name              string
	Address           string
	Profession        string
	ReasonForChoosing string
}

var students = []Student{
	{
		Name:              "Muhammad Jauhari",
		Address:           "Jl. Sudirman, No.234",
		Profession:        "Backend Engineer",
		ReasonForChoosing: "Membuat RestAPI dengan Golang.",
	},
	{
		Name:              "Robbi Nugroho Purnomo",
		Address:           "Jl. Sukabangun 2, No.420",
		Profession:        "Sales Marketing",
		ReasonForChoosing: "Belajar Go Programming dari Basic.",
	},
	{
		Name:              "Muhammad Ihsan Mudhlari",
		Address:           "Jl. Sukawinatan, No.69",
		Profession:        "IT Support",
		ReasonForChoosing: "Menambah wawasan di dunia Programming.",
	},
	{
		Name:              "Muhammad Angga Oktaharisetia",
		Address:           "Jl. Basuki Rahmat, No.77",
		Profession:        "Software Engineer",
		ReasonForChoosing: "Membangun Microservices dengan Golang.",
	},
	{
		Name:              "Dinar Iswahyudi",
		Address:           "Jl. Veteran, No.17",
		Profession:        "College Student",
		ReasonForChoosing: "Mengenal fundamental Go Programming.",
	},
	{
		Name:              "Jaka Purwa Baruna",
		Address:           "Jl. Kavling 1, No.66",
		Profession:        "Sales Marketing",
		ReasonForChoosing: "Menambah wawasan di dunia Programming.",
	},
	{
		Name:              "M. Fajar Wahyudi",
		Address:           "Jl. Kavling 2, No.100",
		Profession:        "Sales Marketing",
		ReasonForChoosing: "Belajar Go Programming dari Basic.",
	},
	{
		Name:              "Ardika Agus Wijaya",
		Address:           "Jl. Mataram 1, No.007",
		Profession:        "Frontline Teller",
		ReasonForChoosing: "Membangun Microservices dengan Golang.",
	},
	{
		Name:              "Satria Adi Nugraha",
		Address:           "Jl. Talang Kelapa, No.456",
		Profession:        "Debt Payment Collector",
		ReasonForChoosing: "Mengenal fundamental Go Programming.",
	},
}

func findBiodata(absent int) Student {
	return students[absent-1]
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Run command: go run biodata.go [absent-number]")
		return
	}

	studentAbsent := args[0]
	convAbsent, err := strconv.Atoi(studentAbsent)
	if err != nil {
		fmt.Println("Error to convert absent number.")
		return
	}

	if convAbsent == 0 {
		fmt.Println("Invalid absent number.")
		return
	}

	studentBiodata := findBiodata(convAbsent)

	fmt.Printf("📌 Biodata dengan No.%s :\n", studentAbsent)
	fmt.Printf("Nama      : %s\n", studentBiodata.Name)
	fmt.Printf("Alamat 	  : %s\n", studentBiodata.Address)
	fmt.Printf("Pekerjaan : %s\n", studentBiodata.Profession)
	fmt.Printf("Alasan memilih kelas Golang : %s\n", studentBiodata.ReasonForChoosing)
}
