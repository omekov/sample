import axios, { AxiosRequestConfig, AxiosResponse } from 'axios'
import { Dispatch } from 'redux'
import { SET_MESSAGE } from '@/redux/types'

export const NewRequest = ({ url, method, data }: AxiosRequestConfig): Promise<AxiosResponse> => {
  const headers = { 'Content-Type': 'application/json' }
  if (method == 'GET') {
    return axios.request({ url, method, headers })
  } else {
    return axios.request({ url, method, data, headers })
  }
}

export const handlerError = (error: any, type: string, dispatch: Dispatch) => {
  const message =
    (error.response &&
      error.response.data &&
      error.response.data.error)
  error.toString();
  dispatch({
    type: type,
  });

  dispatch({
    type: SET_MESSAGE,
    payload: message,
  });

  return Promise.reject();
}