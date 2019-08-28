-- Sql
CREATE DATABASE `Leetcode`;
USE `Leetcode`;

DROP TABLE IF EXISTS `627_Salary`;
CREATE TABLE `627_Salary` (
  `id` int(4) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `sex` varchar(1) NOT NULL,
  `salary` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `627_Salary` VALUES
    (1,'A','m',2500),
    (2,'B','f',1500),
    (3,'C','m',5500),
    (4,'D','f',500);

-- 交换所有的 f 和 m 值

-- Debug
UPDATE 627_Salary set sex = (CASE sex WHEN 'm' THEN 'f' ELSE 'm' END);

-- Answer
UPDATE salary set sex = (CASE sex WHEN 'm' THEN 'f' ELSE 'm' END);
