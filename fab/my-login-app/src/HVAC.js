// HVAC.js

import React, {useEffect, useState} from 'react';
import './HVAC.css';
import {useNavigate} from "react-router-dom"; // Import CSS file

const HVAC = () => {

    const [user, setUser] = useState(null);
    const [error, setError] = useState('');
    const [accountType, setAccountType] = useState('')
    const navigate = useNavigate(); // Instantiate useNavigate hook
    const [isNavVisible, setIsNavVisible] = useState(false);

    useEffect(() => {
        const token = sessionStorage.getItem('token');
        const url = 'http://localhost:8081/hvac';

        if (!token) {
            navigate('/'); // Redirect to login page if token is not present
            return;
        }

        fetch(url, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
        })
            .then(response => response.json())
            .then(response => {
                if (response && response.data) {
                    setUser(response.data.user);
                    setAccountType(response.data.accountType);
                    sessionStorage.setItem('accountType', response.data.accountType);
                } else {
                    setError('Unexpected response from server');
                }
            })
            .catch(error => {
                console.log('Fetch operation error:', error)
            });
    }, [navigate]);
    return (
        // Use the class name that matches your CSS file
        <div className="hvacPage">
            <h2>HVAC Page</h2>
            <p>Contents of the HVAC Page</p>
        </div>
    );
};

export default HVAC;
