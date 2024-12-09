// themeUtils.tsx

import { THEME_CONFIG } from '../constants/themeConfig';

// テーマ設定を適用する関数
export const applyTheme = () => {
    document.documentElement.style.setProperty('--primary-color', THEME_CONFIG.COLORS.PRIMARY);
    document.documentElement.style.setProperty('--secondary-color', THEME_CONFIG.COLORS.SECONDARY);
    document.documentElement.style.setProperty('--background-color', THEME_CONFIG.COLORS.BACKGROUND);
    document.documentElement.style.setProperty('--text-color', THEME_CONFIG.COLORS.TEXT);
    document.documentElement.style.setProperty('--main-font', THEME_CONFIG.FONTS.MAIN);
    document.documentElement.style.setProperty('--heading-font', THEME_CONFIG.FONTS.HEADING);
    document.documentElement.style.setProperty('--header-height', THEME_CONFIG.SIZES.HEADER_HEIGHT);
    document.documentElement.style.setProperty('--footer-height', THEME_CONFIG.SIZES.FOOTER_HEIGHT);
};

// 使用例:
// applyTheme();