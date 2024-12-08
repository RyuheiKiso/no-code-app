import { useState, useEffect } from 'react';

interface RestAPIClientProps {
  url: string;
  requestData: any;
  method: 'GET' | 'POST' | 'PUT' | 'DELETE';
}

const useRestAPIClient = ({ url, requestData, method }: RestAPIClientProps) => {
  const [data, setData] = useState<any | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(url, {
          method,
          headers: {
            'Content-Type': 'application/json',
          },
          body: method !== 'GET' ? JSON.stringify(requestData) : null,
        });

        if (!response.ok) {
          throw new Error(`Error: ${response.statusText}`);
        }

        const responseData = await response.json();
        setData(responseData);
      } catch (err) {
        if (err instanceof Error) {
          setError(err.message);
        } else {
          setError('不明なエラーが発生しました');
        }
      }
    };

    fetchData();
  }, [url, requestData, method]);

  return { data, error };
};

export { useRestAPIClient };