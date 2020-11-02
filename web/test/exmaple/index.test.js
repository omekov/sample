const { sum, nativeNull } = require('./index')

describe('test case 1', () => {
    test('Sum should return sum of two values', () => {
        expect(sum(4, 5)).toBe(9)
        expect(sum(1, 1)).toEqual(2)
    })

    test('Sum should return value correctly cparing to other', () => {
        expect(sum(3, 2)).toBeGreaterThan(4)
        expect(sum(6, 3)).toBeGreaterThanOrEqual(9)
        expect(sum(9, 1)).toBeLessThan(11)
        expect(sum(4, 4)).toBeLessThanOrEqual(8)
    })

    test('Sum should sum 2float values correctly', () => {
        expect(sum(0.1, 0.2)).toBeCloseTo(0.3)
    })

})

describe('Native null function', () => {
    test('should return false value null', () => {
        expect(nativeNull()).toBe(null)
        expect(nativeNull()).toBeNull()
        expect(nativeNull()).toBeFalsy()
        expect(nativeNull()).toBeDefined()
        expect(nativeNull()).not.toBeTruthy()
    })
})