#include <bits/stdc++.h>
using namespace std;
class sinhvien {
	public:
		string code;
		string name, gender;
		double age, maths, physics, chemis;
		void enter();
		void output();
		void edit(string ten);
		void del(string ten);
		sinhvien();
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
	public:	
		string classname;
		map <string, sinhvien> listsvlop;
		void seeallclass();
		void addclass();
		void addstudentclass(string tenlop);
		void studentinclass(string tenlop);
		void delstudentclass(string tenlop);
		lop();
		lop(string classname);
		void setclass(string tenlop) {
			this->classname = tenlop;
			
		}
	
};

map <string, sinhvien> listsv;
map <string, lop> listlop;


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
	setsv(masv, ten, gioitinh, tuoi, toan, ly, hoa);

	
}
	

void sinhvien::output() {	
	cout << setw(6) << "Ma SV" << setw(8) << "Ten" << setw(-10) << "Gioi tinh" << setw(-12) << "Tuoi" << setw(-9) << "Diem toan" << setw(-10) << "Diem ly" << setw(-10) << "Diem hoa" << endl;
	for (pair <string, sinhvien> x: listsv) {
	cout << setw(-14) << (x.second).code << setw(-9) << (x.second).name << setw(-12) << (x.second).gender << setw(-14) << (x.second).age << setw(-12) << (x.second).maths << setw(-12) << (x.second).physics << setw(-12) << (x.second).chemis << endl; 
	
	}
	
}

void sinhvien::edit(string ten) {
	if (listsv.find(ten) != listsv.end()) {
		listsv.erase(ten);
		string masv, gioitinh;
		double tuoi, toan, ly, hoa;
		cout << "Nhap lai ma SV: " << endl;
		getline(cin, masv);
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
		listsv[ten].setsv(masv, ten, gioitinh, tuoi, toan, ly, hoa);
	}
	else {
		cout << "SV ko co tren he thong." << endl;
	}

}

void sinhvien::del(string ten) {
	if (listsv.find(ten) != listsv.end()) {
		listsv.erase(ten);
	}
	else {
		cout << "Sinh vien nay khong co tren he thong." << endl;
	}
}
	


void lop::seeallclass() {
	for (pair <string, lop> x:listlop) {
		cout << "Danh sach cac lop:" << endl;
		cout << x.first << endl;
	}
}


void lop::addclass() {
	string tenlop;
	cout << "Nhap ten lop: " << endl;
	getline(cin, tenlop);
	setclass(tenlop);
	
	
	
	
	
}

void lop::addstudentclass(string tenlop) {
	if (listlop.find(tenlop) != listlop.end()) {
		string ten;
		cout << "Hay them hoc sinh vao lop " << tenlop << ": " << endl;
		getline(cin, ten);
		if (listsv.find(ten) != listsv.end()) {
			(listlop[tenlop].listsvlop).insert({ten, listsv[ten]});
		}
		else {
				cout << "Khong co hoc sinh nay tren he thong." << endl;
		}
					
	}
	
	else {
			cout << "Lop nay chua co tren he thong. Hay them lop vao de them hoc sinh." << endl;
	}
}


void lop::studentinclass(string tenlop) {
	cout << "Cac hoc sinh trong lop: " << tenlop << endl;
	for (pair <string, lop> x:listlop) {
		if (x.first == tenlop) {
			for (pair <string, sinhvien> y:((x.second).listsvlop)) {
				cout << y.first << endl;

			}
		}
		else {
			cout << "Khong co lop nay" << endl;
		}
		}
	}
	
void delstudentclass(string tenlop) {
	if (listlop.find(tenlop) != listlop.end()) {
		string ten;
		cout << "Hay xoa hoc sinh trong lop " << tenlop << ": " << endl;
		cin >> ten;
		if (listlop[tenlop].listsvlop.find(ten) != (listlop[tenlop]).listsvlop.end()) {
			(listlop[tenlop]).listsvlop.erase(ten);
		}
		else {
			cout << "Hoc sinh nay khong co trong lop." << endl;
		}
	}
	else {
		cout << "Khong co lop nay tren he thong." << endl;
	}
	
}

	
int main()
{
    int input;
    string code;
    sinhvien sv;
    lop lp;
    string classname;
    while (true)
    {
        cout << "HE THONG QUAN LY SINH VIEN\n";
        cout << "CHUC NANG : \n";
        cout << "1. Them thong tin sinh vien." << endl;
        cout << "2. Dua ra thong tin sinh vien" << endl;
        cout << "3. Xoa thong tin sinh vien." << endl;
        cout << "4. Chinh sua thong tin sinh vien." << endl;
        cout << "5. Them lop hoc." << endl;
        cout << "6. Xem danh sach lop hoc." << endl;
        cout << "7. Them sinh vien vao lop." << endl;
        cout << "8. Xem danh sach sinh vien trong lop." << endl;
        cout << "9. Xoa sinh vien khoi lop hoc." << endl;
        cout << "0. Thoat." << endl;
        cout << "Hay chon chuc nang: " ; 
        cin >> input;
        if (input == 1) {
        	string ten;
        	sv.enter();
        	ten = sv.name;
        	listsv.insert({ten, sv});
        	break;
		}
		if (input == 2) {
			sv.output();
			break;
		}
		if (input == 3) {
			cout << "Nhap ma SV ban muon xoa: ";
			getline(cin,code);
			sv.del(code);
			break;
		}
		if (input == 4) {
			cout << "Nhap ma SV ban muon sua: ";
			getline(cin,code);
			sv.edit(code);
			break;
		}
		if (input == 5) {
			lp.addclass();
			string tlop;
        	lp.addclass();
        	tlop = lp.classname;
        	listlop.insert({tlop, lp});
			break;
		}
		if (input == 6) {
			lp.seeallclass();
			break;
		}
		if (input == 7) {
			cout << "Ten lop ma ban muon them hoc sinh: ";
			lp.addstudentclass(classname);
			break;
		}
		if (input == 8) {
			cout << "Ten lop ma ban muon xem hoc sinh: ";
			getline(cin, classname);
			lp.studentinclass(classname);
			break;
		}
		if (input == 9) {
			cout << "Ten lop ma ban muon xoa hoc sinh: ";
			getline(cin, classname);
			lp.delstudentclass(classname);
			break;
		}
	}
}


