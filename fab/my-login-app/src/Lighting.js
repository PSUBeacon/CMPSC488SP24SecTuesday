import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { Link } from 'react-router-dom'; // Import Link from react-router-dom for navigation
import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import logoImage from './logo.webp'; 
import houseImage from './houseImage.jpg';
import notificationIcon from './notification.png'
import accountIcon from './account.png'
import settingsIcon from './settings.png'
import menuIcon from './menu.png'
import bedroomIcon from './bedroomIcon.jpg'
import livingroomIcon from './livingroomIcon.jpg'
import lightbulbIcon from './lightbulbIcon.png'
import placeholderImage from './placeholderImage.jpg'; // Replace with the path to your placeholder image
import placeholderImage2 from './placeholderImage2.jpg'; // Replace with the path to your placeholder image
import './Lighting.css';
import axios from "axios";

const Lighting = () => {
  const navigate = useNavigate();
  const [isLightOn, setIsLightOn] = useState(false);
  const [isNavVisible, setIsNavVisible] = useState(false);
  const [isAccountPopupVisible, setIsAccountPopupVisible] = useState(false);
  const [dimmerValue, setDimmerValue] = useState(75); // State for dimmer value
  const handleTurnOn = () => {
    setIsLightOn(true);
    const serverUrl = 'https://beacon-cs2024.digital/api/lighting';

    // Define the body of the request based on your Go server's expected input.
    const requestBody = {
      uuid: '417293',
      status: true,
      brightness: 100,
    };
    const token = localStorage.getItem('token')
    // Send a POST request to turn the light on.
    axios.post(serverUrl, requestBody, {headers: {'Authorization': `Bearer ${token}`}})
        .then(response => {
          // Check if response is successful (status code 2xx)
          if (response.status >= 200 && response.status < 300) {
            console.log('Request succeeded');
            // Access response data if available
            if (response.data) {
              console.log(response.data);
            }
          } else {
            // Handle unsuccessful response
            console.error('Request failed with status:', response.status);
          }
        })
        .catch(error => {
          // Handle error
          console.error('There was an error!', error);
        });
  };

  const handleTurnOff = () => {
    setIsLightOn(false);
    const serverUrl = 'https://beacon-cs2024.digital/api/lighting';

    // Define the body of the request based on your Go server's expected input.
    const requestBody = {
      uuid: '417293',
      status: false,
      brightness: 100,
    };

    const token = localStorage.getItem('token')
    // Send a POST request to turn the light on.
    axios.post(serverUrl, requestBody,{headers: {'Authorization': `Bearer ${token}`}})
        .then(response => {
          console.log(response.data);
        })
        .catch(error => {
          console.error('There was an error!', error);
        });
  };

  const handleDimmerChange = (event) => {
    setDimmerValue(parseInt(event.target.value)); // Parse to integer
  };

  const goToSettings = () => {
    navigate('/settings');
  };
  const toggleNav = () => { // Add this function
    setIsNavVisible(!isNavVisible);
  };

  const toggleAccountPopup = () => {
    setIsAccountPopupVisible(!isAccountPopupVisible);
  };

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
            <span id='menuText'>March 05, 2024</span>
          </div>
          <div>
            <span id='menuText'>11:48 AM</span>
          </div>
          <div>
          <div style={{ position: 'relative' }}>
          <img src={settingsIcon} alt="Settings" style={{ marginRight: '10px' }} id="menuIcon" onClick={goToSettings} />
  <button onClick={toggleAccountPopup} style={{ background: 'none', border: 'none', padding: 0, cursor: 'pointer' }}>
    <img src={accountIcon} alt="account" style={{ marginRight: '10px' }} id = "menuIcon2"/>
  </button>
            {/* <AccountPopup isVisible={isAccountPopupVisible} onClose={() => setIsAccountPopupVisible(false)} /> */}
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
                <Link to="/dashboard" style={{ color: 'white', textDecoration: 'none' }}>
                  <i className="fas fa-home" style={{ marginRight: '10px' }}></i>
                  Overview
                </Link>
              </li>
              <li className="nav-item"style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/security" style={{ color: 'white', textDecoration: 'none' }}>
                  <i className="fas fa-lock" style={{ marginRight: '10px' }}></i>
                  Security
                </Link>
              </li>
              <li className="nav-item"style={{ backgroundColor: '#08192B', margin: '0.5rem 0', padding: '0.5rem', borderLeft: '3px solid #0294A5' }}>
                <Link to="/lighting" style={{ color: '#50BCC0', textDecoration: 'none' }}>
                  <i className="fas fa-lightbulb" style={{ marginRight: '10px' }}></i>
                  Lighting
                </Link>
              </li>
              <li className="nav-item"style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/preferences" style={{ color: 'white', textDecoration: 'none' }}>
                  <i className="fas fa-sliders-h" style={{ marginRight: '10px' }}></i>
                  Preferences
                </Link>
              </li>
              <li className="nav-item"style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/hvac" style={{ color: 'white', textDecoration: 'none' }}>
                  <i className="fas fa-thermometer-half" style={{ marginRight: '10px' }}></i>
                  HVAC
                </Link>
              </li>
              <li className="nav-item"style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/appliances" style={{ color: 'white', textDecoration: 'none' }}>
                  <i className="fas fa-blender" style={{ marginRight: '10px' }}></i>
                  Appliances
                </Link>
              </li>
              <li className="nav-item"style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/energy" style={{ color: 'white', textDecoration: 'none' }}>
                  <i className="fas fa-bolt" style={{ marginRight: '10px' }}></i>
                  Energy
                </Link>
              </li>
            </ul>
          </nav>
        </aside>

    
        <main style={{ flex: '1', padding: '1rem', display: 'flex', flexDirection: 'column', alignItems: 'center', backgroundColor: '#0E2237', width: '100%'}}>
  
  {/* Content Block */}
  
  <div className="contentBlock" style={{ display: 'flex', justifyContent: 'space-around', width: '100%', flexWrap: 'wrap' }}>
            {/* Lights Control Section */}
            <div className="lightsControl" style={{ display: 'flex', flexDirection: 'row', justifyContent: 'center', gap: '20px', width: '100%', flexWrap: 'wrap' }}>
              {/* Room Selection */}
              <div className="roomSelection" style={{ flexBasis: '48%'}}>
              <h3 className="centered-title">Selecting a Room</h3>
                <div className="RoomCards" style={{ display: 'flex', flexWrap: 'wrap', justifyContent: 'space-around',  padding:'0px' }}>
                  {/* Room cards */}
                  <div className="card" style={{ width: '40%', marginBottom: '20px' }}><img class="images" src={bedroomIcon} alt="Room 1" /></div>
                  <div className="card" style={{ width: '40%', marginBottom: '20px' }}><img class="images" src={bedroomIcon} alt="Room 2" /></div>
                  <div className="card" style={{ width: '40%' }}><img class="images" src={livingroomIcon} alt="Room 3" /></div>
                </div>
              </div>
            </div>

            {/* Light Selection */}
            <div className="lightSelection">
              <h3>Selecting A Light</h3>
              <div className="lightCards">
                {/* Replace placeholder content with your light cards */}
                <div className="card"><img src={lightbulbIcon} alt="Light 1" /></div>
                <div className="card"><img src={lightbulbIcon} alt="Light 2" /></div>
                <div className="card"><img src={lightbulbIcon} alt="Light 3" /></div>
                <div className="card"><img src={lightbulbIcon} alt="Light 3" /></div>
              </div>
            </div>

            {/* Light Dimmer Control */}
            <div className="dimmerControl">
              <input type="range" id="dimmer" name="dimmer" min="0" max="100" value={dimmerValue} onChange={handleDimmerChange} />
              <label htmlFor="dimmer">{dimmerValue}%</label>
            </div>
                {/* Light Dimmer Control */}
                <div className="dimmerControl">
                  <input
                    type="range"
                    id="dimmer"
                    name="dimmer"
                    min="0"
                    max="100"
                    value={dimmerValue}
                    onChange={(e) => setDimmerValue(e.target.value)}
                  />
                  <label htmlFor="dimmer">{dimmerValue}%</label>
                </div>
    {/* Turn On/Off Button */}
    <div className="lightControls">
      <button onClick={isLightOn ? handleTurnOff : handleTurnOn}>
        {isLightOn ? 'Turn Off' : 'Turn On'}
      </button>
    </div>
            </div>
        </main>
      </div>
    </div>
  );
};

export default Lighting;