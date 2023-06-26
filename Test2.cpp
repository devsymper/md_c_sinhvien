#include <bits/stdc++.h>
#include <mysql_driver.h>
#include <mysql_connection.h>

using std::cout;
using std::cin;
using std::getline;
using std::string;
using std::endl;
using std::to_string;

int main()
{
    int input;
    

			
	while (true) {
		sql::mysql::MySQL_Driver *driver;
		sql::Connection *con;
		driver = sql::mysql::get_mysql_driver_instance();
		con = driver->connect("tcp://127.0.0.1:3306", "root", "12345678");
		
    	
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
        	cout << "10. Tra cuu hoc sinh." << endl;
        	cout << "0. Thoat." << endl;
        	cout << "Hay chon chuc nang: " ; 
        	cin >> input;
        	if (input == 1) {
        		string ten, gioitinh, masv;
			float tuoi, toan, ly, hoa;
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
			string strtuoi = to_string(tuoi);
			string strtoan = to_string(toan);
			string strly = to_string(ly);
			string strhoa = to_string(hoa);
			string x = masv + ", " + ten + ", " + gioitinh + ", " + strtuoi + ", " + strtoan + ", " + strly + ", " + strhoa;
			
			sql::Statement *stmt;
			stmt = con->createStatement();
			stmt->execute("USE test");
			string insertQuery = "INSERT INTO student(MaSV, TenSV, GioiTinh, Tuoi, Toan, Ly, Hoa) VALUES " + x;
			stmt->execute(insertQuery);
			delete stmt;
					
        		break;
		}	
		if (input == 2) {
			sql::Statement *stmt;
			stmt = con->createStatement();
			stmt->execute("USE test");
			stmt->execute("SELECT * FROM student;");	
			delete stmt;
			break;
		}
		if (input == 3) {
			string code;
			cout << "Nhap ma SV cua sinh vien ma ban muon xoa: ";
			cin >> code;
			sql::Statement *stmt;
			stmt = con->createStatement();
			stmt->execute("USE test");
			string deleteQuery = "DELETE FROM student WHERE MaSV = '" + code + "'";
			stmt->execute(deleteQuery);
			delete stmt;
			break;
		}
		if (input == 4) {
			cout << "Nhap ma sinh vien ban muon sua thong tin: ";
			string ten, gioitinh;
			float tuoi, toan, ly, hoa;
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
			string strtuoi = to_string(tuoi);
			string strtoan = to_string(toan);
			string strly = to_string(ly);
			string strhoa = to_string(hoa);
			sql::Statement *stmt;
			stmt = con->createStatement();
			stmt->execute("USE test");
			string updateQuery1 = "UPDATE student SET TenSV = " + ten;
			string updateQuery2 = "UPDATE student SET GioiTinh = " + gioitinh;
			string updateQuery3 = "UPDATE student SET Tuoi = " + strtuoi;
			string updateQuery4 = "UPDATE student SET Toan = " + strtoan;
			string updateQuery5 = "UPDATE student SET Ly = " + strly;
			string updateQuery6 = "UPDATE student SET Hoa = " + strhoa;
			stmt->execute(updateQuery1);
			stmt->execute(updateQuery2);
			stmt->execute(updateQuery3);
			stmt->execute(updateQuery4);
			stmt->execute(updateQuery5);
			stmt->execute(updateQuery6);
			delete stmt;
			break;
		}
		if (input == 5) {
			string lop, ma;
			cout << "Nhap ten lop ban muon them";
			getline(cin, lop);
			cout << "Nhap ma sinh vien cua 1 hoc sinh trong lop: ";
			getline(cin, ma);
			sql::Statement *stmt;
			stmt = con->createStatement();
       			stmt->execute("USE " (test));
        		stmt->execute("SELECT * FROM student;");
        		string updateQuery = "UPDATE student SET Lop = '" + lop + "' WHERE MaSV = '" + ma + "'";
        		stmt->execute(updateQuery);
			break;
		}
		if (input == 6) {
			sql::Statement *stmt;
			stmt = con->createStatement();
			stmt->execute("USE test");
			stmt->execute("SELECT DISTINCT Lop FROM student;");
			break;
		}
		if (input == 7) {
			string classname, ten;
			cout << "Ten lop ma ban muon them hoc sinh: ";
			getline(cin, classname);
			cout << "Hoc sinh ban muon them vao lop: ";
			getline(cin, ten);
			sql::Statement *stmt;
			stmt = con->createStatement();
			stmt->execute("USE test");
			string updateQuery = "UPDATE student SET Lop = '" + classname + "' WHERE TenSV = '" + ten + "'";
			stmt->execute(updateQuery);
			break;
		}
		if (input == 8) {
			string classname;
			cout << "Ten lop ma ban muon xem hoc sinh: ";
			getline(cin, classname);
			sql::Statement *stmt;
			stmt = con->createStatement();
			stmt->execute("USE test");
			string selectQuery = "SELECT * FROM student WHERE Lop = '" + classname + "'";
			stmt->execute(selectQuery);
			break;
		}
		if (input == 9) {
			string classname, ten;
			cout << "Ten lop ma ban muon xoa hoc sinh: ";
			getline(cin, classname);
			cout << "Ten hoc sinh ban muon xoa khoi lop: ";
			getline(cin, ten);
			sql::Statement *stmt;
			stmt = con->createStatement();
			stmt->execute("USE test");
			string updateQuery = "UPDATE student SET Lop = 'NoClass' WHERE TenSV = '" + ten + "'";
			stmt->execute(updateQuery);
			break;
		}
		if (input == 10) {
			string search, updateQuery;
			cout << "Nhap ten de tra cuu:" << endl;
			cin >> search;
			sql::Statement *stmt;
			stmt = con->createStatement();
			stmt->execute("USE test")
			updateQuery = "SELECT * FROM student WHERE TenSV LIKE '%" + search + "%'";
			stmt->execute(updateQuery);
		}
		if (input == 0) {
			return 0;
		}
	delete con;
	
	return 0;		
		
	}
    
}
