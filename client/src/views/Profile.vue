<template>
    <div class="container">
        <div class="wrap">

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