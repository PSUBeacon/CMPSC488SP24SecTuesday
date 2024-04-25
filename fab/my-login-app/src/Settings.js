import React, { useState, useEffect } from 'react';
import logoImage from './logo.webp'; // Adjust the path as needed
import './Settings.css'; // Adjust the path as necessary based on your file structure
import { Link } from 'react-router-dom'; // Import Link component

const SettingsPage = () => {
  // Initialize theme state from local storage if available, otherwise default to 'dark'
  const [theme, setTheme] = useState(localStorage.getItem('theme') || 'dark');
  const [selectedNav, setSelectedNav] = useState('Theme'); // State variable for selected navigation item

  // Function to set the theme and save it to local storage
  const changeTheme = (newTheme) => {
    setTheme(newTheme);
    localStorage.setItem('theme', newTheme); // Save new theme to local storage
  };

  // Styles that change with the theme
  const topNavStyle = {
    backgroundColor: theme === 'dark' ? '#081624' : '#F0F1F3',
    color: theme === 'dark' ? 'white' : 'black',
    padding: '0.5rem 1rem',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between',
  };

  const sideNavStyle = {
    backgroundColor: theme === 'dark' ? '#173350' : '#F0F1F3',
    color: theme === 'dark' ? 'white' : 'black',
    width: '250px',
    padding: '1rem',
  };

  // useEffect to load theme from local storage on mount
  useEffect(() => {
    const storedTheme = localStorage.getItem('theme');
    if (storedTheme) {
      setTheme(storedTheme);
    }
  }, []);

  // Render content based on the selected navigation item
  const renderContent = () => {
    switch (selectedNav) {
      // ... other cases remain unchanged
      case 'Theme':
        return (
          <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', height: '100%', width: '100%' }}>
            <label htmlFor="theme-select" style={{ fontSize: '1.5rem', marginBottom: '10px' }}>Choose a theme:</label>
            <select
              id="theme-select"
              value={theme}
              onChange={(e) => changeTheme(e.target.value)}
              style={{ width: '80%', height: '50px', fontSize: '1.2rem', textAlign: 'center', margin: '20px 0' }}
            >
              <option value="dark">Dark Mode</option>
              <option value="light">Light Mode</option>
            </select>
          </div>
        );
      case 'Notification Settings':
        return <div>Notification Settings Content Here</div>;
      case 'Account Settings':
        return <div>Account Settings Content Here</div>;
      default:
        return <div>Select a category</div>;
    }
  };

  return (
    <div style={{ display: 'flex', minHeight: '100vh', flexDirection: 'column' }}>
      <nav style={topNavStyle}>
        <img src={logoImage} alt="Logo" id='circle' style={{ height: '50px' }} />
        <span style={{ fontSize: '1.5rem', fontWeight: 'bold' }}>Settings</span>
      </nav>

      <div style={{ display: 'flex', flex: '1' }}>
        <aside style={sideNavStyle}>
          <nav>
            <ul style={{ listStyleType: 'none', padding: 0 }}>
            <li className="settings-nav-item" onClick={() => setSelectedNav('Dashboard')}>
<Link to="/dashboard" style={{ color: 'inherit', textDecoration: 'none' }}>Dashboard</Link> {/* Link to the dashboard */}
</li>
              <li className="settings-nav-item" onClick={() => setSelectedNav('Manage Users')}>Manage Users</li>
              <li className="settings-nav-item" onClick={() => setSelectedNav('Theme')}>Theme</li>
              <li className="settings-nav-item" onClick={() => setSelectedNav('Notification Settings')}>Notification Settings</li>
              <li className="settings-nav-item" onClick={() => setSelectedNav('Account Settings')}>Account Settings</li>
            </ul>
          </nav>
        </aside>

        <div style={{ flex: '1', padding: '1rem', backgroundColor: theme === 'dark' ? '#0E2237' : '#FFFFFF', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
          {renderContent()}
        </div>
      </div>
    </div>
  );
};

export default SettingsPage;