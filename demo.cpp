#include <mysql_driver.h>
#include <mysql_connection.h>

int main() {
    sql::mysql::MySQL_Driver *driver;
    sql::Connection *con;

    // Kết nối với cơ sở dữ liệu MySQL
    driver = sql::mysql::get_mysql_driver_instance();
    con = driver->connect("tcp://127.0.0.1:3306", "username", "password");

    // Chèn dữ liệu
    {
        sql::Statement *stmt;
        stmt = con->createStatement();
        stmt->execute("USE database_name"); // Thay thế database_name bằng tên cơ sở dữ liệu 
                                        // thay thế các column và value
        std::string insertQuery = "INSERT INTO table_name (column1, column2, column3) VALUES ('value1', 'value2', 'value3')";
        stmt->execute(insertQuery);

        delete stmt;
    }
    // Cập nhật dữ liệu
    {
        sql::Statement *stmt;
        stmt = con->createStatement();
        stmt->execute("USE database_name"); // Thay thế database_name bằng tên cơ sở dữ liệu 

        std::string updateQuery = "UPDATE table_name SET column1 = 'new_value' WHERE condition";  // thay thế condition bằng điều kiện xóa
        stmt->execute(updateQuery);

        delete stmt;
    }

    // Xóa dữ liệu
    {
        sql::Statement *stmt;
        stmt = con->createStatement();
        stmt->execute("USE database_name"); // Thay thế database_name bằng tên cơ sở dữ liệu 

        std::string deleteQuery = "DELETE FROM table_name WHERE condition";  // thay thế condition bằng điều kiện xóa
        stmt->execute(deleteQuery);

        delete stmt;
    }
    // Đóng kết nối
    delete con;

    return 0;
}
