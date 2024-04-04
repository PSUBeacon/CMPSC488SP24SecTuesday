import React, {useState, useEffect} from 'react';
import {useNavigate} from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import logoImage from '../img/logo.webp';
import {Table} from 'react-bootstrap';
import placeholderImage from '../img/placeholderImage.jpg'; // Replace with the path to your placeholder image
import placeholderImage2 from '../img/placeholderImage2.jpg'; // Replace with the path to your placeholder image

import {
    faMicrophone, // Placeholder icon, replace with the actual icon for the microwave
    faOtter, // Placeholder icon, replace with the actual icon for the oven
    faIceCream, // Placeholder icon, replace with the actual icon for the fridge
    faSnowflake, // Placeholder icon, replace with the actual icon for the freezer
    faBreadSlice, // Placeholder icon, replace with the actual icon for the toaster
    faSoap // Placeholder icon, replace with the actual icon for the dishwasher
} from '@fortawesome/free-solid-svg-icons';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from "../components/Header";
import Sidebar from "../components/Sidebar";

// Define the Dashboard component using a functional component pattern
const Energy = () => {
    document.title = 'BEACON | Energy';

    const [accountType, setAccountType] = useState('');
    const [isNavVisible, setIsNavVisible] = useState(false);
    const navigate = useNavigate(); // Instantiate useNavigate hook
    const [user, setUser] = useState(null);
    const [error, setError] = useState('');

    const CollectionMapping = {
        dishwasher: 'Appliances',
        toaster: 'Appliances',
        oven: 'Appliances',
        fridge: 'Appliances',
        microwave: 'Appliances',
        hvac: 'HVAC',
        solarpanel: 'Energy',
        securitysystem: 'Security',
    };

    const handleGoToAppliances = () => {
        navigate('/appliances'); // Adjust the route as necessary
    };

    const handleEnergyStatusChange = async (uuid, key, status) => {
        const token = sessionStorage.getItem('token');
        if (!token) {
            navigate('/');
            return;
        }

        const category = CollectionMapping[key.toLowerCase()] || key;

        const requestBody = {
            uuid: uuid,
            name: category,
            apptype: key,
            function: "Status",
            change: status,
        };

        try {
            const response = await fetch('http://localhost:8081/energy', {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(requestBody),
            });
            if (!response.ok) throw new Error('Network response was not ok');
            const updatedData = await response.json();
            setData(updatedData);
        } catch (error) {
            console.error('Failed to update energy status:', error);
        }
    };

    const [data, setData] = useState({});

    useEffect(() => {
        const token = sessionStorage.getItem('token');
        if (!token) {
            navigate('/');
            return;
        }
        const fetchData = async () => {
            try {
                const response = await fetch('http://localhost:8081/energy/GetEnergy', {
                    method: 'POST',
                    headers: {
                        'Authorization': `Bearer ${token}`,
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({}),
                });
                if (!response.ok) throw new Error('Network response was not ok');
                const jsonData = await response.json();
                setData(jsonData);
            } catch (error) {
                console.error('Failed to fetch data:', error);
            }
        };
        fetchData();
    }, [navigate]);


    // This is the JSX return statement where we lay out our component's HTML structure
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
                    backgroundColor: '#0E2237'
                }}>
                    <h2 style={{color: 'white'}}>Devices Energy Usage</h2>
                    {Object.keys(data).length > 0 ? Object.entries(data).map(([key, appliances]) => (
                        <div key={key} style={{alignItems: 'center', width: '70%'}}>
                            <h3>{key}</h3>
                            <Table striped bordered hover variant="dark"
                                   style={{marginTop: '20px', backgroundColor: "#173350"}}>
                                <thead>
                                <tr>
                                    <th>Device Name</th>
                                    <th>Status</th>
                                    <th>Energy Consumption, kWh</th>
                                    <th>Action</th>
                                </tr>
                                </thead>
                                <tbody>
                                {appliances.map((appliance, index) => (
                                    <tr key={index}>
                                        <td style={{width: '25%'}}>{key} - {appliance.Location}</td>
                                        {/* NEED TO USE DEVICE NAME-LABEL */}
                                        <td style={{width: '25%'}}>{appliance.Status ? "On" : "Off"}</td>
                                        <td style={{width: '25%'}}>{appliance.EnergyConsumption}</td>
                                        <td style={{width: '25%'}}>
                                            {/* Implement actual toggle functionality as needed */}
                                            <button
                                                onClick={() => handleEnergyStatusChange(appliance.UUID, key, appliance.Status)}>
                                                Turn On/Off
                                            </button>
                                        </td>
                                    </tr>
                                ))}
                                </tbody>
                            </Table>
                        </div>
                    )) : <p>Loading...</p>}
                </main>
            </div>
        </div>
    );
};


export default Energy;