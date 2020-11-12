import { Dispatch } from "redux"
import axios, { AxiosResponse } from "axios"
import { Tokens, Credential, SET_MESSAGE, SIGNUP_FAIL, SIGNUP_SUCCESS, SIGNIN_FAIL, SIGNIN_SUCCESS, SIGNREFRESH_SUCCESS, SIGNREFRESH_FAIL, SIGNOUT } from '@/redux/types'
import SignService from "@/services/sign.service"
import { handlerError } from "@/services/request.service"

const SIGNUP_SUCCESS_TEXT = 'Регистрация прошла успешно'
const SIGNIN_SUCCESS_TEXT = 'Авторизация прошла успешно'
const REFRESH_SUCCESS_TEXT = 'Токен успешно обновлен'

export const signUp = async (data: any) => (dispatch: Dispatch) => {
  return SignService.signUp(data)
    .then((response: AxiosResponse) => {
      if (response.status == 201) {
        dispatch({
          type: SIGNUP_SUCCESS
        })
        dispatch({
          type: SET_MESSAGE,
          payload: SIGNUP_SUCCESS_TEXT
        })
        return Promise.resolve()
      }
    }, (error) => handlerError(error, SIGNUP_FAIL, dispatch))
}

export const signIn = async (data: Credential) => (dispatch: Dispatch) =>
  SignService.signIn(data)
    .then((response) => {
      if (response.status == 200) {
        dispatch({
          type: SIGNIN_SUCCESS,
          payload: { customer: response.data }
        })
        const headerValue = `Bearer ${response.data.accessToken}`
        axios.defaults.headers.common['Authorization'] = headerValue
        localStorage.setItem("access_token", JSON.stringify(response.data.accessToken));
        localStorage.setItem("refresh_token", JSON.stringify(response.data.refreshToken));
        return Promise.resolve()
      }
    }, (error) => handlerError(error, SIGNIN_FAIL, dispatch))


export const refresh = (data: Tokens) => async (dispatch: Dispatch) => {
  return SignService.refresh(data)
    .then((response: AxiosResponse<Tokens>) => {
      dispatch({
        type: SIGNREFRESH_SUCCESS,
        payload: response.data,
      })
      return Promise.resolve()
    },
      (error) => handlerError(error, SIGNREFRESH_FAIL, dispatch)
    )
}

export const logout = () => async (dispatch: Dispatch) => {
  SignService.logout();
  dispatch({
    type: SIGNOUT,
  });
};
