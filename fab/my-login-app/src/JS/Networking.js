import React, {useState, useEffect} from 'react';
import {useNavigate} from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import placeholderImage from '../img/placeholderImage.jpg'; // Replace with the path to your placeholder image
import placeholderImage2 from '../img/placeholderImage2.jpg'; // Replace with the path to your placeholder image
import {Table} from 'react-bootstrap';
import Header from "../components/Header";
import Sidebar from "../components/Sidebar";

// Define the Dashboard component using a functional component pattern
const Networking = () => {
    document.title = 'BEACON | Logs';
    const [isAccountPopupVisible, setIsAccountPopupVisible] = useState(false);
    const [isNavVisible, setIsNavVisible] = useState(false);
    const [dashboardMessage, setDashboardMessage] = useState('');
    const [accountType, setAccountType] = useState('')
    const [user, setUser] = useState(null);
    const [error, setError] = useState('');
    const [Logs, setLogs] = useState([]);
    const [currentView, setCurrentView] = useState('iotLogs'); // New state to track current view
    // Function to change view
    const changeView = (view) => {
        setCurrentView(view);
    };
    const navigate = useNavigate(); // Instantiate useNavigate hook
    // Define your IoT Logs data similar to how you have solarPanel data
    let iotLogs = [
        // {DeviceID: '37675', Function: 'Temperature', "Change": '73', "Time": "00:00:000"}
    ];

    const token = sessionStorage.getItem('token');
    const url = "http://192.168.8.117:8081/networking/GetNetLogs";

    useEffect(() => {

            const token = sessionStorage.getItem('token');
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
                .then(response => {
                    if (!response.ok) {
                        throw new Error('Network response was not ok');
                    }
                    return response.json();
                })
                .then(data => {
                    // Store the fetched data into networkLogs
                    iotLogs.push(...data);
                    setLogs(data)
                    console.log('IOT logs:', iotLogs);// You can process or log the data here
                })
                .catch(error => {
                    // Handle errors
                    console.error('Error fetching logs:', error);
                });
        },
        []
    );
    const iotLogsTable = (
        <div>
            <h2 style={{color: '#173350'}}>IOT Logs</h2>
            <Table striped bordered hover variant="dark" style={{marginTop: '20px', backgroundColor: "#173350"}}>
                <thead>
                <tr>
                    <th>DeviceID</th>
                    <th>Function</th>
                    <th>Change</th>
                    <th>Time</th>
                </tr>
                </thead>
                <tbody>
                {Logs.map((log, index) => (
                    <tr key={index}>
                        <td>{log.DeviceID}</td>
                        <td>{log.Function}</td>
                        <td>{log.Change}</td>
                        <td>{log.Time}</td>
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


    const [cameraView, setCameraView] = useState('livingroom'); // Default camera view

    // Object that holds the URLs for your camera feeds
    const cameraFeeds = {
        livingroom: placeholderImage, // Replace with the actual camera feed URL or image for the living room
        kitchen: placeholderImage2, // Replace with the actual camera feed URL or image for the kitchen
    };

    // This is the JSX return statement where we layout our component's HTML structure
    return (
        <div style={{display: 'flex', minHeight: '100vh', flexDirection: 'column'}}>
            <Header accountType={accountType}/>
            <div style={{display: 'flex', flex: '1'}}>
                <Sidebar isNavVisible={isNavVisible}/>
                <main style={{
                    flex: '1',
                    padding: '1rem',
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center'
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