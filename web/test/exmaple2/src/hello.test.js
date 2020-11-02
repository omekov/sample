import React from 'react'
import { render, unmountComponentAtNode } from 'react-dom'
import { act } from 'react-dom/test-utils'
import Hello from './hello'

let container = null
beforeEach(() => {
    container = document.createElement("div")
    document.body.appendChild(container)
})

afterEach(() => {
    unmountComponentAtNode(container)
    container.remove()
    container = null
})

it('render with or without a name', () => {
    act(() => {
        render(<Hello/>, container)
    })
    expect(container.textContent).toBe("Hey, stranger");
    act(() => {
        render(<Hello name="Jenny" />, container)
    })
    expect(container.textContent).toBe("Hello, Jenny!") 
})
