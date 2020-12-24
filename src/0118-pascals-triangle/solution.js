/**
 * @param {number} numRows
 * @return {number[][]}
 */
var generate = function(numRows) {
    var result = [];
    if (!numRows) {
        return result;
    }
    // 循环每一行
    for (var i = 0; i < numRows; i++) {
        var row = [];
        // 循环每一列(注意, j <= i, 第 1 行 1 个元素, 第 2 行 2 个元素...)
        for (var j = 0; j <= i; j++) {
            if (j == 0 || j == i) {
                // 如果是第一个元素 或者 最后一个元素, 边界是 1
                row.push(1);
            } else {
                // 注意加法规则
                // (i,j)(当前元素) = (i-1,j-1)(上一行当前位置 - 1) + (i-1,j)(上一行当前位置)
                row.push(result[i - 1][j - 1] + result[i - 1][j]);
            }
        }
        // 每一行结束, 将当前行结果扔入数组
        result.push(row);
    }
    return result;
}

module.exports = generate;