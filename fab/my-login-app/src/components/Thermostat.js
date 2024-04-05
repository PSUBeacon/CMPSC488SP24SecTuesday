import React, {useCallback, useEffect, useState} from 'react';
import '../CSS/Thermostat.css';
import axios from "axios";

let func;
let temperature, setTemperature, mode, setMode, fanSpeed, setFanSpeed;
let INITuuid;

const sendServerRequest = () => {
    const requestBody = {
        UUID: INITuuid,
        Name: "HVAC",
        AppType: "HVAC",
        Function: func,
        Change: func === "Mode" ? mode :
            func === "Temperature" ? JSON.stringify(temperature) :
                func === "FanSpeed" ? JSON.stringify(fanSpeed) :
                    mode === "off" ? "false" : "true",
    };

    try {
        const response = axios.post('http://localhost:8081/hvac/updateHVAC', requestBody, {
            headers: {
                'Content-Type': 'application/json',
            },
        });
    } catch (error) {
        console.error('Failed to send settings to server:', error);
    }
};

const ModeToggle = ({initialMode, onModeChange, initialStatus}) => {
    const [mode, setMode] = useState(initialStatus == 'false' ? 'off' : initialMode);

    const handleModeChange = (newMode) => {
        setMode(newMode);
        onModeChange(newMode);
        func = mode === "cool" || mode === "heat" ? "Mode" : "Status";
        sendServerRequest();
    };

    return (
        <div className="mode-toggle">
            <span className="mode-label">Mode:</span>
            <button
                className={`mode-button ${mode === 'cool' ? 'active' : ''}`}
                onClick={() => handleModeChange('cool')}
            >
                Cool
            </button>
            <button
                className={`mode-button ${mode === 'off' ? 'active' : ''}`}
                onClick={() => handleModeChange('off')}
            >
                Off
            </button>
            <button
                className={`mode-button ${mode === 'heat' ? 'active' : ''}`}
                onClick={() => handleModeChange('heat')}
            >
                Heat
            </button>
        </div>
    );
};

const FanSpeedToggle = ({initialFanSpeed, onFanSpeedChange}) => {
    const [fanSpeed, setFanSpeed] = useState(initialFanSpeed);

    const handleFanSpeedChange = (newFanSpeed) => {
        setFanSpeed(newFanSpeed);
        onFanSpeedChange(newFanSpeed);
        func = "FanSpeed";
        sendServerRequest();
    };

    return (
        <div className="fan-speed-toggle">
            <span className="mode-label">Fan:</span>
            <button
                className={`fan-speed-button ${fanSpeed === 1 ? 'active' : ''}`}
                onClick={() => handleFanSpeedChange(1)}
            >
                Low
            </button>
            <button
                className={`fan-speed-button ${fanSpeed === 2 ? 'active' : ''}`}
                onClick={() => handleFanSpeedChange(2)}
            >
                Medium
            </button>
            <button
                className={`fan-speed-button ${fanSpeed === 3 ? 'active' : ''}`}
                onClick={() => handleFanSpeedChange(3)}
            >
                High
            </button>
        </div>
    );
};

const NestThermostat = ({
                            uuid,
                            initialTemperature,
                            initialMode,
                            initialHumidity,
                            initialFanSpeed,
                            initialStatus,
                            initialPU
                        }) => {
    [temperature, setTemperature] = useState(initialTemperature || 76);
    [mode, setMode] = useState(initialMode);
    [fanSpeed, setFanSpeed] = useState(initialFanSpeed);
    INITuuid = uuid;

    const minTemp = 60;
    const maxTemp = 90;

    const handleTemperatureChange = useCallback(
        (newTemp) => {
            setTemperature(newTemp);
        },
        [setTemperature]
    );

    const handleInputChange = useCallback(
        (event) => {
            const inputTemp = parseInt(event.target.value, 10);
            if (!isNaN(inputTemp) && inputTemp >= minTemp && inputTemp <= maxTemp) {
                setTemperature(inputTemp);
            }
        },
        [setTemperature, minTemp, maxTemp]
    );

    const handleDecrement = useCallback(
        () => {
            func = "Temperature"
            handleTemperatureChange(Math.max(temperature - 1, minTemp));
            sendServerRequest();
        },
        [temperature, minTemp, handleTemperatureChange]
    );

    const handleIncrement = useCallback(
        () => {
            func = "Temperature"
            handleTemperatureChange(Math.min(temperature + 1, maxTemp));
            sendServerRequest();
        },
        [temperature, maxTemp, handleTemperatureChange]
    );

    return (
        <div className="thermostat-container">
            <div className="thermostat">
                <ModeToggle initialMode={initialMode} onModeChange={setMode} initialStatus={initialStatus}/>
                <FanSpeedToggle initialFanSpeed={initialFanSpeed} onFanSpeedChange={setFanSpeed}/>
                <svg viewBox="0 0 100 100" className="thermostat-dial">
                    <circle cx="50" cy="50" r="45" className="dial-background"/>
                    <path
                        d={`M 50 50 L 5 50 A 45 45 0 ${temperature > 180 ? 1 : 0} 1 ${
                            Math.cos((temperature - 90) * (Math.PI / 180)) * 45 + 50
                        } ${Math.sin((temperature - 90) * (Math.PI / 180)) * 45 + 50} Z`}
                        className={'dial-progress-' + mode}
                    />
                    <text x="50" y="50" className="temperature-text">
                        {temperature}
                    </text>
                    {Array.from({length: 60}, (_, i) => (
                        <line
                            key={i}
                            x1="50"
                            y1="5"
                            x2="50"
                            y2="10"
                            transform={`rotate(${i * 6}, 50, 50)`}
                            className="dial-tick"
                        />
                    ))}
                </svg>
                <div className="temperature-controls">
                    <input
                        type="number"
                        value={temperature}
                        min={minTemp}
                        max={maxTemp}
                        onChange={handleInputChange}
                        className="temperature-input"
                        style={{display: 'none'}}
                    />
                    <button onClick={handleDecrement} className="temperature-button">
                        -
                    </button>
                    <button onClick={handleIncrement} className="temperature-button">
                        +
                    </button>
                    <div className="humidity-display">Humidity: {initialHumidity}%</div>
                    <div className="humidity-display">Power Usage: {initialPU}kW</div>
                </div>
            </div>
        </div>
    );
};

export default NestThermostat;
