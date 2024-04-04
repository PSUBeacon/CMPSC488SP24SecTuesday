import React, {useState, useEffect} from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import doorLockIcon from '../img/doorLockIcon.png';
import Header from "../components/Header";
import Sidebar from "../components/Sidebar";
import HomeAlarm from "../components/SecuritySystem";
import '../CSS/Security.css';
import PadlockAnimation from "../components/Lock";

const Security = () => {
    document.title = 'BEACON | Security';

    const [isNavVisible, setIsNavVisible] = useState(false);
    const [accountType, setAccountType] = useState('');
    const [lockStates, setLockStates] = useState({
        '502857': 'locked',
        '502858': 'locked',
    });
    const [securityStatus, setSecurityStatus] = useState(['Armed', 'Disarmed']);
    const [securityCode, setSecurityCode] = useState(['', '', '', '']);
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

    const handleSecurityCodeInput = (digit, index) => {
        setSecurityCode(prevCode => {
            const newCode = [...prevCode];
            newCode[index] = digit;
            return newCode;
        });
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
                const updatedDeviceData = {...deviceData};
                Object.keys(updatedDeviceData).forEach(device => {
                    if (data[device]) {
                        sessionStorage.setItem(device, JSON.stringify(data[device]));
                        updatedDeviceData[device] = data[device];
                    }
                });
                setDeviceData(updatedDeviceData);
                setAccountType(data.accountType);
                sessionStorage.setItem('accountType', data.accountType);
            })
            .catch(error => console.error('Fetch operation error:', error));
    }, []);

    const handleToggleLock = (uuid) => {
        const isLocking = lockStates[uuid] === 'locked' ? 'unlocked' : 'locked';

        const token = sessionStorage.getItem('token');
        if (!token) {
            console.error("Authorization token not found.");
            return;
        }

        const serverUrl = 'http://localhost:8081/security';
        const requestBody = {
            uuid: uuid,
            name: "Security",
            apptype: "Security",
            function: "LockStatus",
            change: isLocking,
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
                if (response.ok) {
                    console.log(`Lock ${uuid} has been ${isLocking === 'locked' ? 'locked' : 'unlocked'} successfully.`);
                    setLockStates(prevStates => ({...prevStates, [uuid]: isLocking}));
                } else {
                    throw new Error(`Failed to toggle the lock ${uuid} to ${isLocking === 'locked' ? 'locked' : 'unlocked'} with status: ${response.status}`);
                }
            })
            .catch(error => {
                console.error(`There was an error toggling the lock ${uuid} to ${isLocking === 'locked' ? 'locked' : 'unlocked'}:`, error);
            });
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
                    backgroundColor: '#0E2237',
                    width: '100%'
                }}>
                    <div style={{
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center',
                        width: '100%',
                        marginTop: '2rem'
                    }}>
                        <div className="doorSelection Lock" style={{marginRight: '2rem'}}>
                            <div style={{display: 'flex', flexDirection: 'column', alignItems: 'center'}}>
                                <h3 className="centered-title">Back Door</h3>
                                <img className="lockImage" src={doorLockIcon} alt="Lock Icon"
                                     style={{width: '100px', height: 'auto'}}/>
                                <PadlockAnimation
                                    isLocked={lockStates['502857'] === 'locked'}
                                    handleToggleLock={() => handleToggleLock('502857')}
                                />
                            </div>
                        </div>
                        <div className="contentBlock" style={{margin: '5rem'}}>
                            <HomeAlarm/>
                        </div>
                        <div className="doorSelection1 Lock" style={{}}>
                            <div style={{display: 'flex', flexDirection: 'column', alignItems: 'center'}}>
                                <h3 className="centered-title">Front Door</h3>
                                <img className="lockImage" src={doorLockIcon} alt="Lock Icon"
                                     style={{width: '100px', height: 'auto'}}/>

                                <PadlockAnimation
                                    isLocked={lockStates['502858'] === 'locked'}
                                    handleToggleLock={() => handleToggleLock('502858')}
                                />
                            </div>
                        </div>
                    </div>
                </main>
            </div>
        </div>
    );
};

export default Security;