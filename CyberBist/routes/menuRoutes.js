const express = require('express');
const router = express.Router();
const MenuItem = require('../models/MenuItem');

// GET /api/menu - Retrieve all menu items for frontend display
router.get('/', async (req, res) => {
  try {
    const items = await MenuItem.find();
    res.json(items);
  } catch (err) {
    res.status(500).json({ message: 'Failed to retrieve menu buffer state.', error: err.message });
  }
});

// POST /api/menu - Seed or add a new menu item
router.post('/', async (req, res) => {
  try {
    const newItem = await MenuItem.create(req.body);
    res.status(201).json(newItem);
  } catch (err) {
    res.status(400).json({ message: 'Invalid payload schema.', error: err.message });
  }
});

module.exports = router;