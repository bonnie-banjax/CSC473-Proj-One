const mongoose = require('mongoose');

const orderItemSchema = new mongoose.Schema({
  menuItem: {
    type: mongoose.Schema.Types.ObjectId,
    ref: 'MenuItem',
    required: true
  },
  name: { type: String, required: true },
  price: { type: Number, required: true },
  quantity: { type: Number, required: true }
});

const orderSchema = new mongoose.Schema({
  items: [orderItemSchema],
  totalPrice: { type: Number, required: true },
  customerInfo: {
    designation: { type: String, required: true }, // Name
    commNode: { type: String, required: true },    // Email
    manifestations: { type: String }                // Message/Notes
  },
  status: { type: String, default: 'QUEUED' }
}, { timestamps: true });

module.exports = mongoose.model('Order', orderSchema);