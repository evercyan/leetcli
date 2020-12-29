

const solution = require('./solution.js')

let cases = [
    {
        "inputs": [
            "10", "9"
        ],
        "expects": [
           	"19"
        ],
    }
];

cases.forEach(function(item, i) {
    test('test-' + i, () => {
		let ret = solution(item['inputs'][0], item['inputs'][1])
		let expected = item['expects'][0]
        expect(ret).toBe(expected)
    })
});