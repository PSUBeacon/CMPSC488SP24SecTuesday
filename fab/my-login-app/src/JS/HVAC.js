// HVAC.js
import '../CSS/HVAC.css'; // Import CSS file
import React, {useState, useEffect} from 'react';
import {useNavigate} from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import Header from "../components/Header";
import Sidebar from "../components/Sidebar";
import NestThermostat from "../components/Thermostat";
import axios from 'axios';

const HVAC = () => {
    document.title = 'BEACON | HVAC';

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

                                <NestThermostat initialTemperature={floor.Temperature}
                                                initialFanSpeed={floor.FanSpeed} initialMode={floor.Mode}
                                                initialHumidity={floor.Humidity} initialPU={floor.EnergyConsumption}/>
                            </div>
                        ))}
                    </div>
                </main>
            </div>
        </div>
    );
};

export default HVAC;
