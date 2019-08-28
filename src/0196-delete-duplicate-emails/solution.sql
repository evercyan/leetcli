-- Sql
CREATE DATABASE IF NOT EXISTS `Leetcode`;
USE `Leetcode`;

DROP TABLE IF EXISTS `196_Person`;
CREATE TABLE `196_Person` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Email` varchar(255) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `196_Person` VALUES
(1, "john@example.com"),
(2, "bob@example.com"),
(3, "john@example.com");

```
编写一个 SQL 查询，来删除 Person 表中所有重复的电子邮箱，重复的邮箱里只保留 Id 最小 的那个。

+----+------------------+
| Id | Email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
| 3  | john@example.com |
+----+------------------+
Id 是这个表的主键。
例如，在运行你的查询语句之后，上面的 Person 表应返回以下几行:

+----+------------------+
| Id | Email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
+----+------------------+
```

-- Debug
-- 多表联系删除, 可以可以, 学习到了
DELETE t1 FROM 196_Person t1, 196_Person t2 WHERE t1.Email = t2.Email AND t1.Id > t2.Id;


-- Answer
-- 1348 ms
DELETE t1 FROM Person t1, Person t2 WHERE t1.Email = t2.Email AND t1.Id > t2.Id;
