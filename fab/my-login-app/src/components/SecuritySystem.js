// HomeAlarm.js

import React, {useState} from 'react';
import '../CSS/SecuritySystem.css';
import PadlockAnimation from './Lock'

function HomeAlarm() {
    const [inputCode, setInputCode] = useState('');
    const [displayCode, setDisplayCode] = useState('');


    const handleNumberPress = (number) => {
        if (inputCode.length < 6) {
            setInputCode(inputCode + number);
            setDisplayCode(displayCode + '*');
        }
    };

    const handleArmHome = () => {
        console.log('ARM HOME pressed with code:', inputCode);
        // Add logic for when ARM HOME is pressed
    };

    const handleArmAway = () => {
        console.log('ARM AWAY pressed with code:', inputCode);
        // Add logic for when ARM AWAY is pressed
    };

    const handleClear = () => {
        setInputCode('');
        setDisplayCode('');
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
                    <button className="ArmKey" onClick={handleArmHome}>ARM</button>
                    <button className="DisarmKey" onClick={handleArmAway}>DISARM</button>
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
