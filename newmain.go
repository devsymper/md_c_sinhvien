package main

import (
	"fmt"
)

type Phone interface {
	insertPhone(name *string)
	removePhone(name *string)
	updatePhone(name *string)
	searchPhone(name *string)
	sort()
}

type PhoneBook struct {
	Phone
	PhoneList map[string]string
}

func (p *PhoneBook) insertPhone(name *string) {
	var name1 string
	name1 = *name
	value, ok := p.PhoneList[name1]
	if ok {
		var check string
		fmt.Println("Ten nay da co trong so dien thoai.")
		fmt.Println("Kiem tra lai so dien thoai cua nguoi nay (Co/Khong).")
		fmt.Scanln(&check)
		if check == "Co" {
			fmt.Println("Nhap so dien thoai cua nguoi nay: ")
			var NewNumber string
			fmt.Scanln(&NewNumber)
			if NewNumber == value {
				fmt.Scanln("So dien thoai giong voi tren danh ba.")
			} else {
				var SDT string = NewNumber + "," + value
				p.PhoneList[name1] = SDT
			}
		} else {
			fmt.Println("Nguoi nay khong co trong danh ba.")
		}
	} else {
		var NewPhone string
		fmt.Println("Them so dien thoai cua nguoi nay: ")
		fmt.Scanln(NewPhone)
		p.PhoneList[name1] = NewPhone
		fmt.Println("Quay ve man hinh chinh.")
	}
}

func (p *PhoneBook) removePhone(name *string) {
	var name1 string
	value, ok := p.PhoneList[name1]
	if ok {
		delete(p.PhoneList, name1)
		fmt.Println("Da xoa thanh cong so dien thoai: ", value)
	} else {
		fmt.Println("Nguoi dung ", name1, "khong co trong danh ba.")
		fmt.Println("Quay ve man hinh chinh.")
	}
}

func (p *PhoneBook) updatePhone(name *string) {
	var name1 string
	name1 = *name
	value, ok := p.PhoneList[name1]
	if ok {
		fmt.Println("Nhap so dien thoai moi cua nguoi nay: ")
		var NewNumber string
		fmt.Scanln(&NewNumber)
		fmt.Scanln("So dien thoai cu cua nguoi nay la: ", value, ", ban co muon sua? (Co/Khong)")
		var ans string
		fmt.Scanln(&ans)
		if ans == "Co" {
			p.PhoneList[name1] = NewNumber
		} else {
			fmt.Println("Quay ve man hinh chinh.")
		}
	} else {
		fmt.Println("Nguoi nay khong co trong danh ba.")
		fmt.Println("Quay ve man hinh chinh.")
	}
}

func (p *PhoneBook) searchPhone(name *string) {
	var name1 string
	name1 = *name
	value, ok := p.PhoneList[name1]
	if ok {
		fmt.Println(name1, ": ", value)
	} else {
		fmt.Println("Nguoi nay khong co trong danh ba.")
	}
}

func (p *PhoneBook) sort() {
	for k, v := range p.PhoneList {
		fmt.Println(k, ": ", v)
	}
}

func main() {
	var p PhoneBook
	var input int
	for true {
		fmt.Println("He thong kiem soat danh ba:")
		fmt.Println("CHUC NANG")
		fmt.Println("1. Them SDT vao danh ba")
		fmt.Println("2. Xoa SDT khoi danh ba")
		fmt.Println("3. Chinh SDT trong danh ba")
		fmt.Println("4. Tim kiem SDT")
		fmt.Println("5. Xem danh ba")
		fmt.Println("0. Thoat")
		fmt.Println("Nhap chuc nang ban muon su dung:")
		fmt.Scanln(&input)
		if input == 1 {
			var name string
			fmt.Println("Nhap ten cua nguoi ban muon them vao danh ba: ")
			fmt.Scanln(&name)
			p.insertPhone(&name)
		}
		if input == 2 {
			var name string
			fmt.Println("Nhap ten nguoi dung ban muon xoa khoi danh ba: ")
			fmt.Scanln(&name)
			p.removePhone(&name)
		}
		if input == 3 {
			var name string
			fmt.Println("Nhap ten nguoi dung ban muon sua so: ")
			fmt.Scanln(&name)
			p.updatePhone(&name)
		}
		if input == 4 {
			var name string
			fmt.Println("Nhap ten nguoi dung ban muon tim: ")
			fmt.Scanln(&name)
			p.searchPhone(&name)
		}
		if input == 5 {
			p.sort()
		}
		if input == 0 {
			break
		}
	}
}
