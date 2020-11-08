import {
  SIGNUP_SUCCESS,
  SIGNUP_FAIL,
  SIGNIN_SUCCESS,
  SIGNIN_FAIL,
  SIGNOUT,
  SignState,
  SignActionTypes,
} from '../types'

const accessToken = localStorage.getItem('access_token')
const refreshToken = localStorage.getItem('refresh_token')

const initialState = accessToken
  ? { isLoggedIn: true, accessToken, refreshToken }
  : { isLoggedIn: false, accessToken: null, refreshToken: null }

export default function (state = initialState, action: any) {
  const { type, payload } = action;

  switch (type) {
    case SIGNUP_SUCCESS:
      return {
        ...state,
        isLoggedIn: false,
      }
    case SIGNUP_FAIL:
      return {
        ...state,
        isLoggedIn: false,
      }
    case SIGNIN_SUCCESS:
      return {
        ...state,
        isLoggedIn: true,
        accessToken: payload.accessToken,
        refreshToken: payload.refreshToken,
      }
    case SIGNIN_FAIL:
      return {
        ...state,
        isLoggedIn: false,
        accessToken: null,
        refreshtoken: null,
      }
    case SIGNOUT:
      return {
        ...state,
        isLoggedIn: false,
        accessToken: null,
        refreshtoken: null,
      }
    default:
      return state
  }
}