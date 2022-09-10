package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	var average float64

	// TODO ambil masing masing score pada struct Student
	for _, value := range s.score {
		// TODO ubah value ke bentuk float dan tambahkan setiap skor dengan nilai average
		average += float64(value)
	}

	// TODO bagi total nilai dengan banyak student
	average /= float64(len(s.score))

	return average
}

func (s Student) Min() (min int, name string) {
	// TODO simpan nama dan nilai terkecil awal dengan nilai pada index pertama
	min = s.score[0]
	name = s.name[0]

	// TODO iterasi field score untuk pengecekan
	for i, value := range s.score {
		// TODO cek apabila value kurang dari min, maka ganti nilai min dengan value dan simpan nama sutdent tersebut
		if value < min {
			min = value
			name = s.name[i]
		}
	}

	return
}

func (s Student) Max() (max int, name string) {
	// TODO simpan nama dan nilai terbesar awal dengan nilai pada index pertama
	max = s.score[0]
	name = s.name[0]

	// TODO iterasi field score untuk pengecekan
	for i, value := range s.score {
		// TODO cek apabila value lebih besar dari max, maka tukar nilai min dengan value dan simpan nama student tersebut
		if value > max {
			max = value
			name = s.name[i]
		}
	}

	return
}

func main() {
	var a = Student{}

	for i := 1; i < 6; i++ {
		var name string
		fmt.Print("Input " + string(rune(i)) + "Student's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Student is", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Student is "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Student is "+nameMin+" (", scoreMin, ")")
}
