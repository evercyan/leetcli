/**
 * @param {string} s
 * @return {number}
 */
var firstUniqChar = function(s) {
    var s_map = {};
    for (var i = 0; i < s.length; i++) {
        if (typeof(s_map[s[i]]) == 'undefined') {
            s_map[s[i]] = 0;
        }
        s_map[s[i]]++;
    }
    for (var i = 0; i < s.length; i++) {
        if (s_map[s[i]] == 1) {
            return i;
        }
    }
    return -1;
};

module.exports = firstUniqChar;