import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom'; // Import useNavigate
import 'bootstrap/dist/css/bootstrap.min.css';
import loginImage from './loginImage.png';
import logoImage from './logo.webp'; 

const LoginForm = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate(); // Instantiate useNavigate

  const handleSubmit = async (event) => {
    event.preventDefault();
  
    // Specify the full URL for the Go server
    const serverUrl = 'http://localhost:8081/login';
  
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
    <div className="container-fluid h-100" style={{ minHeight: '100vh' }}>
      <div className="row no-gutter">
        {/* The image column that takes 70% of the page */}
        <div 
            className="col-md-8 d-none d-md-flex align-items-center justify-content-center p-0" 
            style={{ 
              background: `url(${loginImage}) no-repeat center center`, 
              backgroundSize: 'cover',
              minHeight: '100vh'
          }}
        >
          {/* This div will contain the background image */}
        </div>
  
        {/* The login form column that takes 30% of the page */}
        <div className="col-md-4 d-flex align-items-center" style={{ backgroundColor: '#0E2237', padding: 0 }}>
          {/* Actual login form centered within the form column */}
          <div className="w-100">
            <div className="mx-auto" style={{ maxWidth: '320px' }}>
            <img src={logoImage} alt="Logo" className="mb-4" id = "circle" style={{ maxWidth: '150px', display: 'block', marginLeft: 'auto', marginRight: 'auto' }} />
              <h3 className="display-4 mb-5 text-center text-white">Sign in</h3>
              {error && <div className="alert alert-danger" role="alert">{error}</div>}
              <form onSubmit={handleSubmit}>
              <div>
                <label htmlFor="username" className="text-white">Username</label>
                <div className="form-group mb-4">
                  <input
                    type="text"
                    id="username"
                    className="form-control border-0 shadow-sm px-4"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                  />
                </div>
              </div>
              <div>
                <label htmlFor="password" className="text-white">Password</label>
                <div className="form-group mb-4">
                  <input
                    type="password"
                    id="password"
                    className="form-control border-0 shadow-sm px-4 text-primary"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                  />
                </div>
              </div>
                <div className="form-group mb-5">
                <button type="submit" className="btn btn-primary btn-block text-uppercase shadow-sm" style={{ width: '100%', backgroundColor: '#50BCC0', borderColor: '#50BCC0' }}>Sign in</button>
                </div>
                <div className="text-center mt-5">
                <div style={{ marginBottom: '50px' }}>
                  <a href="#" className="text-white mr-3" style={{ textDecoration: 'none' }}>Forgot your password?</a>
                </div>
                <div style={{ marginBottom: '50px' }}>
                  <a href="#" className="text-white mr-3" style={{ textDecoration: 'none' }}>Don't Have An Account?</a>
                  <button type="button" className="btn btn-link text-uppercase font-weight-bold shadow-none" style={{ color: '#50BCC0', textDecoration: 'none', fontWeight: 'bold' }}>Sign up</button>
                </div>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
  
};

export default LoginForm;

