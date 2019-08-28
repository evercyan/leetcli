-- Sql
CREATE DATABASE IF NOT EXISTS `Leetcode`;
USE `Leetcode`;

DROP TABLE IF EXISTS `183_Customers`;
CREATE TABLE `183_Customers` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `183_Customers` VALUES
(1, "Joe"),
(2, "Henry"),
(3, "Sam"),
(4, "Max");

DROP TABLE IF EXISTS `183_Orders`;
CREATE TABLE `183_Orders` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `CustomerId` int(11) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `183_Orders` VALUES
(1, 3),
(2, 1);


-- 找出所有从不订购任何东西的客户

-- Debug
SELECT t1.Name AS Customers FROM 183_Customers AS t1 LEFT JOIN 183_Orders AS t2 ON t1.Id = t2.CustomerId WHERE t2.CustomerId IS NULL;

SELECT Name AS Customers FROM 183_Customers WHERE Id NOT IN(SELECT CustomerId FROM 183_Orders);

-- Answer

-- 937 ms
SELECT t1.Name AS Customers FROM Customers AS t1 LEFT JOIN Orders AS t2 ON t1.Id = t2.CustomerId WHERE t2.CustomerId IS NULL;

-- 538 ms
SELECT Name AS Customers FROM Customers WHERE Id NOT IN(SELECT CustomerId FROM Orders);
