import React, { useState, useEffect } from 'react';
import Header from './components/Header';
import MenuGrid from './components/MenuGrid';
import CartSidebar from './components/CartSidebar';
import ImageSlider from './components/ImageSlider';
import ContactForm from './components/ContactForm';

export default function App() {
  const [theme, setTheme] = useState('theme-cyberpunk');
  const [cart, setCart] = useState([]);
  const [isCartOpen, setIsCartOpen] = useState(false);

  // Synchronize theme class on body element
  useEffect(() => {
    document.body.className = theme;
  }, [theme]);

  const toggleTheme = () => {
    setTheme(prev => prev === 'theme-cyberpunk' ? 'theme-corporate' : 'theme-cyberpunk');
  };

  const handleAddToCart = (item) => {
    setCart(prevCart => {
      const existing = prevCart.find(i => i._id === item._id);
      if (existing) {
        return prevCart.map(i =>
          i._id === item._id ? { ...i, quantity: i.quantity + 1 } : i
        );
      }
      return [...prevCart, { ...item, quantity: 1 }];
    });
  };

  const handleRemoveFromCart = (id) => {
    setCart(prevCart => prevCart.filter(item => item._id !== id));
  };

  const handleClearCart = () => setCart([]);

  const totalCartCount = cart.reduce((count, item) => count + item.quantity, 0);

  return (
    <React.Fragment>
      <Header
        cartCount={totalCartCount}
        onOpenCart={() => setIsCartOpen(true)}
        theme={theme}
        onToggleTheme={toggleTheme}
      />

      <section id="home">
        <div>
          <h1>CYBERNETIC NUTRIENT MATRIX</h1>
          <p>// CONSUME. SUBMIT. RENDER STATE.</p>
        </div>
      </section>

      <section id="about">
        <div className="container">
          <h2 className="section-title">SYSTEM_LOG.TXT</h2>
          <div className="about-content">
            <p>This application now runs on genuine React state synchronization, powered by an Express API and MongoDB persistence layer. The Trojan-Horse imperative DOM mutations have been completely purged.</p>
          </div>
        </div>
      </section>

      <MenuGrid onAddToCart={handleAddToCart} />

      <ImageSlider />

      <ContactForm />

      <CartSidebar
        isOpen={isCartOpen}
        onClose={() => setIsCartOpen(false)}
        cart={cart}
        onRemoveItem={handleRemoveFromCart}
        onClearCart={handleClearCart}
      />

      <footer>
        <div className="container">
          <div className="footer-grid">
            <div className="footer-col">
              <h3>MONOLITH CONGLOMERATE</h3>
              <p>We abstract your hunger via modular array loops.</p>
            </div>
            <div className="footer-col">
              <h3>Uptime Parameters</h3>
              <p>Cycles 1-4: 1600 - 2200 Hours</p>
              <p>Cycles 5-7: 1200 - 2300 Hours</p>
            </div>
            <div className="footer-col">
              <h3>Signal Channels</h3>
              <div className="social-links">
                <a href="#link">DeepWeb/FB</a>
                <a href="#link">NeuralNet/IG</a>
                <a href="#link">X-Sector</a>
              </div>
            </div>
          </div>
          <div className="footer-bottom">
            <p>© 2026 NEO-MONOLITH AUTOMATIONS. Rebuilt inside React. Still zero shame.</p>
          </div>
        </div>
      </footer>
    </React.Fragment>
  );
}