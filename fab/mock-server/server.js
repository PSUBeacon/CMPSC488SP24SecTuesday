const express = require('express');
const jwt = require('jsonwebtoken');
const cors = require('cors'); // Require the cors library
const app = express();
const PORT = 3001;

app.use(cors()); // Enable CORS for all routes
app.use(express.json());

const users = [
  { id: 1, username: 'user1', password: 'pass1' }, // Example user
];

const SECRET_KEY = 'your_secret_key'; // Keep secret keys safe and out of version control in real applications

app.post('/login', (req, res) => {
  const { username, password } = req.body; //This line extracts the username and password from the request body, which were sent in the login request.
  const user = users.find(u => u.username === username && u.password === password);

  if (user) {
    const token = jwt.sign({ userId: user.id }, SECRET_KEY, { expiresIn: '1h' }); //1 hour time limit
    res.json({ token }); //If the user is found and the token is successfully created, this line sends the JWT token back to the client in a JSON response.
  } else {
    res.status(401).json({ error: 'Invalid username or password' });
  }
});

app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});

app.get('/', (req, res) => {
    res.send('Express server is running!');
  });
  
