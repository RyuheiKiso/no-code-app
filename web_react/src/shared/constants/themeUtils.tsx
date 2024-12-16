// themeUtils.tsx

import { THEME_CONFIG } from '../constants/themeConfig';

// テーマ設定を適用する関数
export const applyTheme = (themeName: 'LIGHT' | 'DARK') => {
    const theme = THEME_CONFIG[themeName];
    document.documentElement.style.setProperty('--primary-color', theme.COLORS.PRIMARY);
    document.documentElement.style.setProperty('--secondary-color', theme.COLORS.SECONDARY);
    document.documentElement.style.setProperty('--background-color', theme.COLORS.BACKGROUND);
    document.documentElement.style.setProperty('--text-color', theme.COLORS.TEXT);
    document.documentElement.style.setProperty('--main-font', theme.FONTS.MAIN);
    document.documentElement.style.setProperty('--heading-font', theme.FONTS.HEADING);
    document.documentElement.style.setProperty('--header-height', theme.SIZES.HEADER_HEIGHT);
    document.documentElement.style.setProperty('--footer-height', theme.SIZES.FOOTER_HEIGHT);
};