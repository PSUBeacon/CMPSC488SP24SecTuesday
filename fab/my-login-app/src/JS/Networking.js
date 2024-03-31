import React, {useState, useEffect} from 'react';
import {useNavigate} from 'react-router-dom';
import {Link} from 'react-router-dom'; // Import Link from react-router-dom for navigation
import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import logoImage from '../img/logo.webp';
import houseImage from '../img/houseImage.jpg';
import notificationIcon from '../img/notification.png'
import settingsIcon from '../img/settings.png'
import accountIcon from '../img/account.png'
import menuIcon from '../img/menu.png'
import placeholderImage from '../img/placeholderImage.jpg'; // Replace with the path to your placeholder image
import placeholderImage2 from '../img/placeholderImage2.jpg'; // Replace with the path to your placeholder image
import {Table} from 'react-bootstrap';
import Header from "../components/Header";
import Sidebar from "../components/Sidebar";

// Define the Dashboard component using a functional component pattern
const Networking = () => {

    const navigate = useNavigate(); // Instantiate useNavigate hook
    const [isAccountPopupVisible, setIsAccountPopupVisible] = useState(false);
    const [isNavVisible, setIsNavVisible] = useState(false);
    const [dashboardMessage, setDashboardMessage] = useState('');
    const [accountType, setAccountType] = useState('')
    const [user, setUser] = useState(null);
    const [error, setError] = useState('');
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
        {Time: '00:00', Activity: 'Light Turned On', User: 'x'},
        {Time: '00:00', Activity: 'Light Turned Off', User: 'x'},
        {Time: '00:00', Activity: 'Light Turned Off', User: 'x'},
        {Time: '00:00', Activity: 'Light Turned Off', User: 'x'},
        {Time: '00:00', Activity: 'Light Turned Off', User: 'x'},
        // Add more IoT device log entries as needed
    ];
// Function to change view
    const changeView = (view) => {
        setCurrentView(view);
    };
    const iotLogsTable = (
        <div>
            <h2 style={{color: '#173350'}}>IOT Logs</h2>
            <Table striped bordered hover variant="dark" style={{marginTop: '20px', backgroundColor: "#173350"}}>
                <thead>
                <tr>
                    <th>Time</th>
                    <th>Activity</th>
                    <th>User</th>
                </tr>
                </thead>
                <tbody>
                {iotLogs.map((log, index) => (
                    <tr key={index}>
                        <td>{log.Time}</td>
                        <td>{log.Activity}</td>
                        <td>{log.User}</td>
                    </tr>
                ))}
                </tbody>
            </Table>
        </div>
    );

// Define your Network Logs data
    const networkLogs = [
        {PortSentFrom: 'Port 80', PortSentTo: 'Port 20'},
        {PortSentFrom: 'Port 80', PortSentTo: 'Port 20'},
        {PortSentFrom: 'Port 80', PortSentTo: 'Port 20'},
        {PortSentFrom: 'Port 80', PortSentTo: 'Port 20'},
        {PortSentFrom: 'Port 80', PortSentTo: 'Port 20'},
        // Add more network log entries as needed
    ];

    const networkLogsTable = (
        <div>
            <h2 style={{color: '#173350'}}>Network Logs</h2>
            <Table striped bordered hover variant="dark" style={{marginTop: '20px', backgroundColor: "#173350"}}>
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
        const token = sessionStorage.getItem('token');
        const url = 'http://localhost:8081/networking';

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

    const [cameraView, setCameraView] = useState('livingroom'); // Default camera view

    // Object that holds the URLs for your camera feeds
    const cameraFeeds = {
        livingroom: placeholderImage, // Replace with the actual camera feed URL or image for the living room
        kitchen: placeholderImage2, // Replace with the actual camera feed URL or image for the kitchen
        // Add more camera feeds as needed
    };

    // This is the JSX return statement where we layout our component's HTML structure
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
                    backgroundColor: '#0E2237'
                }}>

                    {/* Widgets Container */}
                    <div style={{width: '100%', maxWidth: '1700px'}}>
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
                                backgroundColor: '#173350',
                                borderRadius: '10px',
                                overflow: 'hidden',
                                flexBasis: '100%',
                                padding: '12px',
                                marginTop: "70px"
                            }}>
                                {/* Conditionally render the table based on the current view */}
                                {currentView === 'iotLogs' && iotLogsTable}
                                {currentView === 'networkLogs' && networkLogsTable}


                                {/* Camera View Buttons */}
                                <div style={{
                                    position: 'absolute',
                                    top: '10px',
                                    left: '10px',
                                    display: 'flex',
                                    gap: '5px'
                                }}>
                                    <button onClick={() => changeView('iotLogs')} style={{
                                        padding: '5px',
                                        color: 'white',
                                        backgroundColor: currentView === 'iotLogs' ? '#0294A5' : '#08192B'
                                    }}>IOT Logs
                                    </button>
                                    <button onClick={() => changeView('networkLogs')} style={{
                                        padding: '5px',
                                        color: 'white',
                                        backgroundColor: currentView === 'networkLogs' ? '#0294A5' : '#08192B'
                                    }}>Network Logs
                                    </button>
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