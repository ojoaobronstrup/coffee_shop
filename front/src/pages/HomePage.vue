<template>
    <el-container class="container">
        <header class="header">
            <el-avatar :src="this.$store.state.profileImage" :size="50" />
            <div class="header-location">
                <el-icon color="#AF795D" :size="20">
                    <LocationFilled />
                </el-icon>
                <h4>Brasil</h4>
            </div>
            <el-icon :size="20">
                <Bell />
            </el-icon>
        </header>
        <body>
            <h2>Bom dia, {{ this.$store.state.username }}</h2>
            <el-autocomplete/>
        </body>
    </el-container>
</template>

<script>
export default {
    data() {
        return {
            latitude: "",
            longitude: ""
        }
    },

    mounted() {
        this.GetUserPosition()
    },

    methods: {
        GetUserPosition() {
            if (navigator.geolocation) {
                navigator.geolocation.getCurrentPosition(
                    (position) => {
                        this.latitude = position.latitude
                        this.longitude = position.longitude
                    },
                    (error) => {
                        console.error(error)
                    }
                )
            } else {
                console.error("The browser don't have support to geolocation")
            }
        }
    }
}
</script>

<style>
.container {
    height: 100vh;
    width: 100vw;
    display: flex;
    flex-direction: column;
}

.header {
    display: flex;
    align-items: center;
    justify-content: space-around;
    width: 100vw;
}

.header-location {
    display: flex;
    align-items: center;
}
</style>