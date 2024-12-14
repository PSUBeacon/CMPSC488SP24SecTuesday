import React, {useState, useEffect} from 'react';
import logoImage from '../img/logo.webp';
import '../CSS/Settings.css'; // Adjust the path as necessary based on your file structure
import {Link} from 'react-router-dom'; // Import Link component
import accountIcon from '../img/account.png';
import {useNavigate} from 'react-router-dom';
import axios from "axios"; // Import useHistory for navigation

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
    const userAccountType = sessionStorage.getItem('accountType');
    const [isLoading, setIsLoading] = useState(false);
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
    const [newUserUsername, setNewUserUsername] = useState('');
    const [newUserPassword, setNewUserPassword] = useState('');
    const [newuser, setNewuser] = useState([]);
    const [users, setUsers] = useState([]);

    useEffect(() => {
        const fetchUsers = async () => {
            try {
                const response = await fetch('https://beacon-cs2024.digital/api/settings/GetUsers', {
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

        if (userAccountType === 'admin') {
            setSelectedNav('Manage Users')
            fetchUsers().then(r => console.log('Users fetched successfully:', r)).catch(e => console.error('Fetch users error:', e));
        } else {
            setSelectedNav('Account Info')
        }
    }, [token, userAccountType]);

    // Handler to add a new user (this will need to be more complex in a real app)
    const addUserHandler = async () => {
        try {
            const serverUrl = 'https://beacon-cs2024.digital/api/signup';
            const response = await axios.post(serverUrl, {
                firstName: newUserFirstName,
                lastName: newUserLastName,
                username: newUserUsername,
                password: newUserPassword,
            });

            console.log('User added successfully:', response.data);
            // setNewuser([{
            //     firstName: newUserFirstName,
            //     lastName: newUserLastName,
            //     username: newUserUsername,
            //     password: newUserPassword,
            // }]);

            setUsers(prevUsers => [...prevUsers, {
                firstName: newUserFirstName,
                lastName: newUserLastName,
                username: newUserUsername,
                password: newUserPassword,
                role: newUserRole
            }]);

            setShowAddUserForm(false);
            setNewUserFirstName('');
            setNewUserLastName('');
            setNewUserUsername('');
            setNewUserPassword('');


            navigate('/settings')

        } catch (error) {
            if (error.response && error.response.data && error.response.data.error) {
                setError(error.response.data.error);
            } else {
                setError('Failed to add user. Username already exists in the system');
            }
            console.error('Add user error:', error.toJSON());
        }
    };

    // const deleteUserHandler = async (username) => {
    //     try {
    //         setIsLoading(true);
    //         const serverUrl = `https://beacon-cs2024.digital/api/users/${username}`;
    //         await axios.delete(serverUrl, {
    //             headers: {
    //                 'Authorization': `Bearer ${token}`,
    //             },
    //         });
    //         removeUserHandler(username);
    //     } catch (error) {
    //         if (error.response && error.response.data && error.response.data.error) {
    //             setError(error.response.data.error);
    //         } else {
    //             setError('Failed to delete user. Please try again later.');
    //         }
    //         console.error('Delete user error:', error);
    //     } finally{
    //         setIsLoading(false);
    //     }
    // };
    // Handler to remove a user by id
// Handler to remove a user by username
    const removeUserHandler = async (username) => {
        try {
            setIsLoading(true);
            const serverUrl = `https://beacon-cs2024.digital/api/users/${username}`;
            await axios.delete(serverUrl, {
                headers: {
                    'Authorization': `Bearer ${token}`,
                },
            });
            // Remove the user from the state without making another request
            setUsers((prevUsers) => prevUsers.filter((user) => user.username !== username));
        } catch (error) {
            if (error.response && error.response.data && error.response.data.error) {
                setError(error.response.data.error);
            } else {
                setError('Failed to delete user. Please try again later.');
            }
            console.error('Delete user error:', error);
        } finally {
            setIsLoading(false);
        }
    };

    const changeUserRoleHandler = async (username, newRole) => {
        try {
            const serverUrl = `https://beacon-cs2024.digital/api/users/${username}/role`;
            await axios.post(serverUrl, { role: newRole }, {
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json',
                },
            });
            setUsers((prevUsers) =>
                prevUsers.map((user) =>
                    user.username === username ? { ...user, role: newRole } : user
                )
            );
        } catch (error) {
            if (error.response && error.response.data && error.response.data.error) {
                setError(error.response.data.error);
            } else {
                setError('Failed to update user role. Please try again later.');
            }
            console.error('Change user role error:', error);
        }
    };

    const addUserFormHandler = (event) => {
        event.preventDefault();
        addUserHandler().then(r => console.log('User added successfully:', r)).catch(e => console.error('Add user error:', e));
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
        sessionStorage.setItem(name, value);
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
                                        type="text"
                                        placeholder="Username"
                                        value={newUserUsername}
                                        onChange={(e) => setNewUserUsername(e.target.value)}
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
                                        {/*<option value="admin">Admin</option>*/}
                                        <option value="user">User</option>
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
                                                onChange={(e) => changeUserRoleHandler(user.username, e.target.value)}
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
                                            </select>
                                            <button onClick={() => removeUserHandler(user.username)} disabled={isLoading} style={{
                                                backgroundColor: '#50BCC0',
                                                color: 'black',
                                                border: 'none',
                                                padding: '5px 10px',
                                                borderRadius: '5px',
                                                zIndex: 3,
                                                opacity: isLoading ? 0.5 : 1,
                                                cursor: isLoading ? 'not-allowed' : 'pointer'
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