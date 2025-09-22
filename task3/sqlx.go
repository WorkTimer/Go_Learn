package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func SqlxExample() {
	// 数据库连接字符串（根据 docker-compose.yml 配置）
	dsn := "user=postgres password=123456 dbname=students sslmode=disable host=localhost port=5432"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ================================
	// 题目1：使用SQL扩展库进行查询
	// ================================
	// 创建 employees 表并插入测试数据
	err = createEmployeesTable(db)
	if err != nil {
		log.Fatal(err)
	}

	// 查询技术部所有员工
	employees, err := getEmployeesByDepartment(db, "技术部")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("技术部员工: %+v\n", employees)

	// 查询工资最高的员工
	topEmployee, err := getHighestPaidEmployee(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("工资最高员工: %+v\n", topEmployee)

	// ================================
	// 题目2：实现类型安全映射
	// ================================
	// 创建 books 表并插入测试数据
	err = createBooksTable(db)
	if err != nil {
		log.Fatal(err)
	}
	// 查询价格大于50元的书籍
	books, err := getBooksByPrice(db, 50.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("价格大于50元的书籍: %+v\n", books)
}

// ================================
// 题目1相关结构体和函数
// ================================
// Employee 结构体定义
type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

// 创建 employees 表并插入测试数据
func createEmployeesTable(db *sqlx.DB) error {
	// 创建表
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS employees (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		department VARCHAR(50) NOT NULL,
		salary INTEGER NOT NULL
	)`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	// 清空表并插入测试数据
	_, err = db.Exec("DELETE FROM employees")
	if err != nil {
		return err
	}

	insertSQL := `
	INSERT INTO employees (name, department, salary) VALUES 
	('张三', '技术部', 8000),
	('李四', '技术部', 9000),
	('王五', '销售部', 6000),
	('赵六', '技术部', 10000),
	('钱七', '人事部', 7000),
	('孙八', '技术部', 8500)`

	_, err = db.Exec(insertSQL)
	return err
}

// 查询指定部门的所有员工
func getEmployeesByDepartment(db *sqlx.DB, department string) ([]Employee, error) {
	var employees []Employee
	query := "SELECT id, name, department, salary FROM employees WHERE department = $1"
	err := db.Select(&employees, query, department)
	return employees, err
}

// 查询工资最高的员工
func getHighestPaidEmployee(db *sqlx.DB) (Employee, error) {
	var employee Employee
	query := "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1"
	err := db.Get(&employee, query)
	return employee, err
}

// ================================
// 题目2相关结构体和函数
// ================================
// Book 结构体定义
type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

// 创建 books 表并插入测试数据
func createBooksTable(db *sqlx.DB) error {
	// 创建表
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		author VARCHAR(50) NOT NULL,
		price DECIMAL(10,2) NOT NULL
	)`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	// 清空表并插入测试数据
	_, err = db.Exec("DELETE FROM books")
	if err != nil {
		return err
	}

	insertSQL := `
	INSERT INTO books (title, author, price) VALUES 
	('Go语言圣经', 'Alan Donovan', 89.50),
	('算法导论', 'Thomas Cormen', 128.00),
	('设计模式', 'Gang of Four', 45.00),
	('深入理解计算机系统', 'Randal Bryant', 95.00),
	('编程珠玑', 'Jon Bentley', 35.00),
	('代码整洁之道', 'Robert Martin', 65.00)`

	_, err = db.Exec(insertSQL)
	return err
}

// 查询价格大于指定值的书籍
func getBooksByPrice(db *sqlx.DB, minPrice float64) ([]Book, error) {
	var books []Book
	query := "SELECT id, title, author, price FROM books WHERE price > $1"
	err := db.Select(&books, query, minPrice)
	return books, err
}
