

const solution = require('./solution.js')

let cases = [
    {
        "inputs": [
            "a", "b",
        ],
        "expects": [
           	false,
        ],
    },
    {
        "inputs": [
            "aa", "ab",
        ],
        "expects": [
           	false,
        ],
    },
    {
        "inputs": [
            "aa", "aab",
        ],
        "expects": [
              true,
        ],
    },
];

cases.forEach(function(item, i) {
    test('test-' + i, () => {
		let ret = solution(item['inputs'][0], item['inputs'][1])
		let expected = item['expects'][0]
        expect(ret).toBe(expected)
    })
});