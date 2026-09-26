async function loadProducts() {
  const response = await fetch("/api/products");
  const cakes = await response.json();

  const grid = document.getElementById("cake-grid");
  cakes.forEach(cake => {
    grid.innerHTML += renderProductCard(cake);
  });
}

loadProducts();