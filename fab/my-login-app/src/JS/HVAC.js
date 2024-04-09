// HVAC.js
import '../CSS/HVAC.css'; // Import CSS file
import React, {useState, useEffect} from 'react';
import {useNavigate} from 'react-router-dom';
//import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import Header from "../components/Header";
import Sidebar from "../components/Sidebar";
import NestThermostat from "../components/Thermostat";
import axios from 'axios';
import {faAlignCenter} from "@fortawesome/free-solid-svg-icons";

const HVAC = () => {
    document.title = 'BEACON | HVAC';


    const navigate = useNavigate(); // Instantiate useNavigate hook
    const [user, setUser] = useState(null);
    const [error, setError] = useState('');
    const [isNavVisible, setIsNavVisible] = useState(false);
    const [accountType, setAccountType] = useState('');
    const [thermostats, setThermostats] = useState([]);
    const [floorData, setFloorData] = useState([]);
    const [currentIndex, setCurrentIndex] = useState(0); // Track the current index to display

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

    const handleNext = () => {
        setCurrentIndex((prevIndex) => (prevIndex + 1) % floorData.length);
    };

    const handlePrevious = () => {
        setCurrentIndex((prevIndex) => {
            const newIndex = prevIndex - 1;
            return newIndex < 0 ? floorData.length - 1 : newIndex;
        });
    };

    const handleThermostatChange = (updatedThermostat) => {
        const newFloorData = [...floorData];
        newFloorData[currentIndex] = updatedThermostat;
        setFloorData(newFloorData);
    };

    useEffect(() => {
        const token = sessionStorage.getItem('token');
        const url = 'http://localhost:8081/hvac/';

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

    const [updatedFloorData, setUpdatedFloorData] = useState(floorData);

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
                    <h1 style={{color: 'white'}}>HVAC</h1>
                    <div className="hvac-data-container">
                        {floorData.length > 0 && (
                            <div key={floorData[currentIndex].UUID}>
                                <p>{floorData[currentIndex].Location}</p>
                                <NestThermostat
                                    thermostat={floorData[currentIndex]}
                                    index={currentIndex}
                                    onThermostatChange={handleThermostatChange}
                                />
                            </div>
                        )}
                    </div>
                    <div className={'navigation'}>
                        <button onClick={handlePrevious} className={'submitButton'}>Previous</button>
                        <button onClick={handleNext} className={'submitButton'}>Next</button>
                    </div>
                </main>
            </div>
        </div>
    );
};

export default HVAC;