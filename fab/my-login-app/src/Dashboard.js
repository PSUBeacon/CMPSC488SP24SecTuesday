import React from 'react';
import { Link } from 'react-router-dom'; // Import Link from react-router-dom for navigation
import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import logoImage from './logo.webp'; 
import houseImage from './houseImage.jpg';
import notificationIcon from './notification.png'
//import settingsIcon from './settings.png'
//import accountIcon from './account.png'

// Define the Dashboard component using a functional component pattern
const Dashboard = () => {
  // This is the JSX return statement where we layout our component's HTML structure
  return (
    <div style={{ display: 'flex', minHeight: '100vh', flexDirection: 'column', backgroundColor: '#081624' }}>
      {/* Top Navbar */}
      <nav style={{ backgroundColor: '#081624', color: 'white', padding: '0.5rem 1rem' }}>
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
          <div style={{ display: 'flex', alignItems: 'center' }}>
            <img src={logoImage} alt="Logo" style={{ marginRight: '10px'}} id='circle'/> {/* Adjust the height as needed */}
            <span id = 'menuText'>Beacon</span>
          </div>
          <div>
            <span id='menuText'>March 05, 2024</span>
          </div>
          <div>
            <span id='menuText'>11:48 AM</span>
          </div>
          <div>
          <img src={notificationIcon} alt="notifications" style={{ marginRight: '10px'}} id="menuIcon"/>
          </div>
        </div>
      </nav>

      {/* Side Navbar and Dashboard Content */}
      <div style={{ display: 'flex', flex: '1' }}>
        {/* Side Navbar */}
        <aside style={{ backgroundColor: '#001F3F', color: 'white', width: '250px', padding: '1rem' }}>
          <div class="houseInfo">
          <div><img src={houseImage} alt="Logo" style={{ marginRight: '10px'}} id='circle2'/></div>
          <div>My House</div>
          <div>State College, PA 16801</div>
          </div>
          <nav>
            <ul style={{ listStyle: 'none', padding: 0 }}>
              {/* Apply active style to 'Overview' since it's the current page */}
              <li style={{ backgroundColor: '#08192B', margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/" style={{ color: 'white', textDecoration: 'none' }}>
                  <i className="fas fa-home" style={{ marginRight: '10px' }}></i>
                  Overview
                </Link>
              </li>
              <li style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/security" style={{ color: 'white', textDecoration: 'none' }}>
                  <i className="fas fa-lock" style={{ marginRight: '10px' }}></i>
                  Security
                </Link>

              </li>
              <li style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/lighting" style={{ color: 'white', textDecoration: 'none' }}>
                  <i className="fas fa-lightbulb" style={{ marginRight: '10px' }}></i>
                  Lighting
                </Link>

              </li>
              <li style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/preferences" style={{ color: 'white', textDecoration: 'none' }}>
                  <i className="fas fa-sliders-h" style={{ marginRight: '10px' }}></i>
                  Preferences
                </Link>

              </li>
              <li style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/hvac" style={{ color: 'white', textDecoration: 'none' }}>
                  <i className="fas fa-thermometer-half" style={{ marginRight: '10px' }}></i>
                  HVAC
                </Link>

              </li>
              <li style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/Appliances" style={{ color: 'white', textDecoration: 'none' }}>
                  <i className="fas fa-blender" style={{ marginRight: '10px' }}></i>
                  Appliances
                </Link>

              </li>
              <li style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/energy" style={{ color: 'white', textDecoration: 'none' }}>
                  <i className="fas fa-bolt" style={{ marginRight: '10px' }}></i>
                  Energy
                </Link>
              </li>
            </ul>
          </nav>
        </aside>

        {/* Main Content */}
        <main style={{ flex: '1', padding: '1rem', backgroundColor: 'white' }}>
          <h2>Dashboard</h2>
          <p>Welcome to your dashboard!</p>
          {/* Dashboard widgets and content go here */}
        </main>
      </div>
    </div>
  );
};
export default Dashboard;