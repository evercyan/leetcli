-- Sql
CREATE DATABASE IF NOT EXISTS `Leetcode`;
USE `Leetcode`;

DROP TABLE IF EXISTS `184_Employee`;
CREATE TABLE `184_Employee` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) NOT NULL,
  `Salary` int(11) NOT NULL,
  `DepartmentId` int(11),
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `184_Employee` VALUES
(1, "Joe", 70000, 1),
(2, "Henry", 80000, 2),
(3, "Sam", 60000, 2),
(4, "Max", 90000, 1);

DROP TABLE IF EXISTS `184_Department`;
CREATE TABLE `184_Department` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `184_Department` VALUES
(1, "IT"),
(2, "Sales");

```
Employee 表包含所有员工信息，每个员工有其对应的 Id, salary 和 department Id。

+----+-------+--------+--------------+
| Id | Name  | Salary | DepartmentId |
+----+-------+--------+--------------+
| 1  | Joe   | 70000  | 1            |
| 2  | Henry | 80000  | 2            |
| 3  | Sam   | 60000  | 2            |
| 4  | Max   | 90000  | 1            |
+----+-------+--------+--------------+
Department 表包含公司所有部门的信息。

+----+----------+
| Id | Name     |
+----+----------+
| 1  | IT       |
| 2  | Sales    |
+----+----------+
编写一个 SQL 查询，找出每个部门工资最高的员工。例如，根据上述给定的表格，Max 在 IT 部门有最高工资，Henry 在 Sales 部门有最高工资。

+------------+----------+--------+
| Department | Employee | Salary |
+------------+----------+--------+
| IT         | Max      | 90000  |
| Sales      | Henry    | 80000  |
+------------+----------+--------+
```

-- Debug
SELECT t1.Name AS Department, t2.Name AS Employee, t2.Salary AS Salary FROM 184_Department AS t1 LEFT JOIN 184_Employee AS t2 ON t1.Id = t2.DepartmentId WHERE Salary >= (SELECT MAX(Salary) FROM 184_Employee WHERE DepartmentId = t1.Id);


-- Answer
-- 900 ms
SELECT t1.Name AS Department, t2.Name AS Employee, t2.Salary AS Salary FROM Department AS t1 LEFT JOIN Employee AS t2 ON t1.Id = t2.DepartmentId WHERE Salary >= (SELECT MAX(Salary) FROM Employee WHERE DepartmentId = t1.Id);
