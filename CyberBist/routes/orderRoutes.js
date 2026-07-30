const express = require('express');
const router = express.Router();
const Order = require('../models/Order');

// POST /api/orders - Process and persist customer checkout
router.post('/', async (req, res) => {
  try {
    const { items, totalPrice, customerInfo } = req.body;

    if (!items || items.length === 0) {
      return res.status(400).json({ message: 'Buffer empty: Cannot process empty order.' });
    }

    const order = new Order({
      items,
      totalPrice,
      customerInfo
    });

    const savedOrder = await order.save();
    res.status(201).json({
      success: true,
      message: 'Telemetry broadcasted and persisted to database.',
      orderId: savedOrder._id
    });
  } catch (err) {
    res.status(500).json({ message: 'Failed to write order state to database.', error: err.message });
  }
});

// GET /api/orders - (Optional) View system order log
router.get('/', async (req, res) => {
  try {
    const orders = await Order.find().sort({ createdAt: -1 });
    res.json(orders);
  } catch (err) {
    res.status(500).json({ message: 'Failed to retrieve order logs.', error: err.message });
  }
});

module.exports = router;