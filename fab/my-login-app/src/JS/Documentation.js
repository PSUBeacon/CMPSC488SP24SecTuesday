import React, {useState, useEffect} from 'react';

import {useNavigate} from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css'; // Ensure Bootstrap CSS is imported to use its grid system and components
import 'bootstrap/dist/js/bootstrap.bundle.min';

import 'bootstrap/dist/css/bootstrap.min.css';
import Header from "../components/Header";
import Sidebar from "../components/Sidebar";

// Define the Dashboard component using a functional component pattern
const Documentation = () => {
    const [accountType, setAccountType] = useState('');
    const [isNavVisible, setIsNavVisible] = useState(false);
    // This is the JSX return statement where we lay out our component's HTML structure
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
                    
                    <div class = "col-12 col-lg-8">
                        
                        <h1 class = "text-center">How to use BEACON</h1>

                        <div class="accordion" id="DOC">

                            {/* User Manual Tab */}
                            <div class="accordion-item">
                                    <h1 class="accordion-header" id="heading-DOC:title">
                                        <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#body-DOC:title" aria-expanded="true" aria-controls="body-DOC:title">
                                            <b>User Manual</b>
                                        </button>
                                    </h1>
                                    <div id="body-DOC:title" class="accordion-collapse collapse show" aria-labelledby="heading-DOC:title">
                                        <div class="accordion-body">
                                            <p>
                                                This is a user manual for the BEACON Smart Home System. 
                                                Printable version [Link Here]
                                            </p>
                                        </div>
                                    </div>
                            </div>

                            {/* Sign-In Tab */}
                            <div class="accordion-item">
                                <h2 class="accordion-header" id="heading-DOC:sign-in">
                                    <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#body-DOC:sign-in" aria-expanded="true" aria-controls="body-DOC:sign-in">
                                        <b>Sign-In</b>
                                    </button>
                                </h2>
                                <div id="body-DOC:sign-in" class="accordion-collapse collapse show" aria-labelledby="heading-DOC:sign-in">
                                    <div class="accordion-body">
                                        <p>
                                            If you already have an account, type in your username and password and then hit
                                            the sign in button. This will take you to the Dashboard page. If you do not
                                            have an account hit the "Sign Up" button next to the Don't Have an Account link
                                            which will take you to the sign up page.
                                        </p>
                                    </div>
                                </div>
                            </div>

                            {/* Sign-Up Tab */}
                            <div class="accordion-item">
                                <h2 class="accordion-header" id="heading-DOC:sign-up">
                                    <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#body-DOC:sign-up" aria-expanded="true" aria-controls="body-DOC:sign-up">
                                        <b>Sign-Up</b>
                                    </button>
                                </h2>
                                <div id="body-DOC:sign-up" class="accordion-collapse collapse show" aria-labelledby="heading-DOC:sign-up">
                                    <div class="accordion-body">
                                        <p>
                                            On the Sign Up page fill out all the fields with the correct information. Once
                                            you have typed in everything click "Sign Up". Once you have signed up, you can
                                            go back to the sign in page and log in with your newly created account
                                            credentials.
                                        </p>
                                    </div>
                                </div>
                            </div>

                            {/* Dashboard Tab */}
                            <div class="accordion-item">
                                <h2 class="accordion-header" id="heading-DOC:dash">
                                    <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#body-DOC:dash" aria-expanded="true" aria-controls="body-DOC:dash">
                                        <b>Dashboard</b>
                                    </button>
                                </h2>
                                <div id="body-DOC:dash" class="accordion-collapse collapse show" aria-labelledby="heading-DOC:dash">
                                    <div class="accordion-body">
                                    All Users
                                    <ul>
                                        <li>
                                            To view the security footage go to the overview page and click on "R1",
                                            "R2", and "R3" tabs to view the different rooms.
                                        </li>
                                        <li>
                                            To view the backdoor and frontdoor locks and its status go to the
                                            overview page. To view more about the locks, click the "See More"
                                            button.
                                        </li>
                                        <li>
                                            To view the light and its status go to the lights section on the
                                            overview page. To view the other lights for the other rooms click the
                                            "See More" button to the left.
                                        </li>
                                        <li>
                                            To view status by units, go to the overview page and view it at the
                                            bottom left of the page. The power, temperature, humidity, and security
                                            is listed out. These values cannot be changed and is only for viewing.
                                        </li>
                                        <li>
                                            To view the appliances, go to the overview page and view the appliances
                                            section at the bottom right of the page. To see more appliances and
                                            their statuses click the "See More" button you see to the left.
                                        </li>
                                    </ul>
                                    </div>
                                </div>
                            </div>

                            {/* Lighting Tab */}
                            <div class="accordion-item">
                                <h2 class="accordion-header" id="heading-DOC:light">
                                    <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#body-DOC:light" aria-expanded="true" aria-controls="body-DOC:light">
                                        <b>Lighting</b>
                                    </button>
                                </h2>
                                <div id="body-DOC:light" class="accordion-collapse collapse show" aria-labelledby="heading-DOC:light">
                                    <div class="accordion-body">
                                        All Users
                                        <ul>
                                            <li>To view the lights go to the lighting page.</li>
                                            <li>To view what lights are in the room, click on the room of your choice.</li>
                                            <li>Click on the drop down menu, select the room, and view the lights within it.</li>
                                            <li>To turn a light on or off, click on a room from the displayed list list and select the light from the drop down. Then click on the turn on/off button next
                                                to the light you want to change.
                                            </li>
                                        </ul>
                                        Admin Only
                                        <ul>
                                            <li>To add a light, fill out the form with the "Room Name" and the "Light Name" and then click the "Add Light" button.</li>
                                        </ul>
                                    </div>
                                </div>
                            </div>

                            {/* Security Tab */}
                            <div class="accordion-item">
                                <h2 class="accordion-header" id="heading-DOC:security">
                                    <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#body-DOC:security" aria-expanded="true" aria-controls="body-DOC:security">
                                        <b>Security</b>
                                    </button>
                                </h2>
                                <div id="body-DOC:security" class="accordion-collapse collapse show" aria-labelledby="heading-DOC:security">
                                    <div class="accordion-body">
                                        All Users
                                        <ul>
                                            <li>To view the locks go to the security page. You can view the current status of the front and back door lock.</li>
                                        </ul>
                                        Admin Only
                                        <ul>
                                            <li>To arm or disarm the locks type in the security code.</li>
                                        </ul>
                                    </div>
                                </div>
                            </div>

                            {/* Energy Tab */}
                            <div class="accordion-item">
                                <h2 class="accordion-header" id="heading-DOC:energy">
                                    <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#body-DOC:energy" aria-expanded="true" aria-controls="body-DOC:energy">
                                        <b>Energy</b>
                                    </button>
                                </h2>
                                <div id="body-DOC:energy" class="accordion-collapse collapse show" aria-labelledby="heading-DOC:energy">
                                    <div class="accordion-body">
                                        All Users
                                        <ul>
                                            <li>
                                                To view the status and energy consumption of devices/appliances, navigate to
                                                the energy page by clicking "Energy" on the side navigation menu.
                                            </li>
                                        </ul>
                                        Admin
                                        <ul>
                                            <li>
                                                To turn on or off an appliance/device, navigate to the desired table and choose
                                                "Turn On/Off".
                                            </li>
                                        </ul>
                                    </div>
                                </div>
                            </div>

                            {/* HVAC Tab */}
                            <div class="accordion-item">
                                <h2 class="accordion-header" id="heading-DOC:hvac">
                                    <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#body-DOC:hvac" aria-expanded="true" aria-controls="body-DOC:hvac">
                                        <b>HVAC</b>
                                    </button>
                                </h2>
                                <div id="body-DOC:hvac" class="accordion-collapse collapse show" aria-labelledby="heading-DOC:hvac">
                                    <div class="accordion-body">
                                        All Users
                                        <ul>
                                            <li>
                                                To view the temperature and current HVAC status of the home go to the HVAC page
                                                by clicking on the tab called HVAC on the side nav bar.
                                            </li>
                                        </ul>
                                        Admin
                                        <ul>
                                            <li>
                                                Mode: To change HVAC mode, choose between "Heat", "Cool", or "Off".
                                            </li>
                                            <li>
                                                Temperature: To set the temperature for the desired floor, use the "+" and "-"
                                                respectively to increase and decrease the temperature value.
                                            </li>
                                            <li>
                                                Fan Speed: To change the fan speed, choose between "High", "Medium", or "Low".
                                            </li>
                                        </ul>
                                    </div>
                                </div>
                            </div>

                            {/* Appliances Tab */}
                            <div class="accordion-item">
                                <h2 class="accordion-header" id="heading-DOC:appliances">
                                    <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#body-DOC:appliances" aria-expanded="true" aria-controls="body-DOC:appliances">
                                        <b>Appliances</b>
                                    </button>
                                </h2>
                                <div id="body-DOC:appliances" class="accordion-collapse collapse show" aria-labelledby="heading-DOC:appliances">
                                    <div class="accordion-body">
                                        All Users
                                        <ul>
                                            <li>
                                                To view the the full list of appliances and their status, navigate to the
                                                appliances page by clicking "Appliances" on the side navigation menu.
                                            </li>
                                        </ul>
                                    </div>
                                </div>
                            </div>
                            
                            {/* Settings Tab */}
                            <div class="accordion-item">
                                <h2 class="accordion-header" id="heading-DOC:settings">
                                    <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#body-DOC:settings" aria-expanded="true" aria-controls="body-DOC:settings">
                                        <b>Settings</b>
                                    </button>
                                </h2>
                                <div id="body-DOC:settings" class="accordion-collapse collapse show" aria-labelledby="heading-DOC:settings">
                                    <div class="accordion-body">
                                        <ul>
                                            <li>
                                                On the settings page you will be able to view information such as the users,
                                                notifications, and your own account information. You will also be given the
                                                option to sign out of the web page. To navigate through the settings page, use
                                                the side navigation on the left to switch tabs to the desired content.
                                            </li>
                                        </ul>
                                        All Users
                                        <ul>
                                            <li>
                                                Dashboard: Returns to the dashboard home screen.
                                            </li>
                                            <li>
                                                Manage Users: To view all of the users on the website.
                                            </li>
                                            <li>
                                                Notifications: To view the most recent activity.
                                            </li>
                                            <li>
                                                Account Settings: Allows the users to view their account information.
                                            </li>
                                            <li>
                                                Sign Out: Signs out and returns back to the log in page.
                                            </li>
                                        </ul>
                                        Admin
                                        <ul>
                                            <li>
                                                Manage Users: Allows the admin to view and manage all of the users on the website (add, remove, and change permissions).
                                                <ul>
                                                    <li>
                                                        To add a new user, choose "Add User" and fill in the required
                                                        information followed by the "Add" button.
                                                    </li>
                                                    <li>
                                                        To remove a user, click the "Remove" button next to the user you wish to remove.
                                                    </li>
                                                    <li>
                                                        To change the permissions of a user, click on the drop down menu and
                                                        choose between the 2 roles (admin and user).
                                                    </li>
                                                </ul>
                                            </li>
                                            <li>
                                                Notifications: To view the most recent activity.
                                            </li>
                                        </ul>
                                    </div>
                                </div>
                            </div>

                        </div>
        
                    </div>
                </main>
            </div>
            
        </div>
        
    );
    
};


export default Documentation;