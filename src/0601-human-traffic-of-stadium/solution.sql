-- Sql
CREATE DATABASE `Leetcode`;
USE `Leetcode`;

DROP TABLE IF EXISTS `601_stadium`;
CREATE TABLE `601_stadium` (
  `id` int(4) NOT NULL AUTO_INCREMENT,
  `date` varchar(100) NOT NULL,
  `people` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

INSERT INTO `601_stadium` VALUES
    (1,'2017-01-01',10),
    (2,'2017-01-02',109),
    (3,'2017-01-03',150),
    (4,'2017-01-04',99),
    (5,'2017-01-05',145),
    (6,'2017-01-06',1455),
    (7,'2017-01-07',199),
    (8,'2017-01-08',188);

-- 找出高峰期时段, 要求连续三天及以上, 并且每天人流量均不少于100.

-- Debug
SELECT s1.* FROM 601_stadium AS s1, 601_stadium AS s2, 601_stadium as s3
WHERE
(
    (s1.id + 1 = s2.id AND s1.id + 2 = s3.id)
    OR
    (s1.id - 1 = s2.id AND s1.id + 1 = s3.id)
    OR
    (s1.id - 2 = s2.id AND s1.id - 1 = s3.id)
)
AND s1.people >= 100
AND s2.people >= 100
AND s3.people >= 100
GROUP BY s1.id


-- Answer
SELECT s1.* FROM stadium AS s1, stadium AS s2, stadium as s3
WHERE
(
    (s1.id + 1 = s2.id AND s1.id + 2 = s3.id)
    OR
    (s1.id - 1 = s2.id AND s1.id + 1 = s3.id)
    OR
    (s1.id - 2 = s2.id AND s1.id - 1 = s3.id)
)
AND s1.people >= 100
AND s2.people >= 100
AND s3.people >= 100
GROUP BY s1.id
