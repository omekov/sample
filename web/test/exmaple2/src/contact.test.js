import React from "react";
import { render, unmountComponentAtNode } from "react-dom";
import { act } from "react-dom/test-utils";
import Contact from "./contact";

jest.mock("./map", () => {
    return function DummyMap(props) {
        return (
            <div data-testid="map">
                {props.center.lat}:{props.center.long}
            </div>
        )
    }
})

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

it("should render contact information", () => {
    const center = { lat: 0, long: 0 };
    act(() => {
        render(
            <Contact 
                name="Joni Beez"
                email="test@example.com"
                site="http://test.com"
                center={center}
            />,
            container
        )
    })
    expect(container.querySelector("[data-testid='email']").getAttribute("href")).toEqual("mailto:test@example.com")
    expect(container.querySelector('[data-testid="site"]').getAttribute("href")).toEqual("http://test.com")
    expect(container.querySelector('[data-testid="map"]').textContent).toEqual("0:0")
})

