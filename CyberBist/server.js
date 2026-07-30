const express = require('express');
const cors = require('cors');
const path = require('path'); // Node built-in module
require('dotenv').config();
const connectDB = require('./config/db.js');

const app = express();

// Connect to MongoDB
connectDB();

// Middleware
app.use(cors());
app.use(express.json());

// 1. API Routes
app.use('/api/menu', require('./routes/menuRoutes'));
app.use('/api/orders', require('./routes/orderRoutes'));

// 2. Serve static files from Vite's build output folder
app.use(express.static(path.join(__dirname, 'dist')));

// 3. SPA Fallback: Serve index.html for any unhandled routes
// Replace app.get('*', ...) with:
// Replace app.get('*', ...) with:
app.get(/.*$/, (req, res) => {
  res.sendFile(path.join(__dirname, 'dist', 'index.html'));
});

const PORT = process.env.PORT || 5000;
app.listen(PORT, () => {
  console.log(`[SYS_LOG]: Production server running on port ${PORT}`);
});