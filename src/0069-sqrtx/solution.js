/**
 * @param {number} x
 * @return {number}
 */
var mySqrt = function(x) {
    var low = 0,
        high = x;
    while (low <= high) {
        var mid = parseInt((low + high) / 2);
        if (mid * mid == x) {
            return mid;
        } else if (mid * mid < x) {
            low = mid + 1
        } else {
            high = mid - 1;
        }
    }
    return high;
};
