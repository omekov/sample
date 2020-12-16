import thunk from 'redux-thunk'
import { createStore, applyMiddleware, Store } from 'redux'
import { composeWithDevTools } from "redux-devtools-extension"
import rootReducer from './reducers'
import { AppState } from './types'
const middleware = [thunk]


const store: Store<AppState> = createStore(
  rootReducer,
  composeWithDevTools(applyMiddleware(...middleware))
)


export default store