import React, { useState, useEffect } from 'react';
import MenuItemCard from './MenuItemCard';

export default function MenuGrid({ onAddToCart }) {
  const [menuItems, setMenuItems] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch('http://localhost:5000/api/menu')
      .then(res => {
        if (!res.ok) throw new Error('NETWORK RESPONSE ERROR');
        return res.json();
      })
      .then(data => {
        setMenuItems(data);
        setLoading(false);
      })
      .catch(err => {
        console.error('Menu fetch failed:', err);
        setError('FAILED TO INGEST MENU BUFFER');
        setLoading(false);
      });
  }, []);

  if (loading) return <p className="section-title">[FETCHING NUTRIENT MATRIX...]</p>;
  if (error) return <p className="section-title" style={{ color: 'red' }}>[{error}]</p>;

  return (
    <section id="menu">
      <div className="container">
        <h2 className="section-title">Sludge Allocations</h2>
        <div className="menu-grid">
          {menuItems.map(item => (
            <MenuItemCard key={item._id} item={item} onAddToCart={onAddToCart} />
          ))}
        </div>
      </div>
    </section>
  );
}