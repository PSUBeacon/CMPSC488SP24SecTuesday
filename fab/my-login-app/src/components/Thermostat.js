import React, {useCallback, useEffect, useState} from 'react';
import '../CSS/Thermostat.css';
import axios from "axios";

let func;
let temperature, setTemperature, mode, setMode, fanSpeed, setFanSpeed, inituuid, setInituuid;


const sendServerRequest = () => {
    const requestBody = {
        UUID: inituuid,
        Name: "HVAC",
        AppType: "HVAC",
        Function: func,
        Change: func === "Mode" ? mode :
            func === "Temperature" ? temperature :
                func === "FanSpeed" ? JSON.stringify(fanSpeed) :
                    mode === "off" ? "false" : "true",
    };

    try {
        console.log(requestBody)
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
    [mode, setMode] = useState(initialStatus === 'false' ? 'off' : initialMode);

    const handleModeChange = async (newMode) => {
        await setMode(newMode);
        await onModeChange(newMode);
        func = newMode === "cool" || newMode === "heat" ? "Mode" : "Status";
        await sendServerRequest();
    };

    return (
        <div className="mode-toggle">
            <span className="mode-label"></span>
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
    [fanSpeed, setFanSpeed] = useState(initialFanSpeed);

    const handleFanSpeedChange = async (newFanSpeed) => {
        await setFanSpeed(newFanSpeed);
        await onFanSpeedChange(newFanSpeed);
        func = "FanSpeed";
        await sendServerRequest();
    };

    return (
        <div className="fan-speed-toggle">
            <span className="mode-label"></span>
            <button
                className={`fan-speed-button ${fanSpeed === 30 ? 'active' : ''}`}
                onClick={() => handleFanSpeedChange(30)}
            >
                Low
            </button>
            <button
                className={`fan-speed-button ${fanSpeed === 50 ? 'active' : ''}`}
                onClick={() => handleFanSpeedChange(50)}
            >
                Medium
            </button>
            <button
                className={`fan-speed-button ${fanSpeed === 90 ? 'active' : ''}`}
                onClick={() => handleFanSpeedChange(90)}
            >
                High
            </button>
        </div>
    );
};


const NestThermostat = ({
                            thermostat
                        }) => {
    [temperature, setTemperature] = useState(thermostat.Temperature || 76);
    [mode, setMode] = useState(thermostat.Status === 'false' ? 'off' : thermostat.Mode);
    [fanSpeed, setFanSpeed] = useState(thermostat.FanSpeed);
    [inituuid, setInituuid] = useState(thermostat.UUID);
    const [dimmerValue, setDimmerValue] = useState(thermostat.Temperature);

    const minTemp = 60;
    const maxTemp = 90;

    const handleTemperatureChange =
        async (newTemp) => {
            func = "Temperature";
            await setTemperature(newTemp);
            console.log(func + ": " + temperature)
            await sendServerRequest()
        };

    const handleInputChange = useCallback(
        async (event) => {
            const inputTemp = parseInt(event.target.value, 10);
            if (!isNaN(inputTemp) && inputTemp >= minTemp && inputTemp <= maxTemp) {
                setTemperature(inputTemp);
            }
            await handleTemperatureChange(Math.max(temperature - 1, minTemp));

        },
        [setTemperature, minTemp, maxTemp]
    );

    const handleDecrement = useCallback(
        async () => {
            func = "Temperature"
            await handleTemperatureChange(Math.max(temperature - 1, minTemp));
            await sendServerRequest();
        },
        [temperature, minTemp, handleTemperatureChange]
    );

    const handleIncrement = useCallback(
        async () => {
            func = "Temperature"
            await handleTemperatureChange(Math.min(temperature + 1, maxTemp));
            await sendServerRequest();
        },
        [temperature, maxTemp, handleTemperatureChange]
    );

    return (
        <div className="thermostat">
            <ModeToggle initialMode={thermostat.Mode} onModeChange={setMode} initialStatus={thermostat.Status}/>
            <FanSpeedToggle initialFanSpeed={thermostat.FanSpeed} onFanSpeedChange={setFanSpeed}/>
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
                {/*<button onClick={handleDecrement} className="temperature-button">*/}
                {/*    -*/}
                {/*</button>*/}
                {/*<input*/}
                {/*    type="number"*/}
                {/*    value={temperature}*/}
                {/*    min={minTemp}*/}
                {/*    max={maxTemp}*/}
                {/*    onChange={handleInputChange}*/}
                {/*    className="temperature-input"*/}
                {/*    style={{display: 'none'}}*/}
                {/*/>*/}
                {/*<button onClick={handleIncrement} className="temperature-button">*/}
                {/*    +*/}
                {/*</button>*/}

                <input
                    type="range"
                    id="dimmer"
                    name="dimmer"
                    min="60"
                    max="90"
                    value={temperature}
                    onChange={(e) => {
                        const newValue = (e.target.value);
                        setTemperature(newValue)
                    }}
                    onMouseUp={(e) => {

                        // handleInputChange(e);
                        handleTemperatureChange(temperature);
                        console.log(temperature)

                    }}
                    style={{
                        WebkitAppearance: 'none',
                        width: '100%',
                        height: '15px',
                        background: 'White',
                        outline: 'none',
                        opacity: '1',
                        transition: 'opacity .2s',
                        borderRadius: '5px',
                    }}
                />
                <label htmlFor="dimmer"
                       style={{
                           color: '#fff',
                           marginTop: '5px'
                       }}>{temperature}</label>
                <div className="humidity-display">Humidity: {thermostat.Humidity}%</div>
                <div className="humidity-display">Power Usage: {thermostat.EnergyConsumpstion}kW</div>
            </div>
        </div>
    );
};

export default NestThermostat;
