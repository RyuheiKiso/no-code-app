// actionTypes.tsx

// 認証関連のアクションタイプ
export const AUTH = {
  LOGIN_REQUEST: 'LOGIN_REQUEST',
  LOGIN_SUCCESS: 'LOGIN_SUCCESS',
  LOGIN_FAILURE: 'LOGIN_FAILURE',
  LOGOUT: 'LOGOUT',
};

// プロフィール関連のアクションタイプ
export const PROFILE = {
  FETCH_PROFILE_REQUEST: 'FETCH_PROFILE_REQUEST',
  FETCH_PROFILE_SUCCESS: 'FETCH_PROFILE_SUCCESS',
  FETCH_PROFILE_FAILURE: 'FETCH_PROFILE_FAILURE',
};

/*
使用例:

import { AUTH, PROFILE } from './actionTypes';

// アクションクリエーターの例
const loginRequest = () => ({
  type: AUTH.LOGIN_REQUEST,
});

const fetchProfileSuccess = (profile) => ({
  type: PROFILE.FETCH_PROFILE_SUCCESS,
  payload: profile,
});
*/