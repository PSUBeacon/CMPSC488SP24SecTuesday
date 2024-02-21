import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom'; // Import useNavigate

const LoginForm = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate(); // Instantiate useNavigate

  const handleSubmit = async (event) => {
    event.preventDefault();
    const serverUrl = 'http://localhost:8080/login';

    try {
      const response = await axios.post(serverUrl, {
        username,
        password
      });
      const jwtToken = response.data.token; // Assuming the token is returned in a field named 'token'
      console.log('JWT Token:', jwtToken);

      localStorage.setItem('token', jwtToken);

      // Redirect user to the dashboard page
      navigate('/dashboard'); // Use navigate to redirect

    } catch (error) {
      setError('Failed to login. Please check your network connection and credentials.');
      console.error('Login error:', error.toJSON());
      if (error.response) {
          console.error(error.response.data);
          console.error(error.response.status);
          console.error(error.response.headers);
      } else if (error.request) {
          console.error(error.request);
      } else {
          console.error('Error', error.message);
      }
      console.error(error.config);
    }
  };

  return (
    <div>
      <h2>Login</h2>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <form onSubmit={handleSubmit}>
        <div>
          <label>Username:</label>
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </div>
        <div>
          <label>Password:</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
        <button type="submit">Login</button>
      </form>
    </div>
  );
};

export default LoginForm;

