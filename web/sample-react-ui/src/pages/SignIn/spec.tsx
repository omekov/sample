import React from 'react'
import { shallow } from 'enzyme'
import SignIn, { Props } from '.'
import { cleanup } from '@testing-library/react-hooks'
import { findByTestAttr } from '../../../utils'
import { createMemoryHistory } from 'history'
describe('SignIn Component', () => {
  let component: any
  beforeEach(() => {
    component = (props: Props) => shallow(<SignIn {...props} />)
    component = component({ history: createMemoryHistory() })
  })
  afterEach(cleanup)
  it("should render", () => {
    const element = findByTestAttr(component, 'SignInComponent')
    expect(element.length).toBe(1)
  })
})

