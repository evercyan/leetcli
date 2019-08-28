-- Sql
CREATE DATABASE IF NOT EXISTS `Leetcode`;
USE `Leetcode`;

DROP TABLE IF EXISTS `182_Person`;
CREATE TABLE `182_Person` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Email` varchar(255) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `182_Person` VALUES
(1, "a@b.com"),
(2, "c@d.com"),
(3, "a@b.com");


-- 查找 Person 表中所有重复的电子邮箱。

-- Debug
SELECT Email FROM 182_Person GROUP BY Email HAVING COUNT(Email) > 1;

-- Answer
SELECT Email FROM Person GROUP BY Email HAVING COUNT(Email) > 1;
