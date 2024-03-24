<template>
  <h3>{{ info }}</h3>
  <p>{{ someInfo }}</p>

  <!-- Вставка выпадающего меню -->
  <div class="dropdown">
    <button @click="toggleMenu" class="dropdown-toggle">Меню</button>
    <div v-if="isOpen" class="dropdown-menu">
      <button class="menu-item">Пункт 1</button>
      <button class="menu-item">Пункт 2</button>
      <button class="menu-item">Пункт 3</button>
    </div>
  </div>



  <button type="button" @click="userData()">Отправить</button>


  <input type=text v-model="userName" placeholder="Имя2222">
  <input type="password" v-model="userPass" placeholder="Пароль">
  <input type="email" v-model="userEmail" placeholder="Емайл">
  <p> className="error">{{error}}</p>
  <button @click="sendData">Отправить</button>
  <p>{{ userName }}</p>
  <p>{{ userPass }}</p>
  <p>{{ userEmail }}</p>

  <p>{{ users }}</p>

  <div v-if="users.length == 0 ">
    Нет пользователей

  </div>


  <div v-for="(el, index) in users" :key="index">
    <h3>{{el.name}}</h3>
    <p>{{el.email}} - <b>{{el.pass}}</b></p>
  </div>




</template>
<script>
import background from '@/assets/fon.webp';
export default {
  name: 'App',
  mounted() {
    document.body.style.backgroundImage = `url(${background})`;
  },


  data() {
    return {
      info: 'Title222',
      someInfo: 'myAnons',
      error: '',
      users: [],
      userName: '',
      userPass: '',
      userEmail: '',
      isOpen: false, // Для управления видимостью выпадающего меню

    }
  },
  methods: {
    userData(word = 'New text') {
      this.info = word;
    },
    insertData(val) {
      this.userName = val
    },
    sendData(){
      if (this.userName == ''){
        this.error = 'Имя не введено'
        return;
      } else if (this.userPass == '') {
        this.error = 'Пароль не введен'
        return;
      }if (this.userEmail == ''){
        this.error = 'email не введен'
        return;
      }
      this.error = '';
      this.users.push({
        name: this.userName,
        pass: this.userPass,
        email: this.userEmail,

      })
    },
    toggleMenu() {
      this.isOpen = !this.isOpen;
    },

  }
}
</script>

<style scoped>


body {
  background-color: #000000; /* Темный фон */
  background-image: url('~@/assets/fon.webp'); /* Путь к вашему изображению */
  background-size: cover;
  background-attachment: fixed; /* Эффект параллакса для фона */
  color: #e0e0e0; /* Светлый текст */
  font-family: 'Gothic A1', sans-serif;
  background-position: center center;
}

input, button, .dropdown-menu, .menu-item {
  background-color: #333; /* Темный фон для элементов управления */
  border: 1px solid #444; /* Граница для элементов */
  color: #ddd; /* Цвет текста */
  transition: transform 0.3s ease; /* Анимация для трансформации */
}

input::placeholder {
  color: #bbb; /* Цвет плейсхолдера */
}

p, h3, .error {
  color: #ff5555; /* Акцентный цвет для текста и ошибок */
}

h3 {
  font-weight: lighter;
  font-size: 1.5em; /* Увеличиваем размер заголовков */
}

.dropdown {
  position: relative;
  display: inline-block;
}


.dropdown-toggle {
  background-color: #222; /* Темный фон для кнопки меню */
  color: #ddd; /* Цвет текста кнопки меню */
  padding: 10px;
  margin-top: 20px;
  border: none;
  cursor: pointer;
  animation: glow 2s infinite alternate; /* Анимация "пульсации" для кнопки */
}

.dropdown-menu {
  display: block;
  position: absolute;
  left: 0;
  background-color: #333;
  min-width: 160px;
  box-shadow: 0 8px 16px 0 rgba(0,0,0,0.7); /* Более заметная тень */
  z-index: 1;
}

.menu-item {
  color: #ccc;
  padding: 12px 16px;
  text-decoration: none;
  display: block;
  transition: background-color 0.3s; /* Плавное изменение фона при наведении */
}

.menu-item:hover {
  background-color: #555; /* Изменение фона при наведении */
}
input:hover, button:hover, .menu-item:hover {
  transform: scale(1.05); /* Увеличение при наведении */
}

@keyframes glow {
  from {
    box-shadow: 0 0 5px #ff5555;
  }
  to {
    box-shadow: 0 0 20px #ff5555;
  }
}

.error {
  color: #ff5555; /* Яркий красный цвет для ошибок */
}
</style>

