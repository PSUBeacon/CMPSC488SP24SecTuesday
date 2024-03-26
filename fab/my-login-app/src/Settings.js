import React, { useState, useEffect } from 'react';
import logoImage from './logo.webp'; // Adjust the path as needed
import './Settings.css'; // Adjust the path as necessary based on your file structure
import { Link } from 'react-router-dom'; // Import Link component
import accountIcon from './account.png'
import menuIcon from './menu.png'



const SettingsPage = () => {
  // Initialize theme state from local storage if available, otherwise default to 'dark'
  const [theme, setTheme] = useState(localStorage.getItem('theme') || 'dark');
  const [selectedNav, setSelectedNav] = useState('Theme'); // State variable for selected navigation item

  // Function to set the theme and save it to local storage
  const changeTheme = (newTheme) => {
    setTheme(newTheme);
    localStorage.setItem('theme', newTheme); // Save new theme to local storage
  };

  // Styles that change with the theme
  const topNavStyle = {
    backgroundColor: theme === 'dark' ? '#081624' : '#F0F1F3',
    color: theme === 'dark' ? 'white' : 'black',
    padding: '0.5rem 1rem',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between',
  };

  const sideNavStyle = {
    backgroundColor: theme === 'dark' ? '#173350' : '#F0F1F3',
    color: theme === 'dark' ? 'white' : 'black',
    width: '250px',
    padding: '1rem',
  };

  // useEffect to load theme from local storage on mount
  useEffect(() => {
    const storedTheme = localStorage.getItem('theme');
    if (storedTheme) {
      setTheme(storedTheme);
    }
  }, []);



    // Example user data (replace with actual data as necessary)
    const [users, setUsers] = useState([
      { id: 1, name: 'Alice Smith', role: 'Admin' },
      { id: 2, name: 'Bob Johnson', role: 'User' },
      { id: 3, name: 'Charlie Lee', role: 'Child' },
    ]);
  
    // Handler to add a new user (this will need to be more complex in a real app)
    const addUserHandler = () => {
      // Simple example: Add a default new user (you would have more complex logic in a real app)
      const newUser = { id: users.length + 1, name: `New User ${users.length + 1}`, role: 'User' };
      setUsers([...users, newUser]);
    };
  
    // Handler to remove a user by id
    const removeUserHandler = (userId) => {
      setUsers(users.filter(user => user.id !== userId));
    };
  
    // Handler to change a user's role
    const changeUserRoleHandler = (userId, newRole) => {
      setUsers(users.map(user => user.id === userId ? { ...user, role: newRole } : user));
    };

 // Example account data (replace with actual data as necessary)
 const [accountInfo, setAccountInfo] = useState({
  firstName: 'John',
  lastName: 'Doe',
  accountType: 'Owner',
  profilePicture: '/accountIcon.jpg' // Replace with actual path
});

// Handler for changing account info (you'll need a more complex handler for a real app)
const handleAccountInfoChange = (e) => {
  const { name, value } = e.target;
  setAccountInfo(prevState => ({
    ...prevState,
    [name]: value
  }));
};


  // Render content based on the selected navigation item
  const renderContent = () => {
    switch (selectedNav) {
      case 'Manage Users':
    return (
      <div style={{ width: '100%', textAlign: 'center' }}>
        <h3>Manage Users</h3>
        <button onClick={addUserHandler} style={{ margin: '20px', padding: '10px' }}>Add User</button>
        <div>
          {users.map(user => (
            <div key={user.id} style={{ margin: '10px', padding: '10px', backgroundColor: theme === 'dark' ? '#081624' : '#D3D3D3', borderRadius: '5px' }}>
              <span>{user.name} - {user.role}</span>
              <select
                value={user.role}
                onChange={(e) => changeUserRoleHandler(user.id, e.target.value)}
                style={{ margin: '0 10px' }}
              >
                <option value="Admin">Admin</option>
                <option value="User">User</option>
                <option value="Child">Child</option>
              </select>
              <button onClick={() => removeUserHandler(user.id)}>Remove</button>
            </div>
          ))}
        </div>
      </div>
    );

      // ... other cases remain unchanged
      case 'Theme':
        return (
          <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', height: '100%', width: '100%' }}>
            <label htmlFor="theme-select" style={{ fontSize: '1.5rem', marginBottom: '10px' }}>Choose a theme:</label>
            <select
              id="theme-select"
              value={theme}
              onChange={(e) => changeTheme(e.target.value)}
              style={{ width: '80%', height: '50px', fontSize: '1.2rem', textAlign: 'center', margin: '20px 0' }}
            >
              <option value="dark">Dark Mode</option>
              <option value="light">Light Mode</option>
            </select>
          </div>
        );
      case 'Notification Settings':
        return <div>Notification Settings Content Here</div>;

        // Update the 'Account Settings' case in your renderContent function
  case 'Account Settings':
    return (
      <div style={{ width: '100%', textAlign: 'center', padding: '20px' }}>
        <h3>Account Settings</h3>
        <div style={{ marginBottom: '20px' }}>
          <img
            src={accountInfo.profilePicture}
            alt="Profile"
            style={{ width: '100px', height: '100px', borderRadius: '50%' }}
          />
        </div>
        <div style={{ marginBottom: '10px' }}>
          <label>First Name</label>
          <input
            type="text"
            name="firstName"
            value={accountInfo.firstName}
            onChange={handleAccountInfoChange}
          />
        </div>
        <div style={{ marginBottom: '10px' }}>
          <label>Last Name</label>
          <input
            type="text"
            name="lastName"
            value={accountInfo.lastName}
            onChange={handleAccountInfoChange}
          />
        </div>
        <div style={{ marginBottom: '10px' }}>
          <label>Account Type</label>
          <input
            type="text"
            name="accountType"
            value={accountInfo.accountType}
            onChange={handleAccountInfoChange}
            disabled // You might want this to be uneditable, or be a dropdown if editable
          />
        </div>
      </div>
    );
      

    }
  };

  return (
    <div style={{ display: 'flex', minHeight: '100vh', flexDirection: 'column' }}>
      <nav style={topNavStyle}>
        <img src={logoImage} alt="Logo" id='circle' style={{ height: '50px' }} />
        <span style={{ fontSize: '1.5rem', fontWeight: 'bold' }}>Settings</span>
      </nav>

      <div style={{ display: 'flex', flex: '1' }}>
        <aside style={sideNavStyle}>
          <nav>
            <ul style={{ listStyleType: 'none', padding: 0 }}>
            <li className="settings-nav-item" onClick={() => setSelectedNav('Dashboard')}>
<Link to="/dashboard" style={{ color: 'inherit', textDecoration: 'none' }}>Dashboard</Link> {/* Link to the dashboard */}
</li>
              <li className="settings-nav-item" onClick={() => setSelectedNav('Manage Users')}>Manage Users</li>
              <li className="settings-nav-item" onClick={() => setSelectedNav('Theme')}>Theme</li>
              <li className="settings-nav-item" onClick={() => setSelectedNav('Notification Settings')}>Notification Settings</li>
              <li className="settings-nav-item" onClick={() => setSelectedNav('Account Settings')}>Account Settings</li>
            </ul>
          </nav>
        </aside>

        <div style={{ flex: '1', padding: '1rem', backgroundColor: theme === 'dark' ? '#0E2237' : '#FFFFFF', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
          {renderContent()}
        </div>
      </div>
    </div>
  );
};

export default SettingsPage;