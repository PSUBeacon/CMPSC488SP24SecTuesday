// SettingsPage.js
import React from 'react';
import logoImage from './logo.webp'; // Adjust the path as needed
import './Settings.css'; // Adjust the path as necessary based on your file structure

const SettingsPage = () => {
  // You can manage state here to highlight the active section if needed
  
  return (
    <div style={{ display: 'flex', minHeight: '100vh', flexDirection: 'column'}}>
      {/* Top Navbar */}
      <nav style={{ backgroundColor: '#081624', color: 'white', padding: '0.5rem 1rem', display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
        <img src={logoImage} alt="Logo" id="circle" style={{ height: '50px' }} />
        <span style={{ fontSize: '1.5rem', fontWeight: 'bold' }}>Settings</span>
        {/* You can add more top nav items here if needed */}
      </nav>

      {/* Settings Content */}
      <div style={{ display: 'flex', flex: '1' }}>
        {/* Side Navbar */}
        <aside style={{ backgroundColor: '#0E2237', color: 'white', width: '250px', padding: '1rem' }}>
          <nav>
            <ul style={{ listStyleType: 'none', padding: 0 }}>
              <li className="settings-nav-item">Manage Users</li>
              <li className="settings-nav-item">Theme</li>
              <li className="settings-nav-item">Notification Settings</li>
              <li className="settings-nav-item">Account Settings</li>
            </ul>
          </nav>
        </aside>

        {/* Settings Detail Content */}
        <div style={{ flex: '1', padding: '1rem',  backgroundColor: '#081624'}}>
          {/* The content for the selected settings page will go here */}
        </div>
      </div>
    </div>
  );
};

export default SettingsPage;
