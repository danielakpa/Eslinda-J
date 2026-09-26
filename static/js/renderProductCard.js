// Builds one cake card, styled like FNP: image, heart icon, name, price, delivery line
function renderProductCard(cake) {
  return `
    <div class="cake-card">
      <div class="card-image">
        <img src="/static/images/${cake.image}" alt="${cake.name}">
        <span class="wishlist-heart">♡</span>
      </div>
      <button class="customize-btn" onclick="addToCart(${cake.id})">Add to Cart</button>
      <h3 class="cake-name">${cake.name}</h3>
      <p class="cake-price">SAR ${cake.price}</p>
      <p class="delivery-text">Same Day • By 5:00 pm</p>
    </div>
  `;
}