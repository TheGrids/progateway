<template>
    <div class="container">
        <div class="wrap">
            <div class="block">
                <div class="block__wrap">
                    <div class="block__header">Доступные направления</div>
                    <div class="block__content">
                        <ul>
                            <li>- 1С-разработчик</li>
                            <li>- Веб-разработчик</li>
                            <li>- Frontend-разработчик</li>
                            <li>- Разработчик БД</li>
                            <li>- java-разработчик</li>
                        </ul>
                    </div>
                    <div class="block__footer">
                        <button class="btn white" v-on:click="go(1)">Подать заявку на стажировку</button>
                    </div>
                </div>
            </div>

            <div class="block">
                <div class="block__wrap">
                    <div class="block__header">Доступные направления</div>
                    <div class="block__content">
                        <ul>
                            <li>-1С-разработчик</li>
                            <li>- Веб-разработчик</li>
                            <li>- Frontend-разработчик</li>
                            <li>- Разработчик БД</li>
                            <li>- java-разработчик</li>
                        </ul>
                    </div>
                    <div class="block__footer">
                        <button class="btn white"  v-on:click="go(2)">Подать заявку на практику</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "HOME",
    methods: {
        go(id) {
            if(id === 1) this.$router.push('/internship')
            if(id === 2) this.$router.push('/practice')
        }
    }
}
</script>

<style scoped>
.wrap {
    margin-top: 50px;
    display: flex;
    justify-content: center;
    gap: 20px;
}
.block {
    width: 400px;
    background-color: #D8F3DC;
    border-radius: 20px;
}
.block__wrap {
    padding: 20px;
    display: flex;
    flex-direction: column;
    gap: 20px;
}
.block__header {
    text-align: center;
    font-size: 18px;
    font-weight: 500;
}
.block__content {
    
}
.block__content ul {
    display: flex;
    flex-direction: column;
    gap: 10px;
}
.block__content li {
    font-size: 16px;
}

@media screen and (max-width: 768px){ 
    .wrap {
        flex-direction: column;
    }
    .block {
        width: 100%;
    }
}
</style>