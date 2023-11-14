const navItems = document.querySelectorAll('.nav-item');
const infoContainer = document.querySelector('#info-container');

navItems.forEach(item => {
  item.addEventListener('click', () => {
    navItems.forEach(item => item.classList.remove('selected'));
    item.classList.add('selected');

    if (item.innerHTML === 'skills') {
      infoContainer.innerHTML = `
      <ul id="skills-carousel">
          <li><img src="/public/react.svg" alt=""></li>
          <li><img src="/public/language-javascript.svg" alt=""></li>
          <li><img src="/public/language-python.svg" alt=""></li>
          <li><img src="/public/language-html5.svg" alt=""></li>
          <li><img src="/public/language-css3.svg" alt=""></li>
        </ul>
      `;
    } else {
      infoContainer.innerHTML = 'not added yet'
    }
  });
});