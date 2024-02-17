package session_three

import "fmt"

type Student struct {
	Id        string
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func MiniChallengeThree(userInput string) {
	students := []Student{
		{
			Id:        "0",
			Nama:      "Afif",
			Alamat:    "Bandung",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Mendalami Golang",
		},
		{
			Id:        "1",
			Nama:      "Larasati",
			Alamat:    "Jakarta",
			Pekerjaan: "Fullstack Developer",
			Alasan:    "Bosen dengan pekerjaan lama",
		},
	}

	var foundStudents []Student

	for _, student := range students {
		if student.Nama == userInput || student.Id == userInput {
			foundStudents = append(foundStudents, student)
		}
	}

	if len(foundStudents) > 0 {
		for _, data := range foundStudents {
			fmt.Printf("ID: %s\n", data.Id)
			fmt.Printf("Nama: %s\n", data.Nama)
			fmt.Printf("Alamat: %s\n", data.Alamat)
			fmt.Printf("Pekerjaan: %s\n", data.Pekerjaan)
			fmt.Printf("Alasan: %s\n", data.Alasan)
		}
	} else {
		fmt.Println("Data dengan nama/absen tsb tidak tersedia")
	}
}
