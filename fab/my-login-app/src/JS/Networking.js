import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { Table } from 'react-bootstrap';
import Header from "../components/Header";
import Sidebar from "../components/Sidebar";
import 'bootstrap/dist/css/bootstrap.min.css';

const Networking = () => {
    const navigate = useNavigate();
    const [isNavVisible, setIsNavVisible] = useState(false);
    const [accountType, setAccountType] = useState('');
    const [Logs, setLogs] = useState([]);
    const [currentView, setCurrentView] = useState('iotLogs');
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState('');

    document.title = 'BEACON | Logs';
  
    useEffect(() => {
        const token = sessionStorage.getItem('token');
        if (!token) {
            navigate('/');
            return;
        }
        fetchLogs();
    }, [currentView]); // Fetch logs also when currentView changes

    const fetchLogs = () => {
        setIsLoading(true);
        setError('');
        const token = sessionStorage.getItem('token');
        const url = currentView === 'iotLogs' ? "https://beacon-cs2024.digital/api/networking/GetNetLogs" : "https://beacon-cs2024.digital/api/networking/GetNetPcapLogs";

        fetch(url, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json',
            },
        })
            .then(response => {
                console.log(response)
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then(data => {
                console.log(data)
                if (Array.isArray(data)) {
                    setLogs(data);
                } else {
                    throw new Error('Data format is incorrect');
                }
            })
            .catch(error => {
                console.error('Error fetching data:', error);
                setError('Failed to load data');
            })
            .finally(() => setIsLoading(false));
    };

    const changeView = (view) => {
        setCurrentView(view);
    };

    const renderLogsTable = () => (

        <div>
            <h2 style={{color: '#173350'}}>IOT Logs</h2>
            <Table striped bordered hover variant="dark">
            <thead>
                <tr>
                    <th>DeviceID</th>
                    <th>Function</th>
                    <th>Change</th>
                    <th>Time</th>
                </tr>
                </thead>
                <tbody>
                {Logs.map((log, index) => (
                    <tr key={index}>
                        <td>{log.DeviceID}</td>
                        <td>{log.Function}</td>
                        <td>{log.Change}</td>
                        <td>{log.Time}</td>
                    </tr>
                ))}
                </tbody>
            </Table>
        </div>
    );

    const renderPcapLogsTable = () => (
        <div>
            <h2 style={{color: '#173350'}}>Network Logs</h2>
            <Table striped bordered hover variant="dark">
                <thead>
                <tr>
                    <th>Network Type</th>
                    <th>Source MAC</th>
                    <th>Destination MAC</th>
                    <th>Source IP</th>
                    <th>Destination IP</th>
                    <th>Source Port</th>
                    <th>Destination Port</th>
                </tr>
                </thead>
                <tbody>
                {Logs.map((log, index) => (
                    <tr key={index}>
                        <td>{log.network_type}</td>
                        <td>{log.source_mac}</td>
                        <td>{log.destination_mac}</td>
                        <td>{log.source_ip}</td>
                        <td>{log.destination_ip}</td>
                        <td>{log.source_port}</td>
                        <td>{log.destination_port}</td>
                    </tr>
                ))}
                </tbody>
            </Table>
        </div>
    );

    return (
        <div style={{display: 'flex', minHeight: '100vh', flexDirection: 'column', backgroundColor: '#081624' }}>
            <Header accountType={accountType} />
            <div style={{ display: 'flex', flex: '1' }}>
                <Sidebar isNavVisible={isNavVisible} />
                <main style={{ flex: '1', padding: '1rem', display: 'flex', flexDirection: 'column', alignItems: 'center', backgroundColor: '#0E2237' }}>
                    <div style={{ width: '100%', maxWidth: '1700px', display: 'flex', flexDirection: 'row', flexWrap: 'wrap', justifyContent: 'center', gap: '20px', marginBottom: '20px' }}>
                        {/*<div className="camera-widget" style={{ maxWidth: '80%', backgroundColor: '#173350', borderRadius: '10px', overflow: 'hidden', flexBasis: '100%', padding: '12px', marginTop: "70px" }}>*/}
                            <div className="camera-widget" style={{ position: 'relative', maxWidth: '80%', backgroundColor: '#173350', borderRadius: '10px', overflow: 'hidden', flexBasis: '100%', padding: '12px', marginTop: "70px" }}>

                            {isLoading && <div>Loading...</div>}
                            {!isLoading && error && <div>Error: {error}</div>}
                            {!isLoading && !error && currentView === 'iotLogs' && renderLogsTable()}
                            {!isLoading && !error && currentView === 'networkLogs' && renderPcapLogsTable()}
                            <div style={{ position: 'absolute', top: '10px', left: '10px', display: 'flex', gap: '5px' }}>
                                <button onClick={() => changeView('iotLogs')} style={{ padding: '5px', color: 'white', backgroundColor: currentView === 'iotLogs' ? '#0294A5' : '#08192B' }}>IOT Logs</button>
                                <button onClick={() => changeView('networkLogs')} style={{ padding: '5px', color: 'white', backgroundColor: currentView === 'networkLogs' ? '#0294A5' : '#08192B' }}>Network Logs</button>
                            </div>
                        </div>
                    </div>
                </main>
            </div>
        </div>
    );
};

export default Networking;
