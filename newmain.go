package main

import (
	"fmt"
)

type PhoneInterface interface {
	createMap()
	insertPhone(name, phone string)
	removePhone(name string)
	updatePhone(name, newphone string)
	searchPhone(name string)
	sort()
}

type PhoneBook struct {
	PhoneList map[string]string
}

func InitPhone() PhoneInterface {
	return &PhoneBook{}
}

func (p *PhoneBook) createMap() {
	if p.PhoneList == nil {
		p.PhoneList = make(map[string]string)
	}
}

func (p *PhoneBook) insertPhone(name, phone string) {
	value, ok := p.PhoneList[name]
	if ok {
		var check string
		fmt.Println("Ten nay da co trong so dien thoai.")
		fmt.Println("Them so dien thoai phu cua nguoi nay (Co/Khong).")
		fmt.Scanln(&check)
		if check == "Co" {
			fmt.Println("Nhap so dien thoai cua nguoi nay: ")
			var NewNumber string
			fmt.Scanln(&NewNumber)
			if NewNumber == value {
				fmt.Scanln("So dien thoai giong voi tren danh ba.")
			} else {
				var SDT string = NewNumber + "," + value
				p.PhoneList[name] = SDT
			}
		} else {
			fmt.Println("Quay ve man hinh chinh.")
		}
	} else {
		p.PhoneList[name] = phone
		fmt.Println("Da them thanh cong.")
		fmt.Println("Quay ve man hinh chinh.")
	}
}

func (p *PhoneBook) removePhone(name string) {
	value, ok := p.PhoneList[name]
	if ok {
		delete(p.PhoneList, name)
		fmt.Println("Da xoa thanh cong so dien thoai: ", value)
	} else {
		fmt.Println("Nguoi dung khong co trong danh ba.")
		fmt.Println("Quay ve man hinh chinh.")
	}
}

func (p *PhoneBook) updatePhone(name, newphone string) {
	value, ok := p.PhoneList[name]
	if ok {
		fmt.Println("Ban co muon thay SDT: ")
		fmt.Println(value)
		fmt.Println("Voi SDT: ")
		fmt.Println(newphone)
		fmt.Scanln("Co/Khong")
		var ans string
		fmt.Scanln(&ans)
		if ans == "Co" {
			p.PhoneList[name] = newphone
		} else {
			fmt.Println("Quay ve man hinh chinh.")
		}
	} else {
		fmt.Println("Nguoi nay khong co trong danh ba.")
		fmt.Println("Quay ve man hinh chinh.")
	}
}

func (p *PhoneBook) searchPhone(name string) {
	value, ok := p.PhoneList[name]
	if ok {
		fmt.Println(name)
		fmt.Println(value)
	} else {
		fmt.Println("Nguoi nay khong co trong danh ba.")
	}
}

func (p *PhoneBook) sort() {
	for k, v := range p.PhoneList {
		fmt.Println(k)
		fmt.Println(v)
	}
}

func main() {
	var input int
	var p = InitPhone()
	p.createMap()
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
			var phone string
			fmt.Println("Nhap ten cua nguoi ban muon them vao danh ba: ")
			fmt.Scanln(&name)
			fmt.Println("Nhap SDT ban muon them: ")
			fmt.Scanln(&phone)
			p.insertPhone(name, phone)
		}
		if input == 2 {
			var name string
			fmt.Println("Nhap ten nguoi dung ban muon xoa khoi danh ba: ")
			fmt.Scanln(&name)
			p.removePhone(name)
		}
		if input == 3 {
			var name string
			var newphone string
			fmt.Println("Nhap ten nguoi dung ban muon sua so: ")
			fmt.Scanln(&name)
			fmt.Println("Nhap SDT moi cua nguoi nay: ")
			fmt.Scanln(&newphone)
			p.updatePhone(name, newphone)
		}
		if input == 4 {
			var name string
			fmt.Println("Nhap ten nguoi dung ban muon tim: ")
			fmt.Scanln(&name)
			p.searchPhone(name)
		}
		if input == 5 {
			p.sort()
		}
		if input == 0 {
			break
		}
	}
}
