
import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { Link } from 'react-router-dom'; // Import Link from react-router-dom for navigation
import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import logoImage from './logo.webp'; 
import houseImage from './houseImage.jpg';
import settingsIcon from './settings.png';
import accountIcon from './account.png';
import menuIcon from './menu.png';
import bedroomIcon from './bedroomIcon.jpg';
import livingroomIcon from './livingroomIcon.jpg';
import doorLockIcon from './doorLockIcon.png';
import './Security.css'; // Import your CSS file here



const Security = () => {

  const navigate = useNavigate(); // Instantiate useNavigate hook
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

    const [dimmerValue, setDimmerValue] = useState(75); // State to keep track of dimmer value
    const [isLocked, setIsLocked] = useState(false); //need this for the toggle and also the two lines below
    const toggleLock = () => {
      setIsLocked(!isLocked);
    };
    const [selectedRoom, setSelectedRoom] = useState(null);
    // Add a function to handle selecting a room:
    const selectRoom = (roomName) => {
      setSelectedRoom(roomName);
    };
    

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
    <div style={{ display: 'flex', minHeight: '100vh', flexDirection: 'column', backgroundColor: '#081624', width: '100%' }}>
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
              <li className="nav-item"style={{margin: '0.5rem 0', padding: '0.5rem'}}>
                <Link to="/dashboard" style={{ color: 'white', textDecoration: 'none' }}>
                  <i className="fas fa-home" style={{ marginRight: '10px' }}></i>
                  Overview
                </Link>
              </li>
              <li className="nav-item"style={{ backgroundColor: '#08192B', margin: '0.5rem 0', padding: '0.5rem', borderLeft: '3px solid #0294A5' }}>
                <Link to="/security" style={{ color: '#50BCC0', textDecoration: 'none' }}>
                  <i className="fas fa-lock" style={{ marginRight: '10px' }}></i>
                  Security
                </Link>
              </li>
              <li className="nav-item"style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/lighting" style={{ color: 'white', textDecoration: 'none' }}>
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
              <div className="roomSelection" style={{ flexBasis: '48%' }}>
                <h3 className="centered-title">Selecting a Door</h3>
                <div className="roomCards" style={{ display: 'flex', flexWrap: 'wrap', justifyContent: 'space-around', padding: '0px' }}>
                  {/* Room cards */}
                  <div className={selectedRoom === "Room 1" ? "card selected" : "card"} onClick={() => selectRoom("Room 1")} style={{ width: '40%', marginBottom: '20px' }}>
                    <img className="image" src={bedroomIcon} alt="Room 1" />
                  </div>
                  <div className={selectedRoom === "Room 2" ? "card selected" : "card"} onClick={() => selectRoom("Room 2")} style={{ width: '40%', marginBottom: '20px' }}>
                    <img className="image" src={bedroomIcon} alt="Room 2" />
                  </div>
                  
                </div>
              </div>


                {/* Light Selection */}
                <div className="lightSelection" style={{ flexBasis: '48%'}}>
                <h3 className="centered-title">Your Lock</h3>

                <div className="lightCards" style={{display: 'flex', flexWrap: 'wrap', justifyContent: 'space-around', padding:'0px' }}>
                  {/* Lock card */}
                  <div className="card" style={{ width: '100%', maxWidth: '300px', textAlign: 'center', padding: '20px' }}>
                    <img className="lockImage" src={doorLockIcon} alt="Lock Icon" />
                    {/*toggle function*/}
                    <label className="switch"> 
                      <input type="checkbox" checked={isLocked} onChange={toggleLock} />
                      <span className="slider round"></span>
                    </label>
                  </div>
                </div>
      

                
              </div>

              
            </div>
          </div>
        </main>
      </div>
    </div>
  );
};

export default Security;

