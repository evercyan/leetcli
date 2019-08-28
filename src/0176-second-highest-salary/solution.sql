-- Sql
CREATE DATABASE IF NOT EXISTS `Leetcode`;
USE `Leetcode`;

DROP TABLE IF EXISTS `176_Employee`;
CREATE TABLE `176_Employee` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Salary` int(11) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `176_Employee` VALUES
(1, 100),
(2, 200),
(3, 200),
(4, 300);

-- 编写一个 SQL 查询, 获取 Employee 表中第二高的薪水(Salary)

-- Debug
SELECT MAX(Salary) AS SecondHighestSalary FROM 176_Employee WHERE Salary < (SELECT MAX(Salary) FROM 176_Employee);
SELECT IFNULL((SELECT DISTINCT Salary AS SecondHighSalary FROM 176_Employee ORDER BY Salary DESC LIMIT 1, 1), NULL);

-- Answer
SELECT MAX(Salary) AS SecondHighestSalary FROM Employee WHERE Salary < (SELECT MAX(Salary) FROM Employee);
