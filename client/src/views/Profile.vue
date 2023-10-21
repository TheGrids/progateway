<template>
    <div class="container">
        <div class="wrap">
            <div class="user">
                <div class="user__icon">
                    <svg width="96" height="96" viewBox="0 0 96 96" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M48.192 80.64C30.144 80.64 15.552 66.048 15.552 48C15.552 29.952 30.144 15.36 48.192 15.36C66.24 15.36 80.832 29.952 80.832 48C80.832 66.048 66.048 80.64 48.192 80.64ZM48.192 19.2C32.256 19.2 19.392 32.064 19.392 48C19.392 63.936 32.256 76.8 48.192 76.8C64.128 76.8 76.992 63.936 76.992 48C76.992 32.064 63.936 19.2 48.192 19.2Z" fill="black"/>
                        <path d="M29.376 71.616L25.92 70.08C26.88 67.776 29.952 66.432 33.216 64.896C36.48 63.36 40.512 61.632 40.512 59.52V56.64C39.36 55.68 37.44 53.568 37.056 50.496C36.096 49.536 34.56 47.808 34.56 45.504C34.56 44.16 35.136 43.008 35.52 42.24C35.136 40.704 34.752 37.824 34.752 35.52C34.752 28.032 39.936 23.04 48.192 23.04C50.496 23.04 53.376 23.616 54.912 25.344C58.56 26.112 61.632 30.336 61.632 35.52C61.632 38.784 61.056 41.472 60.672 42.816C61.056 43.392 61.44 44.352 61.44 45.504C61.44 48 60.096 49.728 58.944 50.496C58.56 53.568 56.832 55.488 55.68 56.448V59.52C55.68 61.248 59.136 62.592 62.208 63.744C65.856 65.088 69.696 66.624 71.04 69.696L67.392 71.04C66.816 69.504 63.744 68.352 60.864 67.392C56.64 65.856 51.84 64.128 51.84 59.712V54.72L52.8 54.144C52.8 54.144 55.104 52.608 55.104 49.536V48.192L56.256 47.616C56.448 47.616 57.408 47.04 57.408 45.504C57.408 45.12 57.024 44.544 56.832 44.352L56.064 43.584L56.448 42.624C56.448 42.624 57.408 39.552 57.408 35.712C57.408 32.064 55.296 29.376 53.568 29.376H52.416L51.84 28.416C51.84 27.648 50.496 26.88 48.192 26.88C42.24 26.88 38.592 30.144 38.592 35.52C38.592 38.016 39.552 42.24 39.552 42.24L39.744 43.2L38.976 44.16C38.784 44.16 38.4 44.736 38.4 45.504C38.4 46.464 39.552 47.616 40.128 48L40.896 48.576V49.536C40.896 52.416 43.392 53.952 43.392 54.144L44.352 54.72V59.712C44.352 64.32 39.36 66.624 34.752 68.544C32.64 69.312 29.76 70.656 29.376 71.616Z" fill="black"/>
                    </svg>
                </div>
                <div class="user__name">
                    <div class="user__f">Фамилия Имя</div>
                    <div class="user__l">test@gmail.com</div>
                </div>
            </div>


            <div class="requests">
                <div class="requests__header">Заявки</div>
                <div class="requests__content">

                    <table  style="overflow-x:auto;">
                        <thead>
                            <tr>
                                <th>Дата</th>
                                <th>Тип</th>
                                <th>Статус заявки</th>
                                <th>направление</th>
                                <th>Действия</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>
                                    <div class="wwrap">2023-10-20</div>
                                </td>
                                <td><div class="wwrap">Стажировка</div></td>
                                <td><div class="wwrap">Ожидает</div></td>
                                <td><div class="wwrap">FullStack разработчик</div></td>
                                <td><div class="wwrap"><button class="btn white">Просмотр заявки</button></div></td>
                            </tr>
                            <tr>
                                <td><div class="wwrap">2023-10-21</div></td>
                                <td><div class="wwrap">Практика</div></td>
                                <td><div class="wwrap">Подтверждено</div></td>
                                <td><div class="wwrap">FullStack разработчик</div></td>
                                <td><div class="wwrap"><button class="btn white">Просмотр заявки</button></div></td>
                            </tr>

                            
                        </tbody>
                    </table>

                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "Profile"
}
</script>

<style scoped>
.wrap {
    margin-top: 50px;
}
.user {
    display: flex;
    align-items: center;
    gap: 20px;
}
.user__icon {
    display: flex;
}
.user__name {
    display: flex;
    flex-direction: column;
    gap: 5px;
}
.user__f {
    font-size: 18px;
}
.user__l {
    font-size: 16px;
    opacity: .8;
}
.requests {
    display: flex;
    flex-direction: column;
    gap: 20px;
}
.requests__header {
    font-size: 24px;
    font-weight: 500;
    margin-top: 20px;
}
.content__wrap {
    padding: 20px;
    display: flex;
    justify-content: space-between;
}

table {
    width: 100%;
    border-collapse: separate;
    border-spacing: 0 20px;
}

th, td {
    padding: 20px;
    text-align: center;
    font-size: 18px;
}

th {
    opacity: .4;
}

td{
    background-color: #D8F3DC;
}

tr > td:first-child {
    border-top-left-radius: 10px;
    border-bottom-left-radius: 10px;
}

tr > td:last-child {
    border-top-right-radius: 10px;
    border-bottom-right-radius: 10px;
}
@media screen and (max-width: 768px){ 
    table {
        display: block;
    }
}
</style>