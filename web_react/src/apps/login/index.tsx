import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { AUTH_ENDPOINTS } from '../../shared/constants/apiEndpoints';
import { TextField, Button, Typography, Container, Box, Alert } from '@mui/material';
import { useQuicProtoClient } from '../../shared/http/quic/QuicClient';

class LoginResponse {
  token: string;

  constructor(token: string) {
    this.token = token;
  }
}

const Login: React.FC = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const { data, error: quicError, sendRequest } = useQuicProtoClient({
    url: AUTH_ENDPOINTS.LOGIN,
    requestData: { username, password },
    YourResponse: LoginResponse,
  });

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    sendRequest();
  };

  useEffect(() => {
    if (quicError) {
      setError(quicError);
    } else if (data) {
      localStorage.setItem('token', data.token);
      navigate('/');
    }
  }, [data, quicError, navigate]);

  return (
    <Container>
      <Box display="flex" justifyContent="center" alignItems="center" minHeight="100vh">
        <Box
          sx={{
            backdropFilter: 'blur(10px)',
            backgroundColor: 'rgba(255, 255, 255, 0.1)',
            borderRadius: '10px',
            boxShadow: '0 4px 30px rgba(0, 0, 0, 0.1)',
            padding: '20px',
          }}
        >
          <Typography variant="h4" component="h1" gutterBottom>
            ログイン
          </Typography>
          {error && <Alert severity="error">{error}</Alert>}
          <form onSubmit={handleLogin}>
            <TextField
              label="メールアドレス"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              fullWidth
              margin="normal"
            />
            <TextField
              label="パスワード"
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              fullWidth
              margin="normal"
            />
            <Button type="submit" variant="contained" color="primary" fullWidth>
              ログイン
            </Button>
          </form>
        </Box>
      </Box>
    </Container>
  );
};

export default Login;