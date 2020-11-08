import { Credential, SIGNUP_SUCCESS, SET_MESSAGE, SIGNUP_FAIL, SIGNIN_SUCCESS, SIGNIN_FAIL, SIGNOUT } from "../types"
import SignService from "@/services/sign.service"
import { Dispatch } from "redux"
import { useDispatch } from "react-redux"

export const signUp = (data: any) => (dispatch: any) => {
  return SignService.signUp(data)
    .then(data => {
      dispatch({
        type: SIGNUP_SUCCESS
      })
      dispatch({
        type: SET_MESSAGE,
        payload: data
      })
      return Promise.resolve();
    },
      (err) => {
        dispatch({
          type: SIGNUP_FAIL,
        });
        dispatch({
          type: SET_MESSAGE,
          payload: err,
        });
        return Promise.reject();
      })
}

export const signIn = (data: Credential) => (dispatch: Dispatch) => {
  return SignService.signIn(data).then(
    (data) => {
      dispatch({
        type: SIGNIN_SUCCESS,
        payload: data,
      })
      return Promise.resolve()
    },
    (err) => {
      dispatch({
        type: SIGNIN_FAIL,
      });
      dispatch({
        type: SET_MESSAGE,
        payload: err,
      });
      return Promise.reject();
    })
}

export const refresh = (data: any) => (dispatch: any) => {
  return SignService.signIn(data).then(
    (data) => {
      dispatch({
        type: SIGNIN_SUCCESS,
        payload: data,
      })
      return Promise.resolve()
    },
    (err) => {
      dispatch({
        type: SIGNIN_FAIL,
      });
      dispatch({
        type: SET_MESSAGE,
        payload: err,
      });
      return Promise.reject();
    })
}

export const logout = () => (dispatch: any) => {
  SignService.logout();
  dispatch({
    type: SIGNOUT,
  });
};
