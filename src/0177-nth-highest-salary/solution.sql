-- Sql
CREATE DATABASE IF NOT EXISTS `Leetcode`;
USE `Leetcode`;

DROP TABLE IF EXISTS `177_Employee`;
CREATE TABLE `177_Employee` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Salary` int(11) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `177_Employee` VALUES
(1, 100),
(2, 200),
(3, 200),
(4, 400),
(5, 300),
(6, 600),
(7, 900);

-- 编写一个 SQL 查询, 获取 Employee 表中第 n 高的薪水(Salary).

-- Debug
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    DECLARE M INT DEFAULT 0;
    SET M = N - 1;
    RETURN(M);
END
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    DECLARE M INT DEFAULT 0;
    SET M = N - 1;
    RETURN(SELECT IFNULL((SELECT DISTINCT Salary AS SecondHighSalary FROM 177_Employee ORDER BY Salary DESC LIMIT M, 1), NULL));
END

-- Answer
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    DECLARE M INT DEFAULT 0;
    SET M = N - 1;
    RETURN(SELECT IFNULL((SELECT DISTINCT Salary AS SecondHighSalary FROM Employee ORDER BY Salary DESC LIMIT M, 1), NULL));
END
