package main

import (
	"fmt"
)

// slices nó cũng same same của Python nên em hiểu rồi ah
type sinhvien struct {
	code    int
	name    string
	gender  string
	age     int
	math    float32
	physics float32
	chemis  float32
}
type lop struct {
	classname string
	lopsv     []sinhvien
}

func main() {
	for true {
		listsv := map[string]sinhvien{}
		listlop := map[string]lop{}
		fmt.Println("He thong quan li sinh vien")
		fmt.Println("CHUC NANG:")
		fmt.Println("1. Them thong tin sinh vien.")
		fmt.Println("2. Dua ra thong tin sinh vien.")
		fmt.Println("3. Xoa thong tin sinh vien.")
		fmt.Println("4. Chinh sua thong tin sinh vien.")
		fmt.Println("5. Them lop hoc.")
		fmt.Println("6. Xem danh sach lop hoc.")
		fmt.Println("7. Them sinh vien vao lop.")
		fmt.Println("8. Xem danh sach sinh vien trong lop.")
		fmt.Println("9. Xoa sinh vien khoi lop hoc.")
		fmt.Println("10. Tra cuu hoc sinh.")
		fmt.Println("0. Thoat.")
		var input int
		fmt.Println("Hay chon chuc nang: ")
		fmt.Scanln(&input)
		if input == 1 {
			var sv sinhvien
			fmt.Println("Ma SV: ")
			fmt.Scanln(&sv.code)
			fmt.Println("Ten: ")
			fmt.Scanln(&sv.name)
			fmt.Println("Gioi tinh: ")
			fmt.Scanln(&sv.gender)
			fmt.Println("Tuoi: ")
			fmt.Scanln(&sv.age)
			fmt.Println("Toan: ")
			fmt.Scanln(&sv.math)
			fmt.Println("Ly ")
			fmt.Scanln(&sv.physics)
			fmt.Println("Hoa: ")
			fmt.Scanln(&sv.chemis)
			listsv[sv.name] = sv
		}
		if input == 2 {
			for k, v := range listsv {
				fmt.Println("Ma SV: ", v.code)
				fmt.Println("Ten: ", k)
				fmt.Println("Gioi tinh: ", v.gender)
				fmt.Println("Tuoi: ", v.age)
				fmt.Println("Diem toan: ", v.math)
				fmt.Println("Diem ly: ", v.physics)
				fmt.Println("Diem hoa: ", v.chemis)
			}

		}
		if input == 3 {
			var ten string
			fmt.Println("Ten hoc sinh ban muon xoa khoi danh sach: ")
			fmt.Scanln(&ten)
			delete(listsv, ten)

		}
		if input == 4 {
			var ten string
			var sv sinhvien
			fmt.Println("Ten SV ban muon chinh sua: ")
			fmt.Scanln(&ten)
			delete(listsv, ten)
			fmt.Println("Nhap lai thong tin SV")
			fmt.Println("Ma SV: ")
			fmt.Scanln(&sv.code)
			fmt.Println("Ten: ")
			fmt.Scanln(&sv.name)
			fmt.Println("Gioi tinh: ")
			fmt.Scanln(&sv.gender)
			fmt.Println("Tuoi: ")
			fmt.Scanln(&sv.age)
			fmt.Println("Toan: ")
			fmt.Scanln(&sv.math)
			fmt.Println("Ly ")
			fmt.Scanln(&sv.physics)
			fmt.Println("Hoa: ")
			fmt.Scanln(&sv.chemis)
			listsv[sv.name] = sv

		}
		if input == 5 {
			var class lop
			var ten string
			fmt.Println("Ten lop ban muon them: ")
			fmt.Scanln(&class.classname)
			listlop[class.classname] = class
			fmt.Println("Ten hoc sinh ban muon them: ")
			fmt.Scanln(&ten)
			var sv sinhvien
			sv.name = ten
			class.lopsv = append(class.lopsv, sv)
		}
		if input == 6 {
			for k := range listlop {
				fmt.Println("Cac lop co tren he thong:")
				fmt.Println(k)
			}
		}
		if input == 7 {
			var tenlop string
			fmt.Println("Nhap ten lop ban muon them hoc sinh: ")
			fmt.Scanln(&tenlop)
			value0, ok := listlop[tenlop]
			if ok {
				var ten string
				fmt.Println("Nhap ten hoc sinh ban muon them: ")
				fmt.Scanln(&ten)
				value1, ok := listsv[ten]
				if ok {
					value0.lopsv = append(value0.lopsv, value1)
				} else {
					fmt.Println("Hoc sinh khong co tren he thong.")
				}
			} else {
				fmt.Println("Lop khong co tren he thong.")
			}

		}
		if input == 8 {
			var tenlop string
			fmt.Println("Nhap ten lop ban muon xem danh sach hoc sinh: ")
			value, ok := listlop[tenlop]
			if ok {
				for i := 0; i < len(listlop[tenlop].lopsv); i++ {
					fmt.Println("Cac hoc sinh trong lop: ")
					fmt.Println((value.lopsv)[i])
				}
			} else {
				fmt.Println("Lop khong co tren he thong.")
			}

		}
		if input == 9 {
			var tenlop string
			fmt.Println("Nhap ten lop ban muon xem danh sach hoc sinh: ")
			value, ok := listlop[tenlop]
			if ok {
				var ten string
				fmt.Println("Nhap ten hoc sinh ban muon xoa khoi lop:")
				fmt.Scanln(ten)
				for v := range value.lopsv {
					if (value.lopsv[v]).name == ten {
						value.lopsv = append(value.lopsv[:v], value.lopsv[v+1:]...)

					} else {
						fmt.Println("Hoc sinh khong co trong lop.")
					}
				}
			} else {
				fmt.Println("Lop khong ton tai.")
			}
		}
		if input == 0 {
			break
		}
	}
}
