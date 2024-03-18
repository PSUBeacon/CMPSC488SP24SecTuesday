import React, { useEffect, useState } from 'react';

////////////////////////////COPY THIS TO DASHBOARD///////////////////////////////
const Dashboard = () => {
    const [dashboardMessage, setDashboardMessage] = useState('');
    const [accountType ,setAccountType] = useState('')
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


    useEffect(() => {
        const token = localStorage.getItem('token');
        const url = 'http://localhost:8081/dashboard';

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
                // Update state for each device if present in response
                const updatedDeviceData = { ...deviceData };
                Object.keys(updatedDeviceData).forEach(device => {
                    if (data[device]) {
                        localStorage.setItem(device, JSON.stringify(data[device]));
                        updatedDeviceData[device] = data[device];
                    }
                });
                setDeviceData(updatedDeviceData);
                setDashboardMessage(data.message);

                // Store accountType in session storage
                setAccountType(data.accountType);
                sessionStorage.setItem('accountType', data.accountType);
            })
            .catch(error => console.error('Fetch operation error:', error));
    }, []);

    return (
        ////////////////////////////////////////////////
        <div>
            <h2>Dashboard</h2>
            {dashboardMessage && <p>{dashboardMessage}</p>}


            {deviceData.HVAC.Temperature && (
                <p>{deviceData.HVAC.Temperature}</p>
            )}


            {deviceData.SolarPanel.Status && (
                <p>{deviceData.SolarPanel.Status}</p>
            )}

            <p>{accountType}</p>

        </div>
    );
};

export default Dashboard;
