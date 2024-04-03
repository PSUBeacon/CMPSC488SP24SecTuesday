import React, {useState, useEffect} from 'react';
import {useNavigate} from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../CSS/Lighting.css';
import Header from '../components/Header';
import Sidebar from '../components/Sidebar';
import removeIcon from '../img/recycle-bin-icon.png';
import bedroom from '../img/bedroom.png';
import livingRoomImage from '../img/livingRoom.png';
import kitchenImage from '../img/kitchen.png';
import axios from "axios";

const Lighting = () => {
    const navigate = useNavigate();
    const [selectedLight, setSelectedLight] = useState(null);
    const [isNavVisible, setIsNavVisible] = useState(false);
    const [accountType, setAccountType] = useState('');
    const [dimmerValue, setDimmerValue] = useState();
    const [selectedRoom, setSelectedRoom] = useState(null);
    const [isLightOn, setIsLightOn] = useState(false);
    const [error, setError] = useState('');
    const [user, setUser] = useState(null);
    const [lights, setLights] = useState([]);
    const uniqueRoomNames = [...new Set(lights.map(light => light.roomName))];
    const [roomName, setRoomName] = useState('');
    const [lightName, setLightName] = useState('');
    const [lightStates, setLightStates] = useState({
        '417293': false, // Initial state: off
        '417294': false, // Initial state: off
    });


    useEffect(() => {
        if (selectedRoom) {
            const token = sessionStorage.getItem('token');
            if (!token) {
                navigate('/'); // Redirect to login page if token is not present
                return;
            }

            axios.get(`http://localhost:8081/lighting?roomName=${encodeURIComponent(selectedRoom)}`, {
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json',
                },
            })
                .then(response => {
                    setLights(response.data); // Set the fetched lights
                    console.log("data", response.data)
                })
                .catch(error => {
                    console.error('Error fetching lights:', error);
                    setError('Could not fetch lights');
                });
        }
    }, [selectedRoom, navigate]); // Add navigate as a dependency if it's used within the effect

    const handleDimmerChange = (uuid, brightness) => {
        const serverUrl = 'http://localhost:8081/lighting';
        const token = sessionStorage.getItem('token');

        // Prepare the request body
        const requestBody = {
            uuid: uuid,
            name: "Lighting",
            apptype: "Lighting",
            function: "Brightness",
            change: brightness,
        };

        // Send a POST request to toggle the light state
        axios.post(serverUrl, requestBody, {headers: {'Authorization': `Bearer ${token}`}})
            .then(response => {
                if (response.status >= 200 && response.status < 300) {
                    console.log(`Light dimmed successfully:`, response.data);
                } else {
                    console.error(`Failed to dim the light  with status:`, response.status);
                }
            })
            .catch(error => {
                console.error(`There was an error toggling the brightness `, error);
            });

        // Log the action
        setTimeout(() => {
            console.log(`Light has been dimmed.`);
        }, 1000);
    };

    const selectRoom = (roomName) => {
        setSelectedRoom(roomName);
    };


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
    const handleSelectLight = (lightId) => {
        setSelectedLight(lightId);
    };

    const toggleLight = () => {
        lights.forEach(light => handleToggleLight(light.UUID));
        setIsLightOn(!isLightOn);
    };

    useEffect(() => {
        const storedLights = JSON.parse(localStorage.getItem('lights'));
        if (storedLights) {
            setLights(storedLights);
        }
    }, []);

    const handleRemoveLight = (index, roomName) => {
        const updatedLights = lights.filter((light, lightIndex) => {
            return !(lightIndex === index && light.roomName === roomName);
        });
        setLights(updatedLights);
    };


    //combined on and off handlers
    const handleToggleLight = (uuid) => {
        const isTurningOn = !lightStates[uuid];
        setLightStates(prevStates => ({...prevStates, [uuid]: isTurningOn}));

        const serverUrl = 'http://localhost:8081/lighting';
        const token = sessionStorage.getItem('token');

        // Prepare the request body
        const requestBody = {
            uuid: uuid,
            name: "Lighting",
            apptype: "Lighting",
            function: "Status",
            change: isTurningOn ? "true" : "false",
        };

        // Send a POST request to toggle the light state
        axios.post(serverUrl, requestBody, {headers: {'Authorization': `Bearer ${token}`}})
            .then(response => {
                if (response.status >= 200 && response.status < 300) {
                    console.log(`Light ${isTurningOn ? 'turned on' : 'turned off'} successfully:`, response.data);
                } else {
                    console.error(`Failed to toggle the light ${isTurningOn ? 'on' : 'off'} with status:`, response.status);
                }
            })
            .catch(error => {
                console.error(`There was an error toggling the light ${isTurningOn ? 'on' : 'off'}:`, error);
            });

        // Log the action
        setTimeout(() => {
            console.log(`Light ${uuid} has been ${isTurningOn ? 'turned on' : 'turned off'}.`);
        }, 1000);
    };

    const handleFormSubmit = (e) => {
        e.preventDefault();
        if (roomName && lightName) {
            setLights([...lights, {roomName, lightName, isOn: false}]);
            setRoomName('');
            setLightName('');
        }
    };

    return (
        <div style={{display: 'flex', minHeight: '100vh', flexDirection: 'column', backgroundColor: '#081624'}}>
            <Header accountType={accountType}/>
            <div style={{display: 'flex', flex: '1'}}>
                <Sidebar isNavVisible={isNavVisible}/>
                <main style={{flex: '1', padding: '1rem', backgroundColor: '#0E2237', width: '100%'}}>
                    <div className="contentBlock" style={{
                        display: 'flex',
                        flexDirection: 'row',
                        alignItems: 'flex-start',
                        width: '100%',
                        paddingBottom: '60px'
                    }}>
                        <div className="roomSelection" style={{
                            flex: '1',
                            display: 'flex',
                            flexDirection: 'column',
                            alignItems: 'center',
                            marginRight: '20px'
                        }}>
                            <h3 className="centered-title">Selecting a Room</h3>
                            <div className="RoomCards" style={{
                                display: 'flex',
                                flexDirection: 'column',
                                alignItems: 'center',
                                marginTop: '40px'
                            }}>
                                <div className="card" style={{
                                    marginBottom: '20px',
                                    border: selectedRoom === "Bedroom" ? '2px solid #0294A5' : 'none'
                                }} onClick={() => setSelectedRoom("Bedroom")}><img className="images" src={bedroom}
                                                                                   alt="Bedroom"/></div>
                                <div className="card" style={{
                                    marginBottom: '20px',
                                    border: selectedRoom === "Kitchen" ? '2px solid #0294A5' : 'none'
                                }} onClick={() => setSelectedRoom("Kitchen")}><img className="images" src={kitchenImage}
                                                                                   alt="Kitchen"/></div>
                                <div className="card"
                                     style={{border: selectedRoom === "Living Room" ? '2px solid #0294A5' : 'none'}}
                                     onClick={() => setSelectedRoom("Living Room")}><img className="images"
                                                                                         src={livingRoomImage}
                                                                                         alt="Living Room"/></div>
                            </div>
                        </div>
                        <div className="lightControl"
                             style={{flex: '1', display: 'flex', flexDirection: 'column', alignItems: 'center'}}>
                            <div className="formContainer"
                                 style={{width: '100%', marginBottom: '20px', marginTop: '80px'}}>
                                <form onSubmit={handleFormSubmit} className="lightForm">
                                    <div className="formGroup">
                                        <label htmlFor="roomName">Room Name:</label>
                                        <input type="text" id="roomName" value={roomName}
                                               onChange={(e) => setRoomName(e.target.value)}/>
                                    </div>
                                    <div className="formGroup">
                                        <label htmlFor="lightName">Light Name:</label>
                                        <input type="text" id="lightName" value={lightName}
                                               onChange={(e) => setLightName(e.target.value)}/>
                                    </div>
                                    <button type="submit" className="submitButton">Add Light</button>
                                </form>

                                <div className="roomDropdown" style={{
                                    marginBottom: '20px',
                                    width: '72%',
                                    display: 'flex',
                                    alignItems: 'center'
                                }}>
                                    <div style={{marginRight: '10px', display: 'flex', alignItems: 'center'}}>
                                        <select id="selectLight" value={selectedLight}
                                                onChange={(e) => {
                                                    const selectedUUID = e.target.value;
                                                    setSelectedLight(selectedUUID);
                                                    const selectedLightObj = lights.find(light => light.UUID === selectedUUID);
                                                    if (selectedLightObj) {
                                                        setDimmerValue(selectedLightObj.Brightness); // Assuming each light object has a Brightness property
                                                    }
                                                }}
                                                style={{marginLeft: '5px'}}>
                                            <option value="">Select Light</option>
                                            {lights.map((light, index) => (
                                                <option key={index} value={light.UUID}>{light.Location}</option>
                                            ))}
                                        </select>
                                    </div>
                                    <button type="button" className="submitButton" onClick={toggleLight}
                                            style={{marginLeft: '10px'}}>{isLightOn ? 'Turn Off' : 'Turn On'}</button>
                                </div>

                                <div className="dimmerControl"
                                     style={{width: '72%', textAlign: 'center', marginTop: '20px'}}>
                                    <input
                                        type="range"
                                        id="dimmer"
                                        name="dimmer"
                                        min="1"
                                        max="100"
                                        value={Math.round((dimmerValue - 1) / 14 * 99) + 1}
                                        onChange={(e) => {
                                            const newValue = Math.round((e.target.value - 1) / 99 * 14) + 1;
                                            setDimmerValue(newValue);
                                        }}
                                        onMouseUp={(e) => {
                                            if (selectedLight) {
                                                handleDimmerChange(selectedLight, dimmerValue);
                                            }
                                        }}
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
                                    <label htmlFor="dimmer"
                                           style={{
                                               color: '#fff',
                                               marginTop: '5px'
                                           }}>{Math.round((dimmerValue - 1) / 14 * 99) + 1}%</label>
                                </div>
                            </div>
                        </div>
                    </div>
                </main>
            </div>
        </div>
    )
        ;
};

export default Lighting;