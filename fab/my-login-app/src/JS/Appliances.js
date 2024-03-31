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

    const [data, setData] = useState({});

    // Protect Endpoint
    useEffect(() => {
        const token = sessionStorage.getItem('token');
        const url = 'http://localhost:8081/appliances';

        if (!token) {
            navigate('/'); // Redirect to login page if token is not present
            return;
        }

        const fetchData = async () => {
                try {
                        const response = await fetch(url, {
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
                    {Object.keys(data).length > 0 ? (
                        Object.entries(data).map(([key, appliances]) => (
                                <div key={key} style={{alignItems: 'center', width: '70%', marginTop: '20px'}}>
                                        <h3>{key}</h3>
                                        <Table striped bordered hover variant="dark" style={{backgroundColor: "#173350"}}>
                                            <thead>
                                                <tr>
                                                    <th>Device</th>
                                                    {/*<th>Name</th>*/}
                                                    <th>Location</th>
                                                    <th>Status</th>
                                                    <th>Last Used</th>
                                                </tr>
                                            </thead>
                                            <tbody>
                                            {appliances.map((appliance, index) => (
                                                    <tr key={index}>
                                                        {/* Assuming the name is stored in a property called 'Name' */}
                                                        <td style={{width: '25%'}}>{key} - {appliance.UUID}</td>
                                                        {/*<td style={{width: '20%'}}>{appliance.Label ? "Label" : ""}</td>*/}
                                                        <td style={{width: '25%'}}>{appliance.Location}</td>
                                                        <td style={{width: '25%'}}>{appliance.Status ? "On" : "Off"}</td>
                                                        <td style={{width: '25%'}}>{appliance.LastChanged}</td>
                                                    </tr>
                                            ))}
                                            </tbody>
                                        </Table>
                                </div>
                        ))
                    ) : (<p>Loading...</p>)}
                </main>
            </div>
        </div>
    );
};

export default Appliances;