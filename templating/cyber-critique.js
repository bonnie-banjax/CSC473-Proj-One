
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

// UI Link Toggle State Array mutations
hamburgerBtn.addEventListener('click', () => navMenu.classList.toggle('active'));
document.querySelectorAll('#nav-menu a').forEach(link => link.addEventListener('click', () => navMenu.classList.remove('active')));

// Automated Visual Array Translation (Slider Module)
function runSliderUpdate() { slider.style.transform = `translateX(-${currentSlide * 100}%)`; }
document.getElementById('slider-next-btn').addEventListener('click', () => { currentSlide = (currentSlide + 1) % totalSlides; runSliderUpdate(); });
document.getElementById('slider-prev-btn').addEventListener('click', () => { currentSlide = (currentSlide - 1 + totalSlides) % totalSlides; runSliderUpdate(); });

// The Grand State Sync Lie (Cart Component Module)
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

function removeFromCart(id) {
  cart = cart.filter(i => i.id !== id);
  renderSimulacrum();
}

document.getElementById('clear-cart-btn').addEventListener('click', () => { cart = []; renderSimulacrum(); });

// Renders the fake simulation nodes into DOM fragments
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

// Style Changer Button

const toggleSwitch = document.querySelector('#checkbox');

// Establish default state on load
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