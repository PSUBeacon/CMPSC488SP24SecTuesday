import React, { useEffect, useState } from 'react';

const OtherPage = () => {
    //Set and create device variables
    const [hvacStatus, setHvacStatus] = useState('');
    const [securityLocation, setSecurityLocation] = useState('');

    useEffect(() => {
        //Json parser to extract collection and field
        const deviceData = JSON.parse(localStorage.getItem('deviceData'));

        //set parsed data
        setHvacStatus(deviceData.HVAC.Status);
        setSecurityLocation(deviceData.SecuritySystem.Location);


    }, []);

    return (
        <div>
            <h2>Other Page</h2>
            <p>HVAC Status: {hvacStatus}</p>
            <p>Security System Location: {securityLocation}</p>
        </div>
    );
};

export default OtherPage;
