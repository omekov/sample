import { combineReducers } from "redux";
import sign from './sign'
import message from './message'
import { AppState } from "../types";
export default combineReducers<AppState>({
  sign,
  message
})