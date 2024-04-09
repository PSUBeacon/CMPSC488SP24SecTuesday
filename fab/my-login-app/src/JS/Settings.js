import React, {useState, useEffect} from 'react';
import logoImage from '../img/logo.webp';
import '../CSS/Settings.css'; // Adjust the path as necessary based on your file structure
import {Link} from 'react-router-dom'; // Import Link component
import accountIcon from '../img/account.png';
import {useNavigate} from 'react-router-dom'; // Import useHistory for navigation

const SettingsPage = () => {
    document.title = 'BEACON | Settings';
    const [error, setError] = useState('');
    const [selectedNav, setSelectedNav] = useState();
    const navigate = useNavigate();
    const handleSignOut = () => {
        sessionStorage.removeItem('token');
        window.location.href = '/';
    };

    const token = sessionStorage.getItem('token');
    const userFirstName = sessionStorage.getItem('FirstName');
    const userLastName = sessionStorage.getItem('LastName');
    const userAccountType = sessionStorage.getItem('Role');

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
    const [users, setUsers] = useState([]);


    const fetchUsers = async () => {
        try {
            const response = await fetch('http://localhost:8081/settings/GetUsers', {
                method: 'GET',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json',
                },
            });

            if (response.ok) {
                const data = await response.json();
                setUsers(data);
            } else {
                setError('Failed to fetch users');
            }
        } catch (error) {
            console.log('Fetch operation error:', error);
        }
    };

    useEffect(() => {
        if (userAccountType === 'admin') {
            fetchUsers();
        }
    }, [userAccountType]);

    // Handler to add a new user (this will need to be more complex in a real app)
    const addUserHandler = () => {
        // Simple example: Add a default new user (you would have more complex logic in a real app)
        const newUser = {id: users.length + 1, name: `New User ${users.length + 1}`, role: 'User'};
        setUsers([...users, newUser]);
    };

    // Handler to remove a user by id
    const removeUserHandler = (userId) => {
        setUsers(users.filter(user => user.id !== userId));
    };

    // Handler to change a user's role
    const changeUserRoleHandler = (userId, newRole) => {
        setUsers(users.map(user => user.id === userId ? {...user, role: newRole} : user));
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

    function setUserInfo(param) {

    }

    // Handler for changing account info (you'll need a more complex handler for a real app)
    const handleAccountInfoChange = (e) => {
        const {name, value} = e.target;
        setUserInfo(prevState => ({
            ...prevState,
            [name]: value
        }));
    };

    return (
        <div style={{display: 'flex', minHeight: '100vh', flexDirection: 'column'}}>
            <nav style={topNavStyle}>
                <img src={logoImage} alt="Logo" id='circle' style={{height: '50px'}}/>
                <span style={{fontSize: '1.5rem', fontWeight: 'bold'}}>Settings</span>
            </nav>

            <div style={{display: 'flex', flex: '1'}}>
                <aside style={sideNavStyle}>
                    <nav>
                        <ul style={{listStyleType: 'none', padding: 0}}>
                            <li className="settings-nav-item">
                                <Link to="/dashboard" style={{color: 'inherit', textDecoration: 'none'}}>
                                    Dashboard
                                </Link>
                            </li>
                            {userAccountType === 'admin' && (
                                <li className="settings-nav-item" onClick={() => setSelectedNav('Manage Users')}>
                                    Manage Users
                                </li>
                            )}
                            <li className="settings-nav-item" onClick={() => setSelectedNav('Account Info')}>
                                Account Info
                            </li>
                            <li className="settings-nav-item">
                                <button onClick={handleSignOut} style={{
                                    color: 'inherit',
                                    textDecoration: 'none',
                                    border: 'none',
                                    background: 'none',
                                    cursor: 'pointer'
                                }}>
                                    Sign Out
                                </button>
                            </li>
                        </ul>
                    </nav>
                </aside>

                <div style={{
                    flex: '1',
                    padding: '1rem',
                    position: 'relative',
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center'
                }}>
                    {selectedNav === 'Manage Users' && userAccountType === 'admin' && (
                        <div style={{
                            width: '100%',
                            textAlign: 'center',
                            padding: '20px',
                            display: 'flex',
                            flexDirection: 'column',
                            alignItems: 'center',
                            position: 'relative',
                            zIndex: 1
                        }}>
                            <h3 style={{color: '#50BCC0', marginBottom: '20px'}}>Manage Users</h3>
                            <button onClick={() => setShowAddUserForm(!showAddUserForm)} style={{
                                margin: '20px',
                                padding: '10px',
                                backgroundColor: '#50BCC0',
                                color: 'white',
                                border: 'none',
                                borderRadius: '5px',
                                zIndex: 2
                            }}>
                                Add User
                            </button>
                            {showAddUserForm && (
                                <form onSubmit={addUserFormHandler} style={{
                                    padding: '20px',
                                    display: 'flex',
                                    flexDirection: 'column',
                                    alignItems: 'center'
                                }}>
                                    <input
                                        type="text"
                                        placeholder="First Name"
                                        value={newUserFirstName}
                                        onChange={(e) => setNewUserFirstName(e.target.value)}
                                        required
                                        style={{margin: '5px'}}
                                    />
                                    <input
                                        type="text"
                                        placeholder="Last Name"
                                        value={newUserLastName}
                                        onChange={(e) => setNewUserLastName(e.target.value)}
                                        required
                                        style={{margin: '5px'}}
                                    />
                                    <input
                                        type="email"
                                        placeholder="Email"
                                        value={newUserEmail}
                                        onChange={(e) => setNewUserEmail(e.target.value)}
                                        required
                                        style={{margin: '5px'}}
                                    />
                                    <input
                                        type="password"
                                        placeholder="Password"
                                        value={newUserPassword}
                                        onChange={(e) => setNewUserPassword(e.target.value)}
                                        required
                                        style={{margin: '5px'}}
                                    />
                                    <select
                                        value={newUserRole}
                                        onChange={(e) => setNewUserRole(e.target.value)}
                                        required
                                        style={{margin: '5px'}}
                                    >
                                        <option value="admin">Admin</option>
                                        <option value="user">User</option>
                                        <option value="child">Child</option>
                                    </select>
                                    <button type="submit" style={{
                                        margin: '5px',
                                        padding: '10px',
                                        backgroundColor: '#50BCC0',
                                        color: 'black',
                                        border: 'none',
                                        borderRadius: '5px'
                                    }}>
                                        Add
                                    </button>
                                </form>
                            )}
                            <div style={{width: '100%', maxWidth: '400px', overflow: 'auto', zIndex: 2}}>
                                {users.map(user => (
                                    <div key={user.username} style={{
                                        margin: '10px',
                                        padding: '10px',
                                        backgroundColor: '#081624',
                                        borderRadius: '10px',
                                        display: 'flex',
                                        justifyContent: 'space-between',
                                        alignItems: 'center',
                                        zIndex: 2
                                    }}>
                                        <div style={{display: 'flex', alignItems: 'center'}}>
                        <span style={{
                            fontWeight: 'bold',
                            color: 'white',
                            marginRight: '10px'
                        }}>{user.firstName} {user.lastName}</span>
                                            <span style={{color: '#95A4B6'}}>({user.role})</span>
                                        </div>
                                        <div style={{display: 'flex', alignItems: 'center'}}>
                                            <select
                                                value={user.role}
                                                onChange={(e) => changeUserRoleHandler(user.id, e.target.value)}
                                                style={{
                                                    margin: '0 10px',
                                                    padding: '5px',
                                                    borderRadius: '5px',
                                                    border: '1px solid #50BCC0',
                                                    backgroundColor: '#081624',
                                                    color: 'white',
                                                    zIndex: 3
                                                }}
                                            >
                                                <option value="admin">admin</option>
                                                <option value="user">user</option>
                                                <option value="child">child</option>
                                            </select>
                                            <button onClick={() => removeUserHandler(user.id)} style={{
                                                backgroundColor: '#50BCC0',
                                                color: 'black',
                                                border: 'none',
                                                padding: '5px 10px',
                                                borderRadius: '5px',
                                                zIndex: 3
                                            }}>
                                                Remove
                                            </button>
                                        </div>
                                    </div>
                                ))}
                            </div>
                        </div>
                    )}
                    {selectedNav === 'Account Info' && (
                        <div style={{width: '100%', textAlign: 'center', padding: '20px'}}>
                            <h3 style={{marginBottom: '40px'}}>Account Info</h3>
                            <div style={{margin: '10px'}}>
                                <img
                                    src={accountIcon}
                                    alt="Profile"
                                    style={{
                                        width: '100px',
                                        height: '100px',
                                        borderRadius: '50%',
                                        border: '2px solid #50BCC0'
                                    }}
                                />
                            </div>
                            <div style={{
                                display: 'flex',
                                justifyContent: 'center',
                                alignItems: 'center',
                                flexWrap: 'wrap'
                            }}>
                                <div style={{
                                    margin: '10px',
                                    borderRadius: '10px',
                                    backgroundColor: '#081624',
                                    padding: '20px',
                                    minWidth: '200px'
                                }}>
                                    <div style={{marginBottom: '10px'}}>
                                        <label style={{color: '#50BCC0'}}>First Name</label>
                                        <input
                                            type="text"
                                            name="firstName"
                                            value={userFirstName}
                                            onChange={handleAccountInfoChange}
                                            style={{
                                                backgroundColor: 'transparent',
                                                color: 'white',
                                                border: 'none',
                                                borderBottom: '1px solid #50BCC0',
                                                width: '100%',
                                                textAlign: 'center'
                                            }}
                                        />
                                    </div>
                                    <div style={{marginBottom: '10px'}}>
                                        <label style={{color: '#50BCC0'}}>Last Name</label>
                                        <input
                                            type="text"
                                            name="lastName"
                                            value={userLastName}
                                            onChange={handleAccountInfoChange}
                                            style={{
                                                backgroundColor: 'transparent',
                                                color: 'white',
                                                border: 'none',
                                                borderBottom: '1px solid #50BCC0',
                                                width: '100%',
                                                textAlign: 'center'
                                            }}
                                        />
                                    </div>
                                    <div style={{marginBottom: '10px'}}>
                                        <label style={{color: '#50BCC0'}}>Account Type</label>
                                        <input
                                            type="text"
                                            name="accountType"
                                            value={userAccountType}
                                            onChange={handleAccountInfoChange}
                                            disabled
                                            style={{
                                                backgroundColor: 'transparent',
                                                color: 'white',
                                                border: 'none',
                                                borderBottom: '1px solid #50BCC0',
                                                width: '100%',
                                                textAlign: 'center'
                                            }}
                                        />
                                    </div>
                                </div>
                            </div>
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
};

export default SettingsPage;