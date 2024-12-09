// useLanguage.tsx

import { useState } from 'react';
import { LANGUAGES, MESSAGES_EN, MESSAGES_JA } from '../constants/languages';

// 言語管理のカスタムフック
const useLanguage = () => {
    const [language, setLanguage] = useState(LANGUAGES.EN);

    // 言語を切り替える関数
    const switchLanguage = (lang: string) => {
        setLanguage(lang);
    };

    // 現在の言語に応じたメッセージを取得
    const messages = language === LANGUAGES.JA ? MESSAGES_JA : MESSAGES_EN;

    return { language, switchLanguage, messages };
};

export default useLanguage;

// 使用例:
// const { language, switchLanguage, messages } = useLanguage();
// console.log(messages.WELCOME); // 現在の言語に応じたメッセージを表示