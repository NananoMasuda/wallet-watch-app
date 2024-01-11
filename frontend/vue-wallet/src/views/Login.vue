<!-- src/views/Login.vue -->
<template>
    <div>
        <h2>ログイン</h2>
        <form>
            <label for="username">Username:</label>
            <input type="text" id="username" v-model="username" required>
            <br>
            <label for="password">Password:</label>
            <input type="password" id="password" v-model="password" required>
            <br>
            <button type="submit">ログイン</button>
        </form>
        <div>
            <p>登録がまだの方は<a href="/register">こちら</a></p>
        </div>
    </div>
</template>

<script>
    
</script>