import { combineReducers } from "redux";
import auth from './sign'
import message from './message'
export default combineReducers({
  auth,
  message
})