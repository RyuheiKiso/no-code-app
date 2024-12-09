// actionTypes.tsx

// ログインリクエストのアクションタイプ
export const LOGIN_REQUEST = 'LOGIN_REQUEST';

// ログイン成功のアクションタイプ
export const LOGIN_SUCCESS = 'LOGIN_SUCCESS';

// ログイン失敗のアクションタイプ
export const LOGIN_FAILURE = 'LOGIN_FAILURE';

/*
使用例:

import { LOGIN_REQUEST, LOGIN_SUCCESS, LOGIN_FAILURE } from './actionTypes';

// アクションクリエーターの例
const loginRequest = () => ({
  type: LOGIN_REQUEST,
});

const loginSuccess = (user) => ({
  type: LOGIN_SUCCESS,
  payload: user,
});

const loginFailure = (error) => ({
  type: LOGIN_FAILURE,
  payload: error,
});
*/