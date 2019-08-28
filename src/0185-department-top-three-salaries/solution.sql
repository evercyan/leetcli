-- Sql
CREATE DATABASE IF NOT EXISTS `Leetcode`;
USE `Leetcode`;

DROP TABLE IF EXISTS `185_Employee`;
CREATE TABLE `185_Employee` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) NOT NULL,
  `Salary` int(11) NOT NULL,
  `DepartmentId` int(11),
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `185_Employee` VALUES
(1, "Joe", 70000, 1),
(2, "Henry", 80000, 2),
(3, "Sam", 60000, 2),
(4, "Max", 90000, 1),
(5, "Janet", 69000, 1),
(6, "Randy", 85000, 1);

DROP TABLE IF EXISTS `185_Department`;
CREATE TABLE `185_Department` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `185_Department` VALUES
(1, "IT"),
(2, "Sales");

```
Employee 表包含所有员工信息，每个员工有其对应的 Id, salary 和 department Id 。

+----+-------+--------+--------------+
| Id | Name  | Salary | DepartmentId |
+----+-------+--------+--------------+
| 1  | Joe   | 70000  | 1            |
| 2  | Henry | 80000  | 2            |
| 3  | Sam   | 60000  | 2            |
| 4  | Max   | 90000  | 1            |
| 5  | Janet | 69000  | 1            |
| 6  | Randy | 85000  | 1            |
+----+-------+--------+--------------+
Department 表包含公司所有部门的信息。

+----+----------+
| Id | Name     |
+----+----------+
| 1  | IT       |
| 2  | Sales    |
+----+----------+
编写一个 SQL 查询，找出每个部门工资前三高的员工。例如，根据上述给定的表格，查询结果应返回：

+------------+----------+--------+
| Department | Employee | Salary |
+------------+----------+--------+
| IT         | Max      | 90000  |
| IT         | Randy    | 85000  |
| IT         | Joe      | 70000  |
| Sales      | Henry    | 80000  |
| Sales      | Sam      | 60000  |
+------------+----------+--------+
```

-- Debug
SELECT t1.Name AS Department, t2.Name AS Employee, t2.Salary AS Salary FROM 185_Department AS t1 INNER JOIN 185_Employee AS t2 ON t1.Id = t2.DepartmentId WHERE (SELECT COUNT(Salary) FROM 185_Employee WHERE DepartmentId = t1.Id AND Salary > t2.Salary) < 3 ORDER BY Department;

-- 题目似乎略有歧义, 前三并列, 第四算不算
SELECT t1.Name AS Department, t2.Name AS Employee, t2.Salary AS Salary FROM 185_Department AS t1 INNER JOIN 185_Employee AS t2 ON t1.Id = t2.DepartmentId WHERE (SELECT COUNT(DISTINCT(Salary)) FROM 185_Employee WHERE DepartmentId = t1.Id AND Salary >= t2.Salary) <= 3 ORDER BY Department;

-- 测试
INSERT INTO `185_Employee` VALUES
(7, "AAA", 90000, 1);
 -- 查询结果如下
+------------+----------+--------+
| Department | Employee | Salary |
+------------+----------+--------+
| IT         | Joe      |  70000 |
| IT         | Max      |  90000 |
| IT         | Randy    |  85000 |
| IT         | AAA      |  90000 |
| Sales      | Henry    |  80000 |
| Sales      | Sam      |  60000 |
+------------+----------+--------+

-- Answer
-- 1276 ms
SELECT t1.Name AS Department, t2.Name AS Employee, t2.Salary AS Salary FROM Department AS t1 INNER JOIN Employee AS t2 ON t1.Id = t2.DepartmentId WHERE (SELECT COUNT(DISTINCT(Salary)) FROM Employee WHERE DepartmentId = t1.Id AND Salary >= t2.Salary) <= 3 ORDER BY Department;
