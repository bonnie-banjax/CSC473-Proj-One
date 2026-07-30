import React, { useState } from 'react';

export default function Header({ cartCount, onOpenCart, theme, onToggleTheme }) {
  const [isNavOpen, setIsNavOpen] = useState(false);

  return (
    <header>
      <div className="container">
        <a href="#home" className="logo">NEO-MONOLITH // v2.0.0</a>

        <button
          className="hamburger"
          id="hamburger-btn"
          onClick={() => setIsNavOpen(prev => !prev)}
        >
          SYS
        </button>

        <nav id="nav-menu" className={isNavOpen ? 'active' : ''}>
          <a href="#home" onClick={() => setIsNavOpen(false)}>Terminal</a>
          <a href="#menu" onClick={() => setIsNavOpen(false)}>Rations</a>
          <a href="#about" onClick={() => setIsNavOpen(false)}>Log-File</a>
          <a href="#contact" onClick={() => setIsNavOpen(false)}>Geolocate</a>

          <button className="btn cart-trigger-btn" onClick={onOpenCart}>
            [SLOT_ALLOCATION] <span className="cart-count">{cartCount}</span>
          </button>

          <div className="theme-switch-wrapper">
            <span className="theme-label label-cyber">CYBER</span>
            <label className="theme-switch" htmlFor="checkbox">
              <input
                type="checkbox"
                id="checkbox"
                checked={theme === 'theme-corporate'}
                onChange={onToggleTheme}
              />
              <div className="slider-toggle round"></div>
            </label>
            <span className="theme-label label-corp">CORP</span>
          </div>
        </nav>
      </div>
    </header>
  );
}