/**
 * @param {string} num1
 * @param {string} num2
 * @return {string}
 */
var addStrings = function(num1, num2) {
    var num1_len = num1.length;
    var num2_len = num2.length;
    for (var i = 0; i < Math.abs(num2_len - num1_len); i++) {
        if (num2_len > num1_len) {
            num1 = '0' + num1;
        } else {
            num2 = '0' + num2;
        }
    }
    var carry = 0;
    var res = '';
    for (var i = num1.length - 1; i >= 0; i--) {
        var tmp = parseInt(num1[i]) + parseInt(num2[i]) + carry;
        carry = tmp >= 10 ? 1 : 0;
        res = tmp % 10 + res;
    }
    if (carry) {
        res = '1' + res;
    }
    return res;
}

module.exports = addStrings;
