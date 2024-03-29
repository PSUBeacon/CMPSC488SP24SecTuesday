import React, { useState, useEffect } from 'react';
import logoImage from './logo.webp'; // Adjust the path as needed
import './Settings.css'; // Adjust the path as necessary based on your file structure
import { Link } from 'react-router-dom'; // Import Link component
import accountIcon from './account.png'
import menuIcon from './menu.png'
import { Table } from 'react-bootstrap';
import { useNavigate } from 'react-router-dom'; // Import useHistory for navigation


const SettingsPage = () => {

  

  const [selectedNav, setSelectedNav] = useState(); // State variable for selected navigation item
  const navigate = useNavigate(); // Hook to enable redirection
  const handleSignOut = () => {
    // Add your sign-out logic here if necessary, like clearing localStorage or cookies
    // localStorage.removeItem('userToken');
  
    // Redirect to the login page
    window.location.href = '/'; // Replace '/login' with your actual login route
  };

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

  const [showAddUserForm, setShowAddUserForm] = useState(false);
  const [newUserFirstName, setNewUserFirstName] = useState('');
  const [newUserLastName, setNewUserLastName] = useState('');
  const [newUserRole, setNewUserRole] = useState('User');
  const [newUserEmail, setNewUserEmail] = useState('');
  const [newUserPassword, setNewUserPassword] = useState('');



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

    const addUserFormHandler = (event) => {
      event.preventDefault();
      const newUser = {
        id: users.length + 1,
        name: `${newUserFirstName} ${newUserLastName}`,
        role: newUserRole,
        email: newUserEmail,
        password: newUserPassword, // Ensure you hash the password in a real app!
      };
      setUsers([...users, newUser]);
      setShowAddUserForm(false); // Hide the form after adding the user
    
      // Reset form fields
      setNewUserFirstName('');
      setNewUserLastName('');
      setNewUserEmail('');
      setNewUserPassword('');
    };

 // Example account data (replace with actual data as necessary)
 const [accountInfo, setAccountInfo] = useState({
  firstName: 'John',
  lastName: 'Doe',
  accountType: 'Owner',
  profilePicture: '/accountIcon.jpg' // Replace with actual path
});

// Place this array outside of your component if it doesn't change, or in state if it does
const notifications = [
  { Appliance: 'Microwave', DateandTime: '00:00 on mm/dd/yyyy', Status: 'ON/OFF' },
  { Device: 'Oven', DateandTime: '00:00 on mm/dd/yyyy', Status: 'ON/OFF'},
  { Device: 'Fridge', DateandTime: '00:00 on mm/dd/yyyy', Status: 'ON/OFF'},
  { Device: 'Freezer', DateandTime: '00:00 on mm/dd/yyyy', Status: 'ON/OFF' },
  { Device: 'Toaster', DateandTime: '00:00 on mm/dd/yyyy', Status: 'ON/OFF' },
  { Device: 'Dishwasher', DateandTime: '00:00 on mm/dd/yyyy', Status: 'ON/OFF'},
];

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
        <div style={{ width: '100%', textAlign: 'center', padding: '20px', display: 'flex', flexDirection: 'column', alignItems: 'center', position: 'relative', zIndex: 1 }}>
          <h3 style={{ color: '#50BCC0', marginBottom: '20px' }}>Manage Users</h3>
          <button onClick={() => setShowAddUserForm(!showAddUserForm)} style={{ margin: '20px', padding: '10px', backgroundColor: '#50BCC0', color: 'white', border: 'none', borderRadius: '5px', zIndex: 2 }}>Add User</button>
          {showAddUserForm && (
            <form onSubmit={addUserFormHandler} style={{ padding: '20px', display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
              <input
                type="text"
                placeholder="First Name"
                value={newUserFirstName}
                onChange={(e) => setNewUserFirstName(e.target.value)}
                required
                style={{ margin: '5px' }}
              />
              <input
                type="text"
                placeholder="Last Name"
                value={newUserLastName}
                onChange={(e) => setNewUserLastName(e.target.value)}
                required
                style={{ margin: '5px' }}
              />
              <input
                type="email"
                placeholder="Email"
                value={newUserEmail}
                onChange={(e) => setNewUserEmail(e.target.value)}
                required
                style={{ margin: '5px' }}
              />
              <input
                type="password"
                placeholder="Password"
                value={newUserPassword}
                onChange={(e) => setNewUserPassword(e.target.value)}
                required
                style={{ margin: '5px' }}
              />
              <select
                value={newUserRole}
                onChange={(e) => setNewUserRole(e.target.value)}
                required
                style={{ margin: '5px' }}
              >
                <option value="Admin">Admin</option>
                <option value="User">User</option>
                <option value="Child">Child</option>
              </select>
              <button type="submit" style={{ margin: '5px', padding: '10px', backgroundColor: '#50BCC0', color: 'white', border: 'none', borderRadius: '5px' }}>Add</button>
            </form>
          )}
          <div style={{ width: '100%', maxWidth: '400px', overflow: 'auto', zIndex: 2 }}>
            {users.map(user => (
              <div key={user.id} style={{ margin: '10px', padding: '10px', backgroundColor: '#081624', borderRadius: '10px', display: 'flex', justifyContent: 'space-between', alignItems: 'center', zIndex: 2 }}>
                <div style={{ display: 'flex', alignItems: 'center' }}>
                  <span style={{ fontWeight: 'bold', color: 'white', marginRight: '10px' }}>{user.name}</span>
                  <span style={{ color: '#95A4B6' }}>({user.role})</span>
                </div>
                <div style={{ display: 'flex', alignItems: 'center' }}>
                  <select
                    value={user.role}
                    onChange={(e) => changeUserRoleHandler(user.id, e.target.value)}
                    style={{ margin: '0 10px', padding: '5px', borderRadius: '5px', border: '1px solid #50BCC0', backgroundColor: 'transparent', color: 'white', zIndex: 3 }}
                  >
                    <option value="Admin">Admin</option>
                    <option value="User">User</option>
                    <option value="Child">Child</option>
                  </select>
                  <button onClick={() => removeUserHandler(user.id)} style={{ backgroundColor: '#50BCC0', color: 'white', border: 'none', padding: '5px 10px', borderRadius: '5px', zIndex: 3 }}>Remove</button>
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
            <div style={{ margin: '10px', borderRadius: '10px', backgroundColor: '#081624', padding: '20px', minWidth: '200px' }}>
              <div style={{ marginBottom: '10px' }}>
                <label style={{ color: '#50BCC0' }}>First Name</label>
                <input
                  type="text"
                  name="firstName"
                  value={accountInfo.firstName}
                  onChange={handleAccountInfoChange}
                  style={{ backgroundColor: 'transparent', color: 'white', border: 'none', borderBottom: '1px solid #50BCC0', width: '100%', textAlign: 'center' }}
                />
              </div>
              <div style={{ marginBottom: '10px' }}>
                <label style={{ color: '#50BCC0' }}>Last Name</label>
                <input
                  type="text"
                  name="lastName"
                  value={accountInfo.lastName}
                  onChange={handleAccountInfoChange}
                  style={{ backgroundColor: 'transparent', color: 'white', border: 'none', borderBottom: '1px solid #50BCC0', width: '100%', textAlign: 'center' }}
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
                  style={{ backgroundColor: 'transparent', color: 'white', border: 'none', borderBottom: '1px solid #50BCC0', width: '100%', textAlign: 'center' }}
                />
              </div>
            </div>
          </div>
          {/* Sign-out button */}
        <button
          onClick={() => navigate('/login')} // Replace '/login' with your actual login route
          style={{
            marginTop: '20px',
            padding: '10px',
            backgroundColor: '#50BCC0',
            color: 'white',
            border: 'none',
            borderRadius: '5px',
            cursor: 'pointer'
          }}
        >
          Sign Out
        </button>
        </div>
      );

  case 'Notification Settings':
    return (
      <main style={{ flex: '1', padding: '1rem', display: 'flex', flexDirection: 'column', alignItems: 'center', backgroundColor: '#0E2237'}}>
        <h2 style={{ color: 'white' }}>Notifications</h2>
        {/* The UI library's Table component must be imported at the top of your file */}
        <Table striped bordered hover variant="dark" style={{ marginTop: '20px', backgroundColor: "#173350" }}>
          <thead>
            <tr>
              <th>Appliance</th>
              <th>Date and Time</th>
              <th>Status</th>
            </tr>
          </thead>
          <tbody>
            {notifications.map((notification, index) => (
              <tr key={index}>
                <td>{notification.Appliance || notification.Device}</td>
                <td>{notification.DateandTime}</td>
                <td>{notification.Status}</td>
              </tr>
            ))}
          </tbody>
        </Table>
      </main>
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
              <li className="settings-nav-item" onClick={() => setSelectedNav('Notification Settings')}>Notifications</li>
              <li className="settings-nav-item" onClick={() => setSelectedNav('Account Settings')}>Account Settings</li>
              <li className="settings-nav-item" >
<Link to="/" style={{ color: 'inherit', textDecoration: 'none' }}>Logout</Link> {/* Link to the dashboard */}
</li>
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