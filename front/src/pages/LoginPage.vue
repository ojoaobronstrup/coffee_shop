<template>
    <el-container class="login">
        <div class="login-window">
            <div class="login-window-area">
                <el-input class="login-window-area-input" placeholder="Usuário" v-model="username" />
                <el-input class="login-window-area-input" placeholder="Senha" show-password v-model="password" />
            </div>
            <div>
                <EnterButton class="login-window-button" @click="ReqLogin">Entrar</EnterButton>
            </div>
        </div>
    </el-container>
</template>

<script>
import EnterButton from '@/components/EnterButton.vue';
import axios from 'axios';

export default {
    components: {
        EnterButton
    },

    data() {
        return {
            username: undefined,
            password: undefined
        }
    },

    methods: {
        async ReqLogin() {
            const { data, status } = await axios.post("http://localhost:8081/login", {
                username: this.username,
                password: this.password
            })

            if (status === 200) {
                this.$store.state.jwtToken = data.jwtToken
                this.$router.push("/home")
            }
        }
    }
}
</script>

<style>
.login {
    height: 100vh;
    width: 100vw;
    background-image: url(./../../public/assets/coffee_wallpaper.jpeg);
    background-size: cover;
    display: grid;
    grid-template-rows: 1fr 1fr;
}

.login-window {
    height: 15em;
    width: 20em;
    background-color: var(--gray-20);
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-evenly;
    border-radius: 30px;
    grid-row: 2;
    align-self: center;
    justify-self: center;
}

.login-window-button {
    background-color: var(--secondary-brown);
    color: var(--gray-20);
    border: none;
    border-radius: 34px;
    width: 10em;
    height: 2.5em;
    font-size: 20px;
    line-height: 24px;
    font-weight: bold;
    opacity: 0.9;
}

.login-window-area {
    text-align: center;
}

.login-window-area-input {
    width: 20em;
}
</style>