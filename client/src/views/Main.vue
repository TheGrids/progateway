<template>
    <div class="container">
        <div class="main__wrap">
            <div class="main__header">Успешная карьера вместе с нами</div>
            <div class="main__block">
                <div class="main__block__wrap">
                    <div class="main__block-header">Стань частью нашей команды</div>
                    <div class="main__block-content">Открой для себя новые возможности и стань неотъемлемой частью нашей сильной и дружной команды, где твоя мечта становится общей целью.</div>
                    <div class="main__block-footer">
                        <button class="btn white">Подать заявку</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "Main", 
}
</script>

<style scoped>
.main__wrap {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 50px;
}
.main__header {
    margin-top: 50px;
    font-size: 24px;
    font-weight: 500;
}
.main__block {
    width: 400px;
    background-color: #D8F3DC;
    border-radius: 20px;
}
.main__block__wrap {
    padding: 20px;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 25px;
}
.main__block-header {
    font-size: 18px;
    font-weight: 500;
}
.main__block-content {
    font-size: 16px;
    text-align: center;
}
.main__block-footer {
    width: 100%;
}
</style>