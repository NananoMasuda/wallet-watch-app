<!-- src/views/Register.vue -->
<template>
  <div>
      <h1>登録</h1>
      <form @submit.prevent="registerUser">
          <label for="username">
          <input type="text" placeholder="Username" id="username" v-model="username" required>
          </label>
          <br>
          <label for="password">
          <input type="password" placeholder="Username" id="password" v-model="password" required>
          </label>
          <br>
          <button type="submit">登録</button>
      </form>
      <div>
          <p><a href="/login">ログインページに戻る</a></p>
      </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: '',
      password: '',
    };
  },
  methods: {
    registerUser() {
      const data = {
        username: this.username,
        password: this.password,
      };

      fetch('http://localhost:8000/api/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      })
      .then(response => response.json())
      .then(data => {
        console.log('Success:', data);
        this.$router.push('/login');
      })
      .catch((error) => {
        console.error('Error:', error);
      });
    }
  },
};
</script>