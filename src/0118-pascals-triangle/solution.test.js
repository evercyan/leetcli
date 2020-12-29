
const solution = require('./solution.js')

let cases = [{
    "inputs": [
        5
    ],
    "expects": [
        [
            [1],
            [1, 1],
            [1, 2, 1],
            [1, 3, 3, 1],
            [1, 4, 6, 4, 1],
        ],
    ],
}];

cases.forEach(function(item, i) {
    test('test-' + i, () => {
        let ret = solution(item['inputs'][0])
        let expected = item['expects'][0]
        expect(ret).toStrictEqual(expected)
    })
});
