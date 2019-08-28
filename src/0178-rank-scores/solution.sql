-- Sql
CREATE DATABASE IF NOT EXISTS `Leetcode`;
USE `Leetcode`;

DROP TABLE IF EXISTS `178_Scores`;
CREATE TABLE `178_Scores` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Score` varchar(255) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `178_Scores` VALUES
(1, "3.50"),
(2, "3.65"),
(3, "4.00"),
(4, "3.85"),
(5, "4.00"),
(6, "3.65");

-- 根据上述给定的 Scores 表，你的查询应该返回（按分数从高到低排列）：
+-------+------+
| Score | Rank |
+-------+------+
| 4.00  | 1    |
| 4.00  | 1    |
| 3.85  | 2    |
| 3.65  | 3    |
| 3.65  | 3    |
| 3.50  | 4    |
+-------+------+

-- Debug
SELECT t1.Score, COUNT(t2.Score) AS Rank FROM Scores AS t1, (SELECT DISTINCT Score FROM Scores) AS t2 WHERE t1.Score <= t2.Score GROUP BY t1.Id, t1.Score ORDER BY t1.Score DESC;

SELECT t1.Score, (SELECT COUNT(DISTINCT Score) FROM 178_Scores WHERE Score >= t1.Score) AS RANK FROM 178_Scores AS t1 ORDER BY t1.Score DESC;

-- Answer
SELECT t1.Score, (SELECT COUNT(DISTINCT Score) FROM 178_Scores WHERE Score >= t1.Score) AS RANK FROM 178_Scores AS t1 ORDER BY t1.Score DESC;
