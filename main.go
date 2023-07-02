package main

import (
	"fmt"
)

var listsv = map[string]sinhvien{}
var listlop = map[string]lop{}

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

func addsv(masv1 *int, ten1 *string, gioitinh1 *string, tuoi1 *int, toan1 *float32, ly1 *float32, hoa1 *float32) {
	var masv int
	var ten string
	var gioitinh string
	var tuoi int
	var toan float32
	var ly float32
	var hoa float32
	fmt.Println("Ma SV: ")
	fmt.Scanln(&masv)
	fmt.Println("Ten: ")
	fmt.Scanln(&ten)
	fmt.Println("Gioi tinh: ")
	fmt.Scanln(&gioitinh)
	fmt.Println("Tuoi: ")
	fmt.Scanln(&tuoi)
	fmt.Println("Toan: ")
	fmt.Scanln(&toan)
	fmt.Println("Ly ")
	fmt.Scanln(&ly)
	fmt.Println("Hoa: ")
	fmt.Scanln(&hoa)
	masv = *masv1
	ten = *ten1
	gioitinh = *gioitinh1
	tuoi = *tuoi1
	toan = *toan1
	ly = *ly1
	hoa = *hoa1
}

func showsv() {
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

func deletesv(ten *string) {
	delete(listsv, *ten)
}

func addlop(tenlop1 *string, ten1 *string) {
	var tenlop string
	var ten string
	fmt.Println("Ten lop ban muon them: ")
	fmt.Scanln(&tenlop)
	fmt.Println("Ten hoc sinh ban muon them: ")
	fmt.Scanln(&ten)
	tenlop = *tenlop1
	ten = *ten1
}

func showlop() {
	for k := range listlop {
		fmt.Println("Cac lop co tren he thong:")
		fmt.Println(k)
	}
}

func addsvlop() {
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

func showsvlop() {
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

func deletesvlop() {
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
func main() {
	for true {
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
			var masv1 int
			var ten1 string
			var gioitinh1 string
			var tuoi1 int
			var toan1 float32
			var ly1 float32
			var hoa1 float32
			addsv(&masv1, &ten1, &gioitinh1, &tuoi1, &toan1, &ly1, &hoa1)
			masv1 = sv.code
			ten1 = sv.name
			gioitinh1 = sv.gender
			tuoi1 = sv.age
			toan1 = sv.math
			ly1 = sv.physics
			hoa1 = sv.chemis
			listsv[sv.name] = sv
		}
		if input == 2 {
			showsv()
		}
		if input == 3 {
			var ten string
			fmt.Println("Ten hoc sinh ban muon xoa khoi danh sach: ")
			fmt.Scanln(&ten)
			deletesv(&ten)

		}
		if input == 4 {
			var ten string
			fmt.Println("Nhap ten hoc sinh ban muon chinh sua: ")
			fmt.Scanln(&ten)
			deletesv(&ten)
			var sv sinhvien
			var masv1 int
			var ten1 string
			var gioitinh1 string
			var tuoi1 int
			var toan1 float32
			var ly1 float32
			var hoa1 float32
			fmt.Println("Nhap lai thong tin sinh vien: ")
			addsv(&masv1, &ten1, &gioitinh1, &tuoi1, &toan1, &ly1, &hoa1)
			masv1 = sv.code
			ten1 = sv.name
			gioitinh1 = sv.gender
			tuoi1 = sv.age
			toan1 = sv.math
			ly1 = sv.physics
			hoa1 = sv.chemis
			listsv[sv.name] = sv
		}
		if input == 5 {
			var ten1 string
			var tenlop1 string
			var sv sinhvien
			var class lop
			addlop(&tenlop1, &ten1)
			sv.name = ten1
			class.classname = tenlop1
			class.lopsv = append(class.lopsv, sv)
		}
		if input == 6 {
			showlop()
		}
		if input == 7 {
			addsvlop()
		}
		if input == 8 {
			showsvlop()
		}
		if input == 9 {
			deletesvlop()
		}
		if input == 0 {
			break
		}
	}
}
