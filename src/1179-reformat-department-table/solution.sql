# Write your MySQL query statement below
-- sql
CREATE DATABASE IF NOT EXISTS `Leetcode`;

USE `Leetcode`;

DROP TABLE IF EXISTS `1179_Department`;
CREATE TABLE `1179_Department` (
  `id` int(11) NOT NULL,
  `revenue` int(11) NOT NULL,
  `month` varchar(255) NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `1179_Department` VALUES
(1, 8000, 'Jan'),
(2, 9000, 'Jan'),
(3, 10000, 'Feb'),
(1, 7000, 'Feb'),
(1, 6000, 'Mar');

-- 编写一个 SQL 查询来重新格式化表, 使得新的表中有一个部门 id 列和一些对应 每个月 的收入（revenue）列

-- Debug
SELECT
	id,
	SUM( CASE MONTH WHEN 'Jan' THEN revenue ELSE NULL END ) AS 'Jan_Revenue',
	SUM( CASE MONTH WHEN 'Feb' THEN revenue ELSE NULL END ) AS 'Feb_Revenue',
	SUM( CASE MONTH WHEN 'Mar' THEN revenue ELSE NULL END ) AS 'Mar_Revenue',
	SUM( CASE MONTH WHEN 'Apr' THEN revenue ELSE NULL END ) AS 'Apr_Revenue',
	SUM( CASE MONTH WHEN 'May' THEN revenue ELSE NULL END ) AS 'May_Revenue',
	SUM( CASE MONTH WHEN 'Jun' THEN revenue ELSE NULL END ) AS 'Jun_Revenue',
	SUM( CASE MONTH WHEN 'Jul' THEN revenue ELSE NULL END ) AS 'Jul_Revenue',
	SUM( CASE MONTH WHEN 'Aug' THEN revenue ELSE NULL END ) AS 'Aug_Revenue',
	SUM( CASE MONTH WHEN 'Sep' THEN revenue ELSE NULL END ) AS 'Sep_Revenue',
	SUM( CASE MONTH WHEN 'Oct' THEN revenue ELSE NULL END ) AS 'Oct_Revenue',
	SUM( CASE MONTH WHEN 'Nov' THEN revenue ELSE NULL END ) AS 'Nov_Revenue',
	SUM( CASE MONTH WHEN 'Dec' THEN revenue ELSE NULL END ) AS 'Dec_Revenue' 
FROM
	1179_Department 
GROUP BY
	id 
ORDER BY
	id ASC;

-- Answer
SELECT
	id,
	SUM( CASE MONTH WHEN 'Jan' THEN revenue ELSE NULL END ) AS 'Jan_Revenue',
	SUM( CASE MONTH WHEN 'Feb' THEN revenue ELSE NULL END ) AS 'Feb_Revenue',
	SUM( CASE MONTH WHEN 'Mar' THEN revenue ELSE NULL END ) AS 'Mar_Revenue',
	SUM( CASE MONTH WHEN 'Apr' THEN revenue ELSE NULL END ) AS 'Apr_Revenue',
	SUM( CASE MONTH WHEN 'May' THEN revenue ELSE NULL END ) AS 'May_Revenue',
	SUM( CASE MONTH WHEN 'Jun' THEN revenue ELSE NULL END ) AS 'Jun_Revenue',
	SUM( CASE MONTH WHEN 'Jul' THEN revenue ELSE NULL END ) AS 'Jul_Revenue',
	SUM( CASE MONTH WHEN 'Aug' THEN revenue ELSE NULL END ) AS 'Aug_Revenue',
	SUM( CASE MONTH WHEN 'Sep' THEN revenue ELSE NULL END ) AS 'Sep_Revenue',
	SUM( CASE MONTH WHEN 'Oct' THEN revenue ELSE NULL END ) AS 'Oct_Revenue',
	SUM( CASE MONTH WHEN 'Nov' THEN revenue ELSE NULL END ) AS 'Nov_Revenue',
	SUM( CASE MONTH WHEN 'Dec' THEN revenue ELSE NULL END ) AS 'Dec_Revenue' 
FROM
	Department 
GROUP BY
	id 
ORDER BY
	id ASC;
