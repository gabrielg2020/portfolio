const navItems = document.querySelectorAll('.nav-item');
const infoContainer = document.querySelector('#info-container');

navItems.forEach(item => {
  item.addEventListener('click', () => {
    if (item.classList.contains('selected')) {
      infoContainer.innerHTML = `
      <div id="default-wrapper">
        <h1>Hello! I'm Gabriel Guimaraes,</h1>
        <h2>Software Engineer.</h2>
        <p>Click around to learn more!</p>
      </div>`;
      item.classList.remove('selected');
      return;
    }

    if (item.innerHTML === 'download resume') {
      return;
    }

    navItems.forEach(item => item.classList.remove('selected'));
    item.classList.add('selected');

    if (item.innerHTML === 'skills') {
      infoContainer.innerHTML = `
      <ul id="carousel">
          <li><img src="public/react.svg" alt=""></li>
          <li><img src="public/language-javascript.svg" alt=""></li>
          <li><img src="public/language-python.svg" alt=""></li>
          <li><img src="public/language-html5.svg" alt=""></li>
          <li><img src="public/language-css3.svg" alt=""></li>
        </ul>
      `;
    } else if (item.innerHTML === 'projects') {
      infoContainer.innerHTML = ` <div id="project-list">
      <div class="project">
        <div class="top-line">
          <h2>Maze Visualizer</h2>
          <a href="https://github.com/gabrielg2020/A-Level-Computer-Science-NEA-2023"><img src="public/open-in-new.svg" alt=""></a>
        </div>
        <div>
          A highly interactive maze visualizer built with VB.NET and WinForms.
        </div>
      </div>

      <div class="project">
        <div class="top-line">
          <h2>Physics Simulator</h2>
          <a href="https://github.com/gabrielg2020/realtime-physics-sim"><img src="public/open-in-new.svg" alt=""></a>
        </div>
        <div>
          A realtime physics simulator built with React, JavaScript and PixiJS.
        </div>
      </div>
    </div>`;
    } else if (item.innerHTML === 'get in touch') {
      infoContainer.innerHTML = `
      <ul id="carousel">
        <li><a href="https://twitter.com/gabrielg3333"><img src="public/twitter.svg" alt=""></a></li>
        <li><a href="https://discordapp.com/users/598261411709321216"><img src="public/icons8-discord-512.svg" alt=""></a></li>
        <li><a href="mailto:gabriel.mg04@outlook.com"><img src="public/email-outline.svg" alt=""></a></li>
      </ul>`;
    }
  });
});