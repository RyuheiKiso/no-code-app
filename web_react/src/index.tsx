import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';

// ルートDOM要素を取得し、Reactアプリケーションのエントリーポイントを作成
const root = ReactDOM.createRoot(
  // 'root'というIDを持つDOM要素を取得し、型アサーションを使用してHTMLElementとして扱う
  document.getElementById('root') as HTMLElement
);

// Reactアプリケーションをレンダリング
root.render(
  // StrictModeは潜在的な問題を検出するための開発モードのラッパー
  <React.StrictMode>
    {/* Appコンポーネントをレンダリング */}
    <App />
  </React.StrictMode>
);

// パフォーマンス測定のための関数を呼び出し
reportWebVitals();