// HVAC.js
import '../CSS/HVAC.css'; // Import CSS file
import React, {useState, useEffect} from 'react';
import {useNavigate} from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import logoImage from '../img/logo.webp';
import Header from "../components/Header";
import Sidebar from "../components/Sidebar";

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
    const [dashboardMessage, setDashboardMessage] = useState('');
    const [isNavVisible, setIsNavVisible] = useState(false);
    const [accountType, setAccountType] = useState('')
    const [secondFloorMode, setSecondFloorMode] = useState(deviceData.HVAC.secondFloorMode || 'Heat');
    const [basementMode, setBasementMode] = useState(deviceData.HVAC.basementMode || 'Heat');
    const [secondFloorTemp, setSecondFloorTemp] = useState(deviceData.HVAC.secondFloorTemp || '');
    const [basementTemp, setBasementTemp] = useState(deviceData.HVAC.basementTemp || '');
    const [secondFloorHumidity, setSecondFloorHumidity] = useState(deviceData.HVAC.secondFloorHumidity || '');
    const [basementHumidity, setBasementHumidity] = useState(deviceData.HVAC.basementHumidity || '');
    const [secondFloorFanSpeed, setSecondFloorFanSpeed] = useState(deviceData.HVAC.secondFloorFanSpeed || '');
    const [basementFanSpeed, setBasementFanSpeed] = useState(deviceData.HVAC.basementFanSpeed || '');
    const [secondFloorEnergyConsumption, setSecondFloorEnergyConsumption] = useState(deviceData.HVAC.secondFloorEnergyConsumption || '');
    const [basementEnergyConsumption, setBasementEnergyConsumption] = useState(deviceData.HVAC.basementEnergyConsumption || '');
    const [secondFloorHVACStatus, setSecondFloorHVACStatus] = useState(deviceData.HVAC.Status);
    const [basementHVACStatus, setBasementHVACStatus] = useState(deviceData.HVAC.Status);


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
                    {deviceData && (
                        <div className="hvac-data-container" style={{
                            display: 'flex',
                            flexDirection: 'row',
                            justifyContent: 'space-around',
                            width: '100%'
                        }}>
                            {/* First Column */}
                            <div style={{marginRight: '2rem'}}>
                                <h4 style={{textAlign: "center", marginBottom: '20px', color: '#95A4B6'}}>2nd Floor</h4>

                                {/* First Column - mode */}
                                <div className="data-item">
                                    <div className="data-icon">
                                        <i className="fas fa-thermometer-half"></i>
                                    </div>
                                    <div className="data-info">
                                        <p style={{color: '#95A4B6'}}>Mode</p>
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
                                        <p style={{color: '#95A4B6'}}>Temperature</p>
                                        <input
                                            type="number"
                                            value={secondFloorTemp}
                                            onChange={(e) => setSecondFloorTemp(e.target.value)
                                            }
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
                                        <p style={{color: '#95A4B6'}}>Humidity</p>
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
                                        <p style={{color: '#95A4B6'}}>Fan Speed</p>
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
                                        <p style={{color: '#95A4B6'}}>Status</p>
                                        <p>{secondFloorHVACStatus ? "On" : "Off"}</p>
                                        {/* Toggle */}
                                        <label className="toggle" style={{display: 'block', margin: 'auto'}}>
                                            <input
                                                type="checkbox"
                                                checked={secondFloorHVACStatus}
                                                onChange={toggleSecondFloorHVACStatus}
                                            />
                                            <span className="slider"
                                                  style={{background: secondFloorHVACStatus ? '#50BCC0' : 'grey'}}></span>
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
                                        <p style={{color: '#95A4B6'}}>Energy Consumption</p>
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
                                <h4 style={{textAlign: "center", marginBottom: '20px', color: '#95A4B6'}}>First
                                    Floor</h4>

                                <div className="data-item">
                                    <div className="data-icon">
                                        <i className="fas fa-thermometer-half"></i>
                                    </div>
                                    <div className="data-info">
                                        <p style={{color: '#95A4B6'}}>Mode</p>
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
                                        <p style={{color: '#95A4B6'}}>Temperature</p>
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
                                        <p style={{color: '#95A4B6'}}>Humidity</p>
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
                                        <p style={{color: '#95A4B6'}}>Fan Speed</p>
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
                                        <p style={{color: '#95A4B6'}}>Status</p>
                                        <p>{basementHVACStatus ? "On" : "Off"}</p>
                                        {/* Toggle */}
                                        <label className="toggle" style={{display: 'block', margin: 'auto'}}>
                                            <input
                                                type="checkbox"
                                                checked={basementHVACStatus}
                                                onChange={toggleBasementHVACStatus}
                                            />
                                            <span className="slider"
                                                  style={{background: basementHVACStatus ? '#50BCC0' : 'grey'}}></span>
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
                                        <p style={{color: '#95A4B6'}}>Energy Consumption</p>
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
