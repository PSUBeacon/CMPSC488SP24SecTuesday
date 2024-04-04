import React, {useCallback, useState} from 'react';
import '../CSS/Thermostat.css';

const ModeToggle = ({initialMode, onModeChange}) => {
    const [mode, setMode] = useState(initialMode);

    const handleModeChange = (newMode) => {
        setMode(newMode);
        onModeChange(newMode);
    };

    return (
        <div className="mode-toggle">
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
    };

    return (
        <div className="fan-speed-toggle">
            <button
                className={`fan-speed-button ${fanSpeed === 'low' ? 'active' : ''}`}
                onClick={() => handleFanSpeedChange('low')}
            >
                Low
            </button>
            <button
                className={`fan-speed-button ${fanSpeed === 'medium' ? 'active' : ''}`}
                onClick={() => handleFanSpeedChange('medium')}
            >
                Medium
            </button>
            <button
                className={`fan-speed-button ${fanSpeed === 'high' ? 'active' : ''}`}
                onClick={() => handleFanSpeedChange('high')}
            >
                High
            </button>
        </div>
    );
};

const NestThermostat = ({initialTemperature, initialMode, initialHumidity, initialFanSpeed, initialPU}) => {
    const [temperature, setTemperature] = useState(initialTemperature || 76);
    const [mode, setMode] = useState(initialMode);
    const [fanSpeed, setFanSpeed] = useState(initialFanSpeed);

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
            console.log('Decrement button clicked');
            handleTemperatureChange(Math.max(temperature - 1, minTemp));
        },
        [temperature, minTemp, handleTemperatureChange]
    );

    const handleIncrement = useCallback(
        () => {
            console.log('Increment button clicked');
            handleTemperatureChange(Math.min(temperature + 1, maxTemp));
        },
        [temperature, maxTemp, handleTemperatureChange]
    );

    return (
        <div className="thermostat-container">
            <div className="thermostat">
                <ModeToggle initialMode={initialMode} onModeChange={setMode}/>
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
                    <div className="humidity-display">Initial Humidity: {initialHumidity}</div>
                    <div className="humidity-display">Power Usage: {initialPU}</div>
                </div>
            </div>
        </div>
    );
};

export default NestThermostat;