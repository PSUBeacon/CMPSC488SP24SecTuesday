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

const Lighting = () => {
    const navigate = useNavigate();
    const [selectedLight, setSelectedLight] = useState(null);
    const [isNavVisible, setIsNavVisible] = useState(false);
    const [accountType, setAccountType] = useState('');
    const [dimmerValue, setDimmerValue] = useState(75);
    const [selectedRoom, setSelectedRoom] = useState(null);
    const [isLightOn, setIsLightOn] = useState(false);
    const [error, setError] = useState('');
    const [user, setUser] = useState(null);
    const [isAccountPopupVisible, setIsAccountPopupVisible] = useState(false);
    const [lights, setLights] = useState([]);
    const uniqueRoomNames = [...new Set(lights.map(light => light.roomName))];
    const [roomName, setRoomName] = useState('');
    const [lightName, setLightName] = useState('');

    const handleSelectLight = (lightId) => {
        setSelectedLight(lightId);
    };

    const toggleLight = () => {
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

    const handleLightOn = (index) => {
        const updatedLights = [...lights];
        updatedLights[index].isOn = true;
        setLights(updatedLights);
    };

    const handleLightOff = (index) => {
        const updatedLights = [...lights];
        updatedLights[index].isOn = false;
        setLights(updatedLights);
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
                                    <button type="button" className="submitButton"
                                            onClick={toggleLight}>{isLightOn ? 'Turn Off All Lights' : 'Turn On All Lights'}</button>

                                </form>
                                <div className="roomDropdown" style={{marginBottom: '20px', width: '72%'}}>
                                    <label htmlFor="selectRoom">Select Room:</label>
                                    <select id="selectRoom" onChange={(e) => setSelectedRoom(e.target.value)}>
                                        <option value="">Select Room</option>
                                        {uniqueRoomNames.map((room, index) => (
                                            <option key={index} value={room}>{room}</option>
                                        ))}
                                    </select>
                                </div>
                                {selectedRoom && (
                                    <ul className="lightList">
                                        {lights.filter((light) => light.roomName === selectedRoom).map((light, index) => (
                                            <li key={index} className="lightItem">
                                                {light.lightName}
                                                <div>
                                                    <button style={{marginRight: '10px'}}
                                                            onClick={() => handleLightOn(index)}>Turn On
                                                    </button>
                                                    <button onClick={() => handleLightOff(index)}>Turn Off</button>
                                                    <img src={removeIcon} alt="Remove"
                                                         style={{width: '20px', marginLeft: '10px', cursor: 'pointer'}}
                                                         onClick={() => handleRemoveLight(index, selectedRoom)}/>
                                                </div>
                                            </li>
                                        ))}
                                    </ul>
                                )}
                                <div className="dimmerControl"
                                     style={{width: '72%', textAlign: 'center', marginTop: '20px'}}>
                                    <input
                                        type="range"
                                        id="dimmer"
                                        name="dimmer"
                                        min="0"
                                        max="100"
                                        value={dimmerValue}
                                        onChange={(e) => setDimmerValue(e.target.value)}
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
                                           style={{color: '#fff', marginTop: '5px'}}>{dimmerValue}%</label>
                                </div>
                            </div>
                        </div>
                    </div>
                </main>
            </div>
        </div>
    );
};

export default Lighting;