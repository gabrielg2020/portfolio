const navItems = document.querySelectorAll('.nav-item');

navItems.forEach(item => {
  item.addEventListener('click', () => {
    navItems.forEach(item => item.classList.remove('selected'));
    item.classList.add('selected');
  });
});