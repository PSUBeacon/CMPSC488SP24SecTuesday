import React, { useState, useEffect } from 'react';
import logoImage from './logo.webp'; // Adjust the path as needed
import './Settings.css'; // Adjust the path as necessary based on your file structure
import { Link } from 'react-router-dom'; // Import Link component
import accountIcon from './account.png'
import menuIcon from './menu.png'



const SettingsPage = () => {

  const [selectedNav, setSelectedNav] = useState(); // State variable for selected navigation item



  // Styles that change with the theme
  const topNavStyle = {
    backgroundColor: '#081624',
    color: 'white',
    padding: '0.5rem 1rem',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between',
  };

  const sideNavStyle = {
    backgroundColor: '#173350',
    color: 'white',
    width: '250px',
    padding: '1rem',
  };

 



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
            <div style={{ width: '100%', textAlign: 'center', padding: '20px', display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
              <h3 style={{ color: '#50BCC0', marginBottom: '20px' }}>Manage Users</h3>
              <button onClick={addUserHandler} style={{ margin: '20px', padding: '10px', backgroundColor: '#50BCC0', color: 'white', border: 'none', borderRadius: '5px' }}>Add User</button>
              <div style={{ width: '100%', maxWidth: '400px', overflow: 'auto' }}>
                {users.map(user => (
                  <div key={user.id} style={{ margin: '10px', padding: '10px', backgroundColor: '#081624', borderRadius: '10px', display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                    <div>
                      <span style={{ fontWeight: 'bold', color: 'white', marginRight: '10px' }}>{user.name}</span>
                      <span style={{ color: '#95A4B6' }}>({user.role})</span>
                    </div>
                    <div>
                      <select
                        value={user.role}
                        onChange={(e) => changeUserRoleHandler(user.id, e.target.value)}
                        style={{ margin: '0 10px', padding: '5px', borderRadius: '5px', border: '1px solid #50BCC0', backgroundColor: 'transparent', color: 'white' }}
                      >
                        <option value="Admin">Admin</option>
                        <option value="User">User</option>
                        <option value="Child">Child</option>
                      </select>
                      <button onClick={() => removeUserHandler(user.id)} style={{ backgroundColor: '#50BCC0', color: 'white', border: 'none', padding: '5px 10px', borderRadius: '5px' }}>Remove</button>
                    </div>
                  </div>
                ))}
              </div>
            </div>
        
        );
      

     
    case 'Account Settings':
  return (
    <div style={{ width: '100%', textAlign: 'center', padding: '20px' }}>
      <h3 style={{ marginBottom: '40px' }}>Account Settings</h3>
       {/* Profile Picture */}
       <div style={{ margin: '10px' }}>
          <img
            src={accountIcon}
            alt="Profile"
            style={{ width: '100px', height: '100px', borderRadius: '50%', border: '2px solid #50BCC0' }}
          />
          <p style={{ margin: '5px', color: '#50BCC0' }}>Change Picture</p>
        </div>
      <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', flexWrap: 'wrap' }}>
        {/* Account Info */}
        <div style={{ margin: '10px', borderRadius: '10px', backgroundColor: '#081624', padding: '20px', minWidth: '200px' }}>
          <div style={{ marginBottom: '10px' }}>
            <label style={{ color: '#50BCC0' }}>First Name</label>
            <input
              type="text"
              name="firstName"
              value={accountInfo.firstName}
              onChange={handleAccountInfoChange}
              style={{ backgroundColor: 'transparent', color: 'white', border: 'none', borderBottom: '1px solid #50BCC0', width: '100%',  textAlign:'center' }}
            />
          </div>
          <div style={{ marginBottom: '10px' }}>
            <label style={{ color: '#50BCC0' }}>Last Name</label>
            <input
              type="text"
              name="lastName"
              value={accountInfo.lastName}
              onChange={handleAccountInfoChange}
              style={{ backgroundColor: 'transparent', color: 'white', border: 'none', borderBottom: '1px solid #50BCC0', width: '100%', textAlign:'center' }}
            />
          </div>
          <div style={{ marginBottom: '10px' }}>
            <label style={{ color: '#50BCC0' }}>Account Type</label>
            <input
              type="text"
              name="accountType"
              value={accountInfo.accountType}
              onChange={handleAccountInfoChange}
              disabled
              style={{ backgroundColor: 'transparent', color: 'white', border: 'none', borderBottom: '1px solid #50BCC0', width: '100%',  textAlign:'center' }}
            />
          </div>
        </div>
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
            <li className="settings-nav-item" >
<Link to="/dashboard" style={{ color: 'inherit', textDecoration: 'none' }}>Dashboard</Link> {/* Link to the dashboard */}
</li>
              <li className="settings-nav-item" onClick={() => setSelectedNav('Manage Users')}>Manage Users</li>
              <li className="settings-nav-item" onClick={() => setSelectedNav('Notification Settings')}>Notification Settings</li>
              <li className="settings-nav-item" onClick={() => setSelectedNav('Account Settings')}>Account Settings</li>
            </ul>
          </nav>
        </aside>

          {/* Main content area with a gradient background and translucent pattern overlay */}
        <div
          style={{
            flex: '1',
            padding: '1rem',
            backgroundImage: 'linear-gradient(to bottom, #0E2237, #081624)',
            position: 'relative',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
          }}
        >
          {/* Translucent pattern overlay */}
          <div
            style={{
              position: 'absolute',
              top: 0,
              left: 0,
              width: '100%',
              height: '100%',
              backgroundImage: 'url("https://www.transparenttextures.com/patterns/always-grey.png")',
              opacity: 0.3,
            }}
          ></div>
          {/* Render content based on selected navigation item */}
          {renderContent()}
        </div>
      </div>
    </div>
  );
};

export default SettingsPage;