
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

   const [dishwasher, setDishwasher] = useState([
    { Device: "Dishwasher", Name: "N/A", Location: 'N/A', Status: 'ON', LastUsed: 'MM/DD/YY 00:00', WashTime: "00:00" },
    { Device: "Dishwasher", Name: "N/A", Location: 'N/A', Status: 'ON', LastUsed: 'MM/DD/YY 00:00', WashTime: "00:00" },
  ]);
  

  //this const is also for the freezer
  const [fridge, setFridge] = useState([ // <-- Initialize fridge state and setter
    { Device: "Fridge", Name: "N/A", Location: 'N/A', Status: 'ON', LastUsed: 'MM/DD/YY 00:00', Temp: "°F", EnergySavingMode: "KW"},
    { Device: "Fridge", Name: "N/A", Location: 'N/A', Status: 'ON', LastUsed: 'MM/DD/YY 00:00', Temp: "°F", EnergySavingMode: "KW" },
    { Device: "Freezer", Name: "N/A", Location: 'N/A', Status: 'ON', LastUsed: 'MM/DD/YY 00:00', Temp: "°F", EnergySavingMode: "KW" },
    { Device: "Freezer", Name: "N/A", Location: 'N/A', Status: 'ON', LastUsed: 'MM/DD/YY 00:00', Temp: "°F", EnergySavingMode: "KW" },
  ]);

  const [hvac, setHvac] = useState([
    { Device: "HVAC", Name: "N/A", Location: 'N/A', Status: 'ON', LastUsed: 'MM/DD/YY 00:00' },
    { Device: "HVAC", Name: "N/A", Location: 'N/A', Status: 'ON', LastUsed: 'MM/DD/YY 00:00' },
  ]);
  

  const [lighting, setLighting] = useState([
    { Device: "Lighting", Name: "N/A", Location: 'N/A', Status: 'ON', LastUsed: 'MM/DD/YY 00:00' },
    { Device: "Lighting", Name: "N/A", Location: 'N/A', Status: 'ON', LastUsed: 'MM/DD/YY 00:00' },
  ]);
  

  const [microwave, setMicrowave] = useState([
    { Device: "Microwave", Name: "N/A", Location: 'N/A', Status: 'ON', LastUsed: 'MM/DD/YY 00:00', Power: "KW", StopTime: "00:00" },
    { Device: "Microwave", Name: "N/A", Location: 'N/A', Status: 'ON', LastUsed: 'MM/DD/YY 00:00', Power: "KW", StopTime: "00:00" },
  ]);  
  

  const [toaster, setToaster] = useState([
    { Device: "Toaster", Name: "N/A", Location: 'N/A', Status: 'ON', LastUsed: 'MM/DD/YY 00:00', Temp: "°F", StopTime: "00:00" },
    { Device: "Toaster", Name: "N/A", Location: 'N/A', Status: 'ON', LastUsed: 'MM/DD/YY 00:00', Temp: "°F", StopTime: "00:00" },
  ]);
  
  

  const [oven, setOven] = useState([
    { Device: "Oven", Name: "N/A", Location: 'N/A', Status: 'ON', LastUsed: 'MM/DD/YY 00:00', Temp: "°F", StopTime: "00:00"  },
    { Device: "Oven", Name: "N/A", Location: 'N/A', Status: 'ON', LastUsed: 'MM/DD/YY 00:00', Temp: "°F", StopTime: "00:00"  },
  ]);  
  
  
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
        {dishwasher.map((dishwasherItem, index) => (
  <tr key={index}>
    <td>{dishwasherItem.Device}</td>
    <td>{dishwasherItem.Name}</td>
    <td>{dishwasherItem.Location}</td>
    <td>
      <button
        onClick={() => {
          const updatedDishwasher = [...dishwasher];
          updatedDishwasher[index] = {
            ...dishwasherItem,
            Status: dishwasherItem.Status === 'ON' ? 'OFF' : 'ON'
          };
          setDishwasher(updatedDishwasher);
        }}
      >
        {dishwasherItem.Status}
      </button>
    </td>
    <td>{dishwasherItem.LastUsed}</td>
    <td>
      <input
        type="text"
        value={dishwasherItem.WashTime}
        onChange={(e) => {
          const newWashTime = e.target.value;
          const updatedDishwasher = [...dishwasher];
          updatedDishwasher[index] = {
            ...updatedDishwasher[index],
            WashTime: newWashTime
          };
          setDishwasher(updatedDishwasher);
        }}
      />
    </td>
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
        {fridge.map((fridgeItem, index) => (
  <tr key={index}>
    <td>{fridgeItem.Device}</td>
    <td>{fridgeItem.Name}</td>
    <td>{fridgeItem.Location}</td>
    {/* Replace the current display of "Status" with a button */}
    <td>
      <button
        onClick={() => {
          const updatedFridge = [...fridge];
          updatedFridge[index] = {
            ...fridgeItem,
            Status: fridgeItem.Status === 'ON' ? 'OFF' : 'ON'
          };
          setFridge(updatedFridge);
        }}
      >
        {fridgeItem.Status}
      </button>
    </td>
    <td>{fridgeItem.LastUsed}</td>
    <td>
      <input
        type="text"
        value={fridgeItem.Temp} // Bind the value of the input field to the state
        onChange={(e) => {
          const newTemp = e.target.value;
          const updatedFridge = [...fridge]; // Create a copy of the array
          updatedFridge[index] = { ...fridgeItem, Temp: newTemp }; // Update the specific item in the copied array
          setFridge(updatedFridge); // Update the state with the new array
        }}
      />
    </td>
    <td>
      <input
        type="text"
        value={fridgeItem.EnergySavingMode} // Bind the value of the input field to the state
        onChange={(e) => {
          const newEnergySavingMode = e.target.value;
          const updatedFridge = [...fridge]; // Create a copy of the array
          updatedFridge[index] = { ...fridgeItem, EnergySavingMode: newEnergySavingMode }; // Update the specific item in the copied array
          setFridge(updatedFridge); // Update the state with the new array
        }}
      />
    </td>
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
        {hvac.map((hvacItem, index) => (
  <tr key={index}>
    <td>{hvacItem.Device}</td>
    <td>{hvacItem.Name}</td>
    <td>{hvacItem.Location}</td>
    {/* Replace the current display of "Status" with a button */}
    <td>
      <button
        onClick={() => {
          const updatedHvac = [...hvac];
          updatedHvac[index] = {
            ...hvacItem,
            Status: hvacItem.Status === 'ON' ? 'OFF' : 'ON'
          };
          setHvac(updatedHvac);
        }}
      >
        {hvacItem.Status}
      </button>
    </td>
    <td>{hvacItem.LastUsed}</td>
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
        {lighting.map((lightingItem, index) => (
  <tr key={index}>
    <td>{lightingItem.Device}</td>
    <td>{lightingItem.Name}</td>
    <td>{lightingItem.Location}</td>
    {/* Replace the current display of "Status" with a button */}
    <td>
      <button
        onClick={() => {
          const updatedLighting = [...lighting];
          updatedLighting[index] = {
            ...lightingItem,
            Status: lightingItem.Status === 'ON' ? 'OFF' : 'ON'
          };
          setLighting(updatedLighting);
        }}
      >
        {lightingItem.Status}
      </button>
    </td>
    <td>{lightingItem.LastUsed}</td>
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
        {microwave.map((microwaveItem, index) => (
  <tr key={index}>
    <td>{microwaveItem.Device}</td>
    <td>{microwaveItem.Name}</td>
    <td>{microwaveItem.Location}</td>
    {/* Replace the current display of "Status" with a button */}
    <td>
      <button
        onClick={() => {
          const updatedMicrowave = [...microwave];
          updatedMicrowave[index] = {
            ...microwaveItem,
            Status: microwaveItem.Status === 'ON' ? 'OFF' : 'ON'
          };
          setMicrowave(updatedMicrowave);
        }}
      >
        {microwaveItem.Status}
      </button>
    </td>
    <td>{microwaveItem.LastUsed}</td>
    <td>
      <input
        type="text"
        value={microwaveItem.Power} // Bind the value of the input field to the state
        onChange={(e) => {
          const newPower = e.target.value;
          const updatedMicrowave = [...microwave]; // Create a copy of the array
          updatedMicrowave[index] = { ...microwaveItem, Power: newPower }; // Update the specific item in the copied array
          setMicrowave(updatedMicrowave); // Update the state with the new array
        }}
      />
    </td>
    <td>
      <input
        type="text"
        value={microwaveItem.StopTime} // Bind the value of the input field to the state
        onChange={(e) => {
          const newStopTime = e.target.value;
          const updatedMicrowave = [...microwave]; // Create a copy of the array
          updatedMicrowave[index] = { ...microwaveItem, StopTime: newStopTime }; // Update the specific item in the copied array
          setMicrowave(updatedMicrowave); // Update the state with the new array
        }}
      />
    </td>
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
        {toaster.map((toasterItem, index) => (
  <tr key={index}>
    <td>{toasterItem.Device}</td>
    <td>{toasterItem.Name}</td>
    <td>{toasterItem.Location}</td>
    {/* Replace the current display of "Status" with a button */}
    <td>
      <button
        onClick={() => {
          const updatedToaster = [...toaster];
          updatedToaster[index] = {
            ...toasterItem,
            Status: toasterItem.Status === 'ON' ? 'OFF' : 'ON'
          };
          setToaster(updatedToaster);
        }}
      >
        {toasterItem.Status}
      </button>
    </td>
    <td>{toasterItem.LastUsed}</td>
    <td>
      <input
        type="text"
        value={toasterItem.Temperature} // Bind the value of the input field to the state
        onChange={(e) => {
          const newTemperature = e.target.value;
          const updatedToaster = [...toaster]; // Create a copy of the array
          updatedToaster[index] = { ...toasterItem, Temperature: newTemperature }; // Update the specific item in the copied array
          setToaster(updatedToaster); // Update the state with the new array
        }}
      />
    </td>
    <td>
      <input
        type="text"
        value={toasterItem.StopTime} // Bind the value of the input field to the state
        onChange={(e) => {
          const newStopTime = e.target.value;
          const updatedToaster = [...toaster]; // Create a copy of the array
          updatedToaster[index] = { ...toasterItem, StopTime: newStopTime }; // Update the specific item in the copied array
          setToaster(updatedToaster); // Update the state with the new array
        }}
      />
    </td>
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
        {oven.map((ovenItem, index) => (
  <tr key={index}>
    <td>{ovenItem.Device}</td>
    <td>{ovenItem.Name}</td>
    <td>{ovenItem.Location}</td>
    {/* Replace the current display of "Status" with a button */}
    <td>
      <button
        onClick={() => {
          const updatedOven = [...oven];
          updatedOven[index] = {
            ...ovenItem,
            Status: ovenItem.Status === 'ON' ? 'OFF' : 'ON'
          };
          setOven(updatedOven);
        }}
      >
        {ovenItem.Status}
      </button>
    </td>
    <td>{ovenItem.LastUsed}</td>
    <td>
      <input
        type="text"
        value={ovenItem.Temperature} // Bind the value of the input field to the state
        onChange={(e) => {
          const newTemperature = e.target.value;
          const updatedOven = [...oven]; // Create a copy of the array
          updatedOven[index] = { ...ovenItem, Temperature: newTemperature }; // Update the specific item in the copied array
          setOven(updatedOven); // Update the state with the new array
        }}
      />
    </td>
    <td>
      <input
        type="text"
        value={ovenItem.StopTime} // Bind the value of the input field to the state
        onChange={(e) => {
          const newStopTime = e.target.value;
          const updatedOven = [...oven]; // Create a copy of the array
          updatedOven[index] = { ...ovenItem, StopTime: newStopTime }; // Update the specific item in the copied array
          setOven(updatedOven); // Update the state with the new array
        }}
      />
    </td>
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