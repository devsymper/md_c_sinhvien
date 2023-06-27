#include <bits/stdc++.h>
#include <mysql_driver.h>
#include <mysql_connection.h>
#include <iostream>
#include <string>
#include <stdlib.h>
#include "mysql_connection.h"
#include <cppconn/driver.h>
#include <cppconn/exception.h>
#include <cppconn/resultset.h>
#include <cppconn/statement.h>

using namespace std;

int main()
{
    	int input;
    	sql::Driver *driver;
	sql::Connection *con;
	sql::Statement *stmt;
	sql::ResultSet *res;
	sql::PreparedStatement *pstmt;
	driver = get_driver_instance();
	con = driver->connect("tcp://127.0.0.1:3306", "root", "12345678");
	con->setSchema("test");
	stmt = con->createStatement();
    

			
	while (true) {
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
			string tuoi, toan, ly, hoa;
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
			string x = masv + ", " + ten + ", " + gioitinh + ", " + tuoi + ", " + toan + ", " + ly + ", " + hoa;
			string insertQuery = "INSERT INTO student(MaSV, TenSV, GioiTinh, Tuoi, Toan, Ly, Hoa) VALUES " + x;
			stmt->execute(insertQuery);
			delete stmt;
        		break;
		}	
		if (input == 2) {
			stmt->execute("USE test");
			stmt->execute("SELECT * FROM student;");	
			delete stmt;
			break;
		}
		if (input == 3) {
			string code;
			cout << "Nhap ma SV cua sinh vien ma ban muon xoa: ";
			cin >> code;
			string deleteQuery = "DELETE FROM student WHERE MaSV = '" + code + "'";
			stmt->execute(deleteQuery);
			delete stmt;
			break;
		}
		if (input == 4) {
			cout << "Nhap ma sinh vien ban muon sua thong tin: ";
			string ten, gioitinh;
			string tuoi, toan, ly, hoa;
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
			string updateQuery1 = "UPDATE student SET TenSV = " + ten;
			string updateQuery2 = "UPDATE student SET GioiTinh = " + gioitinh;
			string updateQuery3 = "UPDATE student SET Tuoi = " + tuoi;
			string updateQuery4 = "UPDATE student SET Toan = " + toan;
			string updateQuery5 = "UPDATE student SET Ly = " + ly;
			string updateQuery6 = "UPDATE student SET Hoa = " + hoa;
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
        		stmt->execute("SELECT * FROM student;");
        		string updateQuery = "UPDATE student SET Lop = '" + lop + "' WHERE MaSV = '" + ma + "'";
        		stmt->execute(updateQuery);
        		delete stmt;
			break;
		}
		if (input == 6) {
			stmt->execute("SELECT DISTINCT Lop FROM student;");
			delete stmt;
			break;
		}
		if (input == 7) {
			string classname, ten;
			cout << "Ten lop ma ban muon them hoc sinh: ";
			getline(cin, classname);
			cout << "Hoc sinh ban muon them vao lop: ";
			getline(cin, ten);
			string updateQuery = "UPDATE student SET Lop = '" + classname + "' WHERE TenSV = '" + ten + "'";
			stmt->execute(updateQuery);
			delete stmt;
			break;
		}
		if (input == 8) {
			string classname;
			cout << "Ten lop ma ban muon xem hoc sinh: ";
			getline(cin, classname);
			string selectQuery = "SELECT * FROM student WHERE Lop = '" + classname + "'";
			stmt->execute(selectQuery);
			delete stmt;
			break;
		}
		if (input == 9) {
			string classname, ten;
			cout << "Ten lop ma ban muon xoa hoc sinh: ";
			getline(cin, classname);
			cout << "Ten hoc sinh ban muon xoa khoi lop: ";
			getline(cin, ten);
			string updateQuery = "UPDATE student SET Lop = 'NoClass' WHERE TenSV = '" + ten + "'";
			stmt->execute(updateQuery);
			delete stmt;
			break;
		}
		if (input == 10) {
			string search, updateQuery;
			cout << "Nhap ten de tra cuu:" << endl;
			cin >> search;
			updateQuery = "SELECT * FROM student WHERE TenSV LIKE '%" + search + "%'";
			stmt->execute(updateQuery);
			delete stmt;
			
		}
		if (input == 0) {
			return 0;
		}
	delete con;
	
	return 0;		
		
	}
    
}
