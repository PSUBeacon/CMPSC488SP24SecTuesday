import React from 'react';
import ReactDOM from 'react-dom/client';
import './CSS/index.css';
import App from './JS/App';
import reportWebVitals from './JS/reportWebVitals';
import 'bootstrap/dist/css/bootstrap.min.css';
import './CSS/login.css';
import './CSS/Dashboard.css';
import './CSS/Appliances.css';
import './CSS/App.css'

import './CSS/Header.css'

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <React.StrictMode>
        <div className="App">
            <App/>
        </div>
    </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
