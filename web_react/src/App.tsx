// Reactをインポート
import React from 'react';
// React Routerのコンポーネントをインポート
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
// スタイルシートをインポート
import './App.css';
// Loginコンポーネントをインポート
import Login from './apps/login/index';
// ルート定数をインポート
import { LOGIN_ROUTE } from './shared/constants/routes';

// Appコンポーネントを定義
function App() {
  return (
    // ルーターコンポーネントでアプリ全体をラップ
    <Router>
      {/* アプリのメインコンテナ */}
      <div className="App">
        {/* ルートを定義 */}
        <Routes>
          {/* ログインルートを定義 */}
          <Route path={LOGIN_ROUTE} element={<Login />} />
        </Routes>
      </div>
    </Router>
  );
}

// Appコンポーネントをエクスポート
export default App;