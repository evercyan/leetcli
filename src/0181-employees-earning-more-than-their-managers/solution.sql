-- Sql
CREATE DATABASE IF NOT EXISTS `Leetcode`;
USE `Leetcode`;

DROP TABLE IF EXISTS `181_Employee`;
CREATE TABLE `181_Employee` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) NOT NULL,
  `Salary` int(11) NOT NULL,
  `ManagerId` int(11),
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `181_Employee` VALUES
(1, "Joe", 70000, 3),
(2, "Henry", 80000, 4),
(3, "Sam", 60000, NULL),
(4, "Max", 90000, NULL);


-- 编写一个 SQL 查询, 该查询可以获取收入超过他们经理的员工的姓名.

-- Debug
SELECT Name AS Employee FROM 181_Employee AS t1 WHERE ManagerId > 0 AND t1.Salary > (SELECT salary FROM 181_Employee WHERE Id = t1.ManagerId);
SELECT t1.Name AS Employee FROM 181_Employee AS t1 ,181_Employee AS t2 WHERE t1.ManagerId = t2.Id AND t1.Salary > t2.Salary;

-- Answer
-- 872 ms
SELECT Name AS Employee FROM Employee AS t1 WHERE ManagerId > 0 AND t1.Salary > (SELECT salary FROM Employee WHERE Id = t1.ManagerId);
-- 486 ms
SELECT t1.Name AS Employee FROM Employee AS t1 ,Employee AS t2 WHERE t1.ManagerId = t2.Id AND t1.Salary > t2.Salary;
