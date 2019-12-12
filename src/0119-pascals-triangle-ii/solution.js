/**
 * @param {number} rowIndex
 * @return {number[]}
 */
var getRow = function (rowIndex) {
    if (rowIndex == 0) {
        return [1];
    }
    // 保存上一行的结果
    var lastRow = [1, 1];
    // 从第 2 行开始循环行
    for (var i = 2; i <= rowIndex; i++) {
        // 每一行保存新的结果
        var result = [];
        // 循环列
        for (var j = 0; j <= i; j++) {
            if (j == 0 || j == i) {
                result.push(1);
            } else {
                result.push(lastRow[j - 1] + lastRow[j]);
            }
        }
        // 将每一行得到的结果赋值给 lastRow, 预备给下一行使用
        lastRow = result;
    }
    // 如果循环结果, 没有下一行, lastRow 就是最后一行
    return lastRow;
};
