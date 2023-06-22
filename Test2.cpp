#include <bits/stdc++.h>
using namespace std;
class sinhvien {
	private:
		string code;
		string name, gender;
		double age, maths, physics, chemis;
	public:
		void enter();
		void output();
		void edit(string masv);
		void del(string masv);
		sinhvien(string masv, string ten, string gioitinh, double tuoi, double toan, double ly, double hoa);
		void setsv(string masv, string ten, string gioitinh, double tuoi, double toan, double ly, double hoa) {
			this->code = masv;
			this->name = ten;
			this->gender = gioitinh;
			this->age = tuoi;
			this->maths = toan;
			this->physics = ly;
			this->chemis = hoa;
		}
};
		


class lop {
	private:
		string classname;
	public:
		vector <sinhvien> listsvlop;
		void seeallclass();
		void addclass();
		void addstudentclass();
		void studentinclass();
		void delstudentclass();
		lop(string classname);
		void setclass(string tenlop) {
			classname = tenlop;
			
		}
	
};

vector <sinhvien> listsv;
vector <lop> listlop;
vector <lop>:: iterator check;


void sinhvien::enter() {
	string ten, gioitinh, masv;
	double tuoi, toan, ly, hoa;
	cout << "Ma SV: " << endl;
	getline(cin, masv);
	cout << "Ten: " << endl;
	getline(cin, ten);
	cout << "Gioi tinh: " << endl;
	getline(cin, gioitinh);
	cout << "Tuoi: " << endl;
	cin >> tuoi;
	cout << "Diem toan: " << endl;
	cin >> toan;
	cout << "Diem ly: " << endl;
	cin >> ly;
	cout << "Diem hoa: " << endl;
	cin >> hoa;
	listsv.push_back(sinhvien(masv, ten, gioitinh, tuoi, toan, ly, hoa));

	
}
	

void sinhvien::output() {	
	cout << setw(6) << "Ma SV" << setw(8) << "Ten" << setw(-10) << "Gioi tinh" << setw(-12) << "Tuoi" << setw(-9) << "Diem toan" << setw(-10) << "Diem ly" << setw(-10) << "Diem hoa" << endl;
	for (int i = 0; i < listsv.size(); i++) {
	cout << setw(-14) << listsv[i].code << setw(-9) << listsv[i].name << setw(-12) << listsv[i].gender << setw(-14) << listsv[i].age << setw(-12) << listsv[i].maths << setw(-12) << listsv[i].physics << setw(-12) << listsv[i].chemis << endl; }
	
}

void sinhvien::edit(string masv) {
	for (int i = 0; i < listsv.size(); i++) {
		if (listsv[i].code == masv) {
			string ten, gioitinh;
			double tuoi, toan, ly, hoa;
			cout << "Nhap lai ten: " << endl;
			getline(cin, ten);
			cout << "Nhap lai gioi tinh: " << endl;
			getline(cin, gioitinh);
			cout << "Nhap lai tuoi: " << endl;
			cin >> tuoi;
			cout << "Nhap lai diem toan: " << endl;
			cin >> toan;
			cout << "Nhap lai diem ly: " << endl;
			cin >> ly;
			cout << "Nhap lai diem hoa: " << endl;
			cin >> hoa;
			listsv[i].setsv(masv, ten, gioitinh, tuoi, toan, ly, hoa);
		}
		else {
			cout << "Ma SV ko co tren he thong." << endl;
		}
	}
}

void sinhvien::del(string masv) {
	for (int i = 0; i < listsv.size(); i++) {
		if (listsv[i].code == masv) {
			listsv.erase(listsv.begin() + i);
		}
		else {
			cout << "Ma SV ko co tren he thong." << endl;
		}
	
	}
	
}

void lop::seeallclass() {
	for (int i = 0; i < listlop.size(); i++) {
		cout << "Danh sach cac lop:" << endl;
		cout << listlop[i].classname << endl;
	}
}


void lop::addclass() {
	string tenlop;
	cout << "Nhap ten lop: " << endl;
	getline(cin, tenlop);
	if (check == listlop.end()) {
		listlop.push_back(lop(tenlop));
	}
	
	
}
void lop::addstudentclass() {
	
}




