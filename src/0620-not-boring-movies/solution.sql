-- Sql
CREATE DATABASE IF NOT EXISTS `Leetcode`;
USE `Leetcode`;

DROP TABLE IF EXISTS `620_cinema`;
CREATE TABLE `620_cinema` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `movie` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `rating` float(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `620_cinema` VALUES
(1, "War", "great 3D", 8.9),
(2, "Science", "fiction", 8.5),
(3, "irish", "boring", 6.2),
(4, "Ice song", "Fantacy", 8.6),
(5, "House card", "Interesting", 9.1);


```
某城市开了一家新的电影院，吸引了很多人过来看电影。该电影院特别注意用户体验，专门有个 LED显示板做电影推荐，上面公布着影评和相关电影描述。
作为该电影院的信息部主管，您需要编写一个 SQL查询，找出所有影片描述为非 boring (不无聊) 的并且 id 为奇数 的影片，结果请按等级 rating 排列。

例如，下表 cinema:

+---------+-----------+--------------+-----------+
|   id    | movie     |  description |  rating   |
+---------+-----------+--------------+-----------+
|   1     | War       |   great 3D   |   8.9     |
|   2     | Science   |   fiction    |   8.5     |
|   3     | irish     |   boring     |   6.2     |
|   4     | Ice song  |   Fantacy    |   8.6     |
|   5     | House card|   Interesting|   9.1     |
+---------+-----------+--------------+-----------+
对于上面的例子，则正确的输出是为：

+---------+-----------+--------------+-----------+
|   id    | movie     |  description |  rating   |
+---------+-----------+--------------+-----------+
|   5     | House card|   Interesting|   9.1     |
|   1     | War       |   great 3D   |   8.9     |
+---------+-----------+--------------+-----------+
```

-- Debug
SELECT * FROM 620_cinema WHERE description != "boring" AND id % 2 = 1 ORDER BY rating DESC;

-- Answer
-- 250 ms
SELECT * FROM cinema WHERE description != "boring" AND id % 2 = 1 ORDER BY rating DESC;
