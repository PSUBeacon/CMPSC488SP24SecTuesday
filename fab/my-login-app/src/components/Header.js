import React, {useState, useEffect} from 'react';
import {useNavigate} from 'react-router-dom';
import logoImage from '../img/logo.webp';
import accountIcon from '../img/account.png';
import settingsIcon from '../img/settings.png';
import menuIcon from '../img/menu.png';

const Header = ({accountType}) => {
    const [date, setDate] = useState(new Date());
    const [isNavVisible, setIsNavVisible] = useState(false);
    const [isAccountPopupVisible, setIsAccountPopupVisible] = useState(false);
    const navigate = useNavigate();
    const fname = sessionStorage.getItem('lastname')
    const lname = sessionStorage.getItem('firstname')

    // Toggles the navigation menu's visibility
    const toggleNav = () => {
        setIsNavVisible(!isNavVisible);
    };

    // Navigate to the settings page
    const goToSettings = () => {
        navigate('/settings');
    };

    // Handles the user sign-out process
    const signOut = () => {
        sessionStorage.removeItem('token');
        setIsAccountPopupVisible(false); // Close the account popup
        navigate('/'); // Navigate to home or sign-in page
    };

    // Toggles the account information popup's visibility
    const toggleAccountPopup = () => {
        setIsAccountPopupVisible(!isAccountPopupVisible);
    };

    // Set up a timer to update the date every minute
    useEffect(() => {
        const timerId = setInterval(() => setDate(new Date()), 60000);
        return () => clearInterval(timerId);
    }, []);

    // Format date and time
    const formattedDate = date.toLocaleDateString('en-US', {month: 'long', day: 'numeric', year: 'numeric'});
    const formattedTime = date.toLocaleTimeString('en-US', {hour: '2-digit', minute: '2-digit'});

    return (
        <nav className="topNav" style={{backgroundColor: '#081624', color: 'white', padding: '0.5rem 1rem'}}>
            <div style={{display: 'flex', justifyContent: 'space-between', alignItems: 'center'}}>
                <div style={{display: 'flex', alignItems: 'center'}}>
                    <img src={menuIcon} alt="Menu" onClick={toggleNav} className="hamburger-menu"/>
                    <img src={logoImage} alt="Logo" style={{marginRight: '10px'}} id='circle'/>
                    <span id='menuText2'>Beacon</span>
                </div>
                <div>
                    <span id='menuText'>{formattedDate}</span>
                </div>
                <div>
                    <span id='menuText'>{formattedTime}</span>
                </div>
                <div style={{position: 'relative'}}>
                    <img src={settingsIcon} alt="Settings" style={{marginRight: '10px'}} id="menuIcon"
                         onClick={goToSettings}/>
                    <button onClick={toggleAccountPopup}
                            style={{background: 'none', border: 'none', padding: 0, cursor: 'pointer'}}>
                        <img src={accountIcon} alt="account" style={{marginRight: '10px'}} id="menuIcon2"/>
                    </button>
                    {isAccountPopupVisible && (
                        <div className="accountPop" style={{
                            position: 'absolute',
                            top: '100%',
                            right: '0',
                            backgroundColor: '#08192B',
                            padding: '20px',
                            zIndex: 100,
                            color: 'white',
                            borderRadius: '2px',
                        }}>
                            <p>{fname} {lname}</p> {/* Dynamically replace with actual user name */}
                            {accountType && <p>{accountType}</p>}
                            <button onClick={signOut} className="signout">Sign Out</button>
                        </div>
                    )}
                </div>
            </div>
        </nav>
    );
};

export default Header;
