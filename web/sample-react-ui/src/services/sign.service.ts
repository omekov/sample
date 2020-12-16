import { Tokens, Credential, CreateCustomer } from '@/redux/types'
import { NewRequest } from './request.service'
const API_URL = 'http://localhost:9090/'

const signUp = (data: CreateCustomer) => NewRequest({
  url: API_URL + "signup",
  method: "POST",
  data,
})

const signIn = (data: Credential) => NewRequest({
  url: API_URL + "signin",
  method: "POST",
  data
})


const refresh = (data: Tokens) => NewRequest({
  url: API_URL + "refresh",
  method: "POST",
  data
})

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
