<!-- src/views/Login.vue -->
<template>
  <div>
      <h2>ログイン</h2>
      <form @submit.prevent="loginUser">
          <label for="username">Username:</label>
          <input type="text" id="username" v-model="username" required>
          <br>
          <label for="password">Password:</label>
          <input type="password" id="password" v-model="password" required>
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

<style scoped>
.form-item {
  margin: 0 auto;
  text-align: center;
}

label {
  display: block;
}

input {
  width: 50%;
  padding: .5em;
  font: inherit;
}

button {
  padding: 0.5em;
  margin: 1em;
}
</style>