const Lodash = require('./sync')
describe('Lodash: compact', () => {
    let _ = new Lodash()
    let array
    // Перед каждом тестом
    beforeEach(() => {
        array = [false, 43, '', 0, true, null, 'hello']
    })
    // После всех тестов
    afterAll(() => {
        _ = new Lodash()
    })
    test('should be defined', () => {
        expect(_.compact).toBeDefined()
        expect(_.compact).not.toBeUndefined()
    })
    test('should remove falsy values from array', () => {
        const result = [43, true, 'hello']
        expect(_.compact(array)).toEqual(result)
    })
    test('should NOT remove falsy values', () => {
        expect(_.compact(array)).not.toContain(false)
        expect(_.compact(array)).not.toContain(0)
        expect(_.compact(array)).not.toContain('')
        expect(_.compact(array)).not.toContain(null)
        expect(_.compact(array)).not.toContain(undefined)
    })
})


describe('Lodash: groupBy', () => {
    let _ = new Lodash()
    test('should be defined', () => {
        expect(_.groupBy).toBeDefined()
        expect(_.groupBy).not.toBeUndefined()
    })
    test('should group array items by Math.floor', () => {
        const array = [2.2, 2.4, 4.2, 3.1]
        const result = {
            2: [2.2, 2.4],
            4: [4.2],
            3: [3.1]
        }
        expect(_.groupBy(array, Math.floor)).toEqual(result)
    })
    test('should group array items by length', () => {
        const array = ['one', 'two', 'three']
        const result = {
            5: ['three'],
            3: ['one', 'two']
        }
        expect(_.groupBy(array, 'length')).toEqual(result)
    })
    test('should NOT return array', () => {
        expect(_.groupBy([], Math.trunc)).not.toBeInstanceOf(Array)
    })
    
    
})