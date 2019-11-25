/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function(s) {
    // 去掉最右空格
    s = s.replace(/(\s*$)/g, '');
    if (s == '') {
        return 0;
    }
    var s_list = s.split(' ');
    return s_list[s_list.length - 1].length;
}
