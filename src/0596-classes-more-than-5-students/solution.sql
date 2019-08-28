-- Sql
CREATE DATABASE IF NOT EXISTS `Leetcode`;
USE `Leetcode`;

DROP TABLE IF EXISTS `596_courses`;
CREATE TABLE `596_courses` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `student` varchar(255) NOT NULL,
  `class` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `596_courses` VALUES
(1, "A", "Math"),
(2, "B", "English"),
(3, "C", "Math"),
(4, "D", "Biology"),
(5, "E", "Math"),
(6, "F", "Computer"),
(7, "G", "Math"),
(8, "H", "Math"),
(9, "I", "Math");


```
有一个 courses 表 ，有: student (学生) 和 class (课程)。

请列出所有超过或等于5名学生的课。

例如,表:

+---------+------------+
| student | class      |
+---------+------------+
| A       | Math       |
| B       | English    |
| C       | Math       |
| D       | Biology    |
| E       | Math       |
| F       | Computer   |
| G       | Math       |
| H       | Math       |
| I       | Math       |
+---------+------------+
应该输出:

+---------+
| class   |
+---------+
| Math    |
+---------+
Note:
学生在每个课中不应被重复计算。
```

-- Debug
-- COUNT(DISTINCT(student) 学生去重
SELECT class FROM 596_courses GROUP BY class HAVING COUNT(DISTINCT(student)) >= 5;

-- Answer
SELECT class FROM courses GROUP BY class HAVING COUNT(DISTINCT(student)) >= 5;
