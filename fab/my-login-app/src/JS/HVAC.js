// HVAC.js
import '../CSS/HVAC.css'; // Import CSS file
import React, {useState, useEffect} from 'react';
import {useNavigate} from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import logoImage from '../img/logo.webp';
import Header from "../components/Header";
import Sidebar from "../components/Sidebar";
import axios from 'axios';

const HVAC = () => {

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
    const navigate = useNavigate(); // Instantiate useNavigate hook
    const [user, setUser] = useState(null);
    const [error, setError] = useState('');
    const [isNavVisible, setIsNavVisible] = useState(false);
    const [accountType, setAccountType] = useState('');
    const [secondFloorMode, setSecondFloorMode] = useState(deviceData.HVAC.secondFloorMode || 'Heat');
    const [basementMode, setBasementMode] = useState(deviceData.HVAC.basementMode || 'Heat');
    const [secondFloorTemp, setSecondFloorTemp] = useState(deviceData.HVAC.secondFloorTemp || '');
    const [basementTemp, setBasementTemp] = useState(deviceData.HVAC.basementTemp || '');
    const [secondFloorFanSpeed, setSecondFloorFanSpeed] = useState(deviceData.HVAC.secondFloorFanSpeed || '');
    const [basementFanSpeed, setBasementFanSpeed] = useState(deviceData.HVAC.basementFanSpeed || '');
    const [secondFloorHVACStatus, setSecondFloorHVACStatus] = useState(deviceData.HVAC.Status);
    const [basementHVACStatus, setBasementHVACStatus] = useState(deviceData.HVAC.Status);

    const [floorData, setFloorData] = useState([]);


    useEffect(() => {
        const token = sessionStorage.getItem('token');
        if (!token) {
            navigate('/');
            return;
        }

        axios.get(`http://localhost:8081/hvac/GetHVAC`, {
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
        })
            .then(response => {
                setFloorData(response.data);
                console.log("data", response.data)
            })
            .catch(error => {
                console.error('Error fetching floor:', error);
                setError('Could not fetch floor');
            });
    }, [navigate]);

    useEffect(() => {
        const token = sessionStorage.getItem('token');
        const url = 'http://localhost:8081/hvac';

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
                    backgroundImage: 'linear-gradient(to bottom, #0E2237, #081624)',
                    position: 'relative'
                }}>
                    {/* Translucent pattern overlay */}
                    <div style={{
                        position: 'absolute',
                        top: 0,
                        left: 0,
                        width: '100%',
                        height: '100%',
                        backgroundImage: 'url("https://www.transparenttextures.com/patterns/always-grey.png")',
                        opacity: 0.3
                    }}></div>

                    <h1 style={{color: 'white', marginBottom: '2rem'}}>HVAC</h1>
                    <div className="hvac-data-container" style={{
                        display: 'flex',
                        flexDirection: 'row',
                        justifyContent: 'space-around',
                        width: '100%'
                    }}>
                        {floorData.map((floor, index) => (
                            <div key={index} style={{marginRight: index === 0 ? '2rem' : 0}}>
                                <h4 style={{
                                    textAlign: "center",
                                    marginBottom: '20px',
                                    color: '#95A4B6'
                                }}>{floor.location}</h4>

                                {/* Mode */}
                                <div className="data-item">
                                    <div className="data-icon">
                                        <i className="fas fa-thermometer-half"></i>
                                    </div>
                                    <div className="data-info">
                                        <p style={{color: '#95A4B6'}}>Mode</p>
                                        <select
                                            value={floor.Mode}
                                            onChange={(e) => {
                                                const updatedFloorData = floorData.map((f, i) =>
                                                    i === index ? {...f, Mode: e.target.value} : f
                                                );
                                                setFloorData(updatedFloorData);
                                            }}
                                            style={{
                                                color: '#95A4B6',
                                                backgroundColor: '#08192B',
                                                border: 'none',
                                                borderRadius: '5px',
                                                padding: '5px',
                                                zIndex: 10,
                                                position: 'relative'
                                            }}
                                        >
                                            <option value="Heat">Heat</option>
                                            <option value="Cool">Cool</option>
                                        </select>
                                    </div>
                                </div>

                                {/* Temperature */}
                                <div className="data-item">
                                    <div className="data-icon">
                                        <i className="fas fa-thermometer-half"></i>
                                    </div>
                                    <div className="data-info">
                                        <p style={{color: '#95A4B6'}}>Temperature</p>
                                        <input
                                            type="number"
                                            value={floor.Temperature}
                                            onChange={(e) => {
                                                const updatedFloorData = floorData.map((f, i) =>
                                                    i === index ? {...f, Temperature: e.target.value} : f
                                                );
                                                setFloorData(updatedFloorData);
                                            }}
                                            style={{
                                                zIndex: '1',
                                                position: 'relative',
                                            }}
                                            disabled={false}
                                            readOnly={false}
                                        />
                                    </div>
                                </div>

                                {/* Humidity */}
                                <div className="data-item">
                                    <div className="data-icon">
                                        <i className="fas fa-tint"></i>
                                    </div>
                                    <div className="data-info">
                                        <p style={{color: '#95A4B6'}}>Humidity</p>
                                        <p style={{color: '#95A4B6'}}>{floor.Humidity}%</p>
                                    </div>
                                </div>

                                {/* Fan Speed */}
                                <div className="data-item">
                                    <div className="data-icon">
                                        <i className="fas fa-fan"></i>
                                    </div>
                                    <div className="data-info">
                                        <p style={{color: '#95A4B6'}}>Fan Speed</p>
                                        <select
                                            value={floor.FanSpeed}
                                            onChange={(e) => {
                                                const updatedFloorData = floorData.map((f, i) =>
                                                    i === index ? {...f, FanSpeed: e.target.value} : f
                                                );
                                                setFloorData(updatedFloorData);
                                            }}
                                            style={{
                                                color: '#95A4B6',
                                                backgroundColor: '#08192B',
                                                border: 'none',
                                                borderRadius: '5px',
                                                padding: '5px',
                                                zIndex: 2,
                                                position: 'relative',
                                            }}
                                        >
                                            <option value="Low">Low</option>
                                            <option value="Medium">Medium</option>
                                            <option value="High">High</option>
                                        </select>
                                    </div>
                                </div>

                                {/* Status */}
                                <div className="data-item">
                                    <div className="data-icon">
                                        <i className={floor.Status ? "fas fa-power-on" : "fas fa-power-off"}></i>
                                    </div>
                                    <div className="data-info">
                                        <p style={{color: '#95A4B6'}}>Status</p>
                                        <p>{floor.Status ? "On" : "Off"}</p>
                                        <label className="toggle" style={{display: 'block', margin: 'auto'}}>
                                            <input
                                                type="checkbox"
                                                checked={floor.Status}
                                                onChange={(e) => {
                                                    const updatedFloorData = floorData.map((f, i) =>
                                                        i === index ? {...f, Status: e.target.checked} : f
                                                    );
                                                    setFloorData(updatedFloorData);
                                                }}
                                            />
                                            <span className="slider"
                                                  style={{background: floor.Status ? '#50BCC0' : 'grey'}}></span>
                                        </label>
                                    </div>
                                </div>

                                {/* Energy Consumption */}
                                <div className="data-item">
                                    <div className="data-icon">
                                        <i className="fas fa-bolt"></i>
                                    </div>
                                    <div className="data-info">
                                        <p style={{color: '#95A4B6'}}>Energy Consumption</p>
                                        <p style={{color: '#95A4B6'}}>{floor.EnergyConsumption} kWh</p>
                                    </div>
                                </div>
                            </div>
                        ))}
                    </div>
                </main>
            </div>
        </div>
    );
};

export default HVAC;
