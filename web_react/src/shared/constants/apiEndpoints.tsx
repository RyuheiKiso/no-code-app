// apiEndpoints.tsx

// APIのベースURL
export const API_BASE_URL = process.env.REACT_APP_API_BASE_URL || 'https://api.example.com';

// 認証エンドポイント
export const AUTH_ENDPOINTS = {
  LOGIN: `${API_BASE_URL}/auth/login`,
  REGISTER: `${API_BASE_URL}/auth/register`,
};

// ユーザープロフィールエンドポイント
export const USER_PROFILE_ENDPOINT = `${API_BASE_URL}/user/profile`;

// 使用例:
// fetch(AUTH_ENDPOINTS.LOGIN, { method: 'POST', body: JSON.stringify({ username, password }) })
//   .then(response => response.json())
//   .then(data => console.log(data));