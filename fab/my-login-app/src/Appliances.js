
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
const Appliances = () => {

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

  const dishwasher = [
    { Device: "Dishwasher", Name: "N/A", Location: 'N/A', Status: 'ON/OFF', LastUsed: 'MM/DD/YY 00:00', WashTime: "00:00" },
    { Device: "Dishwasher", Name: "N/A", Location: 'N/A', Status: 'ON/OFF', LastUsed: 'MM/DD/YY 00:00', WashTime: "00:00" },
  ];

  //this const is also for the freezer
  const fridge = [
    { Device: "Fridge", Name: "N/A", Location: 'N/A', Status: 'ON/OFF', LastUsed: 'MM/DD/YY 00:00', Temp: "Degrees", EnergySavingMode: "KW"},
    { Device: "Fridge", Name: "N/A", Location: 'N/A', Status: 'ON/OFF', LastUsed: 'MM/DD/YY 00:00', Temp: "Degrees", EnergySavingMode: "KW" },
    { Device: "Freezer", Name: "N/A", Location: 'N/A', Status: 'ON/OFF', LastUsed: 'MM/DD/YY 00:00', Temp: "Degrees", EnergySavingMode: "KW" },
    { Device: "Freezer", Name: "N/A", Location: 'N/A', Status: 'ON/OFF', LastUsed: 'MM/DD/YY 00:00', Temp: "Degrees", EnergySavingMode: "KW" },
  ];

  const hvac = [
    { Device: "HVAC", Name: "N/A", Location: 'N/A', Status: 'ON/OFF', LastUsed: 'MM/DD/YY 00:00' },
    { Device: "HVAC", Name: "N/A", Location: 'N/A', Status: 'ON/OFF', LastUsed: 'MM/DD/YY 00:00' },
  ];

  const lighting = [
    { Device: "Lighting", Name: "N/A", Location: 'N/A', Status: 'ON/OFF', LastUsed: 'MM/DD/YY 00:00' },
    { Device: "Lighting", Name: "N/A", Location: 'N/A', Status: 'ON/OFF', LastUsed: 'MM/DD/YY 00:00' },
  ];

  const microwave = [
    { Device: "Microwave", Name: "N/A", Location: 'N/A', Status: 'ON/OFF', LastUsed: 'MM/DD/YY 00:00', Power: "KW", StopTime: "00:00" },
    { Device: "Microwave", Name: "N/A", Location: 'N/A', Status: 'ON/OFF', LastUsed: 'MM/DD/YY 00:00', Power: "KW", StopTime: "00:00" },
  ];

  const toaster = [
    { Device: "Toaster", Name: "N/A", Location: 'N/A', Status: 'ON/OFF', LastUsed: 'MM/DD/YY 00:00', Temp: "Degrees", StopTime: "00:00" },
    { Device: "Toaster", Name: "N/A", Location: 'N/A', Status: 'ON/OFF', LastUsed: 'MM/DD/YY 00:00', Temp: "Degrees", StopTime: "00:00" },
  ];

  const oven = [
    { Device: "Oven", Name: "N/A", Location: 'N/A', Status: 'ON/OFF', LastUsed: 'MM/DD/YY 00:00', Temp: "Degrees", StopTime: "00:00"  },
    { Device: "Oven", Name: "N/A", Location: 'N/A', Status: 'ON/OFF', LastUsed: 'MM/DD/YY 00:00', Temp: "Degrees", StopTime: "00:00"  },
  ];
  
  const navigate = useNavigate(); // Instantiate useNavigate hook
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
  
  const CameraWidget = () => {
    const [cameraView, setCameraView] = useState('livingroom'); // Default view
    
    const cameraFeeds = {
      livingroom: placeholderImage, // Replace with actual video feed or image
      kitchen: placeholderImage, // Replace with actual video feed or image for kitchen
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
          {/* Add more buttons for additional camera views as needed */}
        </div>
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
              <li className="nav-item"style={{ backgroundColor: '#08192B', margin: '0.5rem 0', padding: '0.5rem', borderLeft: '3px solid #0294A5' }}>
                <Link to="/appliances" style={{  color: '#50BCC0', textDecoration: 'none' }}>
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

{/*dishwasher table*/}    
<main style={{ flex: '1', padding: '1rem', display: 'flex', flexDirection: 'column', alignItems: 'center', backgroundColor: '#0E2237'}}>
<h2 style={{ color: 'white' }}>Appliances</h2>
<div style={{ alignSelf: 'flex-start', width: '100%' }}>
    <h3 style={{ color: 'white', marginLeft: '1rem' }}>Dishwasher</h3>
  </div>
<Table striped bordered hover variant="dark" style={{ marginTop: '20px', backgroundColor: "#173350" }}>
        <thead>
          <tr>
            <th>Device</th>
            <th>Name</th>
            <th>Location</th>
            <th>Status</th>
            <th>Last Used</th>
            <th>Wash Time</th>
          </tr>
        </thead>
        <tbody>
          {dishwasher.map((dishwasher, index) => (
            <tr key={index}>
              <td>{dishwasher.Device}</td>
              <td>{dishwasher.Name}</td>
              <td>{dishwasher.Location}</td>
              <td>{dishwasher.Status}</td>
              <td>{dishwasher.LastUsed}</td>
              <td>{dishwasher.WashTimer}</td>
            </tr>
          ))}
        </tbody>
      </Table>

{/*fridge table*/}
      <div style={{ alignSelf: 'flex-start', width: '100%' }}>
    <h3 style={{ color: 'white', marginLeft: '1rem', marginTop: '10px' }}>Fridge</h3>
  </div>
<Table striped bordered hover variant="dark" style={{ marginTop: '20px', backgroundColor: "#173350" }}>
        <thead>
          <tr>
            <th>Device</th>
            <th>Name</th>
            <th>Location</th>
            <th>Status</th>
            <th>Last Used</th>
            <th>Temperature</th>
            <th>Energy Saving Mode</th>
          </tr>
        </thead>
        <tbody>
          {fridge.map((fridge, index) => (
            <tr key={index}>
              <td>{fridge.Device}</td>
              <td>{fridge.Name}</td>
              <td>{fridge.Location}</td>
              <td>{fridge.Status}</td>
              <td>{fridge.LastUsed}</td>
              <td>{fridge.Temp}</td>
              <td>{fridge.EnergySavingMode}</td>
            </tr>
          ))}
        </tbody>
      </Table>

{/*hvac table*/}
<div style={{ alignSelf: 'flex-start', width: '100%' }}>
    <h3 style={{ color: 'white', marginLeft: '1rem', marginTop: '10px' }}>HVAC</h3>
  </div>
<Table striped bordered hover variant="dark" style={{ marginTop: '20px', backgroundColor: "#173350" }}>
        <thead>
          <tr>
            <th>Device</th>
            <th>Name</th>
            <th>Location</th>
            <th>Status</th>
            <th>Last Used</th>
          </tr>
        </thead>
        <tbody>
          {hvac.map((hvac, index) => (
            <tr key={index}>
              <td>{hvac.Device}</td>
              <td>{hvac.Name}</td>
              <td>{hvac.Location}</td>
              <td>{hvac.Status}</td>
              <td>{hvac.LastUsed}</td>
            </tr>
          ))}
        </tbody>
      </Table>

{/*Lighting table*/}
<div style={{ alignSelf: 'flex-start', width: '100%' }}>
    <h3 style={{ color: 'white', marginLeft: '1rem', marginTop: '10px' }}>Lighting</h3>
  </div>
<Table striped bordered hover variant="dark" style={{ marginTop: '20px', backgroundColor: "#173350" }}>
        <thead>
          <tr>
            <th>Device</th>
            <th>Name</th>
            <th>Location</th>
            <th>Status</th>
            <th>Last Used</th>
          </tr>
        </thead>
        <tbody>
          {lighting.map((lighting, index) => (
            <tr key={index}>
              <td>{lighting.Device}</td>
              <td>{lighting.Name}</td>
              <td>{lighting.Location}</td>
              <td>{lighting.Status}</td>
              <td>{lighting.LastUsed}</td>
            </tr>
          ))}
        </tbody>
      </Table>

{/*Microwave table*/}
<div style={{ alignSelf: 'flex-start', width: '100%' }}>
    <h3 style={{ color: 'white', marginLeft: '1rem', marginTop: '10px' }}>Microwave</h3>
  </div>
<Table striped bordered hover variant="dark" style={{ marginTop: '20px', backgroundColor: "#173350" }}>
        <thead>
          <tr>
            <th>Device</th>
            <th>Name</th>
            <th>Location</th>
            <th>Status</th>
            <th>Last Used</th>
            <th>Power</th>
            <th>Stop Time</th>
          </tr>
        </thead>
        <tbody>
          {microwave.map((microwave, index) => (
            <tr key={index}>
              <td>{microwave.Device}</td>
              <td>{microwave.Name}</td>
              <td>{microwave.Location}</td>
              <td>{microwave.Status}</td>
              <td>{microwave.LastUsed}</td>
              <td>{microwave.Power}</td>
              <td>{microwave.StopTime}</td>
            </tr>
          ))}
        </tbody>
      </Table>

{/*Toaster table*/}
<div style={{ alignSelf: 'flex-start', width: '100%' }}>
    <h3 style={{ color: 'white', marginLeft: '1rem', marginTop: '10px' }}>Toaster</h3>
  </div>
<Table striped bordered hover variant="dark" style={{ marginTop: '20px', backgroundColor: "#173350" }}>
        <thead>
          <tr>
            <th>Device</th>
            <th>Name</th>
            <th>Location</th>
            <th>Status</th>
            <th>Last Used</th>
            <th>Temperature</th>
            <th>Stop Time</th>
          </tr>
        </thead>
        <tbody>
          {toaster.map((toaster, index) => (
            <tr key={index}>
              <td>{toaster.Device}</td>
              <td>{toaster.Name}</td>
              <td>{toaster.Location}</td>
              <td>{toaster.Status}</td>
              <td>{toaster.LastUsed}</td>
              <td>{toaster.Temp}</td>
              <td>{toaster.StopTime}</td>
            </tr>
          ))}
        </tbody>
      </Table>

{/*Oven table*/}
<div style={{ alignSelf: 'flex-start', width: '100%' }}>
    <h3 style={{ color: 'white', marginLeft: '1rem', marginTop: '10px' }}>Oven</h3>
  </div>
<Table striped bordered hover variant="dark" style={{ marginTop: '20px', backgroundColor: "#173350" }}>
        <thead>
          <tr>
            <th>Device</th>
            <th>Name</th>
            <th>Location</th>
            <th>Status</th>
            <th>Last Used</th>
            <th>Temperature</th>
            <th>Stop Time</th>
          </tr>
        </thead>
        <tbody>
          {oven.map((oven, index) => (
            <tr key={index}>
              <td>{oven.Device}</td>
              <td>{oven.Name}</td>
              <td>{oven.Location}</td>
              <td>{oven.Status}</td>
              <td>{oven.LastUsed}</td>
              <td>{oven.Temp}</td>
              <td>{oven.StopTime}</td>
            </tr>
          ))}
        </tbody>
      </Table>


  
</main>

      </div>
    </div>
  );
};


export default Appliances;