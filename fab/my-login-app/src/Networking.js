
import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { Link } from 'react-router-dom'; // Import Link from react-router-dom for navigation
import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import logoImage from './logo.webp'; 
import houseImage from './houseImage.jpg';
import notificationIcon from './notification.png'
import settingsIcon from './settings.png'
import accountIcon from './account.png'
import menuIcon from './menu.png'
import placeholderImage from './placeholderImage.jpg'; // Replace with the path to your placeholder image
import placeholderImage2 from './placeholderImage2.jpg'; // Replace with the path to your placeholder image
import { Table } from 'react-bootstrap';

// Define the Dashboard component using a functional component pattern
const Networking = () => {



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

  const [currentView, setCurrentView] = useState('iotLogs'); // New state to track current view
  
  // Define your IoT Logs data similar to how you have solarPanel data
const iotLogs = [
  { Time: '00:00', Activity: 'Light Turned On'},
  { Time: '00:00', Activity: 'Light Turned Off' },
  { Time: '00:00', Activity: 'Light Turned Off' },
  { Time: '00:00', Activity: 'Light Turned Off' },
  { Time: '00:00', Activity: 'Light Turned Off' },
  // Add more IoT device log entries as needed
];
// Function to change view
const changeView = (view) => {
  setCurrentView(view);
};
const iotLogsTable = (
  <div>
    <h2 style={{ color: '#173350' }}>IOT Logs</h2>
    <Table striped bordered hover variant="dark" style={{ marginTop: '20px', backgroundColor: "#173350" }}>
      <thead>
        <tr>
          <th>Time</th>
          <th>Activity</th>
        </tr>
      </thead>
      <tbody>
        {iotLogs.map((log, index) => (
          <tr key={index}>
            <td>{log.Time}</td>
            <td>{log.Activity}</td>
          </tr>
        ))}
      </tbody>
    </Table>
  </div>
);

// Define your Network Logs data
const networkLogs = [
  { PortSentFrom: 'Port 80', PortSentTo: 'Port 20' },
  { PortSentFrom: 'Port 80', PortSentTo: 'Port 20' },
  { PortSentFrom: 'Port 80', PortSentTo: 'Port 20' },
  { PortSentFrom: 'Port 80', PortSentTo: 'Port 20' },
  { PortSentFrom: 'Port 80', PortSentTo: 'Port 20' },
  // Add more network log entries as needed
];

const networkLogsTable = (
  <div>
    <h2 style={{ color: '#173350' }}>Network Logs</h2>
    <Table striped bordered hover variant="dark" style={{ marginTop: '20px', backgroundColor: "#173350" }}>
      <thead>
        <tr>
          <th>Port Sent From</th>
          <th>Port Sent To</th>
        </tr>
      </thead>
      <tbody>
        {networkLogs.map((log, index) => (
          <tr key={index}>
            <td>{log.PortSentFrom}</td>
            <td>{log.PortSentTo}</td>
          </tr>
        ))}
      </tbody>
    </Table>
  </div>
);


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
    // Add more camera feeds as needed
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
              <li className="nav-item"style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
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
              <li className="nav-item"style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/lighting" style={{ color: 'white', textDecoration: 'none' }}>
                  <i className="fas fa-lightbulb" style={{ marginRight: '10px' }}></i>
                  Lighting
                </Link>
              </li>
              <li className="nav-item"style={{ backgroundColor: '#08192B', margin: '0.5rem 0', padding: '0.5rem', borderLeft: '3px solid #0294A5' }}>
                <Link to="/networking" style={{ color: '#50BCC0', textDecoration: 'none'}}>
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

    
<main style={{ flex: '1', padding: '1rem', display: 'flex', flexDirection: 'column', alignItems: 'center', backgroundColor: '#0E2237'}}>
  
  {/* Widgets Container */}
  <div style={{ width: '100%', maxWidth: '1700px' }}>
    {/* Widgets Row */}
    <div style={{ display: 'flex', flexDirection: 'row', flexWrap: 'wrap', justifyContent: 'center', gap: '20px', marginBottom: '20px' }}>
        {/* Camera Widget */}
        <div className="camera-widget" style={{ position: 'relative', maxWidth: '60%', backgroundColor: '#173350', borderRadius: '10px', overflow: 'hidden', flexBasis: '100%', padding: '12px' }}>
          {/* Conditionally render the table based on the current view */}
          {currentView === 'iotLogs' && iotLogsTable}
          {currentView === 'networkLogs' && networkLogsTable}
          
          
              
          {/* Camera View Buttons */}
          <div style={{ position: 'absolute', top: '10px', left: '10px', display: 'flex', gap: '5px'}}>
          <button onClick={() => changeView('iotLogs')} style={{ padding: '5px', color:'white',backgroundColor: currentView === 'iotLogs' ? '#0294A5' : '#08192B' }}>IOT Logs</button>
          <button onClick={() => changeView('networkLogs')} style={{ padding: '5px', color:'white', backgroundColor: currentView === 'networkLogs' ? '#0294A5' : '#08192B' }}>Network Logs</button>
            {/* Add more buttons for additional camera views */}
          </div>
      </div>
      
      
      
      
    </div>
  </div>
</main>

      </div>
    </div>
  );
};
export default Networking;