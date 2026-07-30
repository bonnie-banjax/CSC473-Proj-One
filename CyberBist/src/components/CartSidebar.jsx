import React, { useState } from 'react';

export default function CartSidebar({ isOpen, onClose, cart, onRemoveItem, onClearCart }) {
  const [submitting, setSubmitting] = useState(false);

  const totalPrice = cart.reduce((sum, item) => sum + (item.price * item.quantity), 0);

  const handleCheckout = async () => {
    if (cart.length === 0) return;

    setSubmitting(true);
    const orderPayload = {
      items: cart.map(i => ({
        menuItem: i._id,
        name: i.name,
        price: i.price,
        quantity: i.quantity
      })),
      totalPrice,
      customerInfo: {
        designation: "Terminal User",
        commNode: "user@cybernet.local",
        manifestations: "Instant dispatch request"
      }
    };

    try {
      const response = await fetch('http://localhost:5000/api/orders', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(orderPayload)
      });

      if (response.ok) {
        alert('Telemetry broadcasted and persisted to database.');
        onClearCart();
        onClose();
      } else {
        alert('Order transmission error.');
      }
    } catch (err) {
      console.error('Checkout error:', err);
      alert('System communications offline.');
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div className={`cart-sidebar ${isOpen ? 'open' : ''}`} id="cart-sidebar">
      <div className="cart-header">
        <h3>MANIFEST REQUISITION BUFFER</h3>
        <button className="cart-close" onClick={onClose}>🎚️</button>
      </div>

      <div className="cart-items">
        {cart.length === 0 ? (
          <p style={{ textAlign: 'center', color: '#555', fontSize: '0.8rem', marginTop: '20px' }}>
            [BUFFER EMPTY: NO INGESTION SCHEDULED]
          </p>
        ) : (
          cart.map(item => (
            <div key={item._id} className="cart-item">
              <div className="cart-item-info">
                <h4>{item.name}</h4>
                <span>{item.price.toFixed(2)} CR x {item.quantity}</span>
              </div>
              <button
                className="remove-item-btn"
                onClick={() => onRemoveItem(item._id)}
              >
                [DISCARD]
              </button>
            </div>
          ))
        )}
      </div>

      <div className="cart-footer">
        <div className="cart-total-row">
          <span>CREDIT DEBT:</span>
          <span>{totalPrice.toFixed(2)} creds</span>
        </div>
        <div className="cart-actions" style={{ display: 'flex', gap: '10px' }}>
          <button className="btn btn-clear" onClick={onClearCart}>Purge</button>
          <button
            className="btn"
            onClick={handleCheckout}
            disabled={cart.length === 0 || submitting}
          >
            {submitting ? 'TRANSMITTING...' : 'Commit Order'}
          </button>
        </div>
      </div>
    </div>
  );
}