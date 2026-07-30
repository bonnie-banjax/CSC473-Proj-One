import React, { useEffect, useState } from 'react';

export default function App() {
  useEffect(() => {
    // 1. Paste your exact Vanilla JS code completely unmodified
    let cart = [];
    let currentSlide = 0;

    const hamburgerBtn = document.getElementById('hamburger-btn');
    const navMenu = document.getElementById('nav-menu');
    const slider = document.getElementById('image-slider');
    const totalSlides = document.querySelectorAll('.slide').length;
    const cartSidebar = document.getElementById('cart-sidebar');
    const cartItemsContainer = document.getElementById('cart-items-container');
    const cartTotalPrice = document.getElementById('cart-total-price');
    const cartIndicator = document.getElementById('cart-indicator');

    hamburgerBtn.addEventListener('click', () => navMenu.classList.toggle('active'));
    document.querySelectorAll('#nav-menu a').forEach(link =>
      link.addEventListener('click', () => navMenu.classList.remove('active'))
    );

    function runSliderUpdate() { slider.style.transform = `translateX(-${currentSlide * 100}%)`; }
    document.getElementById('slider-next-btn').addEventListener('click', () => { currentSlide = (currentSlide + 1) % totalSlides; runSliderUpdate(); });
    document.getElementById('slider-prev-btn').addEventListener('click', () => { currentSlide = (currentSlide - 1 + totalSlides) % totalSlides; runSliderUpdate(); });

    document.getElementById('open-cart-btn').addEventListener('click', () => cartSidebar.classList.add('open'));
    document.getElementById('close-cart-btn').addEventListener('click', () => cartSidebar.classList.remove('open'));

    document.querySelectorAll('.add-to-cart-btn').forEach(btn => {
      btn.addEventListener('click', (e) => {
        const id = e.target.getAttribute('data-id');
        const name = e.target.getAttribute('data-name');
        const price = parseFloat(e.target.getAttribute('data-price'));
        const cachedItem = cart.find(i => i.id === id);
        if (cachedItem) { cachedItem.quantity++; } else { cart.push({ id, name, price, quantity: 1 }); }
        renderSimulacrum();
      });
    });

    // CRITICAL FIX: Bind to window so the vanilla inline string onclick works
    window.removeFromCart = function(id) {
      cart = cart.filter(i => i.id !== id);
      renderSimulacrum();
    };

    document.getElementById('clear-cart-btn').addEventListener('click', () => { cart = []; renderSimulacrum(); });

    function renderSimulacrum() {
      cartItemsContainer.innerHTML = '';
      let debitCounter = 0;
      let nodeTally = 0;

      if (cart.length === 0) {
        cartItemsContainer.innerHTML = '<p style="text-align:center; color:#555; font-size:0.8rem; margin-top:20px;">[BUFFER EMPTY: NO INGESTION SCHEDULED]</p>';
      } else {
        cart.forEach(item => {
          debitCounter += item.price * item.quantity;
          nodeTally += item.quantity;
          const row = document.createElement('div');
          row.classList.add('cart-item');
          row.innerHTML = `
            <div class="cart-item-info">
              <h4>${item.name}</h4>
              <span>${item.price.toFixed(2)} CR x ${item.quantity}</span>
            </div>
            <button class="remove-item-btn" onclick="removeFromCart('${item.id}')">[DISCARD]</button>
          `;
          cartItemsContainer.appendChild(row);
        });
      }
      cartTotalPrice.textContent = `${debitCounter.toFixed(2)} creds`;
      cartIndicator.textContent = nodeTally;
    }

    const toggleSwitch = document.querySelector('#checkbox');
    if (document.body.classList.contains('theme-corporate')) {
      toggleSwitch.checked = true;
    } else {
      document.body.classList.add('theme-cyberpunk');
    }

    function switchTheme(e) {
      if (e.target.checked) {
        document.body.classList.remove('theme-cyberpunk');
        document.body.classList.add('theme-corporate');
      } else {
        document.body.classList.remove('theme-corporate');
        document.body.classList.add('theme-cyberpunk');
      }
    }
    toggleSwitch.addEventListener('change', switchTheme, false);

    // Form fix
    const form = document.querySelector('.contact-form');
    const formHandler = (event) => {
      event.preventDefault();
      alert('Telemetry broadcasted to corporate routers.');
      form.reset();
    };
    form.addEventListener('submit', formHandler);

    // Cleanup listeners on unmount so React doesn't duplicate them
    return () => {
      form.removeEventListener('submit', formHandler);
      delete window.removeFromCart;
    };
  }, []);

  // 2. Your exact HTML string with class changed to className and fixed typos
  return (
    <>
      <header>
        <div className="container">
          <a href="mono-culture.html" className="logo">NEO-MONOLITH // v1.0.4</a>
          <button className="hamburger" id="hamburger-btn">SYS</button>
          <nav id="nav-menu">
            <a href="#home">Terminal</a>
            <a href="#menu">Rations</a>
            <a href="#about">Log-File</a>
            <a href="#contact">Geolocate</a>
            <button className="btn cart-trigger-btn" id="open-cart-btn">
              [SLOT_ALLOCATION] <span className="cart-count" id="cart-indicator">0</span>
            </button>
            <div className="theme-switch-wrapper">
              <span className="theme-label label-cyber">CYBER</span>
              <label className="theme-switch" htmlFor="checkbox">
                <input type="checkbox" id="checkbox" />
                <div className="slider-toggle round"></div>
              </label>
              <span className="theme-label label-corp">CORP</span>
            </div>
          </nav>
        </div>
      </header>

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
            <p>This section features generic, heartfelt copy optimized for search algorithms. Every restaurant layout on the web utilizes this exact layout pattern. It took an LLM approximately 0.4 seconds to generate this entire business structure. 94% of all local business websites are structural duplicates operating under the illusion of unique identity. You are browsing inside a deterministic cage of standard boxes styled to look distinct. Enjoy your simulated dining paradigm.</p>
          </div>
        </div>
      </section>

      <section id="menu">
        <div className="container">
          <h2 className="section-title">Sludge Allocations</h2>
          <div className="menu-grid">
            <div className="menu-card">
              <img className="menu-img" src="https://images.unsplash.com/photo-1568901346375-23c9450c58cd?auto=format&fit=crop&w=500&q=80" alt="Synthetic Carbs" />
              <div className="menu-info">
                <div className="menu-title-row">
                  <h3>Synth-Wagyu Cuboid</h3>
                  <span className="menu-price">18.00 creds</span>
                </div>
                <p className="menu-desc">Reconstituted beef proteins suspended in artificial lipid-emulsion on high-density carbohydrate bun. Enhanced with synthetic truffle flavor-code.</p>
                <button className="btn add-to-cart-btn" data-id="1" data-name="Synth-Wagyu Cuboid" data-price="18.00">Download Mass</button>
              </div>
            </div>
            <div className="menu-card">
              <img className="menu-img" src="https://images.unsplash.com/photo-1513104890138-7c749659a591?auto=format&fit=crop&w=500&q=80" alt="Algae Discs" />
              <div className="menu-info">
                <div className="menu-title-row">
                  <h3>Thermal Dough Disc</h3>
                  <span className="menu-price">16.50 creds</span>
                </div>
                <p className="menu-desc">Flatbread mechanism loaded with coagulated cow lipids and high-sodium plant extract fluids. Broiled at 450°C inside corporate kilns.</p>
                <button className="btn add-to-cart-btn" data-id="2" data-name="Thermal Dough Disc" data-price="16.50">Download Mass</button>
              </div>
            </div>
            <div className="menu-card">
              <img className="menu-img" src="https://images.unsplash.com/photo-1546069901-ba9599a7e63c?auto=format&fit=crop&w=500&q=80" alt="Omega Matrix" />
              <div className="menu-info">
                <div className="menu-title-row">
                  <h3>Heavy-Metal Salmon Base</h3>
                  <span className="menu-price">21.00 creds</span>
                </div>
                <p className="menu-desc">Slices of aquaculture flesh loaded with Omega-3 and trace ocean microplastics, stacked perfectly upon bleached carbohydrate grains.</p>
                <button className="btn add-to-cart-btn" data-id="3" data-name="Heavy-Metal Salmon Base" data-price="21.00">Download Mass</button>
              </div>
            </div>
          </div>
        </div>
      </section>

      <section id="gallery">
        <div className="container">
          <h2 className="section-title">Optic Buffers</h2>
          <div className="slider-container">
            <div className="slider" id="image-slider">
              <img className="slide" src="https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&w=800&q=80" alt="Img" />
              <img className="slide" src="https://images.unsplash.com/photo-1555396273-367ea4eb4db5?auto=format&fit=crop&w=800&q=80" alt="Img" />
              <img className="slide" src="https://images.unsplash.com/photo-1414235077428-338989a2e8c0?auto=format&fit=crop&w=800&q=80" alt="Img" />
              <img className="slide" src="https://images.unsplash.com/photo-1559339352-11d035aa65de?auto=format&fit=crop&w=800&q=80" alt="Img" />
              <img className="slide" src="https://images.unsplash.com/photo-1514933651103-005eec06c04b?auto=format&fit=crop&w=800&q=80" alt="Img" />
              <img className="slide" src="https://images.unsplash.com/photo-1544025162-d76694265947?auto=format&fit=crop&w=800&q=80" alt="Img" />
            </div>
            <button className="slider-btn prev" id="slider-prev-btn">◀</button>
            <button className="slider-btn next" id="slider-next-btn">▶</button>
          </div>
        </div>
      </section>

      <section id="contact">
        <div className="container">
          <h2 className="section-title">Telemetry Upload</h2>
          <div className="contact-wrapper">
            <form className="contact-form">
              <div className="form-group">
                <label>Designation (Name)</label>
                <input type="text" required />
              </div>
              <div className="form-group">
                <label>Comm Net-Node (Email)</label>
                <input type="email" required />
              </div>
              <div className="form-group">
                <label>Manifest Manifestations (Message)</label>
                <textarea rows={5} required></textarea>
              </div>
              <button type="submit" className="btn">Broadcast Signal</button>
            </form>
            <div className="map-container">
              <iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3022.6175392743824!2d-73.9878531242371!3d40.74844447138905!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x89c259a9b3117469%3A0xd134e199a405a163!2sEmpire%20State%20Building!5e0!3m2!1sen!2sus!4v1710000000000!5m2!1sen!2sus" allowFullScreen={true} loading="lazy" title="map"></iframe>
            </div>
          </div>
        </div>
      </section>

      <div className="cart-sidebar" id="cart-sidebar">
        <div className="cart-header">
          <h3>MANIFEST REQUISITION BUFFER</h3>
          <button className="cart-close" id="close-cart-btn">🎚️</button>
        </div>
        <div className="cart-items" id="cart-items-container"></div>
        <div className="cart-footer">
          <div className="cart-total-row">
            <span>CREDIT DEBT:</span>
            <span id="cart-total-price">0.00 creds</span>
          </div>
          <div className="cart-actions">
            <button className="btn btn-clear" id="clear-cart-btn">Purge Buffer State</button>
          </div>
        </div>
      </div>

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
            <p>© 2026 NEO-MONOLITH AUTOMATIONS. Built inside React. Still zero shame.</p>
          </div>
        </div>
      </footer>
    </>
  );
}