<template>
    <div class="container">
        <div class="wrap">{{this.message}}</div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: "Activate",
    data() {
        return {
            message: 'Загрузка...'
        }
    },
    mounted() {
        axios.get(process.env.VUE_APP_API + 'activate/' + this.$route.params.token).then(res => {
            this.message = res.data.msg
            setTimeout(function() {
                window.location.href = '/login';
            }, 2000);
        }).catch(err => {
            this.message = err.response.data.msg
        })
    }
}
</script>

<style scoped>
.wrap {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 50px;
    font-size: 24px;
}
</style>