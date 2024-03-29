
import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { Link } from 'react-router-dom'; // Import Link from react-router-dom for navigation
import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import logoImage from './logo.webp'; 
import houseImage from './houseImage.jpg';
import doorLockIcon from './doorLockIcon.png'
import settingsIcon from './settings.png'
import accountIcon from './account.png'
import menuIcon from './menu.png'
import bulbIcon from './bulb-icon.png'
import placeholderImage from './placeholderImage.jpg'; // Replace with the path to your placeholder image
import placeholderImage2 from './placeholderImage2.jpg'; // Replace with the path to your placeholder image

// Define the Dashboard component using a functional component pattern
const Dashboard = () => {

 

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
  
  const [cameraView, setCameraView] = useState('livingroom'); // Default camera view
  
  // Object that holds the URLs for your camera feeds
  const cameraFeeds = {
    livingroom: placeholderImage, // Replace with the actual camera feed URL or image for the living room
    kitchen: placeholderImage2, // Replace with the actual camera feed URL or image for the kitchen
    r3: placeholderImage, //adding another tab to the camera views
  };

  const [isAccountPopupVisible, setIsAccountPopupVisible] = useState(false);

  const toggleAccountPopup = () => {
    setIsAccountPopupVisible(!isAccountPopupVisible);
  };


  const CameraWidget = () => {
    
    const cameraFeeds = {
      livingroom: placeholderImage, // Replace with actual video feed or image
      kitchen: placeholderImage, // Replace with actual video feed or image for kitchen
      r3: placeholderImage,
      // Add more camera feeds as needed
    };
    
    return (
      <div className="camera-widget" style={{ position: 'relative', maxWidth: '100%', backgroundColor: '#12232E', borderRadius: '10px', overflow: 'hidden' }}>
        {/* Live Feed */}
        <img src={cameraFeeds[cameraView]} alt="Live feed" style={{ width: '100%', height: 'auto', display: 'block' }} />
        
        {/* Camera View Buttons */}
        <div style={{ position: 'absolute', top: '10px', left: '10px', display: 'flex', gap: '5px' }}>
          <button onClick={() => setCameraView('livingroom')} style={{ padding: '5px', backgroundColor: cameraView === 'livingroom' ? '#4CAF50' : 'transparent' }}>R1</button>
          <button onClick={() => setCameraView('kitchen')} style={{ padding: '5px', backgroundColor: cameraView === 'kitchen' ? '#4CAF50' : 'transparent' }}>R2</button>
          <button onClick={() => setCameraView('r3')} style={{ padding: '5px', backgroundColor: cameraView === 'r3' ? '#4CAF50' : 'transparent' }}>R3</button>
          {/* Add more buttons for additional camera views as needed */}
        </div>
      </div>
    );
  };

 // Locks Widget JSX
 const LocksWidget = () => {
  const [locksStatus, setLocksStatus] = useState({
    frontDoor: 'Locked',
    backDoor: 'Locked',
  });

  const toggleLockStatus = (door) => {
    setLocksStatus(prevStatus => ({
      ...prevStatus,
      [door]: prevStatus[door] === 'Locked' ? 'Unlocked' : 'Locked',
    }));
  };

  // Assuming you have a state to track the light status and its setter function
  const [lightStatus, setLightStatus] = useState('Off');

  const toggleLightStatus = () => {
    // Toggle the light status
    setLightStatus(prevStatus => prevStatus === 'Off' ? 'On' : 'Off');
  };

  return (
    <div>
      <div className="widget" style={{ display: 'flex', flexDirection: 'row', justifyContent: 'center', alignItems: 'center', backgroundColor: '#173350', padding: '20px', paddingRight:"0px", borderRadius: '10px', margin: '0px' }}>
        {/* Front Door Lock */}
        <div style={{ flex: '1', maxWidth:'250px', textAlign: 'center' }}>
          <p style={{color:"#95A4B6", marginRight:'20px'}}>Front Door</p>
          <img src={doorLockIcon} alt="Front Door" style={{ width: '50%', height: 'auto', marginBottom:'10px' }} /> {/* Replace with lock icon */}
          <p style={{color:"#95A4B6", marginRight:'20px'}}>{locksStatus.frontDoor === 'Unlocked' ? 'Unlocked' : 'Locked'}</p>
          <label className="toggle" style={{ display: 'block', margin: 'auto',  }}>
            <input
              type="checkbox"
              checked={locksStatus.frontDoor === 'Unlocked'}
              onChange={() => toggleLockStatus('frontDoor')}
            />
            <span className="slider"></span>
          </label>
        </div>
        
        {/* Back Door Lock */}
        <div style={{ flex: '1', maxWidth:'250px', textAlign: 'center' }}>
          <p style={{color:"#95A4B6", marginRight:'20px'}}>Back Door</p>
          <img src={doorLockIcon} alt="Back Door" style={{ width: '50%', height: 'auto', marginBottom:'10px' }} /> {/* Replace with lock icon */}
          <p style={{color:"#95A4B6", marginRight:'20px'}}>{locksStatus.backDoor === 'Unlocked' ? 'Unlocked' : 'Locked'}</p>
          <label className="toggle" style={{ display: 'block', margin: 'auto' }}>
            <input
              type="checkbox"
              checked={locksStatus.backDoor === 'Unlocked'}
              onChange={() => toggleLockStatus('backDoor')}
            />
            <span className="slider"></span>
          </label>
        </div>
      </div>
      <div>
        <div style={{display: 'flex', flexDirection: 'row', justifyContent: 'space-between', alignItems: 'center', backgroundColor: '#173350', padding: '20px', paddingRight:"0px", borderRadius: '10px', margin: '0px'}}> 
          <div>
        <h3 style={{ color: "#95A4B6", marginBottom: '20px' }}>Lights</h3>
         </div>
          <div>
        <button onClick={() => navigate('/lighting')} style={{ backgroundColor: '#0294A5', color: 'white', border: 'none', borderRadius: '5px', padding: '8px 18px', cursor: 'pointer', display: 'block', marginLeft: 'auto', marginRight: 'auto', fontSize:'11px' }}>
          See More
        </button>
        </div>
        </div>
        <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-around' }}>
          <img src={bulbIcon} alt="Bedroom Light" style={{ width: '60px', height: 'auto', marginRight: '30px', marginBottom: '10px' }} />
          <div style={{ textAlign: 'center' }}>
            <span style={{ color: "#95A4B6", fontSize: '20px' }}>Livingroom light:</span>
            <p style={{ marginTop: '10px' }}>{lightStatus}</p>
          </div>
          <label className="toggle" style={{ display: 'block', margin: 'auto' }}>
            <input
              type="checkbox"
              checked={lightStatus === 'Off'}
              onChange={toggleLightStatus}
            />
            <span className="slider"></span>
          </label>
        </div>
      </div>
    </div>
  );
};



 // Appliances Widget JSX
const AppliancesWidget = () => {
  return (
    <div>
    <div className="widget" style={{ flex: '1', minWidth: '200px', backgroundColor: '#173350', padding: '0px', borderRadius: '1px', margin: '10px', boxSizing: 'border-box', display: 'flex', justifyContent: 'space-around' }}>
      {/* Fridge Status */}
      <div style={{ textAlign: 'center' }}>
        <span style={{ color:"#95A4B6", margin:'6px' }}>Fridge</span>
        <p>{deviceData.Fridge.Status}</p>
      </div>

      {/* Dishwasher Status */}
      <div style={{ textAlign: 'center' }}>
        <span style={{ color:"#95A4B6", margin:'6px' }}>Dishwasher</span>
        <p>{deviceData.Dishwasher.Status}</p>
      </div>

      {/* Oven Status */}
      <div style={{ textAlign: 'center' }}>
        <span style={{ color:"#95A4B6", margin:'6px' }}>Oven</span>
        <p>{deviceData.Oven.Status}</p>
      </div>

      {/* Toaster Status */}
      <div style={{ textAlign: 'center' }}>
        <span style={{ color:"#95A4B6", margin:'6px' }}>Toaster</span>
        <p>{deviceData.Toaster.Status}</p>
      </div>
      
    </div>
   
    </div>
    
  );
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
            <span id='menuText'>{"May 5"}</span>
          </div>
          <div>
            <span id='menuText'>{"5 PM"}</span>
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
              <li className="nav-item"style={{ backgroundColor: '#08192B', margin: '0.5rem 0', padding: '0.5rem', borderLeft: '3px solid #0294A5' }}>
                <Link to="/dashboard" style={{ color: '#50BCC0', textDecoration: 'none' }}>
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
              <li className="nav-item"style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/hvac" style={{ color: '#95A4B6', textDecoration: 'none' }}>
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

    
<main style={{ flex: '1', padding: '1rem', display: 'flex', flexDirection: 'column', alignItems: 'center', backgroundColor: '#0E2237'}}>
  
  {/* Widgets Container */}
  <div style={{ width: '100%', maxWidth: '1200px' }}>
    {/* Widgets Row */}
    <div style={{ display: 'flex', flexDirection: 'row', flexWrap: 'wrap', justifyContent: 'center', gap: '20px', marginBottom: '20px' }}>
      {/* Camera Widget */}
      <div className="camera-widget" style={{ position: 'relative', maxWidth: '60%', backgroundColor: '#173350', borderRadius: '1px', overflow: 'hidden', flexBasis: '100%', padding:'12px' }}>
              {/* Camera Feed */}
              <img src={cameraFeeds[cameraView]} alt="Camera feed" style={{ width: '100%', height: '100%' }} />
              
              {/* Camera View Buttons */}
              <div style={{ position: 'absolute', top: '10px', left: '10px', display: 'flex', gap: '5px' }}>
                <button onClick={() => setCameraView('livingroom')} style={{ padding: '5px', color:'white',backgroundColor: cameraView === 'livingroom' ? '#0294A5' : '#08192B' }}>R1</button>
                <button onClick={() => setCameraView('kitchen')} style={{ padding: '5px', color:'white', backgroundColor: cameraView === 'kitchen' ? '#0294A5' : '#08192B' }}>R2</button>
                <button onClick={() => setCameraView('r3')} style={{ padding: '5px', color: 'white', backgroundColor: cameraView === 'r3' ? '#0294A5' : '#08192B' }}>R3</button>
              </div>
      </div>
      
      {/* Locks Widget */}
      <div className="widget" style={{ flex: '1', backgroundColor: '#173350', padding: '20px', borderRadius: '1px', margin: '10px', boxSizing: 'border-box' }}>
        <div style={{display: 'flex', flexDirection: 'row', justifyContent: 'space-between', alignItems: 'center', backgroundColor: '#173350', padding: '20px', paddingRight:"0px", borderRadius: '10px', margin: '0px'}}>
        <div>
        <h3 style={{color:"#95A4B6"}}>Locks</h3>
        </div>
        <div>
      <button onClick={() => navigate('/security')} style={{ backgroundColor: '#0294A5', color: 'white', border: 'none', borderRadius: '5px', padding: '8px 18px', cursor: 'pointer', display: 'block', marginLeft: 'auto', marginRight: 'auto', fontSize:'11px' }}>
        See More
      </button>
      </div>
      </div>
        <LocksWidget />
      </div>
    </div>

    {/* Another Row for More Widgets */}
    <div style={{ display: 'flex', flexDirection: 'row', flexWrap: 'wrap', justifyContent: 'center', gap: '20px' }}>
      {/* Status by Units Widget */}
      <div className="widget" style={{ flex: '1', minWidth: '290px', backgroundColor: '#173350', padding: '20px', borderRadius: '1px', margin: '10px', boxSizing: 'border-box' }}>
        <h3 style={{ marginBottom: '20px', color:"#95A4B6"}}>Status by Units</h3>
          {/* Content of the status by units widget */}
          <div style={{ display: 'flex', flexDirection: 'row', justifyContent: 'space-around'}}>
          <div style={{ textAlign:"center", margin:'6px'}}>
            <p style={{ color:"#95A4B6"}}>Power</p>
          {deviceData.HVAC.Temperature && (
              <p style={{ fontSize:'22px'}}>{deviceData.SolarPanel.EnergyGeneratedToday}kw</p>
          )}</div>
          <div style={{ textAlign:"center",  margin:'6px'}}>
            <p style={{ color:"#95A4B6"}}>Temperature</p>
          {deviceData.HVAC.Temperature && (
              <p style={{ fontSize:'22px'}}>{deviceData.HVAC.Temperature}</p>
          )}</div>

          <div style={{ textAlign:"center",  margin:'6px'}}>
            <p style={{ color:"#95A4B6"}}> Humidity</p>
          {deviceData.HVAC.Temperature && (
              <p style={{ fontSize:'22px'}}>{deviceData.HVAC.Humidity}</p>
          )}</div>

          <div style={{ textAlign:"center",  margin:'6px'}}>
            <p style={{ color:"#95A4B6"}}>Security</p>
          {deviceData.HVAC.Temperature && (
              <p style={{ color: 'green', fontWeight:'bold', fontSize:'22px'}}>{deviceData.SecuritySystem.Status}</p>
          )}</div>
          </div>

      </div>
      
      {/* Scheduled Activity Widget */}
      <div className="widget" style={{ flex: '1', minWidth: '250px', backgroundColor: '#173350', padding: '20px', borderRadius: '1px', margin: '10px', boxSizing: 'border-box' }}>
      <div style={{display: 'flex', flexDirection: 'row', justifyContent: 'space-between', alignItems: 'center', backgroundColor: '#173350', padding: '20px', paddingRight:"0px", borderRadius: '10px', margin: '0px'}}>
        <div>
        <h3 style={{color:"#95A4B6"}}>Active Appliances</h3>
        </div>
        <div>
      <button onClick={() => navigate('/appliances')} style={{ backgroundColor: '#0294A5', color: 'white', border: 'none', borderRadius: '5px', padding: '8px 18px', cursor: 'pointer', display: 'block', marginLeft: 'auto', marginRight: 'auto', fontSize:'11px' }}>
        See More
      </button>
      </div>
      </div>
        <AppliancesWidget />
      </div>
    </div>
  </div>
</main>

      </div>
    </div>
  );
};
export default Dashboard;