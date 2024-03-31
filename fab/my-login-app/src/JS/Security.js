import React, {useState, useEffect} from 'react';
import {useNavigate} from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import bedroomIcon from '../img/bedroomIcon.jpg';
import doorLockIcon from '../img/doorLockIcon.png';
import '../CSS/Security.css';
import Header from "../components/Header";
import Sidebar from "../components/Sidebar"; // Import your CSS file here
import axios from "axios";

const Security = () => {
    const navigate = useNavigate(); // Instantiate useNavigate hook
    const [isNavVisible, setIsNavVisible] = useState(false);
    const [accountType, setAccountType] = useState('')
    const [dimmerValue, setDimmerValue] = useState(75); // State to keep track of dimmer value
    const [isLocked, setIsLocked] = useState(false); //need this for the toggle and also the two lines below
    const [selectedRoom, setSelectedRoom] = useState(null);
    const [dashboardMessage, setDashboardMessage] = useState('');
    const [lockStates, setLockStates] = useState({
        '502857': false, // Initial state: off
        //'502858': false, // Initial state: off
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

    const toggleLock = () => {
        Object.keys(lockStates).forEach(uuid => {
            handleToggleLock(uuid);
        });
        setIsLocked(!isLocked);
    };


    // Add a function to handle selecting a room:
    const selectRoom = (roomName) => {
        setSelectedRoom(roomName);
    };

    useEffect(() => {
        const token = sessionStorage.getItem('token');
        const url = 'http://localhost:8081/security';

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
                const updatedDeviceData = {...deviceData};
                Object.keys(updatedDeviceData).forEach(device => {
                    if (data[device]) {
                        sessionStorage.setItem(device, JSON.stringify(data[device]));
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

    const handleToggleLock = (uuid) => {
        const isLocking = !lockStates[uuid];
        setLockStates(prevStates => ({...prevStates, [uuid]: isLocking}));

        const token = sessionStorage.getItem('token');
        if (!token) {
            console.error("Authorization token not found.");
            return;
        }

        const serverUrl = 'http://localhost:8081/security';
        // Prepare the request body
        const requestBody = {
            uuid: uuid,
            name: "SecuritySystem",
            apptype: "Security",
            function: "Status",
            change: isLocking ? "true" : "false",
        };

        fetch(serverUrl, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(requestBody),
        })
            .then(response => {
                if (response.status >= 200 && response.status < 300) {
                    console.log(`Lock ${isLocking ? 'locked' : 'unlocked'} successfully:`, response.data);
                } else {
                    console.error(`Failed to toggle the lock ${isLocking ? 'locked' : 'unlocked'} with status:`, response.status);
                }
            })
            .catch(error => {
                console.error(`There was an error toggling the lock ${isLocking ? 'locked' : 'unlocked'}:`, error);
            });

        setTimeout(() => {
            console.log(`Lock ${uuid} has been ${isLocking ? 'locked' : 'unlocked'}.`);
        }, 1000);
    };

    // This is the JSX return statement where we lay out our component's HTML structure
    return (
        <div style={{display: 'flex', minHeight: '100vh', flexDirection: 'column', backgroundColor: '#081624'}}>
            <Header accountType={accountType}/>
            <div style={{display: 'flex', flex: '1'}}>
                <Sidebar isNavVisible={isNavVisible}/>
                <main style={{
                    flex: '1',
                    padding: '1rem',
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                    backgroundColor: '#0E2237',
                    width: '100%'
                }}>
                    {/* Content Block */}
                    <div className="contentBlock" style={{display: 'flex', width: '100%', flexWrap: 'wrap'}}>
                        {/* Lights Control Section */}
                        <div className="lightsControl" style={{
                            display: 'flex',
                            flexDirection: 'row',
                            justifyContent: 'center',
                            gap: '20px',
                            width: '100%',
                            flexWrap: 'wrap'
                        }}>
                            {/*/!* Room Selection *!/*/}
                            {/*<div className="roomSelection"*/}
                            {/*     style={{flexBasis: '100%', maxWidth: '300px', margin: '0 auto'}}>*/}
                            {/*    <h3 className="centered-title">Selecting a Door</h3>*/}
                            {/*    <div className="roomCards" style={{*/}
                            {/*        display: 'flex',*/}
                            {/*        flexDirection: 'column',*/}
                            {/*        alignItems: 'center',*/}
                            {/*        padding: '0px'*/}
                            {/*    }}>*/}
                            {/*        /!* Room cards *!/*/}
                            {/*        <div className={selectedRoom === "Room 1" ? "card selected" : "card"}*/}
                            {/*             onClick={() => selectRoom("Room 1")}*/}
                            {/*             style={{marginBottom: '20px', width: '100%'}}>*/}
                            {/*            <img className="image" src={bedroomIcon} alt="Room 1"/>*/}
                            {/*        </div>*/}
                            {/*        <div className={selectedRoom === "Room 2" ? "card selected" : "card"}*/}
                            {/*             onClick={() => selectRoom("Room 2")}*/}
                            {/*             style={{marginBottom: '20px', width: '100%'}}>*/}
                            {/*            <img className="image" src={bedroomIcon} alt="Room 2"/>*/}
                            {/*        </div>*/}
                            {/*    </div>*/}
                            {/*</div>*/}


                            {/* Light Selection */}
                            <div className="lightSelection" style={{flexBasis: '48%'}}>
                                <h3 className="centered-title">Your Lock</h3>

                                <div className="lightCards" style={{
                                    display: 'flex',
                                    flexWrap: 'wrap',
                                    justifyContent: 'space-around',
                                    padding: '0px'
                                }}>
                                    {/* Lock card */}
                                    <div className="card" style={{
                                        width: '100%',
                                        maxWidth: '300px',
                                        textAlign: 'center',
                                        padding: '20px'
                                    }}>
                                        <img className="lockImage" src={doorLockIcon} alt="Lock Icon"/>
                                        {/*toggle function*/}
                                        <label className="switch">
                                            <input type="checkbox" checked={isLocked} onChange={toggleLock}/>
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

