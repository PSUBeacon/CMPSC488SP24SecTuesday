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

// Define the Dashboard component using a functional component pattern
const Lighting= () => {



  const navigate = useNavigate(); // Instantiate useNavigate hook
  const [selectedLight, setSelectedLight] = useState(null);

  // Function to handle card click
  const handleSelectLight = (lightId) => {
    setSelectedLight(lightId); // Update the selected light state
  };
  
  const [isNavVisible, setIsNavVisible] = useState(false);
    const [dashboardMessage, setDashboardMessage] = useState('');
    const [accountType ,setAccountType] = useState('')
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
    const [dimmerValue, setDimmerValue] = useState(75); // State to keep track of dimmer value    const [roomName, setRoomName] = useState('');
    const [selectedRoom, setSelectedRoom] = useState(null); 
    // State for light on/off toggle
    const [isLightOn, setIsLightOn] = useState(false);
    // Function to toggle light on/off
    const toggleLight = () => {
      setIsLightOn(!isLightOn);
    };

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
  
  // Add new states for the room and light names
  const [roomName, setRoomName] = useState('');
  const [lightName, setLightName] = useState('');
  const [lights, setLights] = useState([]);

  // Function to handle form submission
  const handleFormSubmit = (e) => {
    e.preventDefault(); // Prevent page refresh
    const newLight = { roomName, lightName };
    setLights([...lights, newLight]); // Add new light to the list
  };

  // Function to remove a light
  const handleRemoveLight = (index) => {
    const newLights = [...lights];
    newLights.splice(index, 1);
    setLights(newLights);
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
                <Link to="/networking" style={{ color: 'white', textDecoration: 'none' }}>
                  <i className="fas fa-sliders-h" style={{ marginRight: '10px' }}></i>
                  Networking
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

    
        <main style={{ flex: '1', padding: '1rem', display: 'flex', flexDirection: 'column', alignItems: 'flex-start', backgroundColor: '#0E2237', width: '100%'}}>
  
  {/* Content Block */}
  
  <div className="contentBlock" style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', width: '100%', paddingBottom: '60px'}}>
            {/* Lights Control Section */}
            <div className="lightsControl" style={{ flex: '1', display: 'flex', flexDirection: 'column'}}>
              {/* Room Selection */}
              <div className="roomSelection" style={{ width: '50%', display: 'flex', flexDirection: 'column' }}>
              <h3 className="centered-title">Selecting a Room</h3>
                <div className="RoomCards" style={{ display: 'flex', flexWrap: 'wrap', justifyContent: 'space-around',  padding:'0px' }}>
                  {/* Room cards */}
                  <div className="card" style={{ width: '40%', marginBottom: '20px', border: selectedRoom === "Bedroom 1" ? '2px solid #0294A5' : 'none' }} onClick={() => {setRoomName("Bedroom 1"); setSelectedRoom("Bedroom 1");}}><img class="images" src={bedroomIcon} alt="Room 1" /></div>
                  <div className="card" style={{ width: '40%', marginBottom: '20px', border: selectedRoom === "Bedroom 2" ? '2px solid #0294A5' : 'none' }} onClick={() => {setRoomName("Bedroom 2"); setSelectedRoom("Bedroom 2");}}><img class="images" src={bedroomIcon} alt="Room 2" /></div>
                  <div className="card" style={{ width: '40%', border: selectedRoom === "Living Room" ? '2px solid #0294A5' : 'none' }} onClick={() => {setRoomName("Living Room"); setSelectedRoom("Living Room");}}><img class="images" src={livingroomIcon} alt="Room 3" /></div>
                </div>
                <div className="formContainer" style={{ width: '100%', display: 'flex', flexDirection: 'column' }}>
                  {/* Form to add lights */}
                  <form onSubmit={handleFormSubmit}>
                    <label>
                      Room Name:
                      <input type="text" value={roomName} onChange={(e) => setRoomName(e.target.value)} />
                    </label>
                    <label>
                      Light Name:
                      <input type="text" value={lightName} onChange={(e) => setLightName(e.target.value)} />
                    </label>
                    <button type="submit">Add Light</button>
                  </form>
                  {/* List of lights with remove option */}
                  <ul>
                    {lights.map((light, index) => (
                      <li key={index}>
                        {light.roomName} - {light.lightName}
                        <button onClick={() => handleRemoveLight(index)}>Remove</button>
                      </li>
                  ))}
                </ul>
                </div>
                {/* Light Dimmer Control */}
                <div className="dimmerControl" style={{ width: '30%', marginRight: '75px' }}>
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
                  {/* Turn on/off button */}
                  <button onClick={toggleLight} className="toggleButton">
                    {isLightOn ? 'Turn Off' : 'Turn On'}
                  </button>
                </div>
              </div>
                
                
              

              
            </div>
          </div>
        </main>
      </div>
    </div>
  );
};

export default Lighting;