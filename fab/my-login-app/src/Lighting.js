import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import logoImage from './logo.webp';
import houseImage from './houseImage.jpg';
import notificationIcon from './notification.png';
import lightbulbIcon from './lightbulbIcon.png';
import bedroomIcon from './bedroomIcon.jpg';
import livingroomIcon from './livingroomIcon.jpg';
import './Lighting.css';

const Lighting = () => {
  const [isLightOn, setIsLightOn] = useState(false);

  const handleTurnOn = () => {
    setIsLightOn(true);
    // Add logic to turn the light on
  };

  const handleTurnOff = () => {
    setIsLightOn(false);
    // Add logic to turn the light off
  };

  return (
    <div className="lightingPage" style={{ display: 'flex', minHeight: '100vh', flexDirection: 'column', backgroundColor: '#081624' }}>
      {/* Top Navbar */}
      <nav style={{ backgroundColor: '#081624', color: 'white', padding: '0.5rem 1rem' }}>
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
          <div style={{ display: 'flex', alignItems: 'center' }}>
            <img src={logoImage} alt="Logo" style={{ marginRight: '10px', width: '40px', height: '40px'}} id='circle'/> {/* Adjust the height as needed */}
            <span id='menuText'>Beacon</span>
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
          <div className="houseInfo">
            <div><img src={houseImage} alt="Logo" style={{ marginRight: '10px'}} id='circle2'/></div>
            <div>My House</div>
            <div>State College, PA 16801</div>
          </div>
          <nav>
            <ul style={{ listStyle: 'none', padding: 0 }}>
              {/* Apply active style to 'Overview' since it's the current page */}
              <li style={{margin: '0.5rem 0', padding: '0.5rem' }}>
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
              <li style={{backgroundColor: '#08192B', margin: '0.5rem 0', padding: '0.5rem' }}>
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
                <Link to="/appliances" style={{ color: 'white', textDecoration: 'none' }}>
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
        <main style={{ flex: '1', padding: '1rem', backgroundColor: '#173350'}}>
          <h2>Lighting Page</h2>
          
          {/* Lights Control Section */}
          <div className="lightsControl">
            {/* Room Selection */}
            <div className="roomSelection">
              <h3>Selecting A Room</h3>
              <div className="roomCards">
                {/* Replace placeholder content with your room cards */}
                <div className="card"><img src={bedroomIcon} alt="Room 1" /></div>
                <div className="card"><img src={bedroomIcon} alt="Room 2" /></div>
                <div className="card"><img src={livingroomIcon} alt="Room 3" /></div>
                <div className="card"><img src={livingroomIcon} alt="Room 4" /></div>
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
              <input type="range" id="dimmer" name="dimmer" min="0" max="100" />
              <label htmlFor="dimmer">75%</label>
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
