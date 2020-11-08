import axios, { AxiosRequestConfig, AxiosResponse } from 'axios'
import { Tokens, Credential } from '@/redux/types'

const API_URL = 'http://localhost:8080'
const SIGNUP_SUCCES_TEXT = 'Регистрация прошла успешно'
const SIGNIN_SUCCES_TEXT = 'Авторизация прошла успешно'
const REFRESH_SUCCES_TEXT = 'Токен успешно обновлен'

const signUp = (data: any) => {
  const reqConfig: AxiosRequestConfig = {
    url: API_URL + "signup",
    method: "POST",
    data
  }
  return axios.request(reqConfig)
    .then((response: AxiosResponse) => {
      if (response.status == 201) {
        return SIGNUP_SUCCES_TEXT
      } else {
        return response.data
      }
    })
}

const signIn = (data: Credential) => {
  const reqConfig: AxiosRequestConfig = {
    url: API_URL + "signin",
    method: "POST",
    data
  }
  return axios.request(reqConfig)
    .then((response: AxiosResponse<Tokens>) => {
      if (response.status == 200) {
        const headerValue = `Bearer ${response.data.accessToken}`
        axios.defaults.headers.common['Authorization'] = headerValue
        localStorage.setItem("access_token", JSON.stringify(response.data.accessToken));
        localStorage.setItem("refresh_token", JSON.stringify(response.data.refreshToken));
        return SIGNIN_SUCCES_TEXT
      } else {
        return response.data
      }
    })
}

const refresh = (data: Tokens) => {
  const reqConfig: AxiosRequestConfig = {
    url: API_URL + "refresh",
    method: "POST",
    data
  }
  return axios.request(reqConfig)
    .then((response: AxiosResponse<Tokens>) => {
      if (response.status == 200) {
        const headerValue = `Bearer ${response.data.accessToken}`
        axios.defaults.headers.common['Authorization'] = headerValue
        localStorage.setItem("access_token", JSON.stringify(response.data.accessToken));
        localStorage.setItem("refresh_token", JSON.stringify(response.data.refreshToken));
        return REFRESH_SUCCES_TEXT
      } else {
        return response.data
      }
    })
}

const logout = () => {
  localStorage.removeItem("access_token");
  localStorage.removeItem("refresh_token");
};


export default {
  signUp,
  signIn,
  refresh,
  logout,
};
