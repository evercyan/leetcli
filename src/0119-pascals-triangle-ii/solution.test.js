
const solution = require('./solution.js')

let cases = [{
    "inputs": [
        3
    ],
    "expects": [
        [1, 3, 3, 1],
    ],
}];

cases.forEach(function(item, i) {
    test('test-' + i, () => {
        let ret = solution(item['inputs'][0])
        let expected = item['expects'][0]
        expect(ret).toStrictEqual(expected)
    })
});
