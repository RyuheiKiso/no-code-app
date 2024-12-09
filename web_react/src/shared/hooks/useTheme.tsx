// useTheme.tsx

import { useState } from 'react';
import { THEME_CONFIG } from '../constants/themeConfig';

// テーマ管理のカスタムフック
const useTheme = () => {
    const [theme, setTheme] = useState(THEME_CONFIG.LIGHT);

    // テーマを切り替える関数
    const switchTheme = (themeName: 'LIGHT' | 'DARK') => {
        setTheme(THEME_CONFIG[themeName]);
    };

    return { theme, switchTheme };
};

export default useTheme;

// 使用例:
// const { theme, switchTheme } = useTheme();
// console.log(theme.COLORS.PRIMARY); // 現在のテーマのプライマリカラーを表示