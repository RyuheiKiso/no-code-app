import { onCLS, onFID, onFCP, onLCP, onTTFB, Metric } from 'web-vitals';

// reportWebVitals 関数
// パフォーマンス測定のための関数を登録する
// onPerfEntry: パフォーマンス測定結果を受け取るコールバック関数 (オプション)
const reportWebVitals = (onPerfEntry?: (metric: Metric) => void) => {
  // onPerfEntry が関数として提供されている場合
  if (onPerfEntry && onPerfEntry instanceof Function) {
    // Cumulative Layout Shift (CLS) の測定結果をコールバック関数に渡す
    onCLS(onPerfEntry);
    // First Input Delay (FID) の測定結果をコールバック関数に渡す
    onFID(onPerfEntry);
    // First Contentful Paint (FCP) の測定結果をコールバック関数に渡す
    onFCP(onPerfEntry);
    // Largest Contentful Paint (LCP) の測定結果をコールバック関数に渡す
    onLCP(onPerfEntry);
    // Time to First Byte (TTFB) の測定結果をコールバック関数に渡す
    onTTFB(onPerfEntry);
  }
};

export default reportWebVitals;