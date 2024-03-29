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

 // States for each device
 const [deviceData, setDeviceData] = useState({});

    useEffect(() => {
        // This useEffect is for loading stored data from localStorage
        const storedData = localStorage.getItem('devices');
        if (storedData) {
            setDeviceData(JSON.parse(storedData));
        }
    }, []); // The dependency array is empty, meaning this effect runs once on component mount


    // States for date and time
   const [currentDate, setCurrentDate] = useState(new Date().toLocaleDateString());
   const [currentTime, setCurrentTime] = useState(new Date().toLocaleTimeString());
   const [secondFloorMode, setSecondFloorMode] = useState(deviceData.HVAC?.[0]?.Mode || 'Heat');
    const [basementMode, setBasementMode] = useState(deviceData.HVAC?.[1]?.Mode || 'Heat');
    const [secondFloorTemp, setSecondFloorTemp] = useState(deviceData.HVAC?.[0]?.Temperature ? `${deviceData.HVAC[0].Temperature}°F` : '');
    const [basementTemp, setBasementTemp] = useState(deviceData.HVAC?.[1]?.Temperature ? `${deviceData.HVAC[1].Temperature}°F` : '');
    const [secondFloorHumidity, setSecondFloorHumidity] = useState(deviceData.HVAC?.[0]?.Humidity ? `${deviceData.HVAC[0].Humidity}%` : '');
   const [basementHumidity, setBasementHumidity] = useState(deviceData.HVAC?.[1]?.Humidity ? `${deviceData.HVAC[1].Humidity}%` : '');
   const [secondFloorFanSpeed, setSecondFloorFanSpeed] = useState(deviceData.HVAC?.[0]?.FanSpeed || '')
   const [basementFanSpeed, setBasementFanSpeed] =  useState(deviceData.HVAC?.[1]?.FanSpeed || '')
   const [secondFloorEnergyConsumption, setSecondFloorEnergyConsumption] = useState(deviceData.HVAC?.[0]?.EnergyConsumption ? `${deviceData.HVAC[0].EnergyConsumption}kWh` : '');
   const [basementEnergyConsumption, setBasementEnergyConsumption] =useState(deviceData.HVAC?.[1]?.EnergyConsumption ? `${deviceData.HVAC[1].EnergyConsumption}kWh` : '');
   const [secondFloorHVACStatus, setSecondFloorHVACStatus] = useState(deviceData?.HVAC?.[1]?.Status ? "On" : "Off");
  const [basementHVACStatus, setBasementHVACStatus] = useState(deviceData?.HVAC?.[0]?.Status ? "On" : "Off");


   const updateSecondFloorHVACMode = (newMode) => {
    setDeviceData(prevDeviceData => ({
        ...prevDeviceData,
        HVAC: {
            ...prevDeviceData.HVAC,
            secondFloorMode: newMode,
        },
    }));
};

const updateBasementHVACMode = (newMode) => {
    setDeviceData(prevDeviceData => ({
        ...prevDeviceData,
        HVAC: {
            ...prevDeviceData.HVAC,
            basementMode: newMode,
        },
    }));
};

const updateDeviceData = (key, value) => {
  setDeviceData(prevDeviceData => ({
    ...prevDeviceData,
    HVAC: {
      ...prevDeviceData.HVAC,
      [key]: value,
    },
  }));
};


 
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
    
    
  

   // Toggle function for 2nd floor HVAC status
   const toggleSecondFloorHVACStatus = () => {
    const newStatus = !secondFloorHVACStatus;
    setSecondFloorHVACStatus(newStatus);
    updateDeviceData('secondFloorStatus', newStatus);
};

const toggleBasementHVACStatus = () => {
    const newStatus = !basementHVACStatus;
    setBasementHVACStatus(newStatus);
    updateDeviceData('basementStatus', newStatus);
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
                if (data.devices) { // Check if the devices data is present in the response
                    // Update the entire device data state
                    setDeviceData(data.devices);


                    localStorage.setItem('devices', JSON.stringify(data.devices));
                }
                setDashboardMessage(data.message);

                // Update and store the account type
                setAccountType(data.accountType);
                sessionStorage.setItem('accountType', data.accountType);
            })
            .catch(error => console.error('Fetch operation error:', error));
    }, []); // Ensure the useEffect hook runs only once after the component mounts

      // Function to update HVAC status in device data
  const updateHVACStatus = (status) => {
    setDeviceData(prevDeviceData => ({
      ...prevDeviceData,
      HVAC: {
        ...prevDeviceData.HVAC,
        Status: status,
      },
    }));
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
    <div className="hvac-data-container" style={{ display: 'flex', flexDirection: 'row', justifyContent: 'space-around', width: '100%' }}>
      {/* First Column */}
      <div style={{ marginRight: '2rem' }}>
        <h4 style={{textAlign:"center", marginBottom:'20px', color: '#95A4B6'}}>2nd Floor</h4>
        
        {/* First Column - mode */}
        <div className="data-item">
    <div className="data-icon">
        <i className="fas fa-thermometer-half"></i>
    </div>
    <div className="data-info">
        <p style={{ color: '#95A4B6' }}>Mode</p>
        <select
    value={secondFloorMode}
    onChange={(e) => {
        setSecondFloorMode(e.target.value);
        updateSecondFloorHVACMode(e.target.value);
    }}
            style={{ 
                color: '#95A4B6',
                backgroundColor: '#08192B',
                border: 'none',
                borderRadius: '5px',
                padding: '5px',
                zIndex: 10, // Ensure dropdown is clickable
                position: 'relative' // Required for z-index to work
            }}
        >
            <option value="Heat">Heat</option>
            <option value="Cool">Cool</option>
        </select>
    </div>
</div>
        {/* First Column - temperature */}
           
        <div className="data-item">
  <div className="data-icon">
    <i className="fas fa-thermometer-half"></i>
  </div>
  <div className="data-info">
    <p style={{ color: '#95A4B6' }}>Temperature</p>
    <input 
      type="number" 
      value={secondFloorTemp} 
      onChange={(e) => setSecondFloorTemp(e.target.value)}
      onBlur={() => updateDeviceData('secondFloorTemp', secondFloorTemp)}
      style={{
        zIndex: '1', // Ensure the input is clickable
        position: 'relative', // Z-index works on positioned elements
        // Add more styles to fit your design
      }}
      disabled={false} // Ensure the input is not disabled
      readOnly={false} // Ensure the input is not read-only
    />
  </div>
</div>
{/* ========================================================================= */}

        {/* First Column - humidity */}
        <div className="data-item">
  <div className="data-icon">
    <i className="fas fa-tint"></i>
  </div>
  <div className="data-info">
    <p style={{ color: '#95A4B6' }}>Humidity</p>
    <input 
      type="number" 
      value={secondFloorHumidity} 
      onChange={(e) => setSecondFloorHumidity(e.target.value)}
      onBlur={() => updateDeviceData('secondFloorHumidity', secondFloorHumidity)}
      style={{
        zIndex: 2, // Ensure the input is clickable
        position: 'relative', // Z-index works on positioned elements
        // Add more styles to fit your design
      }}
      min="0"
      max="100"
      disabled={false} // Ensure the input is not disabled
      readOnly={false} // Ensure the input is not read-only
    />
  </div>
</div>
{/* ========================================================================= */}

        {/* First Column - fan speed */}
        <div className="data-item">
  <div className="data-icon">
    <i className="fas fa-fan"></i>
  </div>
  <div className="data-info">
    <p style={{ color: '#95A4B6'}}>Fan Speed</p>
    <input 
      type="number" 
      value={secondFloorFanSpeed} 
      onChange={(e) => setSecondFloorFanSpeed(e.target.value)}
      onBlur={() => updateDeviceData('secondFloorFanSpeed', secondFloorFanSpeed)}
      style={{
        zIndex: 2, // Ensure the input is clickable
        position: 'relative', // Z-index works on positioned elements
        // Add more styles to fit your design
      }}
      min="0"
      max="100" // Assuming fan speed is a percentage from 0 to 100
      disabled={false} // Ensure the input is not disabled
      readOnly={false} // Ensure the input is not read-only
    />
  </div>
</div>

{/* ========================================================================= */}

        {/* First Column - status */}
        <div className="data-item">
    <div className="data-icon">
        <i className={secondFloorHVACStatus ? "fas fa-power-on" : "fas fa-power-off"}></i>
    </div>
    <div className="data-info">
        <p style={{ color: '#95A4B6' }}>Status</p>
        <p>{secondFloorHVACStatus ? "On" : "Off"}</p>
        {/* Toggle */}
        <label className="toggle" style={{ display: 'block', margin: 'auto' }}>
            <input
                type="checkbox"
                checked={secondFloorHVACStatus}
                onChange={toggleSecondFloorHVACStatus}
            />
            <span className="slider" style={{ background: secondFloorHVACStatus ? '#50BCC0' : 'grey' }}></span>
        </label>
    </div>
</div>

{/* ========================================================================= */}

        {/* First Column - energy consumption */}
        <div className="data-item">
  <div className="data-icon">
    <i className="fas fa-bolt"></i>
  </div>
  <div className="data-info">
    <p style={{ color: '#95A4B6'}}>Energy Consumption</p>
    <input 
      type="number" 
      value={secondFloorEnergyConsumption} 
      onChange={(e) => setSecondFloorEnergyConsumption(e.target.value)}
      onBlur={() => updateDeviceData('secondFloorEnergyConsumption', secondFloorEnergyConsumption)}
      style={{
        zIndex: 2, // Make sure the input is above other elements
        position: 'relative', // Necessary for z-index to take effect
        // Define additional styles here
      }}
      min="0" // Set minimum value as needed
      // You can also add step attribute to control the valid steps
    />
  </div>
</div>

</div>

{/* ============================================================== */}

      {/* Second Column (Basement Floor) */}
      {/* Second Column - Mode */}
      <div>
      <h4 style={{textAlign:"center", marginBottom:'20px', color: '#95A4B6'}}>Basement</h4>

      <div className="data-item">
    <div className="data-icon">
        <i className="fas fa-thermometer-half"></i>
    </div>
    <div className="data-info">
        <p style={{ color: '#95A4B6' }}>Mode</p>
        <select
    value={basementMode}
    onChange={(e) => {
        setBasementMode(e.target.value);
        updateBasementHVACMode(e.target.value);
    }}
            style={{ 
                color: '#95A4B6',
                backgroundColor: '#08192B',
                border: 'none',
                borderRadius: '5px',
                padding: '5px',
                zIndex: 10, // Ensure dropdown is clickable
                position: 'relative' // Required for z-index to work
            }}
        >
            <option value="Heat">Heat</option>
            <option value="Cool">Cool</option>
        </select>
    </div>
</div>
{/* ============================================================== */}

{/* Second Column - Temperature */}
<div className="data-item">
  <div className="data-icon">
    <i className="fas fa-thermometer-half"></i>
  </div>
  <div className="data-info">
    <p style={{ color: '#95A4B6' }}>Temperature</p>
    <input 
      type="number" 
      value={basementTemp} 
      onChange={(e) => setBasementTemp(e.target.value)}
      onBlur={() => updateDeviceData('basementTemp', basementTemp)}
      style={{
        zIndex: '1', // Ensure the input is clickable
        position: 'relative', // Z-index works on positioned elements
        // Add more styles to fit your design
      }}
      disabled={false} // Ensure the input is not disabled
      readOnly={false} // Ensure the input is not read-only
    />
  </div>
</div>
{/* ============================================================== */}

  {/* Second Column - Humidity */}   
<div className="data-item">
  <div className="data-icon">
    <i className="fas fa-tint"></i>
  </div>
  <div className="data-info">
    <p style={{ color: '#95A4B6' }}>Humidity</p>
    <input 
      type="number" 
      value={basementHumidity} 
      onChange={(e) => setBasementHumidity(e.target.value)}
      onBlur={() => updateDeviceData('basementHumidity', basementHumidity)}
      style={{
        zIndex: 2, // Ensure the input is clickable
        position: 'relative', // Z-index works on positioned elements
        // Add more styles to fit your design
      }}
      min="0"
      max="100"
      disabled={false} // Ensure the input is not disabled
      readOnly={false} // Ensure the input is not read-only
    />
  </div>
</div>
{/* ============================================================== */}

{/* Second Column - Fan Speed */}
<div className="data-item">
  <div className="data-icon">
    <i className="fas fa-fan"></i>
  </div>
  <div className="data-info">
    <p style={{ color: '#95A4B6'}}>Fan Speed</p>
    <input 
      type="number" 
      value={basementFanSpeed} 
      onChange={(e) => setBasementFanSpeed(e.target.value)}
      onBlur={() => updateDeviceData('basementFanSpeed', basementFanSpeed)}
      style={{
        zIndex: 2, // Ensure the input is clickable
        position: 'relative', // Z-index works on positioned elements
        // Add more styles to fit your design
      }}
      min="0"
      max="100" // Assuming fan speed is a percentage from 0 to 100
      disabled={false} // Ensure the input is not disabled
      readOnly={false} // Ensure the input is not read-only
    />
  </div>
</div>
{/* ============================================================== */}

{/* Second Column - Status */}
<div className="data-item">
    <div className="data-icon">
        <i className={basementHVACStatus ? "fas fa-power-on" : "fas fa-power-off"}></i>
    </div>
    <div className="data-info">
        <p style={{ color: '#95A4B6' }}>Status</p>
        <p>{basementHVACStatus ? "On" : "Off"}</p>
        {/* Toggle */}
        <label className="toggle" style={{ display: 'block', margin: 'auto' }}>
            <input
                type="checkbox"
                checked={basementHVACStatus}
                onChange={toggleBasementHVACStatus}
            />
            <span className="slider" style={{ background: basementHVACStatus ? '#50BCC0' : 'grey' }}></span>
        </label>
    </div>
</div>

{/* ============================================================== */}

        {/* Second Column - Energy Consumption */}
        <div className="data-item">
  <div className="data-icon">
    <i className="fas fa-bolt"></i>
  </div>
  <div className="data-info">
    <p style={{ color: '#95A4B6'}}>Energy Consumption</p>
    <input 
      type="number" 
      value={basementEnergyConsumption} 
      onChange={(e) => setBasementEnergyConsumption(e.target.value)}
      onBlur={() => updateDeviceData('basementEnergyConsumption', basementEnergyConsumption)}
      style={{
        zIndex: 2, // Make sure the input is above other elements
        position: 'relative', // Necessary for z-index to take effect
        // Define additional styles here
      }}
      min="0" // Set minimum value as needed
      // You can also add step attribute to control the valid steps
    />
  </div>
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
