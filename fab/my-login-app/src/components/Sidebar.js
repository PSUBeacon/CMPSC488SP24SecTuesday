import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import houseImage from '../img/houseImage.jpg';
import '../CSS/SideBar.css';


const Sidebar = ({ isNavVisible }) => {
    const location = useLocation();

    // Helper function to determine if the path matches
    const isPathActive = (path) => {
        return location.pathname === path;
    };

    return (
        <aside className={`side-nav ${isNavVisible ? 'visible' : 'hidden'}`} style={{ backgroundColor: '#0E2237', color: 'white', width: '250px', padding: '1rem' }}>
            <div className="houseInfo">
                <div><img src={houseImage} alt="House" style={{ marginRight: '10px', borderRadius: '50%' }} id='circle2'/></div>
                <div>My House</div>
                <div>State College, PA 16801</div>
            </div>
            <nav>
                <ul style={{ listStyleType: 'none', padding: 0 }}>
                    <li className={`nav-item ${isPathActive('/dashboard') ? 'active' : ''}`} style={{margin: '0.5rem 0', padding: '0.5rem' }}>
                        <Link to="/dashboard" style={{ color: isPathActive('/dashboard') ? '#50BCC0' : '#95A4B6', textDecoration: 'none' }}>
                            <i className="fas fa-home" style={{ marginRight: '10px' }}></i>
                            Overview
                        </Link>
                    </li>
                    <li className={`nav-item ${isPathActive('/security') ? 'active' : ''}`} style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                        <Link to="/security" style={{ color: isPathActive('/security') ? '#50BCC0' : '#95A4B6', textDecoration: 'none' }}>
                            <i className="fas fa-lock" style={{ marginRight: '10px' }}></i>
                            Security
                        </Link>
                    </li>
                    <li className={`nav-item ${isPathActive('/lighting') ? 'active' : ''}`} style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                        <Link to="/lighting" style={{ color: isPathActive('/lighting') ? '#50BCC0' : '#95A4B6', textDecoration: 'none' }}>
                            <i className="fas fa-lightbulb" style={{ marginRight: '10px' }}></i>
                            Lighting
                        </Link>
                    </li>
                    <li className={`nav-item ${isPathActive('/networking') ? 'active' : ''}`} style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                        <Link to="/networking" style={{ color: isPathActive('/networking') ? '#50BCC0' : '#95A4B6', textDecoration: 'none' }}>
                            <i className="fas fa-sliders-h" style={{ marginRight: '10px' }}></i>
                            Networking
                        </Link>
                    </li>
                    <li className={`nav-item ${isPathActive('/hvac') ? 'active' : ''}`} style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                        <Link to="/hvac" style={{ color: isPathActive('/hvac') ? '#50BCC0' : '#95A4B6', textDecoration: 'none' }}>
                            <i className="fas fa-thermometer-half" style={{ marginRight: '10px' }}></i>
                            HVAC
                        </Link>
                    </li>
                    <li className={`nav-item ${isPathActive('/appliances') ? 'active' : ''}`} style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                        <Link to="/appliances" style={{ color: isPathActive('/appliances') ? '#50BCC0' : '#95A4B6', textDecoration: 'none' }}>
                            <i className="fas fa-blender" style={{ marginRight: '10px' }}></i>
                            Appliances
                        </Link>
                    </li>
                    <li className={`nav-item ${isPathActive('/energy') ? 'active' : ''}`} style={{ margin: '0.5rem 0', padding: '0.5rem' }}>
                        <Link to="/energy" style={{ color: isPathActive('/energy') ? '#50BCC0' : '#95A4B6', textDecoration: 'none' }}>
                            <i className="fas fa-bolt" style={{ marginRight: '10px' }}></i>
                            Energy
                        </Link>
                    </li>
                </ul>
            </nav>
        </aside>
    );
};

export default Sidebar;
