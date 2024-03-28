
import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { Link } from 'react-router-dom'; // Import Link from react-router-dom for navigation
import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import logoImage from './logo.webp'; 
import houseImage from './houseImage.jpg';
import { Table } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import settingsIcon from './settings.png'
import accountIcon from './account.png'
import menuIcon from './menu.png'
import placeholderImage from './placeholderImage.jpg'; // Replace with the path to your placeholder image
import placeholderImage2 from './placeholderImage2.jpg'; // Replace with the path to your placeholder image
import {
  faMicrophone, // Placeholder icon, replace with the actual icon for the microwave
  faOtter, // Placeholder icon, replace with the actual icon for the oven
  faIceCream, // Placeholder icon, replace with the actual icon for the fridge
  faSnowflake, // Placeholder icon, replace with the actual icon for the freezer
  faBreadSlice, // Placeholder icon, replace with the actual icon for the toaster
  faSoap // Placeholder icon, replace with the actual icon for the dishwasher
} from '@fortawesome/free-solid-svg-icons';
import 'bootstrap/dist/css/bootstrap.min.css';
// Define the Dashboard component using a functional component pattern
const Energy = () => {

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

  const energy = [
    { Device: 'Microwave', NetlossEnergy: '%', NetgainEnergy: '%', Battery: '%', Status: 'ON/OFF' },
    { Device: 'Oven', NetlossEnergy: '%', NetgainEnergy: '%', Battery: '%', Status: 'ON/OFF' },
    { Device: 'Fridge', NetlossEnergy: '%', NetgainEnergy: '%', Battery: '%', Status: 'ON/OFF' },
    { Device: 'Freezer', NetlossEnergy: '%', NetgainEnergy: '%', Battery: '%', Status: 'ON/OFF' },
    { Device: 'Toaster', NetlossEnergy: '%', NetgainEnergy: '%', Battery: '%', Status: 'ON/OFF' },
    { Device: 'Dishwasher', NetlossEnergy: '%', NetgainEnergy: '%', Battery: '%', Status: 'ON/OFF' },
    
  ];

  {/*solar panel statistics table*/}
  const solarPanel = [
    { TotalEnergy: 'KW', EnergyUsedToday: 'KW' },
    { TotalEnergy: 'KW', EnergyUsedToday: 'KW'},
    
  ];
  
  const navigate = useNavigate(); // Instantiate useNavigate hook
  const handleGoToAppliances = () => {
    navigate('/appliances'); // Adjust the route as necessary
  };
  const [isNavVisible, setIsNavVisible] = useState(false);

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
      <div class = "accountPop"style={{
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
        <p>Admin</p> {/* Dynamically display user role */}
        <button onClick={signOut} class="signout">Sign Out</button>
      </div>
    );
  };
  
  


  // This is the JSX return statement where we layout our component's HTML structure
  return (
    <div style={{ display: 'flex', minHeight: '100vh', flexDirection: 'column', backgroundColor: '#081624' }}>
      {/* Top Navbar */}
      <nav class="topNav" style={{ backgroundColor: '#081624', color: 'white', padding: '0.5rem 1rem' }}>
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
              <li className="nav-item"style={{margin: '0.5rem 0', padding: '0.5rem'}}>
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
              <li className="nav-item"style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                <Link to="/hvac" style={{ color: '#95A4B6', textDecoration: 'none' }}>
                  <i className="fas fa-thermometer-half" style={{ marginRight: '10px' }}></i>
                  HVAC
                </Link>
              </li>
              <li className="nav-item"style={{margin: '0.5rem 0', padding: '0.5rem'}}>
                <Link to="/appliances" style={{  color: '#95A4B6', textDecoration: 'none' }}>
                  <i className="fas fa-blender" style={{ marginRight: '10px' }}></i>
                  Appliances
                </Link>
              </li>
              <li className="nav-item"style={{ backgroundColor: '#08192B', margin: '0.5rem 0', padding: '0.5rem', borderLeft: '3px solid #0294A5' }}>
                <Link to="/energy" style={{ color: '#50BCC0', textDecoration: 'none' }}>
                  <i className="fas fa-bolt" style={{ marginRight: '10px' }}></i>
                  Energy
                </Link>
              </li>
            </ul>
          </nav>
        </aside>


        <main style={{ flex: '1', padding: '1rem', display: 'flex', flexDirection: 'column', alignItems: 'center', backgroundColor: '#0E2237'}}>
  <h2 style={{ color: 'white' }}>Devices Using Energy</h2>
  <Table striped bordered hover variant="dark" style={{ marginTop: '20px', backgroundColor: "#173350" }}>
    <thead>
      <tr>
        <th>
          Device
          <button onClick={handleGoToAppliances} style={{ marginLeft: '10px', padding: '2px 6px', fontSize: '0.8em', background: '#0294A5', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer' }}>
            See More
          </button>
        </th>
        <th>Net loss Energy</th>
        <th>Net Gain Energy</th>
        <th>Battery %</th>
        <th>Status</th>
      </tr>
    </thead>
    <tbody>
          {energy.map((energy, index) => (
            <tr key={index}>
              <td>{energy.Device}</td>
              <td>{energy.NetlossEnergy}</td>
              <td>{energy.NetgainEnergy}</td>
              <td>{energy.Battery}</td>
              <td>{energy.Status}</td>
            </tr>
          ))}
        </tbody>
      </Table>

      {/*this is the table for the solar panel content*/}
      <h2 style={{ color: 'white' }}>Solar Panel Statistics</h2>
      <Table striped bordered hover variant="dark" style={{ marginTop: '20px', backgroundColor: "#173350" }}>
        <thead>
          <tr>
            <th>Total Energy</th>
            <th>Energy Used Today</th>
          </tr>
        </thead>
        <tbody>
          {solarPanel.map((solarPanel, index) => (
            <tr key={index}>
              
              <td>{solarPanel.TotalEnergy}</td>
              <td>{solarPanel.EnergyUsedToday}</td>
            </tr>
          ))}
        </tbody>
      </Table>
  
</main>

      </div>
    </div>
  );
};


export default Energy;