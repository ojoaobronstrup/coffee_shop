<template>
    <header class="header">
        <el-avatar :src="this.$store.state.profileImage" />
        <div class="header-location">
            <el-icon>
                <Location color="#AF795D" />
            </el-icon>
            <h4>teste</h4>
        </div>
        <el-icon>
            <Bell color="#AF795D" />
        </el-icon>
    </header>
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
.header {
    display: flex;
    align-items: center;
    justify-content: space-around;
}

.header-location {
    display: flex;
    align-items: center;
}
</style>