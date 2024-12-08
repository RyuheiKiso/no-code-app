import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';

// テストケース: 'learn react'リンクがレンダリングされることを確認する
test('renders learn react link', () => {
  // Appコンポーネントをレンダリング
  render(<App />);
  // 'learn react'というテキストを持つ要素を取得
  const linkElement = screen.getByText(/learn react/i);
  // 要素がドキュメント内に存在することを確認
  expect(linkElement).toBeInTheDocument();
});