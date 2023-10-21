<template>
    <div class="container">
        <div class="container__wrap">
            <div class="ab">
                <div class="left">
                    <div class="left__header">Доступные направления для <br> практики</div>
                    <div class="left__content">
                        <ul class="left__list">
                            <li class="list-item">1С-разработчик</li>
                            <li class="list-item">Веб-разработчик</li>
                            <li class="list-item">Frontend-разработчик</li>
                            <li class="list-item">Разработчик баз данных</li>
                            <li class="list-item">java-разработчик</li>
                        </ul>
                    </div>
                </div>
                <div class="right">
                    <div class="right__header">Как оформить практику</div>
                    <ul class="right__list">
                        <li><span class="list-number">1</span>Заполни заявку на практику</li>
                        <li><span class="list-number">2</span>Жди письмо с обратной связью от кураторапо практике в течение недели</li>
                        <li><span class="list-number">3</span>Если заявка на практику одобрена, запросив вузе направление на практику и отправькуратору в ответном письме</li>
                        <li><span class="list-number">4</span>Жди письмо с информацией по практике,приказом на практику, контактами оруководителе практики</li>
                    </ul>
                </div>
            </div>
            <button class="btn green w">Подать заявку </button>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Internship'
}
</script>

<style scoped>
@media screen and (max-width: 768px){  
    .ab {
        flex-direction: column;
    }
    .left {
        min-width: 200px!important;
    }
}
.a {
    padding: 8px 15px;
    border-radius: 10px;
    background: #3AA935;
}
.container {
    background: linear-gradient(180deg, rgba(58, 169, 53, 0.26) 0%, rgba(58, 169, 53, 0.00) 80%);
}
.container__wrap {
    min-height: 700px;
    padding: 20px;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 30px;
}
.ab {
    display: flex;
    justify-content: center;
    gap: 100px;
}
.w {
    max-width: 300px;
}
.left {
    display: flex;
    flex-direction: column;
    gap: 30px;
    min-width: 500px;
}
.left__header {
    text-align: center;
    font-size: 20px;
    font-weight: 600;
}
.left__list {
    width: 100%;
    display: flex;
    flex-direction: column;
    gap: 30px;
}
.list-item {
    display: flex;
    padding: 15px;
    justify-content: center;
    align-items: center;
    gap: 10px;

    border-radius: 10px;
    background: #FFF;

    /* my shadow */
    box-shadow: 0px 5px 20px 0px rgba(0, 0, 0, 0.10);
}
.right {
    display: flex;
    flex-direction: column;
    gap: 30px;
    max-width: 300px;
}
.right__header {
    font-size: 20px;
    font-weight: 600;
}
.right__list {
    display: flex;
    flex-direction: column;
    gap: 25px;
}
.right__list li {
    display: flex; 
    align-items: center;
    gap: 10px;
}
.list-number {
    color: #3AA935;
    font-size: 22px;
    font-weight: 600;
}
</style>