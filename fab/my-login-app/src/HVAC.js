// HVAC.js
import './HVAC.css'; // Import CSS file
import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { Link } from 'react-router-dom'; // Import Link from react-router-dom for navigation
import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import logoImage from './logo.webp'; 
import houseImage from './houseImage.jpg';
import settingsIcon from './settings.png'
import accountIcon from './account.png'
import menuIcon from './menu.png'

const HVAC = () => {

   // States for date and time
   const [currentDate, setCurrentDate] = useState(new Date().toLocaleDateString());
   const [currentTime, setCurrentTime] = useState(new Date().toLocaleTimeString());
 
   useEffect(() => {
       const timer = setInterval(() => {
           setCurrentDate(new Date().toLocaleDateString());
           setCurrentTime(new Date().toLocaleTimeString());
       }, 1000);
 
       // Cleanup on component unmount
       return () => clearInterval(timer);
   }, []);

  const navigate = useNavigate(); // Instantiate useNavigate hook
  const [isNavVisible, setIsNavVisible] = useState(false);
    const [dashboardMessage, setDashboardMessage] = useState('');
    const [accountType ,setAccountType] = useState('')
    
    const toggleHVACStatus = () => {
      setDeviceData(prevDeviceData => ({
        ...prevDeviceData,
        HVAC: {
          ...prevDeviceData.HVAC,
          Status: !prevDeviceData.HVAC.Status,
        },
      }));
    };

 // States for each device
 const [deviceData, setDeviceData] = useState({
  HVAC: {},
  Dishwasher: {},
  Fridge: {},
  Lighting: {},
  Microwave: {},
  Oven: {},
  SecuritySystem: {},
  SolarPanel: {},
  Toaster: {},
});


    useEffect(() => {
        const token = localStorage.getItem('token');
        const url = 'http://localhost:8081/dashboard';

        fetch(url, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
        })
            .then(response => response.json())
            .then(data => {
                console.log(data);
                // Update state for each device if present in response
                const updatedDeviceData = { ...deviceData };
                Object.keys(updatedDeviceData).forEach(device => {
                    if (data[device]) {
                        localStorage.setItem(device, JSON.stringify(data[device]));
                        updatedDeviceData[device] = data[device];
                    }
                });
                setDeviceData(updatedDeviceData);
                setDashboardMessage(data.message);

                // Store accountType in session storage
                setAccountType(data.accountType);
                sessionStorage.setItem('accountType', data.accountType);
            })
            .catch(error => console.error('Fetch operation error:', error));
    }, []);

  const toggleNav = () => {
    setIsNavVisible(!isNavVisible);
  };

  const goToSettings = () => {
    navigate('/settings');
  };

  const signOut = () => {
    // Add your sign-out logic here
    setIsAccountPopupVisible(false); // Close the popup
    navigate('/'); // Use navigate to redirect
  };

  const [isAccountPopupVisible, setIsAccountPopupVisible] = useState(false);

  const toggleAccountPopup = () => {
    setIsAccountPopupVisible(!isAccountPopupVisible);
  };



  
  const AccountPopup = ({ isVisible, onClose }) => {
    if (!isVisible) return null;
  
    return (
      <div className = "accountPop"style={{
        position: 'absolute',
        top: '100%', // Position it right below the button
        right: '0', // Align it with the right edge of the container
        backgroundColor: '#08192B',
        padding: '20px',
        zIndex: 100,
        color: 'white',
        borderRadius: '2px',
        // Add box-shadow or borders as needed for better visibility
      }}>
        <p>John Doe</p> {/* Replace with actual user name */}
          {accountType && <p>{accountType}</p>} {/* Dynamically display user role */}
        <button onClick={signOut} className="signout">Sign Out</button>
      </div>
    );
  };
  



  // This is the JSX return statement where we layout our component's HTML structure
  return (
    <div style={{ display: 'flex', minHeight: '100vh', flexDirection: 'column', backgroundColor: '#081624' }}>
      {/* Top Navbar */}
      <nav className="topNav" style={{ backgroundColor: '#081624', color: 'white', padding: '0.5rem 1rem' }}>
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
          <div style={{ display: 'flex', alignItems: 'center' }}>
          <img src={menuIcon} alt="Menu" onClick={toggleNav} className="hamburger-menu"/>
            <img src={logoImage} alt="Logo" style={{ marginRight: '10px'}} id='circle'/> {/* Adjust the height as needed */}
            <span id = 'menuText2'>Beacon</span>
          </div>
          <div>
          <span id='menuText'>{currentDate}</span>
          </div>
          <div>
            <span id='menuText'>{currentTime}</span>
          </div>
          <div>
          <div style={{ position: 'relative' }}>
          <img src={settingsIcon} alt="Settings" style={{ marginRight: '10px' }} id="menuIcon" onClick={goToSettings} />
  <button onClick={toggleAccountPopup} style={{ background: 'none', border: 'none', padding: 0, cursor: 'pointer' }}>
    <img src={accountIcon} alt="account" style={{ marginRight: '10px' }} id = "menuIcon2"/>
  </button>
  <AccountPopup isVisible={isAccountPopupVisible} onClose={() => setIsAccountPopupVisible(false)} />
</div>
</div>
        </div>
      </nav>

      {/* Side Navbar and Dashboard Content */}
      <div style={{ display: 'flex', flex: '1' }}>
        {/* Side Navbar */}
        <aside className={`side-nav ${isNavVisible ? '' : 'hidden'}`} style={{ backgroundColor: '#0E2237', color: 'white', width: '250px', padding: '1rem' }}>          <div class="houseInfo">
          <div><img src={houseImage} alt="Logo" style={{ marginRight: '10px'}} id='circle2'/></div>
          <div>My House</div>
          <div>State College, PA 16801</div>
          </div>
          <nav>
            <ul style={{ listStyle: 'none', padding: 0 }}>
              {/* Apply active style to 'Overview' since it's the current page */}
              <li className="nav-item"style={{margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/dashboard" style={{ color: '#95A4B6', textDecoration: 'none' }}>
                  <i className="fas fa-home" style={{ marginRight: '10px' }}></i>
                  Overview
                </Link>
              </li>
              <li className="nav-item"style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/security" style={{ color: '#95A4B6', textDecoration: 'none' }}>
                  <i className="fas fa-lock" style={{ marginRight: '10px' }}></i>
                  Security
                </Link>
              </li>
              <li className="nav-item"style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/lighting" style={{ color: '#95A4B6', textDecoration: 'none' }}>
                  <i className="fas fa-lightbulb" style={{ marginRight: '10px' }}></i>
                  Lighting
                </Link>
              </li>
              <li className="nav-item"style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/networking" style={{ color: '#95A4B6', textDecoration: 'none' }}>
                  <i className="fas fa-sliders-h" style={{ marginRight: '10px' }}></i>
                  Networking
                </Link>
              </li>
              <li className="nav-item"style={{ backgroundColor: '#08192B', margin: '0.5rem 0', padding: '0.5rem', borderLeft: '3px solid #0294A5' }}>
                <Link to="/hvac" style={{ color: '#50BCC0', textDecoration: 'none' }}>
                  <i className="fas fa-thermometer-half" style={{ marginRight: '10px' }}></i>
                  HVAC
                </Link>
              </li>
              <li className="nav-item"style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/appliances" style={{ color: '#95A4B6', textDecoration: 'none' }}>
                  <i className="fas fa-blender" style={{ marginRight: '10px' }}></i>
                  Appliances
                </Link>
              </li>
              <li className="nav-item"style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/energy" style={{ color: '#95A4B6', textDecoration: 'none' }}>
                  <i className="fas fa-bolt" style={{ marginRight: '10px' }}></i>
                  Energy
                </Link>
              </li>
            </ul>
          </nav>
        </aside>

    
        <main style={{ flex: '1', padding: '1rem', display: 'flex', flexDirection: 'column', alignItems: 'center', backgroundImage: 'linear-gradient(to bottom, #0E2237, #081624)', position: 'relative' }}>
  {/* Translucent pattern overlay */}
  <div style={{ position: 'absolute', top: 0, left: 0, width: '100%', height: '100%', backgroundImage: 'url("https://www.transparenttextures.com/patterns/always-grey.png")', opacity: 0.3 }}></div>

  <h1 style={{ color: 'white', marginBottom: '2rem' }}>HVAC</h1>
  {deviceData && (
    <div className="hvac-data-container">
      <div className="data-item">
        <div className="data-icon">
          <i className="fas fa-thermometer-half"></i>
        </div>
        <div className="data-info">
          <p style={{ color: '#95A4B6'}}>Temperature</p>
          <p>{deviceData.HVAC.Temperature}°F</p>
        </div>
      </div>
      <div className="data-item">
        <div className="data-icon">
          <i className="fas fa-tint"></i>
        </div>
        <div className="data-info">
          <p style={{ color: '#95A4B6'}}>Humidity</p>
          <p>{deviceData.HVAC.Humidity}%</p>
        </div>
      </div>
      <div className="data-item">
        <div className="data-icon">
          <i className="fas fa-fan"></i>
        </div>
        <div className="data-info">
          <p style={{ color: '#95A4B6'}}>Fan Speed</p>
          <p>{deviceData.HVAC.FanSpeed}%</p>
        </div>
      </div>
      <div className="data-item">
        <div className="data-icon">
          <i className={deviceData.HVAC.Status ? "fas fa-power-on" : "fas fa-power-off"}></i>
        </div>
        <div className="data-info">
          <p style={{ color: '#95A4B6'}}>Status</p>
          <p>{deviceData.HVAC.Status ? "On" : "Off"}</p>
          {/* Toggle */}
          <label className="toggle" style={{ display: 'block', margin: 'auto' }}>
            <input
              type="checkbox"
              checked={deviceData.HVAC.Status}
              onChange={toggleHVACStatus}
            />
            <span className="slider" style={{ background: deviceData.HVAC.Status ? '#50BCC0' : 'grey' }}></span>
          </label>
        </div>
      </div>
      <div className="data-item">
        <div className="data-icon">
          <i className="fas fa-bolt"></i>
        </div>
        <div className="data-info">
          <p style={{ color: '#95A4B6'}}>Energy Consumption</p>
          <p>{deviceData.HVAC.EnergyConsumption}KW</p>
        </div>
      </div>
    </div>
  )}
</main>





      </div>
    </div>
  );
};

export default HVAC;
