

const solution = require('./solution.js')

let cases = [
    {
        "inputs": [
            "leetcode",
        ],
        "expects": [
           	0,
        ],
    },
    {
        "inputs": [
            "loveleetcode",
        ],
        "expects": [
           	2,
        ],
    },
];

cases.forEach(function(item, i) {
    test('test-' + i, () => {
		let ret = solution(item['inputs'][0])
		let expected = item['expects'][0]
        expect(ret).toBe(expected)
    })
});