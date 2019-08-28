-- Sql
CREATE DATABASE IF NOT EXISTS `Leetcode`;
USE `Leetcode`;

DROP TABLE IF EXISTS `197_Weather`;
CREATE TABLE `197_Weather` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `RecordDate` DATE NOT NULL,
  `Temperature` int(11) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `197_Weather` VALUES
(1, "2015-01-01", 10),
(2, "2015-01-02", 25),
(3, "2015-01-03", 20),
(4, "2015-01-04", 30),
(5, "2015-01-06", 20);

```
给定一个 Weather 表，编写一个 SQL 查询，来查找与之前（昨天的）日期相比温度更高的所有日期的 Id。

+---------+------------------+------------------+
| Id(INT) | RecordDate(DATE) | Temperature(INT) |
+---------+------------------+------------------+
|       1 |       2015-01-01 |               10 |
|       2 |       2015-01-02 |               25 |
|       3 |       2015-01-03 |               20 |
|       4 |       2015-01-04 |               30 |
+---------+------------------+------------------+
例如，根据上述给定的 Weather 表格，返回如下 Id:

+----+
| Id |
+----+
|  2 |
|  4 |
+----+
```

-- Debug
SELECT t1.Id FROM 197_Weather AS t1, 197_Weather AS t2 WHERE t1.Temperature > t2.Temperature AND t1.RecordDate = t2.RecordDate + 1;

SELECT t.Id FROM ( SELECT w.*, CASE WHEN @pre_d IS NULL OR (DATEDIFF(w.RecordDate, @pre_d) != 1) OR @pre_t IS NULL OR @pre_t >= w.Temperature THEN @rank := 0 ELSE @rank := 1 END AS Rank, (@pre_d := w.RecordDate) AS cur_d, (@pre_t := w.Temperature) AS cur_t FROM 197_Weather AS w, (SELECT @rank := 0, @pre_d := NULL, @pre_t := NULL) AS a ORDER BY w.RecordDate ASC) AS t WHERE t.Rank = 1;

-- Answer
-- 以下为可行答案, 暂不理解
SELECT t.Id FROM ( SELECT w.*, CASE WHEN @pre_d IS NULL OR (DATEDIFF(w.RecordDate, @pre_d) != 1) OR @pre_t IS NULL OR @pre_t >= w.Temperature THEN @rank := 0 ELSE @rank := 1 END AS Rank, (@pre_d := w.RecordDate) AS cur_d, (@pre_t := w.Temperature) AS cur_t FROM Weather AS w, (SELECT @rank := 0, @pre_d := NULL, @pre_t := NULL) AS a ORDER BY w.RecordDate ASC) AS t WHERE t.Rank = 1;
