function expect(value) {
    return {
        toBe: exp => {
            if (value === exp) {
                console.log('Sucess')
            } else {
                console.error(`Value is ${value}, but expectattion is ${exp}`)
            }
        }
    }
}

const sum = (a, b) => a + b

const nativeNull = () => null

console.log(sum(41, 1))

// expect(sum(41, 1)).toBe(41)

module.exports = { sum, nativeNull }