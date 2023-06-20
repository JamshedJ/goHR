CREATE TABLE employees (
    id              SERIAL PRIMARY KEY,
    first_name      VARCHAR(255) NOT NULL,
    last_name       VARCHAR(255) NOT NULL,
    position_id     INT REFERENCES positions(id),
    department_id   INT REFERENCES departments(id),
    employment_date DATE NOT NULL,
    salary          DECIMAL(10, 2) NOT NULL
);

CREATE TABLE departments (
    id              SERIAL PRIMARY KEY,
    title           VARCHAR(255) NOT NULL,
    department_head VARCHAR(255)
);

CREATE TABLE positions (
    id              SERIAL PRIMARY KEY,
    title           VARCHAR(255) NOT NULL,
    salary          DECIMAL(10, 2) NOT NULL,
    qualification   VARCHAR(255)
);

--CREATE TABLE vacation_requests (
--    id              SERIAL PRIMARY KEY,
--    employee_id     INT REFERENCES employees(id),
--    start_date      DATE NOT NULL,
--    end_date        DATE NOT NULL,
--    reason          VARCHAR(255)
--);
--
--CREATE TABLE sick_leave_requests (
--    id              SERIAL PRIMARY KEY,
--    employee_id     INT REFERENCES employees(id),
--    start_date      DATE NOT NULL,
--    end_date        DATE NOT NULL,
--    reason          VARCHAR(255)
--);

CREATE TABLE users (
    id              SERIAL PRIMARY KEY,
    username        VARCHAR(255) NOT NULL,
    password        VARCHAR(255) NOT NULL,
    role            VARCHAR(50) DEFAULT 'user' -- admin, user
);

CREATE TABLE requests (
    id              SERIAL PRIMARY KEY,
    employee_id     INT REFERENCES employees(id),
    start_date      DATE NOT NULL,
    end_date        DATE NOT NULL,
    reason          VARCHAR(255),
    type_id         INT NOT NULL REFERENCES types(id)
);

CREATE TABLE types (
    id              SERIAL PRIMARY KEY,
    title           VARCHAR(255)  --(sick_leave or vacation)
);
