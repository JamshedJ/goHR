CREATE TABLE employees (
                id INT PRIMARY KEY,
                first_name VARCHAR(100),
                last_name VARCHAR(100),
                position VARCHAR(100),
                department VARCHAR(100),
                employment_date DATE,
                salary DECIMAL(10, 2)
);

CREATE TABLE departments (
                id INT PRIMARY KEY,
                title VARCHAR(100),
                department_head VARCHAR(100)
);

CREATE TABLE positions (
                id INT PRIMARY KEY,
                title VARCHAR(100),
                salary DECIMAL(10, 2),
                qualification VARCHAR(100)
);

CREATE TABLE vacation_requests (
                id INT PRIMARY KEY,
                employee_id INT,
                start_date DATE,
                end_date DATE,
                reason VARCHAR(100),
                FOREIGN KEY (employee_id) REFERENCES employees (id)
);

CREATE TABLE sick_leave_requests (
                id INT PRIMARY KEY,
                employee_id INT,
                start_date DATE,
                end_date DATE,
                reason VARCHAR(100),
                FOREIGN KEY (employee_id) REFERENCES employees (id)
);

CREATE TABLE users (
                id INT PRIMARY KEY,
                username VARCHAR(100),
                password VARCHAR(100),
                role VARCHAR(50)
);
