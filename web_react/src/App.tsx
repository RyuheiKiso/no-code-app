import React from 'react';
import './App.css';

// Appコンポーネント
// Reactアプリケーションのメインコンポーネント
function App() {
  return (
    // アプリケーション全体を囲むdiv要素
    <div className="App">
      {/* アプリケーションのヘッダー */}
      <header className="App-header">
        {/* 編集を促すメッセージ */}
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        {/* Reactの公式サイトへのリンク */}
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;