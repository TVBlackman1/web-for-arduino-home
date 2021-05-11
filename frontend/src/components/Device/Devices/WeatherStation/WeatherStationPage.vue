<template>
<!--    <p>Температура воздуха: {{ device.additional.temperature }} °C</p>-->
<!--    <p>Давление: {{ device.additional.air_pressure }}мм рт ст</p>-->
    <div class="content">
        <p class="category">Текущая температура воздуха:</p>
        <div :class="'val ' + classNames" :style="styles">
            {{ deviceData.additional.temperature }}°C
        </div>

        <div class="more">
            <header>Подробности:</header>
            <p>
                Указанный ГОСТ в зависимости от назначения помещения определяет допустимые значения температуры воздуха в жилых помещениях: от 18 до 25 градусов Цельсия.

                Так, например, если, по общему правилу, температура воздуха в жилой комнате квартире должны быть не менее 18 градусов Цельсия, то температура воздуха в ванной или совмещенном помещении уборной и ванной должна быть 25 градусов Цельсия.

                Допустимая температура воздуха в вестибюле, лестничной клетке, общем коридоре в квартирном доме составляет 16 градусов.

                Батареи должны прогревать помещение до указанных выше величин. В противном случае жильцы вправе рассчитывать на перерасчет платы за коммунальные услуги. Допустимое отклонение температуры воды в батареях в дневное время: 3° С. В ночное время суток (с 00:00 до 05:00) — до 5° С. Большее снижение температуры не допускается за исключением аварийных ситуаций.
            </p>
        </div>

        <p class="category">Влажность:</p>
        <div class="val">
            {{ deviceData.additional.humidity }}%
        </div>

        <div class="more">
            <header>Подробности:</header>
            <p>
                Разобраться, какая влажность должна быть в квартире, помогут разработанные на основе ГОСТ 30494-96 нормы. Согласно соответствующим статьям в СНиП и СанПиН, допустимый уровень зависит от времени года:

                в отопительный сезон нормой влажности в квартире считается значение до 45% при допустимом максимуме 60%;
                когда центральное отопление отключено, нормальной влажностью в квартире считают 30-60% при допустимом максимуме 65%            </p>
        </div>

    </div>
</template>

<script>
    export default {
        name: "WeatherStationPage",
        props: { device: Object },
        data() {
            return {
                deviceData: this.device,
                styles: {
                        "--colorDang": "red",
                        "--color": "black",

                },
                classNames: ""
            }
        },
        created() {
            // setTimeout(() => {
            //     this.deviceData.additional.temperature = 12
            // },  3000)
            // setTimeout(() => {
            //     this.deviceData.additional.air_pressure = 768
            // },  3000)
            // setTimeout(() => {
            //     this.deviceData.additional.temperature = 8
            //     if(this.deviceData.additional.temperature < 10) {
            //         this.classNames = "dang"
            //     }
            // },  6000)
            // setTimeout(() => {
            //     this.deviceData.additional.air_pressure = 771
            // },  6000)
        }
    }
</script>

<style scoped>
@import "../../../../css/pop-menu-style.css";
</style>