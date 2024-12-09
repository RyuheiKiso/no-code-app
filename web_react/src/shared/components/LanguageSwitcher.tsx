// LanguageSwitcher.tsx

import React from 'react';
import useLanguage from '../shared/hooks/useLanguage';
import { LANGUAGES } from '../shared/constants/languages';

const LanguageSwitcher: React.FC = () => {
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const { language, switchLanguage, messages } = useLanguage();

    return (
        <div>
            <p>{messages.WELCOME}</p>
            <button onClick={() => switchLanguage(LANGUAGES.EN)}>English</button>
            <button onClick={() => switchLanguage(LANGUAGES.JA)}>日本語</button>
        </div>
    );
};

export default LanguageSwitcher;

// 使用例:
// <LanguageSwitcher />