<template>
    <div class="container">
        <div class="login__wrap">
            <div class="login__block">
                <form @submit.prevent="submit"  class="login__block__wrap">
                    <div class="login__block-header">Личный кабинет</div>
                    <div class="login__block-content">
                        <div class="input__wrap">
                            <div class="input__icon">
                                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                    <path d="M2 6C2 5.46957 2.21071 4.96086 2.58579 4.58579C2.96086 4.21071 3.46957 4 4 4H20C20.5304 4 21.0391 4.21071 21.4142 4.58579C21.7893 4.96086 22 5.46957 22 6V18C22 18.5304 21.7893 19.0391 21.4142 19.4142C21.0391 19.7893 20.5304 20 20 20H4C3.46957 20 2.96086 19.7893 2.58579 19.4142C2.21071 19.0391 2 18.5304 2 18V6ZM5.519 6L12 11.671L18.481 6H5.519ZM20 7.329L12.659 13.753C12.4766 13.9128 12.2424 14.0009 12 14.0009C11.7576 14.0009 11.5234 13.9128 11.341 13.753L4 7.329V18H20V7.329Z" fill="black"/>
                                </svg>
                            </div>
                            <input class="input" type="email" name="" id="email" v-model="this.email" placeholder="Введите почту" required>
                        </div>
                        <div class="input__wrap">
                            <div class="input__icon">
                                <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                    <path d="M15.75 1.5C14.6959 1.49976 13.6565 1.74642 12.7149 2.22023C11.7734 2.69404 10.9559 3.38181 10.328 4.22844C9.7001 5.07506 9.27925 6.057 9.09917 7.09555C8.91909 8.13411 8.98479 9.20041 9.291 10.209L1.5 18V22.5H6L13.791 14.709C14.7194 14.9908 15.6977 15.0692 16.6591 14.9387C17.6206 14.8083 18.5426 14.4721 19.3624 13.953C20.1821 13.434 20.8804 12.7444 21.4095 11.9311C21.9386 11.1178 22.2862 10.2 22.4285 9.24025C22.5709 8.28049 22.5046 7.30132 22.2343 6.36948C21.964 5.43764 21.496 4.57502 20.8622 3.84042C20.2283 3.10582 19.4436 2.51649 18.5614 2.11261C17.6792 1.70872 16.7203 1.49978 15.75 1.5ZM15.75 13.5C15.2336 13.4997 14.7201 13.4234 14.226 13.2735L13.3657 13.0125L12.7305 13.6478L10.3448 16.0335L9.3105 15L8.25 16.0605L9.28425 17.0947L8.09475 18.2842L7.0605 17.25L6 18.3105L7.03425 19.3447L5.379 21H3V18.621L10.3515 11.2695L10.9875 10.6342L10.7265 9.774C10.4059 8.71724 10.4268 7.58631 10.786 6.54207C11.1453 5.49784 11.8247 4.59347 12.7275 3.95762C13.6304 3.32177 14.7108 2.98681 15.815 3.0004C16.9192 3.01398 17.9911 3.37542 18.878 4.03329C19.765 4.69116 20.4219 5.61197 20.7554 6.66473C21.0888 7.71749 21.0818 8.84859 20.7354 9.89714C20.3889 10.9457 19.7206 11.8583 18.8256 12.5051C17.9305 13.152 16.8543 13.5001 15.75 13.5Z" fill="black"/>
                                    <path d="M16.5 9C17.3284 9 18 8.32843 18 7.5C18 6.67157 17.3284 6 16.5 6C15.6716 6 15 6.67157 15 7.5C15 8.32843 15.6716 9 16.5 9Z" fill="black"/>
                                </svg>
                            </div>
                            <input class="input" type="password" name="" id="pass" v-model="this.password" placeholder="Введите пароль" required>
                        </div>
                    </div>
                    <div class="login__block-footer">
                        <div class="footer-dop">Нет аккаунта? <router-link to="/registration">Зарегистрироваться</router-link></div>
                        <button class="btn white" type="submit">Авторизоваться</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
import userService from '../services/user'

export default {
    name: "Login",
    data() {
        return {
            email: null,
            password: null
        }
    }, 
    methods: {
        submit() {
            userService.userLogin(this.email, this.password).then(res => {
                console.log(res);
                this.$notify({
                    title: "Успешно",
                    text: "Вы успешно авторизовались!",
                    type: "success"
                });
            }).catch(res => {
                console.log(res);
                this.$notify({
                    title: "Внимание",
                    text: res.response.data.msg,
                    type: "warn"
                });
            })
        }
    },
    mounted() {

    }
}
</script>

<style scoped>
.login__wrap {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-top: 50px;
}
.login__block {
    width: 400px;
    background-color: #D8F3DC;
    border-radius: 20px;
}
.login__block__wrap {
    padding: 20px;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 25px;
}
.login__block-header {
    font-size: 18px;
    font-weight: 500;
}
.login__block-content {
    display: flex;
    flex-direction: column;
    gap: 10px;
    width: 100%;
}
.login__block-footer {
    display: flex;
    flex-direction: column;
    gap: 10px;
    width: 100%;
}
</style>