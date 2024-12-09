// ThemeSwitcher.tsx

import React from 'react';
import useTheme from '../hooks/useTheme';

const ThemeSwitcher: React.FC = () => {
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const { theme, switchTheme } = useTheme();

    return (
        <div>
            <button onClick={() => switchTheme('LIGHT')}>ライトテーマ</button>
            <button onClick={() => switchTheme('DARK')}>ダークテーマ</button>
        </div>
    );
};

export default ThemeSwitcher;

// 使用例:
// <ThemeSwitcher />