import React, {useState, useEffect} from 'react';
import {useNavigate} from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import logoImage from '../img/logo.webp';
import {Table} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from "../components/Header";
import Sidebar from "../components/Sidebar";

// Define the Dashboard component using a functional component pattern
const Appliances = () => {

    // States for date and time
    const navigate = useNavigate();
    const [error, setError] = useState('');
    const [accountType, setAccountType] = useState('')
    const [user, setUser] = useState(null);
    const [isNavVisible, setIsNavVisible] = useState(false);


    // URL of the API you want to fetch from
    const apiUrl = 'https://localhost:8081/appliances'

    // Use fetch to get the data from the API
    async function logAppliances() {
        const response = await fetch(apiUrl);
        const appliances = await response.json();
        console.log(appliances);
    }


    const [dishwasher, setDishwasher] = useState([
        {
            Device: "Dishwasher",
            Name: "N/A",
            Location: 'N/A',
            Status: 'ON',
            LastUsed: 'MM/DD/YY 00:00',
            WashTime: "00:00"
        },
        {
            Device: "Dishwasher",
            Name: "N/A",
            Location: 'N/A',
            Status: 'ON',
            LastUsed: 'MM/DD/YY 00:00',
            WashTime: "00:00"
        },
    ]);


    //this const is also for the freezer
    const [fridge, setFridge] = useState([ // <-- Initialize fridge state and setter
        {
            Device: "Fridge",
            Name: "N/A",
            Location: 'N/A',
            Status: 'ON',
            LastUsed: 'MM/DD/YY 00:00',
            Temp: "°F",
            EnergySavingMode: "KW"
        },
        {
            Device: "Fridge",
            Name: "N/A",
            Location: 'N/A',
            Status: 'ON',
            LastUsed: 'MM/DD/YY 00:00',
            Temp: "°F",
            EnergySavingMode: "KW"
        },
        {
            Device: "Freezer",
            Name: "N/A",
            Location: 'N/A',
            Status: 'ON',
            LastUsed: 'MM/DD/YY 00:00',
            Temp: "°F",
            EnergySavingMode: "KW"
        },
        {
            Device: "Freezer",
            Name: "N/A",
            Location: 'N/A',
            Status: 'ON',
            LastUsed: 'MM/DD/YY 00:00',
            Temp: "°F",
            EnergySavingMode: "KW"
        },
    ]);

    const [microwave, setMicrowave] = useState([
        {
            Device: "Microwave",
            Name: "N/A",
            Location: 'N/A',
            Status: 'ON',
            LastUsed: 'MM/DD/YY 00:00',
            Power: "KW",
            StopTime: "00:00"
        },
        {
            Device: "Microwave",
            Name: "N/A",
            Location: 'N/A',
            Status: 'ON',
            LastUsed: 'MM/DD/YY 00:00',
            Power: "KW",
            StopTime: "00:00"
        },
    ]);


    const [toaster, setToaster] = useState([
        {
            Device: "Toaster",
            Name: "N/A",
            Location: 'N/A',
            Status: 'ON',
            LastUsed: 'MM/DD/YY 00:00',
            Temp: "°F",
            StopTime: "00:00"
        },
        {
            Device: "Toaster",
            Name: "N/A",
            Location: 'N/A',
            Status: 'ON',
            LastUsed: 'MM/DD/YY 00:00',
            Temp: "°F",
            StopTime: "00:00"
        },
    ]);


    const [oven, setOven] = useState([
        {
            Device: "Oven",
            Name: "N/A",
            Location: 'N/A',
            Status: 'ON',
            LastUsed: 'MM/DD/YY 00:00',
            Temp: "°F",
            StopTime: "00:00"
        },
        {
            Device: "Oven",
            Name: "N/A",
            Location: 'N/A',
            Status: 'ON',
            LastUsed: 'MM/DD/YY 00:00',
            Temp: "°F",
            StopTime: "00:00"
        },
    ]);

    // Protect Endpoint
    useEffect(() => {
        const token = sessionStorage.getItem('token');
        const url = 'http://localhost:8081/appliances';

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
                    <h2 style={{color: 'white'}}>Appliances</h2>
                    <div style={{alignSelf: 'flex-start', width: '100%'}}>
                        <h3 style={{color: 'white', marginLeft: '1rem'}}>Dishwasher</h3>
                    </div>
                    <Table striped bordered hover variant="dark"
                           style={{marginTop: '20px', backgroundColor: "#173350"}}>
                        <thead>
                        <tr>
                            <th>Device</th>
                            <th>Name</th>
                            <th>Location</th>
                            <th>Status</th>
                            <th>Last Used</th>
                            <th>Wash Time</th>
                        </tr>
                        </thead>
                        <tbody>
                        {dishwasher.map((dishwasherItem, index) => (
                            <tr key={index}>
                                <td>{dishwasherItem.Device}</td>
                                <td>{dishwasherItem.Name}</td>
                                <td>{dishwasherItem.Location}</td>
                                <td>
                                    <button
                                        onClick={() => {
                                            const updatedDishwasher = [...dishwasher];
                                            updatedDishwasher[index] = {
                                                ...dishwasherItem,
                                                Status: dishwasherItem.Status === 'ON' ? 'OFF' : 'ON'
                                            };
                                            setDishwasher(updatedDishwasher);
                                        }}
                                    >
                                        {dishwasherItem.Status}
                                    </button>
                                </td>
                                <td>{dishwasherItem.LastUsed}</td>
                                <td>{dishwasherItem.WashTime}</td>
                            </tr>
                        ))}
                        </tbody>


                    </Table>

                    {/*fridge table*/}
                    <div style={{alignSelf: 'flex-start', width: '100%'}}>
                        <h3 style={{color: 'white', marginLeft: '1rem', marginTop: '10px'}}>Fridge</h3>
                    </div>
                    <Table striped bordered hover variant="dark"
                           style={{marginTop: '20px', backgroundColor: "#173350"}}>
                        <thead>
                        <tr>
                            <th>Device</th>
                            <th>Name</th>
                            <th>Location</th>
                            <th>Status</th>
                            <th>Last Used</th>
                            <th>Temperature</th>
                            <th>Energy Saving Mode</th>
                        </tr>
                        </thead>
                        <tbody>
                        {fridge.map((fridgeItem, index) => (
                            <tr key={index}>
                                <td>{fridgeItem.Device}</td>
                                <td>{fridgeItem.Name}</td>
                                <td>{fridgeItem.Location}</td>
                                {/* Replace the current display of "Status" with a button */}
                                <td>
                                    <button
                                        onClick={() => {
                                            const updatedFridge = [...fridge];
                                            updatedFridge[index] = {
                                                ...fridgeItem,
                                                Status: fridgeItem.Status === 'ON' ? 'OFF' : 'ON'
                                            };
                                            setFridge(updatedFridge);
                                        }}
                                    >
                                        {fridgeItem.Status}
                                    </button>
                                </td>
                                <td>{fridgeItem.LastUsed}</td>
                                <td>{fridgeItem.Temp}</td>
                                <td>{fridgeItem.EnergySavingMode}</td>
                            </tr>
                        ))}

                        </tbody>
                    </Table>

                    {/*Microwave table*/}
                    <div style={{alignSelf: 'flex-start', width: '100%'}}>
                        <h3 style={{color: 'white', marginLeft: '1rem', marginTop: '10px'}}>Microwave</h3>
                    </div>
                    <Table striped bordered hover variant="dark"
                           style={{marginTop: '20px', backgroundColor: "#173350"}}>
                        <thead>
                        <tr>
                            <th>Device</th>
                            <th>Name</th>
                            <th>Location</th>
                            <th>Status</th>
                            <th>Last Used</th>
                            <th>Power</th>
                            <th>Stop Time</th>
                        </tr>
                        </thead>
                        <tbody>
                        {microwave.map((microwaveItem, index) => (
                            <tr key={index}>
                                <td>{microwaveItem.Device}</td>
                                <td>{microwaveItem.Name}</td>
                                <td>{microwaveItem.Location}</td>
                                {/* Replace the current display of "Status" with a button */}
                                <td>
                                    <button
                                        onClick={() => {
                                            const updatedMicrowave = [...microwave];
                                            updatedMicrowave[index] = {
                                                ...microwaveItem,
                                                Status: microwaveItem.Status === 'ON' ? 'OFF' : 'ON'
                                            };
                                            setMicrowave(updatedMicrowave);
                                        }}
                                    >
                                        {microwaveItem.Status}
                                    </button>
                                </td>
                                <td>{microwaveItem.LastUsed}</td>
                                <td>{microwaveItem.Power}</td>
                                <td>{microwaveItem.StopTime}</td>
                            </tr>
                        ))}

                        </tbody>

                    </Table>

                    {/*Toaster table*/}
                    <div style={{alignSelf: 'flex-start', width: '100%'}}>
                        <h3 style={{color: 'white', marginLeft: '1rem', marginTop: '10px'}}>Toaster</h3>
                    </div>
                    <Table striped bordered hover variant="dark"
                           style={{marginTop: '20px', backgroundColor: "#173350"}}>
                        <thead>
                        <tr>
                            <th>Device</th>
                            <th>Name</th>
                            <th>Location</th>
                            <th>Status</th>
                            <th>Last Used</th>
                            <th>Temperature</th>
                            <th>Stop Time</th>
                        </tr>
                        </thead>
                        <tbody>
                        {toaster.map((toasterItem, index) => (
                            <tr key={index}>
                                <td>{toasterItem.Device}</td>
                                <td>{toasterItem.Name}</td>
                                <td>{toasterItem.Location}</td>
                                {/* Replace the current display of "Status" with a button */}
                                <td>
                                    <button
                                        onClick={() => {
                                            const updatedToaster = [...toaster];
                                            updatedToaster[index] = {
                                                ...toasterItem,
                                                Status: toasterItem.Status === 'ON' ? 'OFF' : 'ON'
                                            };
                                            setToaster(updatedToaster);
                                        }}
                                    >
                                        {toasterItem.Status}
                                    </button>
                                </td>
                                <td>{toasterItem.LastUsed}</td>
                                <td>{toasterItem.Temp}</td>
                                <td>{toasterItem.StopTime}</td>
                            </tr>
                        ))}

                        </tbody>

                    </Table>

                    {/*Oven table*/}
                    <div style={{alignSelf: 'flex-start', width: '100%'}}>
                        <h3 style={{color: 'white', marginLeft: '1rem', marginTop: '10px'}}>Oven</h3>
                    </div>
                    <Table striped bordered hover variant="dark"
                           style={{marginTop: '20px', backgroundColor: "#173350"}}>
                        <thead>
                        <tr>
                            <th>Device</th>
                            <th>Name</th>
                            <th>Location</th>
                            <th>Status</th>
                            <th>Last Used</th>
                            <th>Temperature</th>
                            <th>Stop Time</th>
                        </tr>
                        </thead>
                        <tbody>
                        {oven.map((ovenItem, index) => (
                            <tr key={index}>
                                <td>{ovenItem.Device}</td>
                                <td>{ovenItem.Name}</td>
                                <td>{ovenItem.Location}</td>
                                {/* Replace the current display of "Status" with a button */}
                                <td>
                                    <button
                                        onClick={() => {
                                            const updatedOven = [...oven];
                                            updatedOven[index] = {
                                                ...ovenItem,
                                                Status: ovenItem.Status === 'ON' ? 'OFF' : 'ON'
                                            };
                                            setOven(updatedOven);
                                        }}
                                    >
                                        {ovenItem.Status}
                                    </button>
                                </td>
                                <td>{ovenItem.LastUsed}</td>
                                <td>{ovenItem.Temp}</td>
                                <td>{ovenItem.StopTime}</td>
                            </tr>
                        ))}

                        </tbody>

                    </Table>


                </main>

            </div>
        </div>
    );
};

export default Appliances;