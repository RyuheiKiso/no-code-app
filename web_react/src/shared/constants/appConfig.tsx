// appConfig.tsx

export const APP_CONFIG = {
  NAME: 'MyApp',
  VERSION: '1.0.0',
  API_TIMEOUT: 5000,
  ENVIRONMENT: process.env.NODE_ENV || 'development',
  SETTINGS: {
      API_TIMEOUT: 5000,
      ENVIRONMENT: process.env.NODE_ENV || 'development',
  },
};

  // 使用例:
  // console.log(`アプリケーション名: ${APP_CONFIG.NAME}`);
  // console.log(`バージョン: ${APP_CONFIG.VERSION}`);
  // fetch(API_ENDPOINT, { timeout: APP_CONFIG.API_TIMEOUT })
  //   .then(response => response.json())
  //   .then(data => console.log(data));