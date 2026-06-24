
// --- APP STATE ---
let cart = [];
let currentSlide = 0;

// --- DOM ELEMENT SELECTORS ---
const hamburgerBtn = document.getElementById('hamburger-btn');
const navMenu = document.getElementById('nav-menu');
const slider = document.getElementById('image-slider');
const totalSlides = document.querySelectorAll('.slide').length;
const cartSidebar = document.getElementById('cart-sidebar');
const cartItemsContainer = document.getElementById('cart-items-container');
const cartTotalPrice = document.getElementById('cart-total-price');
const cartIndicator = document.getElementById('cart-indicator');

// ==========================================
// MODULE 1: RESPONSIVE INTERFACES & MENUS
// ==========================================
hamburgerBtn.addEventListener('click', () => {
  navMenu.classList.toggle('active');
});

// Automatically collapse mobile menu when clicking a link
document.querySelectorAll('#nav-menu a').forEach(link => {
  link.addEventListener('click', () => navMenu.classList.remove('active'));
});

// ==========================================
// MODULE 2: IMAGE SLIDER COMPONENT
// ==========================================
function updateSliderPosition() {
  slider.style.transform = `translateX(-${currentSlide * 100}%)`;
}

document.getElementById('slider-next-btn').addEventListener('click', () => {
  currentSlide = (currentSlide + 1) % totalSlides;
  updateSliderPosition();
});

document.getElementById('slider-prev-btn').addEventListener('click', () => {
  currentSlide = (currentSlide - 1 + totalSlides) % totalSlides;
  updateSliderPosition();
});

// ==========================================
// MODULE 3: STATEFUL SHOPPING CART CONTROLLER
// ==========================================

// Toggle View State Links
document.getElementById('open-cart-btn').addEventListener('click', () => cartSidebar.classList.add('open'));
document.getElementById('close-cart-btn').addEventListener('click', () => cartSidebar.classList.remove('open'));

// Hook Up "Add To Cart" DOM Buttons
document.querySelectorAll('.add-to-cart-btn').forEach(button => {
  button.addEventListener('click', (e) => {
    const id = e.target.getAttribute('data-id');
    const name = e.target.getAttribute('data-name');
    const price = parseFloat(e.target.getAttribute('data-price'));
    addToCart(id, name, price);
  });
});

// State Mutation: Add Item
function addToCart(id, name, price) {
  const existingItem = cart.find(item => item.id === id);
  if (existingItem) {
    existingItem.quantity += 1;
  } else {
    cart.push({ id, name, price, quantity: 1 });
  }
  renderCart();
}

// State Mutation: Remove Item
function removeFromCart(id) {
  cart = cart.filter(item => item.id !== id);
  renderCart();
}

// State Mutation: Reset / Clear Cart
document.getElementById('clear-cart-btn').addEventListener('click', () => {
  cart = [];
  renderCart();
});

// Render Function: Maps Array State directly onto dynamic HTML DOM components
function renderCart() {
  cartItemsContainer.innerHTML = '';
  let total = 0;
  let totalItemsCount = 0;

  if (cart.length === 0) {
    cartItemsContainer.innerHTML = '<p style="text-align:center; color:#888; margin-top:20px;">Your cart is empty.</p>';
  } else {
    cart.forEach(item => {
      total += item.price * item.quantity;
      totalItemsCount += item.quantity;

      const itemElement = document.createElement('div');
      itemElement.classList.add('cart-item');
      itemElement.innerHTML = `
        <div class="cart-item-info">
          <h4>${item.name}</h4>
          <span>$${item.price.toFixed(2)} x ${item.quantity}</span>
        </div>
        <button class="remove-item-btn" onclick="removeFromCart('${item.id}')">Remove</button>
      `;
      cartItemsContainer.appendChild(itemElement);
    });
  }

  // Sync computed values with visual indicators
  cartTotalPrice.textContent = `$${total.toFixed(2)}`;
  cartIndicator.textContent = totalItemsCount;
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
