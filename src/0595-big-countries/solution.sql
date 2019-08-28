-- Sql
CREATE DATABASE IF NOT EXISTS `Leetcode`;
USE `Leetcode`;

DROP TABLE IF EXISTS `595_World`;
CREATE TABLE `595_World` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `continent` varchar(255) NOT NULL,
  `area` int(11) NOT NULL,
  `population` int(11) NOT NULL,
  `gdp` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `595_World` VALUES
(1, "Afghanistan", "Asia", 652230, 25500100, 20343000),
(2, "Albania", "Europe", 28748, 2831741, 12960000),
(3, "Algeria", "Africa", 2381741, 37100000, 188681000),
(4, "Andorra", "Europe", 468, 78115, 3712000),
(5, "Angola", "Africa", 1246700, 20609294, 100990000);


```
这里有张 World 表

+-----------------+------------+------------+--------------+---------------+
| name            | continent  | area       | population   | gdp           |
+-----------------+------------+------------+--------------+---------------+
| Afghanistan     | Asia       | 652230     | 25500100     | 20343000      |
| Albania         | Europe     | 28748      | 2831741      | 12960000      |
| Algeria         | Africa     | 2381741    | 37100000     | 188681000     |
| Andorra         | Europe     | 468        | 78115        | 3712000       |
| Angola          | Africa     | 1246700    | 20609294     | 100990000     |
+-----------------+------------+------------+--------------+---------------+
如果一个国家的面积超过300万平方公里，或者人口超过2500万，那么这个国家就是大国家。

编写一个SQL查询，输出表中所有大国家的名称、人口和面积。

例如，根据上表，我们应该输出:

+--------------+-------------+--------------+
| name         | population  | area         |
+--------------+-------------+--------------+
| Afghanistan  | 25500100    | 652230       |
| Algeria      | 37100000    | 2381741      |
+--------------+-------------+--------------+
```

-- Debug
SELECT name, population, area FROM 595_World WHERE area > 3000000 OR population > 25000000;
SELECT name, population, area FROM 595_World WHERE area > 3000000 UNION SELECT name, population, area FROM 595_World WHERE population > 25000000;

-- Answer
-- 3669 ms
SELECT name, population, area FROM World WHERE area > 3000000 OR population > 25000000;
--  3340 ms
SELECT name, population, area FROM World WHERE area > 3000000 UNION SELECT name, population, area FROM World WHERE population > 25000000;
