import React from 'react';

export default function MenuItemCard({ item, onAddToCart }) {
  return (
    <div className="menu-card">
      <img className="menu-img" src={item.imageUrl} alt={item.name} />
      <div className="menu-info">
        <div className="menu-title-row">
          <h3>{item.name}</h3>
          <span className="menu-price">{item.price.toFixed(2)} creds</span>
        </div>
        <p className="menu-desc">{item.description}</p>
        <button
          className="btn add-to-cart-btn"
          onClick={() => onAddToCart(item)}
        >
          Download Mass
        </button>
      </div>
    </div>
  );
}