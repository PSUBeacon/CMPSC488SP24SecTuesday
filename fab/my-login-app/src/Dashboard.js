import React from 'react';

// Define the Dashboard component using a functional component pattern
const Dashboard = () => {
  // This is the JSX return statement where we layout our component's HTML structure
    const welcomeMessage = localStorage.getItem('welcomeMessage'); // Retrieve the message
  return (
    <div>
      <h2>Dashboard</h2>
      <p>Welcome to your dashboard!</p>
        {welcomeMessage && <p>{welcomeMessage}</p>} {/* Display the message */}
      {/* You can add more content here such as user information, links, or any other components relevant to your application */}
    </div>
  );
};

// Export the Dashboard component so it can be used in other parts of our application
export default Dashboard;