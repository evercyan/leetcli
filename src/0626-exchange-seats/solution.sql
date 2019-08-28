-- Sql
CREATE DATABASE IF NOT EXISTS `Leetcode`;
USE `Leetcode`;

DROP TABLE IF EXISTS `626_seat`;
CREATE TABLE `626_seat` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `student` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `626_seat` VALUES
(1, "Abbot"),
(2, "Doris"),
(3, "Emerson"),
(4, "Green"),
(5, "Jeames");


```
小美是一所中学的信息科技老师，她有一张 seat 座位表，平时用来储存学生名字和与他们相对应的座位 id。
其中纵列的 id 是连续递增的
小美想改变相邻俩学生的座位。
你能不能帮她写一个 SQL query 来输出小美想要的结果呢？

示例：

+---------+---------+
|    id   | student |
+---------+---------+
|    1    | Abbot   |
|    2    | Doris   |
|    3    | Emerson |
|    4    | Green   |
|    5    | Jeames  |
+---------+---------+
假如数据输入的是上表，则输出结果如下：

+---------+---------+
|    id   | student |
+---------+---------+
|    1    | Doris   |
|    2    | Abbot   |
|    3    | Green   |
|    4    | Emerson |
|    5    | Jeames  |
+---------+---------+
注意：

如果学生人数是奇数，则不需要改变最后一个同学的座位。
```

-- Debug
SELECT (CASE WHEN id%2=0 THEN id-1 WHEN id != maxid THEN id + 1 ELSE id END) AS id, student FROM 626_seat, (SELECT max(id) AS maxid FROM 626_seat) AS max ORDER BY id ASC;

-- Answer
-- 1004 ms
SELECT (CASE WHEN id%2=0 THEN id-1 WHEN id != maxid THEN id + 1 ELSE id END) AS id, student FROM seat, (SELECT max(id) AS maxid FROM seat) AS t ORDER BY id ASC;
