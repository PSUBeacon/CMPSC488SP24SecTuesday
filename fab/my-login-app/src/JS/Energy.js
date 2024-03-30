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
    const [accountType, setAccountType] = useState('');
    const [isNavVisible, setIsNavVisible] = useState(false);
    const navigate = useNavigate(); // Instantiate useNavigate hook


    const energy = [
        {Device: 'Microwave', NetlossEnergy: '%', NetgainEnergy: '%', Battery: '%', Status: 'ON/OFF'},
        {Device: 'Oven', NetlossEnergy: '%', NetgainEnergy: '%', Battery: '%', Status: 'ON/OFF'},
        {Device: 'Fridge', NetlossEnergy: '%', NetgainEnergy: '%', Battery: '%', Status: 'ON/OFF'},
        {Device: 'Freezer', NetlossEnergy: '%', NetgainEnergy: '%', Battery: '%', Status: 'ON/OFF'},
        {Device: 'Toaster', NetlossEnergy: '%', NetgainEnergy: '%', Battery: '%', Status: 'ON/OFF'},
        {Device: 'Dishwasher', NetlossEnergy: '%', NetgainEnergy: '%', Battery: '%', Status: 'ON/OFF'},

    ];

    {/*solar panel statistics table*/
    }
    const solarPanel = [
        {TotalEnergy: 'KW', EnergyUsedToday: 'KW'},
        {TotalEnergy: 'KW', EnergyUsedToday: 'KW'},

    ];

    const handleGoToAppliances = () => {
        navigate('/appliances'); // Adjust the route as necessary
    };

    // Object that holds the URLs for your camera feeds
    const cameraFeeds = {
        livingroom: placeholderImage, // Replace with the actual camera feed URL or image for the living room
        kitchen: placeholderImage2, // Replace with the actual camera feed URL or image for the kitchen
        // Add more camera feeds as needed
    };


    // This is the JSX return statement where we layout our component's HTML structure
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
                    <h2 style={{color: 'white'}}>Devices Using Energy</h2>
                    <Table striped bordered hover variant="dark"
                           style={{marginTop: '20px', backgroundColor: "#173350"}}>
                        <thead>
                        <tr>
                            <th>
                                Device
                                <button onClick={handleGoToAppliances} style={{
                                    marginLeft: '10px',
                                    padding: '2px 6px',
                                    fontSize: '0.8em',
                                    background: '#0294A5',
                                    color: 'white',
                                    border: 'none',
                                    borderRadius: '4px',
                                    cursor: 'pointer'
                                }}>
                                    See More
                                </button>
                            </th>
                            <th>Net loss Energy</th>
                            <th>Net Gain Energy</th>
                            <th>Battery %</th>
                            <th>Status</th>
                        </tr>
                        </thead>
                        <tbody>
                        {energy.map((energy, index) => (
                            <tr key={index}>
                                <td>{energy.Device}</td>
                                <td>{energy.NetlossEnergy}</td>
                                <td>{energy.NetgainEnergy}</td>
                                <td>{energy.Battery}</td>
                                <td>{energy.Status}</td>
                            </tr>
                        ))}
                        </tbody>
                    </Table>

                    {/*this is the table for the solar panel content*/}
                    <h2 style={{color: 'white'}}>Solar Panel Statistics</h2>
                    <Table striped bordered hover variant="dark"
                           style={{marginTop: '20px', backgroundColor: "#173350"}}>
                        <thead>
                        <tr>
                            <th>Total Energy</th>
                            <th>Energy Used Today</th>
                        </tr>
                        </thead>
                        <tbody>
                        {solarPanel.map((solarPanel, index) => (
                            <tr key={index}>

                                <td>{solarPanel.TotalEnergy}</td>
                                <td>{solarPanel.EnergyUsedToday}</td>
                            </tr>
                        ))}
                        </tbody>
                    </Table>

                </main>

            </div>
        </div>
    );
};


export default Energy;