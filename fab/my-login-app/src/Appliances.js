import React, {useEffect, useState} from 'react';
import {useNavigate} from 'react-router-dom';
import {Link} from 'react-router-dom'; // Import Link from react-router-dom for navigation
import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import logoImage from './logo.webp';
import houseImage from './houseImage.jpg';
import {Table} from 'react-bootstrap';
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome';
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
    const [user, setUser] = useState(null);
    const [error, setError] = useState('');
    const [accountType, setAccountType] = useState('')
    const navigate = useNavigate(); // Instantiate useNavigate hook
    const [isNavVisible, setIsNavVisible] = useState(false);

    const appliances = [
        {icon: faMicrophone, name: 'Microwave', location: 'Kitchen', status: 'ON/OFF', lastUsed: 'MM/DD/YY 00:00'},
        {icon: faOtter, name: 'Oven', location: 'Kitchen', status: 'ON/OFF', lastUsed: 'MM/DD/YY 00:00'},
        {icon: faIceCream, name: 'Fridge', location: 'Kitchen', status: 'ON/OFF', lastUsed: 'MM/DD/YY 00:00'},
        {icon: faSnowflake, name: 'Freezer', location: 'Kitchen', status: 'ON/OFF', lastUsed: 'MM/DD/YY 00:00'},
        {icon: faBreadSlice, name: 'Toaster', location: 'Kitchen', status: 'ON/OFF', lastUsed: 'MM/DD/YY 00:00'},
        {icon: faSoap, name: 'Dishwasher', location: 'Kitchen', status: 'ON/OFF', lastUsed: 'MM/DD/YY 00:00'},
    ];

    useEffect(() => {
        const token = sessionStorage.getItem('token');
        const url = 'http://localhost:8081/appliances';

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

    const AccountPopup = ({isVisible, onClose}) => {
        if (!isVisible) return null;

        return (
            <div className="accountPop" style={{
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
                <button onClick={signOut} className="signout">Sign Out</button>
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
            <div className="camera-widget" style={{
                position: 'relative',
                maxWidth: '100%',
                backgroundColor: '#12232E',
                borderRadius: '10px',
                overflow: 'hidden'
            }}>
                {/* Live Feed */}
                <img src={cameraFeeds[cameraView]} alt="Live feed"
                     style={{width: '100%', height: 'auto', display: 'block'}}/>

                {/* Camera View Buttons */}
                <div style={{position: 'absolute', top: '10px', left: '10px', display: 'flex', gap: '5px'}}>
                    <button onClick={() => setCameraView('livingroom')} style={{
                        padding: '5px',
                        backgroundColor: cameraView === 'livingroom' ? '#4CAF50' : 'transparent'
                    }}>R1
                    </button>
                    <button onClick={() => setCameraView('kitchen')} style={{
                        padding: '5px',
                        backgroundColor: cameraView === 'kitchen' ? '#4CAF50' : 'transparent'
                    }}>R2
                    </button>
                    {/* Add more buttons for additional camera views as needed */}
                </div>
            </div>
        );
    };

    // This is the JSX return statement where we layout our component's HTML structure
    return (
        <div style={{display: 'flex', minHeight: '100vh', flexDirection: 'column', backgroundColor: '#081624'}}>
            {/* Top Navbar */}
            <nav className="topNav" style={{backgroundColor: '#081624', color: 'white', padding: '0.5rem 1rem'}}>
                <div style={{display: 'flex', justifyContent: 'space-between', alignItems: 'center'}}>
                    <div style={{display: 'flex', alignItems: 'center'}}>
                        <img src={menuIcon} alt="Menu" onClick={toggleNav} className="hamburger-menu"/>
                        <img src={logoImage} alt="Logo" style={{marginRight: '10px'}}
                             id='circle'/> {/* Adjust the height as needed */}
                        <span id='menuText2'>Beacon</span>
                    </div>
                    <div>
                        <span id='menuText'>March 05, 2024</span>
                    </div>
                    <div>
                        <span id='menuText'>11:48 AM</span>
                    </div>
                    <div>
                        <div style={{position: 'relative'}}>
                            <img src={settingsIcon} alt="Settings" style={{marginRight: '10px'}} id="menuIcon"
                                 onClick={goToSettings}/>
                            <button onClick={toggleAccountPopup}
                                    style={{background: 'none', border: 'none', padding: 0, cursor: 'pointer'}}>
                                <img src={accountIcon} alt="account" style={{marginRight: '10px'}} id="menuIcon2"/>
                            </button>
                            <AccountPopup isVisible={isAccountPopupVisible}
                                          onClose={() => setIsAccountPopupVisible(false)}/>
                        </div>
                    </div>
                </div>
            </nav>

            {/* Side Navbar and Dashboard Content */}
            <div style={{display: 'flex', flex: '1'}}>
                {/* Side Navbar */}
                <aside className={`side-nav ${isNavVisible ? '' : 'hidden'}`}
                       style={{backgroundColor: '#0E2237', color: 'white', width: '250px', padding: '1rem'}}>
                    <div className="houseInfo">
                        <div><img src={houseImage} alt="Logo" style={{marginRight: '10px'}} id='circle2'/></div>
                        <div>My House</div>
                        <div>State College, PA 16801</div>
                    </div>
                    <nav>
                        <ul style={{listStyle: 'none', padding: 0}}>
                            {/* Apply active style to 'Overview' since it's the current page */}
                            <li className="nav-item" style={{margin: '0.5rem 0', padding: '0.5rem'}}>
                                <Link to="/dashboard" style={{color: 'white', textDecoration: 'none'}}>
                                    <i className="fas fa-home" style={{marginRight: '10px'}}></i>
                                    Overview
                                </Link>
                            </li>
                            <li className="nav-item" style={{margin: '0.5rem 0', padding: '0.5rem'}}>
                                <Link to="/security" style={{color: 'white', textDecoration: 'none'}}>
                                    <i className="fas fa-lock" style={{marginRight: '10px'}}></i>
                                    Security
                                </Link>
                            </li>
                            <li className="nav-item" style={{margin: '0.5rem 0', padding: '0.5rem'}}>
                                <Link to="/lighting" style={{color: 'white', textDecoration: 'none'}}>
                                    <i className="fas fa-lightbulb" style={{marginRight: '10px'}}></i>
                                    Lighting
                                </Link>
                            </li>
                            <li className="nav-item" style={{margin: '0.5rem 0', padding: '0.5rem'}}>
                                <Link to="/preferences" style={{color: 'white', textDecoration: 'none'}}>
                                    <i className="fas fa-sliders-h" style={{marginRight: '10px'}}></i>
                                    Preferences
                                </Link>
                            </li>
                            <li className="nav-item" style={{margin: '0.5rem 0', padding: '0.5rem'}}>
                                <Link to="/hvac" style={{color: 'white', textDecoration: 'none'}}>
                                    <i className="fas fa-thermometer-half" style={{marginRight: '10px'}}></i>
                                    HVAC
                                </Link>
                            </li>
                            <li className="nav-item" style={{
                                backgroundColor: '#08192B',
                                margin: '0.5rem 0',
                                padding: '0.5rem',
                                borderLeft: '3px solid #0294A5'
                            }}>
                                <Link to="/appliances" style={{color: '#50BCC0', textDecoration: 'none'}}>
                                    <i className="fas fa-blender" style={{marginRight: '10px'}}></i>
                                    Appliances
                                </Link>
                            </li>
                            <li className="nav-item" style={{margin: '0.5rem 0', padding: '0.5rem'}}>
                                <Link to="/energy" style={{color: 'white', textDecoration: 'none'}}>
                                    <i className="fas fa-bolt" style={{marginRight: '10px'}}></i>
                                    Energy
                                </Link>
                            </li>
                        </ul>
                    </nav>
                </aside>


                <main style={{
                    flex: '1',
                    padding: '1rem',
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                    backgroundColor: '#0E2237'
                }}>
                    <h2 style={{color: 'white'}}>Appliances</h2>
                    <Table striped bordered hover variant="dark"
                           style={{marginTop: '20px', backgroundColor: "#173350"}}>
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
                        {appliances.map((appliance, index) => (
                            <tr key={index}>
                                <td>
                                    <FontAwesomeIcon icon={appliance.icon} style={{color: 'cyan'}}/>
                                </td>
                                <td>{appliance.name}</td>
                                <td>{appliance.location}</td>
                                <td>{appliance.status}</td>
                                <td>{appliance.lastUsed}</td>
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