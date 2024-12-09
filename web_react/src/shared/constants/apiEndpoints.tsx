// apiEndpoints.tsx

// APIのベースURL
export const API_BASE_URL = 'https://api.example.com';

// ログインエンドポイント
export const LOGIN_ENDPOINT = `${API_BASE_URL}/auth/login`;

// 登録エンドポイント
export const REGISTER_ENDPOINT = `${API_BASE_URL}/auth/register`;

// ユーザープロフィールエンドポイント
export const USER_PROFILE_ENDPOINT = `${API_BASE_URL}/user/profile`;

// 使用例:
// fetch(LOGIN_ENDPOINT, { method: 'POST', body: JSON.stringify({ username, password }) })
//   .then(response => response.json())
//   .then(data => console.log(data));