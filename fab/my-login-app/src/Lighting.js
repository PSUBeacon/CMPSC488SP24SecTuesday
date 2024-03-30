import React, {useState, useEffect} from 'react';
import {useNavigate} from 'react-router-dom';
import {Link} from 'react-router-dom'; // Import Link from react-router-dom for navigation
import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import logoImage from './logo.webp';
 
import removeIcon from './recycle-bin-icon.png';   
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


    useEffect(() => {
        const token = sessionStorage.getItem('token');
        const url = 'http://localhost:8081/lighting';

        if (!token) {
            navigate('/'); // Redirect to login page if token is not present
            return;
        }

        fetch(url, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
        })
            .then(response => response.json())
            .then(response => {
                if (response && response.data) {
                    setUser(response.data.user);
                    setAccountType(response.data.accountType);
                    sessionStorage.setItem('accountType', response.data.accountType);
                } else {
                    setError('Unexpected response from server');
                }
            })
            .catch(error => {
                console.log('Fetch operation error:', error)
            });
    }, [navigate]);

  const navigate = useNavigate(); // Instantiate useNavigate hook
  const [selectedLight, setSelectedLight] = useState(null);
    const handleTurnOn = () => {
        setIsLightOn(true);
        const serverUrl = 'http://localhost:8081/lighting';

        // Define the body of the request based on your Go server's expected input.
        const requestBody = {

            uuid: '417293',
            name: "Lighting",
            apptype: "Lighting",
            function: "Status",
            change: "true"

        };

        const token = sessionStorage.getItem('token')
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
        setTimeout(() => {
            console.log('Light turned on');
        }, 1000); //
    };

    const handleTurnOff = () => {
        setIsLightOn(false);
        const serverUrl = 'http://localhost:8081/lighting';

        // Define the body of the request based on your Go server's expected input.
        const requestBody = {
            uuid: '417293',
            name: "Lighting",
            apptype: "Lighting",
            function: "Status",
            change: "false",

        };

        const token = sessionStorage.getItem('token')
        // Send a POST request to turn the light on.
        axios.post(serverUrl, requestBody, {headers: {'Authorization': `Bearer ${token}`}})
            .then(response => {
                console.log(response.data);
            })
            .catch(error => {
                console.error('There was an error!', error);
            });
        setTimeout(() => {
            console.log('Light turned off');
        }, 1000);
    };
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
  


    
    
  useEffect(() => {
    // Retrieve lights from local storage
    const storedLights = JSON.parse(localStorage.getItem('lights'));
    if (storedLights) {
      setLights(storedLights);
    }
  }, []);
     
      
  



  
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
  const [lights, setLights] = useState([]);
  const uniqueRoomNames = [...new Set(lights.map(light => light.roomName))];
  const [roomName, setRoomName] = useState('');
  const [lightName, setLightName] = useState('');


  // Function to handle light removal
const handleRemoveLight = (index, roomName) => {
  const updatedLights = lights.filter((light, lightIndex) => {
    return !(lightIndex === index && light.roomName === roomName);
  });
  setLights(updatedLights);
};

  // Function to handle turning a light on
  const handleLightOn = (index) => {
    const updatedLights = [...lights];
    updatedLights[index].isOn = true;
    setLights(updatedLights);
  };

  // Function to handle turning a light off
  const handleLightOff = (index) => {
    const updatedLights = [...lights];
    updatedLights[index].isOn = false;
    setLights(updatedLights);
  };

  // Function to handle form submission when adding a light
  const handleFormSubmit = (e) => {
    e.preventDefault();
    if (roomName && lightName) {
      setLights([...lights, { roomName, lightName, isOn: false }]);
      setRoomName('');
      setLightName('');
    }
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
              <li className="nav-item"style={{ backgroundColor: '#08192B', margin: '0.5rem 0', padding: '0.5rem', borderLeft: '3px solid #0294A5' }}>
                <Link to="/lighting" style={{ color: '#50BCC0', textDecoration: 'none' }}>
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

    
        <main style={{ flex: '1', padding: '1rem', display: 'flex', flexDirection: 'column', alignItems: 'center', backgroundColor: '#0E2237', width: '100%'}}>
      <div className="contentBlock" style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start', width: '100%', paddingBottom: '60px'}}>
        <div className="roomSelection" style={{ flex: '1', display: 'flex', flexDirection: 'column', alignItems: 'center', marginRight: '20px' }}>
          <h3 className="centered-title">Selecting a Room</h3>
          <div className="RoomCards" style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', marginTop: '40px' }}>
            <div className="card" style={{ marginBottom: '20px', border: selectedRoom === "Bedroom 1" ? '2px solid #0294A5' : 'none' }} onClick={() => {setRoomName("Bedroom 1"); setSelectedRoom("Bedroom 1");}}><img class="images" src={bedroomIcon} alt="Room 1" /></div>
            <div className="card" style={{ marginBottom: '20px', border: selectedRoom === "Bedroom 2" ? '2px solid #0294A5' : 'none' }} onClick={() => {setRoomName("Bedroom 2"); setSelectedRoom("Bedroom 2");}}><img class="images" src={bedroomIcon} alt="Room 2" /></div>
            <div className="card" style={{ border: selectedRoom === "Living Room" ? '2px solid #0294A5' : 'none' }} onClick={() => {setRoomName("Living Room"); setSelectedRoom("Living Room");}}><img class="images" src={livingroomIcon} alt="Room 3" /></div>
          </div>
        </div>
        <div className="lightControl" style={{ flex: '1', display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
          <div className="formContainer" style={{ width: '100%', marginBottom: '20px', marginTop: '80px' }}>
            <form onSubmit={handleFormSubmit} className="lightForm">
              <div className="formGroup">
                <label htmlFor="roomName">Room Name:</label>
                <input type="text" id="roomName" value={roomName} onChange={(e) => setRoomName(e.target.value)} />
              </div>
              <div className="formGroup">
                <label htmlFor="lightName">Light Name:</label>
                <input type="text" id="lightName" value={lightName} onChange={(e) => setLightName(e.target.value)} />
              </div>
              <button type="submit" className="submitButton">Add Light</button>
            </form>
            <div className="roomDropdown" style={{ marginBottom: '20px', width:'72%' }}>
  <label htmlFor="selectRoom">Select Room:</label>
  <select id="selectRoom" onChange={(e) => setSelectedRoom(e.target.value)}>
    <option value="">Select Room</option>
    {uniqueRoomNames.map((room, index) => (
      <option key={index} value={room}>{room}</option>
    ))}
  </select>
</div>
            {selectedRoom && (
              <ul className="lightList">
                {lights
                  .filter((light) => light.roomName === selectedRoom)
                  .map((light, index) => (
                    <li key={index} className="lightItem">
                      {light.lightName}
                      <div>
                        <button style={{ marginRight: '10px' }} onClick={() => handleLightOn(index)}>Turn On</button>
                        <button onClick={() => handleLightOff(index)}>Turn Off</button>
                        <img src={removeIcon} alt="Remove" style={{ width: '20px', marginLeft: '10px', cursor: 'pointer' }} onClick={() => handleRemoveLight(index, selectedRoom)} />
                      </div>
                    </li>
                  ))}
              </ul>
            )}
            <div className="dimmerControl" style={{ width: '72%', textAlign: 'center', marginTop: '20px' }}>
              <input
                type="range"
                id="dimmer"
                name="dimmer"
                min="0"
                max="100"
                value={dimmerValue}
                onChange={(e) => setDimmerValue(e.target.value)}
                style={{
                  WebkitAppearance: 'none',
                  width: '100%',
                  height: '15px',
                  background: '#d3d3d3',
                  outline: 'none',
                  opacity: '0.7',
                  transition: 'opacity .2s',
                  borderRadius: '5px',
                }}
              />
              <label htmlFor="dimmer" style={{ color: '#fff', marginTop: '5px' }}>{dimmerValue}%</label>
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