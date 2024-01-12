<!-- src/views/Login.vue -->
<template>
  <div>
      <h1>ログイン</h1>
      <form @submit.prevent="loginUser">
          <label for="username">
          <input type="text" placeholder="Username" id="username" v-model="username" required>
          </label>
          <br>
          <label for="password">
          <input type="password" placeholder="Password" id="password" v-model="password" required>
          </label>
          <br>
          <button type="submit">ログイン</button>
      </form>
      <div v-if="errorMessage" class="error-message">
        {{ errorMessage }}
      </div>
      <div>
          <p>登録がまだの方は<a href="/register">こちら</a></p>
      </div>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data() {
    return {
      username: '',
      password: '',
      token: '',
      errorMessage: ''
    };
  },
  methods: {
    loginUser() {
      const data = {
        username: this.username,
        password: this.password,
      };
      

      fetch('http://localhost:8000/api/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      })
      .then(tokenResponse => tokenResponse.json())
      .then(tokenResponse => {

        this.token = tokenResponse.token;

        fetch('http://localhost:8000/api/admin/user', {
            method: 'GET',
            headers: {
              'Authorization': `Bearer ${this.token}`,
            }
        })
        .then(loginResponse => {
          if (!loginResponse.ok) {
            throw new Error(`HTTP error! Status: ${loginResponse.status}`);
          }
          
          this.$router.push('/home');

        })
        .catch(error => {
        console.error('Error:', error);
        this.errorMessage = 'アカウントがありませんよ。先に登録してくださいね。';
        })

      })
      .catch((error) => {
        console.error('Error:', error);
      })
    }
  }
}
</script>