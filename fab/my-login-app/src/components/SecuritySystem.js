import React, {useState} from 'react';
import '../CSS/SecuritySystem.css';
import PadlockAnimation from './Lock';

function HomeAlarm() {
    const [inputCode, setInputCode] = useState('');
    const [displayCode, setDisplayCode] = useState('');

    const handleNumberPress = (number) => {
        if (inputCode.length < 6) {
            setInputCode(inputCode + number);
            setDisplayCode(displayCode + '*');
        }
    };

    const handleArm = () => {
        console.log('ARM HOME pressed with code:', inputCode);
        sendCodeToBackend(inputCode, "true");
    };

    const handleDisarm = () => {
        console.log('ARM AWAY pressed with code:', inputCode);
        sendCodeToBackend(inputCode, "false");
    };

    const handleClear = () => {
        setInputCode('');
        setDisplayCode('');
    };

    const sendCodeToBackend = (code, AlarmStatus) => {
        const token = sessionStorage.getItem('token');
        if (!token) {
            console.error("Authorization token not found.");
            return;
        }
        const serverUrl = 'http://192.168.8.117:8081/security/system';
        const requestData = {
            code: code,
            status: AlarmStatus,
        };

        fetch(serverUrl, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(requestData),
        })
            .then(response => {
                if (response.ok) {
                    console.log(`Code ${code} has been sent to the backend successfully.`);
                    // Reset the input code and display code after sending to the backend
                    setInputCode('');
                    setDisplayCode('');
                } else {
                    throw new Error(`Failed to send the code ${code} to the backend with status: ${response.status}`);
                }
            })
            .catch(error => {
                console.error(`There was an error sending the code ${code} to the backend:`, error);
            });
    };

    return (
        <div>
            <div className="home-alarm">
                <h2>Home Alarm</h2>
                <div className="status-indicator">
                    <input
                        type="text"
                        className="code-display"
                        value={displayCode}
                        readOnly
                    />
                </div>
                <div className="buttons-row">
                    <button className="ArmKey" onClick={handleArm}>ARM</button>
                    <button className="DisarmKey" onClick={handleDisarm}>DISARM</button>
                </div>
                <div className="code-entry">
                    <div className="numbers-grid">
                        {[1, 2, 3, 4, 5, 6, 7, 8, 9, '*', 0, '#'].map((number) => (
                            <button className="keyPadKeys" key={number} onClick={() => handleNumberPress(number)}>
                                {number}
                            </button>
                        ))}
                    </div>
                    <button className="clear-button keyPadKeys" onClick={handleClear}>
                        CLEAR
                    </button>
                </div>

            </div>
        </div>
    );
}

export default HomeAlarm;
