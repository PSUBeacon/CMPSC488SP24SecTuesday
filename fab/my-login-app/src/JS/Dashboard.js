import React, {useState, useEffect} from 'react';
import {useNavigate} from 'react-router-dom';
import {Link} from 'react-router-dom'; // Import Link from react-router-dom for navigation
import 'bootstrap/dist/css/bootstrap.min.css';
import doorLockIcon from '../img/doorLockIcon.png';
import bulbIcon from '../img/bulb-icon.png';
import placeholderImage from '../img/placeholderImage.jpg';
import livingRoomFootage from '../img/livingRoomFootage1.gif';
import bedRoomFootage from '../img/BedroomFootage.gif';
import kitchenFootage from '../img/kitchenFootage.gif';
import Header from '../components/Header';
import Sidebar from '../components/Sidebar';
import PadlockAnimation from "../components/Lock";

const Dashboard = () => {
    document.title = 'BEACON | Dashboard';

    const [cameraView, setCameraView] = useState('livingroom'); // Default camera view
    const navigate = useNavigate(); // Instantiate useNavigate hook
    const [isNavVisible, setIsNavVisible] = useState(false);
    const [dashboardMessage, setDashboardMessage] = useState('');
    const [accountType, setAccountType] = useState('')
    const [user, setUser] = useState(null);
    const [error, setError] = useState('');
    const [lockStates, setLockStates] = useState({
        '502857': 'locked',
        '502858': 'locked',
    });
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
        const token = sessionStorage.getItem('token');
        const url = 'http://localhost:8081/dashboard';

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


    // Object that holds the URLs for your camera feeds
    const cameraFeeds = {
        livingroom: livingRoomFootage, // Replace with actual video feed or image
        kitchen: kitchenFootage, // Replace with the actual camera feed URL or image for the kitchen
        bedroom: bedRoomFootage, //adding another tab to the camera views
    };


    const CameraWidget = () => {

        return (
            <div className="camera-widget" style={{
                position: 'relative',
                maxWidth: '100%',
                borderRadius: '10px',
                backgroundColor: '#12232E',
                overflow: 'hidden'
            }}>
                {/* Live Feed Video */}
                <video src={cameraFeeds[cameraView]} alt="Live feed"
                       style={{width: '100%', height: 'auto', display: 'block'}} controls autoplay loop muted>
                    Your browser does not support the video tag.
                </video>

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
                    <button onClick={() => setCameraView('bedroom')} style={{
                        padding: '5px',
                        backgroundColor: cameraView === 'bedroom' ? '#4CAF50' : 'transparent'
                    }}>R3
                    </button>
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

        const handleToggleLock = (door) => {
            setLocksStatus(prevStatus => ({
                ...prevStatus,
                [door]: prevStatus[door] === 'Locked' ? 'Unlocked' : 'Locked',
            }));
        };

        // State to track the light status
        const [lightStatus, setLightStatus] = useState('Off');

        const handleToggleLight = () => {
            const isTurningOn = lightStatus === 'Off';
            setLightStatus(isTurningOn ? 'On' : 'Off');

            const serverUrl = 'http://localhost:8081/lighting';
            const token = sessionStorage.getItem('token');

            // Prepare the request body
            const requestBody = {
                uuid: "417294",
                name: "Lighting",
                apptype: "Lighting",
                function: "Status",
                change: isTurningOn ? "true" : "false",
            };

            // Send the request to the server
            fetch(serverUrl, {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(requestBody),
            })
                .then(response => {
                    if (response.ok) {
                        console.log('Light status updated successfully');
                    } else {
                        console.error('Failed to update light status');
                    }
                })
                .catch(error => {
                    console.error('Error updating light status:', error);
                });
        };

        return (
            <div>
                <div className="widget" style={{
                    display: 'flex',
                    flexDirection: 'row',
                    justifyContent: 'center',
                    alignItems: 'center',
                    backgroundColor: '#173350',
                    padding: '20px',
                    paddingRight: "0px",
                    borderRadius: '10px',
                    margin: '0px'
                }}>
                    {/* Front Door Lock */}
                    <div style={{flex: '1', maxWidth: '250px', textAlign: 'center'}}>
                        <p style={{color: "#95A4B6", marginRight: '20px'}}>Front Door</p>
                        <img className="lockImage" src={doorLockIcon} alt="Lock Icon"
                             style={{width: '100px', height: 'auto'}}/>
                        <PadlockAnimation
                            isLocked={lockStates['502858'] === 'locked'}
                            handleToggleLock={() => handleToggleLock('502858')}/>
                    </div>

                    {/* Back Door Lock */}
                    <div style={{flex: '1', maxWidth: '250px', textAlign: 'center'}}>

                        <p style={{color: "#95A4B6", marginRight: '20px'}}>Back Door</p>
                        <img className="lockImage" src={doorLockIcon} alt="Lock Icon"
                             style={{width: '100px', height: 'auto'}}/>
                        <PadlockAnimation
                            isLocked={lockStates['502858'] === 'locked'}
                            handleToggleLock={() => handleToggleLock('502858')}/>
                    </div>
                </div>
                <div>
                    <div style={{
                        display: 'flex',
                        flexDirection: 'row',
                        justifyContent: 'space-between',
                        alignItems: 'center',
                        backgroundColor: '#173350',
                        padding: '20px',
                        paddingRight: "0px",
                        borderRadius: '10px',
                        margin: '0px'
                    }}>
                        <div>
                            <h3 style={{color: "#95A4B6", marginBottom: '20px'}}>Lights</h3>
                        </div>
                        <div>
                            <button onClick={() => navigate('/lighting')} style={{
                                backgroundColor: '#0294A5',
                                color: 'white',
                                border: 'none',
                                borderRadius: '5px',
                                padding: '8px 18px',
                                cursor: 'pointer',
                                display: 'block',
                                marginLeft: 'auto',
                                marginRight: 'auto',
                                fontSize: '11px'
                            }}>
                                See More
                            </button>
                        </div>
                    </div>
                    <div style={{display: 'flex', alignItems: 'center', justifyContent: 'space-around'}}>
                        <img src={bulbIcon} alt="Bedroom Light"
                             style={{width: '60px', height: 'auto', marginRight: '30px', marginBottom: '10px'}}/>
                        <div style={{textAlign: 'center'}}>
                            <span style={{color: "#95A4B6", fontSize: '20px'}}>Livingroom light:</span>
                            <p style={{marginTop: '10px'}}>{lightStatus}</p>
                        </div>
                        <label className="toggle" style={{display: 'block', margin: 'auto'}}>
                            <input
                                type="checkbox"
                                checked={lightStatus === 'On'}
                                onChange={handleToggleLight}
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
                <div className="widget" style={{
                    flex: '1',
                    minWidth: '200px',
                    padding: '0px',
                    borderRadius: '1px',
                    backgroundColor: '#173350',
                    margin: '10px',
                    boxSizing: 'border-box',
                    display: 'flex',
                    justifyContent: 'space-around'
                }}>
                    {/* Iterate over the deviceData object */}
                    {Object.entries(deviceData).map(([deviceName, deviceInfo]) => (
                        // Only render the device if its status is available and not an empty string
                        deviceInfo.Status && deviceInfo.Status !== '' && (
                            <div key={deviceName} style={{textAlign: 'center'}}>
                                <span style={{color: "#95A4B6", margin: '6px'}}>{deviceName}</span>
                                <p>{deviceInfo.Status}</p>
                                {console.log(deviceInfo.Status)}
                            </div>
                        )
                    ))}
                </div>
            </div>
        );
    };


    // This is the JSX return statement where we layout our component's HTML structure
    return (
        <div style={{display: 'flex', minHeight: '100vh', flexDirection: 'column', backgroundColor: '#081624'}}>
            <Header accountType={accountType}/>
            {/* Side Navbar and Dashboard Content */}
            <div style={{display: 'flex', flex: '1'}}>
                <Sidebar isNavVisible={isNavVisible}/>
                <main style={{
                    flex: '1',
                    padding: '1rem',
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                    backgroundColor: '#173350',
                }}>

                    {/* Widgets Container */}
                    <div style={{width: '100%', maxWidth: '1200px'}}>
                        {/* Widgets Row */}
                        <div style={{
                            display: 'flex',
                            flexDirection: 'row',
                            flexWrap: 'wrap',
                            justifyContent: 'center',
                            gap: '20px',
                            marginBottom: '20px'
                        }}>
                            {/* Camera Widget */}
                            <div className="camera-widget" style={{
                                position: 'relative',
                                maxWidth: '60%',
                                borderRadius: '1px',
                                overflow: 'hidden',
                                backgroundColor: '#173350',
                                flexBasis: '100%',
                                padding: '12px'
                            }}>
                                {/* Camera Feed */}
                                <img src={cameraFeeds[cameraView]} alt="Camera feed"
                                     style={{width: '100%', height: '100%'}}/>

                                {/* Camera View Buttons */}
                                <div style={{
                                    position: 'absolute',
                                    top: '10px',
                                    left: '10px',
                                    display: 'flex',
                                    gap: '5px'
                                }}>
                                    <button onClick={() => setCameraView('livingroom')} style={{
                                        padding: '5px',
                                        color: 'white',
                                        backgroundColor: cameraView === 'livingroom' ? '#0294A5' : '#08192B'
                                    }}>R1
                                    </button>
                                    <button onClick={() => setCameraView('kitchen')} style={{
                                        padding: '5px',
                                        color: 'white',
                                        backgroundColor: cameraView === 'kitchen' ? '#0294A5' : '#08192B'
                                    }}>R2
                                    </button>
                                    <button onClick={() => setCameraView('bedroom')} style={{
                                        padding: '5px',
                                        color: 'white',
                                        backgroundColor: cameraView === 'bedroom' ? '#0294A5' : '#08192B'
                                    }}>R3
                                    </button>
                                </div>
                            </div>

                            {/* Locks Widget */}
                            <div className="widget" style={{
                                flex: '1',
                                padding: '20px',
                                backgroundColor: '#173350',
                                borderRadius: '1px',
                                margin: '10px',
                                boxSizing: 'border-box'
                            }}>
                                <div style={{
                                    display: 'flex',
                                    flexDirection: 'row',
                                    justifyContent: 'space-between',
                                    alignItems: 'center',
                                    backgroundColor: '#173350',
                                    padding: '20px',
                                    paddingRight: "0px",
                                    borderRadius: '10px',
                                    margin: '0px'
                                }}>
                                    <div>
                                        <h3 style={{color: "#95A4B6"}}>Locks</h3>
                                    </div>
                                    <div>
                                        <button onClick={() => navigate('/security')} style={{
                                            backgroundColor: '#0294A5',
                                            color: 'white',
                                            border: 'none',
                                            borderRadius: '5px',
                                            padding: '8px 18px',
                                            cursor: 'pointer',
                                            display: 'block',
                                            marginLeft: 'auto',
                                            marginRight: 'auto',
                                            fontSize: '11px'
                                        }}>
                                            See More
                                        </button>
                                    </div>
                                </div>
                                <LocksWidget/>
                            </div>
                        </div>
                        {/* Another Row for More Widgets */}
                        <div style={{
                            display: 'flex',
                            flexDirection: 'row',
                            flexWrap: 'wrap',
                            justifyContent: 'center',
                            gap: '20px'
                        }}>
                            {/* Status by Units Widget */}
                            <div className="widget" style={{
                                flex: '1',
                                minWidth: '290px',
                                backgroundColor: '#173350',
                                padding: '20px',
                                borderRadius: '1px',
                                margin: '10px',
                                boxSizing: 'border-box'
                            }}>
                                <h3 style={{marginBottom: '20px', color: "#95A4B6"}}>Status by Units</h3>
                                {/* Content of the status by units widget */}
                                <div style={{display: 'flex', flexDirection: 'row', justifyContent: 'space-around'}}>
                                    <div style={{textAlign: "center", margin: '6px'}}>
                                        <p style={{color: "#95A4B6"}}>Power</p>
                                        {deviceData.solarPanel && deviceData.solarPanel[0] && deviceData.solarPanel[0].EnergyGeneratedToday && (
                                            <p style={{fontSize: '22px'}}>{deviceData.solarPanel[0].EnergyGeneratedToday}kw</p>
                                        )}
                                    </div>
                                    <div style={{textAlign: "center", margin: '6px'}}>
                                        <p style={{color: "#95A4B6"}}>Temperature</p>
                                        {deviceData.hvac && deviceData.hvac[0] && deviceData.hvac[0].Temperature && (
                                            <p style={{fontSize: '22px'}}>{deviceData.hvac[0].Temperature}</p>
                                        )}
                                    </div>
                                    <div style={{textAlign: "center", margin: '6px'}}>
                                        <p style={{color: "#95A4B6"}}> Humidity</p>
                                        {deviceData.hvac && deviceData.hvac[0] && deviceData.hvac[0].Humidity && (
                                            <p style={{fontSize: '22px'}}>{deviceData.hvac[0].Humidity}</p>
                                        )}
                                    </div>
                                    <div style={{textAlign: "center", margin: '6px'}}>
                                        <p style={{color: "#95A4B6"}}>Security</p>
                                        {deviceData.securitySystem && deviceData.securitySystem[0] && deviceData.securitySystem[0].Status && (
                                            <p style={{
                                                color: 'green',
                                                fontWeight: 'bold',
                                                fontSize: '22px'
                                            }}>{deviceData.securitySystem[0].Status}</p>
                                        )}
                                    </div>
                                </div>

                            </div>

                            {/* Scheduled Activity Widget */}
                            <div className="widget" style={{
                                flex: '1',
                                minWidth: '250px',
                                backgroundColor: '#173350',
                                padding: '20px',
                                borderRadius: '1px',
                                margin: '10px',
                                boxSizing: 'border-box'
                            }}>
                                <div style={{
                                    display: 'flex',
                                    flexDirection: 'row',
                                    justifyContent: 'space-between',
                                    alignItems: 'center',
                                    backgroundColor: '#173350',
                                    padding: '20px',
                                    paddingRight: "0px",
                                    borderRadius: '10px',
                                    margin: '0px'
                                }}>
                                    <div>
                                        <h3 style={{color: "#95A4B6"}}>Active Appliances</h3>
                                    </div>
                                    <div>
                                        <button onClick={() => navigate('/appliances')} style={{
                                            backgroundColor: '#0294A5',
                                            color: 'white',
                                            border: 'none',
                                            borderRadius: '5px',
                                            padding: '8px 18px',
                                            cursor: 'pointer',
                                            display: 'block',
                                            marginLeft: 'auto',
                                            marginRight: 'auto',
                                            fontSize: '11px'
                                        }}>
                                            See More
                                        </button>
                                    </div>
                                </div>
                                <AppliancesWidget/>
                            </div>
                        </div>
                    </div>
                </main>
            </div>
        </div>
    );
};
export default Dashboard;