

export const SET_MESSAGE = "SET_MESSAGE";
export const CLEAR_MESSAGE = "CLEAR_MESSAGE";


export interface Credential {
  username: string
  password: string
}

export interface CreateCustomer {
  username: string
  password: string
  repeatPassword: string
  firstname: string
}
export interface Tokens {
  accessToken?: string
  refreshToken: string
}
export interface ErrorText {
  error: string
}

export interface SignState {
  accessToken?: string
  refreshToken?: string
  isLoggedIn: boolean
}
export const SIGNUP_SUCCESS = "SIGNUP_SUCCESS";
export const SIGNUP_FAIL = "SIGNUP_FAIL";
export const SIGNIN_SUCCESS = "SIGNIN_SUCCESS";
export const SIGNIN_FAIL = "SIGNIN_FAIL";
export const SIGNREFRESH_SUCCESS = "SIGNREFRESH_SUCCESS";
export const SIGNREFRESH_FAIL = "SIGNREFRESH_FAIL";
export const SIGNOUT = "SIGNOUT";

interface SignUpSuccessAction {
  type: typeof SIGNUP_SUCCESS,
}
interface SignUpFailAction {
  type: typeof SIGNUP_FAIL,
  payload: ErrorText
}
interface SignInSuccessAction {
  type: typeof SIGNIN_SUCCESS,
  payload: Tokens
}
interface SignInFailAction {
  type: typeof SIGNIN_FAIL,
  payload: ErrorText
}
interface SignRefreshSuccessAction {
  type: typeof SIGNREFRESH_SUCCESS,
  payload: Tokens
}
interface SignRefreshFailAction {
  type: typeof SIGNREFRESH_FAIL,
  payload: ErrorText
}
interface SignOutAction {
  type: typeof SIGNOUT,
}

export type SignActionTypes =
  SignUpSuccessAction |
  SignUpFailAction |
  SignInSuccessAction |
  SignInFailAction |
  SignRefreshSuccessAction |
  SignRefreshFailAction |
  SignOutAction