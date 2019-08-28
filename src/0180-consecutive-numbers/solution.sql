-- Sql
CREATE DATABASE IF NOT EXISTS `Leetcode`;
USE `Leetcode`;

DROP TABLE IF EXISTS `180_Logs`;
CREATE TABLE `180_Logs` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Num` int(11) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `180_Logs` VALUES
(1, 1),
(2, 1),
(3, 1),
(4, 2),
(5, 1),
(6, 2),
(7, 2);

-- 编写一个 SQL 查询, 查找所有至少连续出现三次的数字.

-- Debug
SELECT DISTINCT t1.Num ConsecutiveNums FROM 180_Logs AS t1 JOIN 180_Logs AS t2 ON t1.Id = t2.Id - 1 JOIN 180_Logs AS t3 ON t1.Id = t3.Id - 2 WHERE t1.Num = t2.Num AND t1.Num = t3.Num;

-- Answer
SELECT DISTINCT t1.Num ConsecutiveNums FROM Logs AS t1 JOIN Logs AS t2 ON t1.Id = t2.Id - 1 JOIN Logs AS t3 ON t1.Id = t3.Id - 2 WHERE t1.Num = t2.Num AND t1.Num = t3.Num;
