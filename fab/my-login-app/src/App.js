import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LoginForm from './LoginForm';
import Dashboard from './Dashboard';
import SignUp from './SignUp';
import Security from './Security'; // Import Security component
import Lighting from './Lighting'; // Import Lighting component
import Preferences from './Preferences'; // Import Preferences component
import HVAC from './HVAC'; // Import HVAC component
import Appliances from './Appliances'; // Import Appliances component
import Energy from './Energy'; // Import Energy component
import Settings from './Settings'; // Import your settings page component

import 'bootstrap/dist/css/bootstrap.min.css';
import './login.css';


function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LoginForm />} />
        <Route path="/dashboard" element={<Dashboard />}>
        </Route>
        <Route path="/settings" element={<Settings />} />
        <Route path="/security" element={<Security />} />
        <Route path="/lighting" element={<Lighting />} />
        <Route path="/preferences" element={<Preferences />} />
        <Route path="/hvac" element={<HVAC />} />
        <Route path="/appliances" element={<Appliances />} />
        <Route path="/energy" element={<Energy />} />
        <Route path="/signup" element={<SignUp />} />
        
      </Routes>
    </Router>
  );
}

export default App;
