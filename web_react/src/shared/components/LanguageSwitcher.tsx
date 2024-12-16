import React from 'react';
import useLanguage from '../../shared/hooks/useLanguage';
import { LANGUAGES } from '../../shared/constants/languages';

const LanguageSwitcher: React.FC = () => {
  const { language, switchLanguage, messages } = useLanguage();

  return (
    <div>
      <button onClick={() => switchLanguage(LANGUAGES.EN)}>English</button>
      <button onClick={() => switchLanguage(LANGUAGES.JA)}>日本語</button>
      <p>{messages.WELCOME}</p>
    </div>
  );
};

export default LanguageSwitcher;