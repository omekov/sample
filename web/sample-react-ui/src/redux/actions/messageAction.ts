import { SET_MESSAGE, CLEAR_MESSAGE } from "@/redux/types";

export const setMessage = (message: any) => ({
  type: SET_MESSAGE,
  payload: message,
});

export const clearMessage = () => ({
  type: CLEAR_MESSAGE,
});
