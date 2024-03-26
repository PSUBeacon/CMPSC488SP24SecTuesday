import React, { useState } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';
import loginImage from './loginImage.png';
import logoImage from './logo.webp'; 

const SignUp = () => {
  const [email, setEmail] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();
    const serverUrl = 'http://localhost:8080/signup'; // Adjust URL as needed

    try {
      const response = await axios.post(serverUrl, {
        username,
        password
      });
      // Handle successful sign-up
      console.log('Sign up successful:', response.data);
    } catch (error) {
      setError('Failed to sign up. Please try again.');
      console.error('Sign up error:', error.toJSON());
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
              <h3 className="display-4 mb-5 text-center text-white">Sign Up</h3>
              {error && <div className="alert alert-danger" role="alert">{error}</div>}
              <form onSubmit={handleSubmit}>
                <div>
                  <label htmlFor="email" className="text-white">Email</label>
                  <div className="form-group mb-4">
                    <input
                      type="text"
                      id="email"
                      className="form-control border-0 shadow-sm px-4"
                      value={email}
                      onChange={(e) => setEmail(e.target.value)}
                    />
                  </div>
                </div>
                
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
                  <button type="submit" className="btn btn-primary btn-block text-uppercase shadow-sm" style={{ width: '100%', backgroundColor: '#50BCC0', borderColor: '#50BCC0' }}>Sign Up</button>
                </div>
                <div className="text-center mt-5">
                  <div style={{ marginBottom: '50px' }}>
                    <span className="text-white mr-3">Already Have An Account?</span>
                    <Link to="/" className="btn btn-link text-uppercase font-weight-bold shadow-none" style={{ color: '#50BCC0', textDecoration: 'none', fontWeight: 'bold' }}>Sign in</Link> {/* Link to the login page */}
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

export default SignUp;
