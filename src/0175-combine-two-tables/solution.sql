-- sql
CREATE DATABASE IF NOT EXISTS `Leetcode`;

USE `Leetcode`;

DROP TABLE IF EXISTS `175_Person`;
CREATE TABLE `175_Person` (
  `PersonId` int(11) NOT NULL AUTO_INCREMENT,
  `FirstName` varchar(255) NOT NULL,
  `LastName` varchar(255) NOT NULL,
  PRIMARY KEY (`PersonId`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `175_Person` VALUES
(1, 'FristName1', 'LastName1'),
(2, 'FristName2', 'LastName2'),
(3, 'FristName3', 'LastName3'),
(4, 'FristName4', 'LastName4');

DROP TABLE IF EXISTS `175_Address`;
CREATE TABLE `175_Address` (
  `AddressId` int(11) NOT NULL AUTO_INCREMENT,
  `PersonId` int(11) NOT NULL,
  `City` varchar(255) NOT NULL,
  `State` varchar(255) NOT NULL,
  PRIMARY KEY (`AddressId`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `175_Address` VALUES
(1, 1, 'Shanghai', 'SH'),
(2, 2, 'Beijing', 'BJ'),
(3, 2, 'Shenzhen', 'SZ');

-- 编写一个 SQL 查询, 满足条件: 无论 person 是否有地址信息，都需要基于上述两表提供 person 的以下信息:
-- FirstName, LastName, City, State

-- Debug
SELECT t1.FirstName, t1.LastName, t2.City, t2.State FROM 175_Person AS t1 left join 175_Address AS t2 on t1.PersonId = t2.PersonId;

-- Answer
SELECT t1.FirstName, t1.LastName, t2.City, t2.State FROM Person AS t1 left join Address AS t2 on t1.PersonId = t2.PersonId;
